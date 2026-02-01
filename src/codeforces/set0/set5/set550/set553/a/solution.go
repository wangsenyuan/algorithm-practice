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

func drive(reader *bufio.Reader) int {
	var k int
	fmt.Fscan(reader, &k)
	c := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &c[i])
	}
	return solve(c)
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(c []int) int {
	var n int
	for _, v := range c {
		n += v
	}
	k := len(c)
	// n - k 个球可以随便放，剩下的球按照顺序放置, 找到最后
	C := make([][]int, n+1)
	for i := range n + 1 {
		C[i] = make([]int, n+1)
		C[i][0] = 1
		C[i][i] = 1

		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j], C[i-1][j-1])
		}
	}
	res := 1
	for i := k - 1; i >= 0; i-- {
		// 必须选择一个放在最后空的位置上，剩余的v-1个，在其他位置放置
		v := c[i]
		cur := C[n-1][v-1]
		res = mul(res, cur)
		n -= v
	}
	return res
}
