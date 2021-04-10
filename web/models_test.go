package web

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newErrResp(t *testing.T) {
	got := newErrResp("200", "err_br1")
	require.Equal(t, "200", got.Code)
	require.Equal(t, "err_br1", got.Message)
}
