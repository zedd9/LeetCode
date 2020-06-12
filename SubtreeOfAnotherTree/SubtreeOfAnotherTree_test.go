package SubtreeOfAnotherTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)
	node0 := TreeNode{Val: 0}
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}

	node3.Left = &node4
	node3.Right = &node5
	node4.Left = &node1
	node4.Right = &node2
	node1.Left = &node0

	node41 := TreeNode{Val: 1}
	node42 := TreeNode{Val: 2}
	node44 := TreeNode{Val: 4}

	node44.Left = &node41
	node44.Right = &node42

	result := IsSubtree(&node3, &node44)

	assert.Equal(false, result)
}

func TestCase2(t *testing.T) {
	assert := assert.New(t)
	node10 := TreeNode{Val: 1}
	node11 := TreeNode{Val: 1}

	node10.Left = &node11

	node21 := TreeNode{Val: 1}

	result := IsSubtree(&node10, &node21)

	assert.Equal(true, result)
}

func TestCase3(t *testing.T) {
	assert := assert.New(t)
	node11 := TreeNode{Val: 1}
	node12 := TreeNode{Val: 1}
	node13 := TreeNode{Val: 1}
	node14 := TreeNode{Val: 1}
	node15 := TreeNode{Val: 1}

	node13.Left = &node14
	node13.Right = &node15

	node14.Left = &node11
	node15.Left = &node12

	node21 := TreeNode{Val: 1}
	node22 := TreeNode{Val: 2}
	node23 := TreeNode{Val: 3}

	node23.Left = &node21
	node23.Right = &node22

	result := IsSubtree(&node13, &node23)

	assert.Equal(false, result)
}
