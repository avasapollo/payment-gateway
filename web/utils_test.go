package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalRequest(t *testing.T) {
	req := &CaptureReq{
		AuthorizationID: "id_1",
		Amount:          100,
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)
	r, err := http.NewRequest(http.MethodPost, "http://localhost/v1/capture", bytes.NewBuffer(b))
	require.NoError(t, err)
	res := new(CaptureReq)
	err = UnmarshalRequest(r, res)
	require.NoError(t, err)
	require.Equal(t, req, res)
}

func TestWriteResponse(t *testing.T) {
	w := httptest.NewRecorder()
	WriteResponse(w, http.StatusOK, nil)
	require.Equal(t, http.StatusOK, w.Code)
}
