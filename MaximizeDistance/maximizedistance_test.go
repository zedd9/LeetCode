package maximzedistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)

	seats := []int{1, 0, 0, 0, 1, 0, 1}
	output := maxDistToClosest(seats)
	assert.Equal(2, output)

	seats = []int{1, 0, 0, 0}
	output = maxDistToClosest(seats)
	assert.Equal(3, output)
}
