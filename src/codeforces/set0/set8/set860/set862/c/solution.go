package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x int
	fmt.Fscan(reader, &n, &x)

	res := solve(n, x)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int, x int) []int {
	if n == 1 {
		return []int{x}
	}

	if n == 2 {
		if x == 0 {
			return nil
		}
	}

	res := make([]int, n)

	var sum int
	for i := 1; i < n; i++ {
		res[i] = i
		sum ^= i
	}

	if x == sum {
		return res
	}

	sum ^= (n - 1)
	res[n-1] = 0
	y := sum ^ x

	if y >= n {
		res[n-1] = y
		return res
	}
	sum ^= y
	// sum = x
	// sum ^ y = x
	// y < n, 能找到两个数 u ^ v = y, 且u >= n, v >= n
	h := bits.Len(uint(n))
	// 那么要找到三个数 a ^ b ^ c = 0
	// 110000
	// 010001
	// 100001
	// 如果 y = 0
	if y > 0 {
		res[0] = 1<<(h+1) + 1<<(h+2)
		res[y] = 1<<(h+1) + 1
		res[n-1] = 1<<(h+2) + 1
	} else {
		// 1100000
		// 0100011
		// 1000010
		res[0] = 1<<(h+1) + 1<<(h+2)
		res[1] = 1<<(h+1) + 3
		res[n-1] = 1<<(h+2) + 2
	}

	return res
}
