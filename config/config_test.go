package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig("../config.toml")
	fmt.Printf("%v\n", config)
}
