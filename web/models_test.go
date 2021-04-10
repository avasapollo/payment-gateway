package web

import (
	"reflect"
	"testing"
)

func Test_newErrResp(t *testing.T) {
	t.Fatal("TODO")
	type args struct {
		code    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *ErrorResp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newErrResp(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newErrResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
