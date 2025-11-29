package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, k int, res []int) {
	fmt.Fscan(reader, &n, &k)
	res = solve(n, k)
	return
}

func solve(n int, k int) []int {
	if k == 1 {
		return []int{n}
	}
	// 避免k太大，溢出
	if n*2/k < k+1 {
		return nil
	}

	sum := (1 + k) * k / 2

	if sum > n {
		return nil
	}
	// k最多为n的平方根
	// 如果n是个质数
	// 那么就完蛋了

	var g int

	check := func(i int) {
		if i <= n/sum && i*sum <= n {
			g = max(g, i)
		}
	}

	for i := 1; i <= n/i; i++ {
		if n%i == 0 {
			check(i)
			check(n / i)
		}
	}

	// n % g == 0
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = g * (i + 1)
		n -= res[i]
	}
	res[k-1] += n

	return res
}
