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
	// 如果有不在环上的 -1
	// 如果环的大小是偶数，应该是除以2，否则需要环的大小
	// 需要求它们的lcm 吗？ 似乎不大对
	// (2, 1)
	// 如果 t = 2, Owwf -> Owwf -> Owf -> Joo 好像是ok的
	// (2, 3, 4, 1)
	// 如果 t = 4, Owwwwf -> Owwwf -> Owwf -> Owf -> Joo
	// 也是ok的
	// 确实需要lcm
	n := len(a)
	deg := make([]int, n)
	for i := range n {
		a[i]--
		deg[a[i]]++
	}

	for i := range n {
		if deg[i] == 0 {
			// 没人让x出现joo
			return -1
		}
	}
	l := 1

	marked := make([]bool, n)
	for i := range n {
		if !marked[i] {
			j := i
			var cnt int
			for !marked[j] {
				marked[j] = true
				cnt++
				j = a[j]
			}
			if cnt%2 == 0 {
				l = lcm(l, cnt/2)
			} else {
				l = lcm(l, cnt)
			}
		}
	}
	return l
}

func lcm(a, b int) int {
	c := gcd(a, b)
	return a / c * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
