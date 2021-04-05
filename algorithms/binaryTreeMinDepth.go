func minDepth(root *TreeNode) int {
    switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + MinOfTwo(minDepth(root.Left), minDepth(root.Right))
	}
}

func MinOfTwo(a, b int) int {
	if a <= b {
		return a
	}
	return b
}