package prefix

import "fmt"

// 一维前缀和
func prefix() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 计算前缀和
	sum := make([]int, len(a))
	sum[0] = a[0]
	for i := 1; i < len(a); i++ {
		sum[i] = sum[i-1] + a[i]
	}
	fmt.Println(sum) // [1 3 6 10 15 21 28 36]
	// 现在我要询问 [2,4]，那就是打印sum[4] - sum[2-1]
	// 询问[L,R]，打印sum[R] - sum[L-1]
	// 但是对于[0,4]这种，要么单独拎出来，返回sum[4]，要么sum统一错一位
	sum2 := make([]int, len(a)+1) // 加一是为了统一错一位
	for i, v := range a {
		// sum[0]直接为0
		sum2[i+1] = sum2[i] + v
	}
	fmt.Println(sum2) // [0 1 3 6 10 15 21 28 36]
	// 询问[L,R]，打印sum[R+1] - sum[L]
}

// 二维前缀和
func prefix2() {
	// 二维前缀和
	q := [][]int{
		{2, 3, 23, 5},
		{6, 33, 7, 45},
		{23, 4, 9, 56},
		{34, 57, 78, 75},
	}
	fmt.Println(q)
	sum := make([][]int, len(q))
	for i := range sum {
		sum[i] = make([]int, len(q[0]))
	}
	// 先计算两边的
	sum[0][0] = q[0][0]
	for i := 1; i < len(q); i++ {
		sum[i][0] = sum[i-1][0] + q[i][0]
	}
	for i := 1; i < len(q[0]); i++ {
		sum[0][i] = sum[0][i-1] + q[0][i]
	}
	// 计算中间的
	for i := 1; i < len(q); i++ {
		for j := 1; j < len(q[0]); j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + q[i][j]
		}
	}
	for _, v := range sum {
		fmt.Println(v)
	}
	// 计算结果
	calc := func(l1, r1, l2, r2 int) int {
		if l1 == 0 && r1 == 0 {
			return sum[l2][r2]
		}
		if l1 == 0 {
			return sum[l2][r2] - sum[l2][r1-1]
		}
		if r1 == 0 {
			return sum[l2][r2] - sum[l1-1][r2]
		}
		return sum[l2][r2] - sum[l1-1][r2] - sum[l2][r1-1] + sum[l1-1][r1-1]
	}
	fmt.Println(calc(1, 1, 3, 2))
}
