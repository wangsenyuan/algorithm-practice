package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, a := drive(reader)
		s := fmt.Sprintf("%v", a)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (b []int, a []int) {
	var n int
	fmt.Fscan(reader, &n)
	b = make([]int, 2*n)
	for i := range b {
		fmt.Fscan(reader, &b[i])
	}
	a = solve(slices.Clone(b))
	return
}

func solve(b []int) []int {
	n := len(b)
	// n 是偶数
	slices.Sort(b)

	pref := make([][2]int, n+1)

	nums := make(map[int]int)

	for i, v := range b {
		pref[i+1][0] = pref[i][0]
		pref[i+1][1] = pref[i][1]
		pref[i+1][(i+1)&1] += v
		nums[v]++
	}

	x := pref[n][0] - pref[n][1]

	if x >= 1 && nums[x] == 0 {
		// 这个x正好可以作为最后一个
		return append(b, x)
	}

	play := func(i int, x int) []int {
		a := make([]int, n+1)
		copy(a, b)
		copy(a[i+1:], b[i:])
		a[i] = x
		return a
	}

	var suf int
	// 1 2 3 4
	for i := n - 1; i >= 0; i-- {
		// 如果把x放在位置i
		diff := pref[i][0] - pref[i][1]
		// 右边 - 左边 = diff
		if i&1 == 0 {
			suf -= b[i]
			// x要放在左边, 所以i要放入右边, i+1要放入左边
			diff += suf
		} else {
			suf += b[i]
			// x 要放在右边,
			diff -= suf
		}
		x = -diff
		if x >= 1 && nums[x] == 0 {
			return play(i, x)
		}
	}

	// 放在头部
	x = -suf
	return append([]int{x}, b...)
}
