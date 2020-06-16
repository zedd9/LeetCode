package sumuptozero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)

	output := sumZero(4)
	assert.Equal(4, len(output))
	sum := 0
	for _, i := range output {
		sum += i
	}

	assert.Equal(sum, 0)

	output = sumZero(3)
	assert.Equal(3, len(output))
	sum = 0
	for _, i := range output {
		sum += i
	}
	assert.Equal(sum, 0)

	output = sumZero(6)
	assert.Equal(6, len(output))
	sum = 0
	for _, i := range output {
		sum += i
	}
	assert.Equal(sum, 0)
}
