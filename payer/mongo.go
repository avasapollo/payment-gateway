package payer

import (
	"context"
	"time"

	"github.com/avasapollo/payment-gateway/payments"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collTransactions = "transactions"

// assertion the interface Payer to MongoPayer
var _ payments.Payer = (*MongoPayer)(nil)

// Options to customize the MongoDB connection
type Options struct {
	config *Config
}

type Option func(opts *Options)

func WithMongoDBUrl(uri string) Option {
	return func(opts *Options) {
		opts.config.MongoDBUrl = uri
	}
}

func WithMongoDBDatabaseName(dbName string) Option {
	return func(opts *Options) {
		opts.config.DatabaseName = dbName
	}
}

// Config to customize the MongoDB connection
type Config struct {
	MongoDBUrl   string `envconfig:"MONGODB_URL" default:"mongodb://localhost:27017"`
	DatabaseName string `envconfig:"MONGODB_DATABASE_NAME" default:"payment-gateway"`
	MaxPoolSize  uint64 `envconfig:"MONGODB_MAX_POOL_SIZE" default:"50"`
}

// defaultOptions to set the default connection to MongoDB
func defaultOptions() *Options {
	c := new(Config)
	_ = envconfig.Process("", c)
	return &Options{config: c}
}

func clientOptions(config *Config) *options.ClientOptions {
	opts := options.Client()
	opts.SetMaxPoolSize(config.MaxPoolSize)
	opts.ApplyURI(config.MongoDBUrl)
	return opts
}

type MongoPayer struct {
	lgr          *logrus.Entry
	client       *mongo.Client
	database     *mongo.Database
	transactions *mongo.Collection
}

func New(opts ...Option) (*MongoPayer, error) {
	c := defaultOptions()

	for _, opt := range opts {
		opt(c)
	}

	client, err := mongo.Connect(context.Background(), clientOptions(c.config))
	if err != nil {
		return nil, err
	}

	db := client.Database(c.config.DatabaseName)

	return &MongoPayer{
		lgr:          logrus.WithField("pkg", "payer"),
		client:       client,
		database:     db,
		transactions: db.Collection(collTransactions),
	}, nil
}

func (m *MongoPayer) Authorize(ctx context.Context, req *payments.AuthorizeReq) (*payments.Transaction, error) {
	if req.Card.CardNumber == "4000000000000119" {
		return nil, payments.ErrAuthFailed
	}
	dto := toTransactionDtoFromAuthorizeReq(req)
	_, err := m.transactions.InsertOne(ctx, dto)
	if err != nil {
		return nil, err
	}
	return toTransaction(dto), nil
}

func (m *MongoPayer) Void(ctx context.Context, req *payments.VoidReq) (*payments.Transaction, error) {
	now := time.Now().UTC().Truncate(time.Millisecond)
	filter := bson.M{
		"_id":    req.AuthorizationID,
		"status": payments.Authorize.String(),
	}
	updates := bson.M{
		"$set": bson.M{
			"status":     payments.Void.String(),
			"updated_at": now,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	dto := new(TransactionDto)
	if err := m.transactions.FindOneAndUpdate(ctx, filter, updates, opts).Decode(dto); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, payments.ErrVoidFailed
		}
		return nil, err
	}
	return toTransaction(dto), nil
}

func (m *MongoPayer) Capture(ctx context.Context, req *payments.CaptureReq) (*payments.Transaction, error) {
	now := time.Now().UTC().Truncate(time.Millisecond)
	filter := bson.M{
		"_id":                  req.AuthorizationID,
		"status":               bson.M{"$in": []string{payments.Authorize.String(), payments.Capture.String()}},
		"capture_amount.value": bson.M{"$gte": req.Amount},
	}
	updates := bson.M{
		"$set": bson.M{
			"status":     payments.Capture.String(),
			"updated_at": now,
		},
		"$inc": bson.M{
			"capture_amount.value": -req.Amount,
			"refund_amount.value":  req.Amount,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	dto := new(TransactionDto)
	if err := m.transactions.FindOneAndUpdate(ctx, filter, updates, opts).Decode(dto); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, payments.ErrCaptureFailed
		}
		return nil, err
	}
	if dto.CardNumber == "4000000000000259" {
		// TODO: think roll back the operation
		return nil, payments.ErrCaptureFailed
	}
	return toTransaction(dto), nil
}

func (m *MongoPayer) Refund(ctx context.Context, req *payments.RefundReq) (*payments.Transaction, error) {
	now := time.Now().UTC().Truncate(time.Millisecond)
	filter := bson.M{
		"_id":                 req.AuthorizationID,
		"status":              bson.M{"$in": []string{payments.Refund.String(), payments.Capture.String()}},
		"refund_amount.value": bson.M{"$gte": req.Amount},
	}
	updates := bson.M{
		"$set": bson.M{
			"status":     payments.Refund.String(),
			"updated_at": now,
		},
		"$inc": bson.M{
			"refund_amount.value": -req.Amount,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	dto := new(TransactionDto)
	if err := m.transactions.FindOneAndUpdate(ctx, filter, updates, opts).Decode(dto); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, payments.ErrRefundFailed
		}
		return nil, err
	}
	if dto.CardNumber == "4000000000003238" {
		// TODO: think roll back the operation
		return nil, payments.ErrRefundFailed
	}
	return toTransaction(dto), nil
}
