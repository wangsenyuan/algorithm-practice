package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	A := make([][]int, n)
	for i := range n {
		A[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &A[i][j])
		}
	}
	var m int
	fmt.Fscan(reader, &m)
	ops := make([][]int, m)
	for i := range m {
		var k int
		fmt.Fscan(reader, &k)
		if k == 3 {
			ops[i] = []int{3}
		} else {
			var j int
			fmt.Fscan(reader, &j)
			ops[i] = []int{k, j}
		}
	}
	return solve(A, ops)
}

func solve(A [][]int, ops [][]int) string {
	n := len(A)

	// A[i][j] 会在 计算第i行i列, 或者j行j列时,产生贡献
	// 如果 i != j, 它的贡献始终是2次,所以抵消了
	// 只要 i = j 上的才有意义
	var sum int
	for i := range n {
		sum ^= A[i][i]
	}

	var ans []byte
	for _, cur := range ops {
		if cur[0] == 3 {
			ans = append(ans, byte('0'+sum))
		} else {
			sum ^= 1
		}
	}

	return string(ans)
}
