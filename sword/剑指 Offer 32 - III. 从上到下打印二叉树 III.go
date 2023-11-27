package sword

func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	var round int
	for len(q) != 0 {
		var path []int
		j := len(q) - 1
		for i := 0; i <= j; i++ {
			var n *TreeNode
			n = q[0]
			q = q[1:]
			path = append(path, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		if round%2 == 1 {
			for l, r := 0, len(path)-1; l < r; {
				path[l], path[r] = path[r], path[l]
				l++
				r--
			}
		}
		res = append(res, path)
		round++
	}
	return
}
