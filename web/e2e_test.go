// +build integration

package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/avasapollo/payment-gateway/payer"
	"github.com/avasapollo/payment-gateway/payments"
	"github.com/stretchr/testify/require"
)

func Test_E2E(t *testing.T) {
	type result struct {
		AuthorizationID string `json:"authorization_id"`
		Amount          struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"amount"`
	}

	mongoPayer, err := payer.New()
	require.NoError(t, err)
	payment := payments.New(mongoPayer)
	srv := NewServer(payment)
	go srv.ListenAndServe()

	// just to be sure the grpc gateway is up
	time.Sleep(1 * time.Second)

	t.Run("ok health", func(t *testing.T) {
		r, err := http.Get("http://localhost:8080/health")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, r.StatusCode)
	})

	t.Run("ok authorize", func(t *testing.T) {
		req := `
				{
					"card": {
						"name": "Andrea Vasapollo",
						"card_number": "4242424242424242",
						"expire_month": "12",
						"expire_year": "2022",
						"cvv": "123"
					},
					"amount": {
						"value": 100,
						"currency": "EUR"
					}
					
				}
			`
		resp, err := http.Post("http://localhost:8080/v1/authorize", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		result := new(result)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(100), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)
	})

	t.Run("ok void", func(t *testing.T) {
		req := `
				{
					"card": {
						"name": "Andrea Vasapollo",
						"card_number": "4242424242424242",
						"expire_month": "12",
						"expire_year": "2022",
						"cvv": "123"
					},
					"amount": {
						"value": 100,
						"currency": "EUR"
					}
					
				}
			`
		resp, err := http.Post("http://localhost:8080/v1/authorize", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		result := new(result)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(100), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)

		// void
		req = `
				{
					"authorization_id": "` + result.AuthorizationID + `"
				}
			`
		resp, err = http.Post("http://localhost:8080/v1/void", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(100), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)
	})

	t.Run("ok capture", func(t *testing.T) {
		req := `
				{
					"card": {
						"name": "Andrea Vasapollo",
						"card_number": "4242424242424242",
						"expire_month": "12",
						"expire_year": "2022",
						"cvv": "123"
					},
					"amount": {
						"value": 100,
						"currency": "EUR"
					}
					
				}
			`
		resp, err := http.Post("http://localhost:8080/v1/authorize", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		result := new(result)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(100), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)

		// capture
		req = `
				{
					"authorization_id": "` + result.AuthorizationID + `",
                    "amount": 100
				}
			`
		resp, err = http.Post("http://localhost:8080/v1/capture", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(0), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)
	})

	t.Run("ok refund", func(t *testing.T) {
		req := `
				{
					"card": {
						"name": "Andrea Vasapollo",
						"card_number": "4242424242424242",
						"expire_month": "12",
						"expire_year": "2022",
						"cvv": "123"
					},
					"amount": {
						"value": 100,
						"currency": "EUR"
					}
					
				}
			`
		resp, err := http.Post("http://localhost:8080/v1/authorize", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		result := new(result)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(100), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)

		// capture
		req = `
				{
					"authorization_id": "` + result.AuthorizationID + `",
                    "amount": 100
				}
			`
		resp, err = http.Post("http://localhost:8080/v1/capture", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(0), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)

		// refund
		req = `
				{
					"authorization_id": "` + result.AuthorizationID + `",
                    "amount": 100
				}
			`
		resp, err = http.Post("http://localhost:8080/v1/refund", "application/json", bytes.NewBufferString(req))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		require.NotEmpty(t, result.AuthorizationID)
		require.Equal(t, float64(0), result.Amount.Value)
		require.Equal(t, "EUR", result.Amount.Currency)
	})
}
