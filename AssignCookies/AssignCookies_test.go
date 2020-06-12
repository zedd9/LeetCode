package AssignCookies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)

	i1 := []int{1, 2, 3}
	i2 := []int{1, 1}

	output := findContentChildren(i1, i2)
	assert.Equal(1, output)

	i1 = []int{1, 2}
	i2 = []int{1, 2, 3}
	output = findContentChildren(i1, i2)
	assert.Equal(2, output)
}
