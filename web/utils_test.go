package web

import (
	"net/http"
	"testing"
)

func TestUnmarshalRequest(t *testing.T) {
	t.Fatal("TODO")
	type args struct {
		r   *http.Request
		dto interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnmarshalRequest(tt.args.r, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteResponse(t *testing.T) {
	t.Fatal("TODO")

	type args struct {
		w    http.ResponseWriter
		code int
		dto  interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
