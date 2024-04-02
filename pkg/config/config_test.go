package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	// go test pkg/config/*.go -config=$(pwd)/config.toml
	if err := InitConfig(); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "12424", Conf.System.Addr)
}
