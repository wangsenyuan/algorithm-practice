package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) bool {
	freq := make([]int, 4)
	var tot int
	for _, num := range a {
		freq[num%4]++
		tot += num / 4
	}

	m := 4 * tot

	mid := n
	side := n * 2

	if tot < n {
		mid = n - tot
	} else {
		mid = 0
		side -= (tot - n) * 2
	}

	// freq[3]的肯定会浪费一个, 就把它们拆成2和1
	freq[2] += freq[3]
	freq[1] += freq[3]
	freq[3] = 0

	x := min(freq[2], side)
	freq[2] -= x
	side -= x
	m += 2 * x

	// 在浪费一个位置的情况下，尽量和2配对
	x = min(freq[1], freq[2], mid)
	freq[1] -= x
	freq[2] -= x
	m += 4 * x
	mid -= x

	// 2的只能放中间了，但是肯定会浪费2个位置，这两个位置可以免费放置1
	x = min(freq[2], mid)
	freq[2] -= x
	m += 4 * x
	mid -= x

	free := x

	// 还有2的情况， 当作1处理
	freq[1] += freq[2] * 2

	x = min(freq[1], free)
	// 这部分已经被计算了
	freq[1] -= x
	free -= x

	// x优先放置到边上上去
	x = min(freq[1], side)
	freq[1] -= x
	m += 2 * x
	side -= x

	// 剩下的1放中间，但是可以行用两个
	x = min((freq[1]+1)/2, mid)
	freq[1] -= x * 2
	m += 4 * x
	mid -= x

	if freq[1] > 0 || m > 8*n {
		return false
	}
	return true
}
