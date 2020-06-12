package SubtreeOfAnotherTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	return s != nil && (isEqual(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t))
}

func isEqual(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}
