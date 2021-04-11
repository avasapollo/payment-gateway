package payments

import "testing"

func Test_truncateFloat64(t *testing.T) {
	type args struct {
		k float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ok 3.1211 to 3.12",
			args: args{
				k: 3.1211,
			},
			want: 3.12,
		},
		{
			name: "ok 3.1299 to 3.12",
			args: args{
				k: 3.1299,
			},
			want: 3.12,
		},
		{
			name: "ok 10.1299 to 10.12",
			args: args{
				k: 10.1299,
			},
			want: 10.12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateFloat(tt.args.k); got != tt.want {
				t.Errorf("truncateFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
