package config

import (
	"testing"
)

func TestSecret(t *testing.T) {
	SetSecret("hello")
	if secretKey != "hello" || Secret() != "hello" {
		t.Error("config did not store the right secret key")
	}
}
