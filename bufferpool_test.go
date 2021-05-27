package bpool

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuffer(t *testing.T) {
	b := GetBuffer()
	assert.NotNil(t, b)
	assert.Equal(t, 0, b.Len())
}

func TestPutBuffer(t *testing.T) {
	b := bytes.NewBuffer(nil)
	PutBuffer(b)
}