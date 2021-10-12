package api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer(context.Background())

	assert.Equal(t, ":8000", s.s.Addr)
}
