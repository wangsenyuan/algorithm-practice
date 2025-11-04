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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	nums := make(map[int]int)
	var res int
	for i := range n {
		if a[i] == 0 {
			res++
		}
		nums[a[i]]++
	}

	buf := make([]int, n)
	for i := range n {
		// 如果以i为开始位置
		nums[a[i]]--
		for j := range n {
			if i == j || a[i] == 0 && a[j] == 0 {
				continue
			}
			nums[a[j]]--
			u, v := a[i], a[j]
			var it int
			for {
				next := u + v
				if nums[next] == 0 {
					break
				}
				nums[next]--
				buf[it] = next
				it++
				u, v = v, next
			}

			res = max(res, it+2)
			for k := 0; k < it; k++ {
				nums[buf[k]]++
			}
			nums[a[j]]++
		}
		nums[a[i]]++
	}

	return res
}
