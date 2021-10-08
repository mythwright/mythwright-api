package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()

	assert.Equal(t, ":8000", s.s.Addr)
}
