package secondminimum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)

	node11 := TreeNode{Val: 2}
	node12 := TreeNode{Val: 2}
	node13 := TreeNode{Val: 5}
	node16 := TreeNode{Val: 5}
	node17 := TreeNode{Val: 7}

	node11.Left = &node12
	node11.Right = &node13

	node13.Left = &node16
	node13.Right = &node17

	output := findSecondMinimumValue(&node11)
	assert.Equal(5, output)

	node21 := TreeNode{Val: 2}
	node22 := TreeNode{Val: 2}
	node23 := TreeNode{Val: 2}

	node21.Left = &node22
	node22.Right = &node23

	output = findSecondMinimumValue(&node21)
	assert.Equal(-1, output)
}
