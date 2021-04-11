package web

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/avasapollo/payment-gateway/web/proto/v1"
	"google.golang.org/grpc/reflection"

	// Update
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	lgr                  *logrus.Entry
	grpcServer           *grpc.Server
	httpServer           *http.ServeMux
	paymentGatewayServer v1.PaymentGatewayServer
}

func NewServer(payment Payment) *Server {
	return &Server{
		lgr:                  logrus.WithField("pkg", "web"),
		grpcServer:           grpc.NewServer(),
		httpServer:           http.NewServeMux(),
		paymentGatewayServer: NewPaymentGatewayService(payment),
	}
}

func (srv *Server) ListenAndServe() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	jsonpb := &runtime.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	v1.RegisterPaymentGatewayServer(srv.grpcServer, srv.paymentGatewayServer)
	reflection.Register(srv.grpcServer)

	go srv.grpcServer.Serve(lis)
	go waitForGracefulShutdown(srv.grpcServer)

	if err := v1.RegisterPaymentGatewayHandlerFromEndpoint(ctx, mux, "localhost:50051", opts); err != nil {
		return err
	}
	srv.httpServer.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/proto/v1/payment-gateway.swagger.json")
	})
	srv.httpServer.Handle("/", mux)
	return http.ListenAndServe(":8080", srv.httpServer)
}

func waitForGracefulShutdown(srv *grpc.Server) {
	lgr := logrus.WithField("pkg", "web")
	lgr.Info("Grpc messaging server started ...")
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.GracefulStop()

	lgr.Info("Shutting down grpc messaging server.")
	os.Exit(0)
}
