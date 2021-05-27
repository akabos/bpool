package bpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByteSlice(t *testing.T) {
	p := GetByteSlice()
	assert.Equal(t, DefaultByteSlicePoolCapacity, cap(p))
}

func TestPutByteSlice(t *testing.T) {
	p := make([]byte, 0, 0)
	PutByteSlice(p)

	p = GetByteSlice()
	assert.Equal(t, 0, cap(p))
}