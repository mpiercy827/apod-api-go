package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	srv := New()
	assert.Equal(t, srv.Addr, ":4000")
}
