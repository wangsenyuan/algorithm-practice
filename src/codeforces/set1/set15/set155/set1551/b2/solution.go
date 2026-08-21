package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (k int, a []int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(k, a)
	return
}

func solve(k int, a []int) []int {
	pos := make(map[int][]int)
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}

	// 超过k的只有k
	var todo []int

	for _, v := range pos {
		if len(v) > k {
			v = v[:k]
		}
		todo = append(todo, v...)
	}

	best := len(todo) / k * k

	ans := make([]int, len(a))

	for i := range best {
		ans[todo[i]] = i%k + 1
	}

	return ans
}
