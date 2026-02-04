package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(x, y, a)
}

func solve(x int, y int, a []int) int {
	slices.Sort(a)
	var res int
	for _, v := range a {
		// 一共有v个题目
		if v&1 == 1 && y > 0 {
			// 避免浪费2页纸中的一半
			y--
			v--
		}
		x1 := min(x, (v+1)/2)
		x -= x1
		v -= 2 * x1
		if v > 0 {
			if y < v {
				break
			}
			y -= v
		}
		res++
	}
	return res
}
