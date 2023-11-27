package sword

func levelOrder2(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		var path []int
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			path = append(path, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, path)
	}
	return
}
