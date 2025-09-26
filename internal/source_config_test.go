package internal

import "testing"

func TestSourceConfig(t *testing.T) {
	c, _ := NewSourceConfig("qweqweqwe")

	_ = c.ComposeQueryParams()
	_ = c.ComposeConnectionUrl()
}

func TestNewSourceConfig(t *testing.T) {
	NewSourceConfig("asdasd")
}
