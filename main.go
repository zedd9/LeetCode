package main

import (
	"fmt"

	"github.com/zedd9/LeetCode/SubtreeOfAnotherTree"
)

func main() {
	node0 := SubtreeOfAnotherTree.TreeNode{Val: 0}
	node1 := SubtreeOfAnotherTree.TreeNode{Val: 1}
	node2 := SubtreeOfAnotherTree.TreeNode{Val: 2}
	node3 := SubtreeOfAnotherTree.TreeNode{Val: 3}
	node4 := SubtreeOfAnotherTree.TreeNode{Val: 4}
	node5 := SubtreeOfAnotherTree.TreeNode{Val: 5}

	node3.Left = &node4
	node3.Right = &node5
	node4.Left = &node1
	node4.Right = &node2
	node1.Left = &node0

	node41 := SubtreeOfAnotherTree.TreeNode{Val: 1}
	node42 := SubtreeOfAnotherTree.TreeNode{Val: 2}
	node44 := SubtreeOfAnotherTree.TreeNode{Val: 4}

	node44.Left = &node41
	node44.Right = &node42

	result := SubtreeOfAnotherTree.IsSubtree(&node3, &node44)
	fmt.Print(result)
}
