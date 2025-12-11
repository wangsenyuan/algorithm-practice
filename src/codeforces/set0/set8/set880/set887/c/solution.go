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
	a := make([]int, 24)
	for i := range 24 {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) bool {

	check := func() bool {
		for i := 0; i < 24; i += 4 {
			for j := range 4 {
				if a[i+j] != a[i] {
					return false
				}
			}
		}
		return true
	}

	// 一次旋转，改变了8个面
	groups := [][]int{
		{1, 3, 5, 7, 9, 11, 24, 22},
		{2, 4, 6, 8, 10, 12, 23, 21},
		{13, 14, 5, 6, 17, 18, 21, 22},
		{15, 16, 7, 8, 19, 20, 23, 24},
		{1, 2, 18, 20, 12, 11, 15, 13},
		{3, 4, 17, 19, 10, 9, 16, 14},
	}

	rotate := func(g []int) {
		c0, c1 := a[g[0]-1], a[g[1]-1]
		for i := 0; i < len(g)-2; i += 2 {
			a[g[i]-1] = a[g[i+2]-1]
			a[g[i+1]-1] = a[g[i+3]-1]
		}
		a[g[6]-1] = c0
		a[g[7]-1] = c1
	}

	for _, g := range groups {
		for j := range 4 {
			rotate(g)
			if j%2 == 0 && check() {
				return true
			}
		}
	}

	return false
}
