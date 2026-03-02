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

	var res int
	for d := 20; d >= 0; d-- {
		for i := 0; i < n; i++ {
			if (a[i]>>d)&1 == 0 {
				continue
			}
			var sum int
			pos := make(map[int]int)
			pos[0] = i - 1
			for i < n && (a[i]>>d)&1 == 1 {
				v := a[i] >> d
				sum ^= v
				if j, ok := pos[sum]; ok {
					res = max(res, i-j)
				} else {
					pos[sum] = i
				}
				i++
			}
		}
	}
	return res
}
