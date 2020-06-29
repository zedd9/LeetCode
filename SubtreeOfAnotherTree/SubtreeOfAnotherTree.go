package SubtreeOfAnotherTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	return (isEqual(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t))
}

func isEqual(s *TreeNode, t *TreeNode) bool {
	return s.Val == t.Val && isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}
