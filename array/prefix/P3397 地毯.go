package prefix

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n, m int
	fmt.Fscanln(in, &n, &m)
	sub := make([][]int, n+1)
	for i := range sub {
		sub[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		var x1, x2, y1, y2 int
		fmt.Fscanln(in, &x1, &y1, &x2, &y2)
		sub[x1-1][y1-1] += 1
		sub[x1-1][y2] -= 1
		sub[x2][y1-1] -= 1
		sub[x2][y2] += 1
	}
	for i := 1; i < n; i++ {
		sub[i][0] += sub[i-1][0]
	}
	for j := 1; j < n; j++ {
		sub[0][j] += sub[0][j-1]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			sub[i][j] += -sub[i-1][j-1] + sub[i-1][j] + sub[i][j-1]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(out, "%d ", sub[i][j])
		}
		fmt.Fprintln(out)
	}
	out.Flush()
}
