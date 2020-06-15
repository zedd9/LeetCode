package secondminimum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	secondNum := -1
	findSecond(root.Left, root.Val, &secondNum)
	findSecond(root.Right, root.Val, &secondNum)

	return secondNum
}

func findSecond(node *TreeNode, FirstValue int, SecondValue *int) {
	if node == nil {
		return
	}

	if node.Val > FirstValue {
		if *SecondValue == -1 {
			*SecondValue = node.Val
		} else if *SecondValue > node.Val {
			*SecondValue = node.Val
		}
	}

	findSecond(node.Left, FirstValue, SecondValue)
	findSecond(node.Right, FirstValue, SecondValue)
}
