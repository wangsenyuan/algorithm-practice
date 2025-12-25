package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, x := drive(reader)
	s := fmt.Sprintf("%v", x)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (y []int, x []int) {
	var n int
	fmt.Fscan(reader, &n)
	y = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &y[i])
	}
	x = solve(y)
	return
}

func solve(y []int) []int {
	// 假设最大值时w, 那么w必然是某个y[i]的前缀
	// 那么此时，其他所有的数，的前缀 <= w, 且必须有n个位置room
	slices.Sort(y)
	// 越大的数，留出的空间越大
	slices.Reverse(y)

	var ws []int
	for _, v := range y {
		for v > 0 {
			ws = append(ws, v)
			v >>= 1
		}
	}

	slices.Sort(ws)
	ws = slices.Compact(ws)

	check := func(mid int) bool {
		marked := make(map[int]bool)
		for _, v := range y {
			for v > 0 {
				if v <= ws[mid] && !marked[v] {
					marked[v] = true
					break
				}
				v >>= 1
			}
			// 这个数没有空间了
			if v == 0 {
				return false
			}
			if len(marked) >= len(y) {
				return true
			}
		}
		return false
	}

	l, r := 0, len(ws)

	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	marked := make(map[int]bool)

	var x []int

	for _, v := range y {
		for v > 0 {
			if v <= ws[l] && !marked[v] {
				x = append(x, v)
				marked[v] = true
				break
			}
			v >>= 1
		}
		if len(x) == len(y) {
			break
		}
	}

	return x
}
