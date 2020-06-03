package Binarygap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase(t *testing.T) {
	assert := assert.New(t)

	result := BinaryGap(22)
	assert.Equal(result, 2)

	result = BinaryGap(5)
	assert.Equal(result, 2)

	result = BinaryGap(6)
	assert.Equal(result, 1)

	result = BinaryGap(8)
	assert.Equal(result, 0)
}
