package constructrectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)

	output := constructRectangle(4)
	assert.Equal(2, len(output))
	assert.Equal(2, output[0])
	assert.Equal(2, output[1])

	output = constructRectangle(3)
	assert.Equal(2, len(output))
	assert.Equal(3, output[0])
	assert.Equal(1, output[1])

	output = constructRectangle(6)
	assert.Equal(2, len(output))
	assert.Equal(3, output[0])
	assert.Equal(2, output[1])
}
