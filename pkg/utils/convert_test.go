package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "english", args: args{s: "string"}, want: []byte("string")},
		{name: "chinese", args: args{s: "中文"}, want: []byte("中文")},
		{name: "empty str", args: args{s: ""}, want: nil},
		{name: "operation ", args: args{s: "+-*/"}, want: []byte("+-*/")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, StringToBytes(tt.args.s))
		})
	}
}
