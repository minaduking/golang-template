package sample_test

import (
	"golang-template/sample"
	"testing"
)

func TestReturnString(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success returing string",
			args: args{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := sample.ReturnString()
			if req != "Hello, World!" {
				t.Errorf("miss error")
			}
		})
	}
}
