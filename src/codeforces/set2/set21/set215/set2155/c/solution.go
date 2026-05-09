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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
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
	wizards := make([]int, n)

	play := func(s int) bool {
		wizards[0] = s
		for i := 1; i < n; i++ {
			if a[i-1] == a[i] {
				// 方向必须相反
				wizards[i] = 1 - wizards[i-1]
			} else if a[i-1]+1 == a[i] {
				// 前面必须指向右边
				if wizards[i-1] == 1 {
					return false
				}
				wizards[i] = 0
			} else if a[i-1]-1 == a[i] {
				// 必须指向左边
				if wizards[i-1] == 0 {
					return false
				}
				wizards[i] = 1
			} else {
				return false
			}
		}
		return true
	}

	suf := make([]int, n+1)
	check := func() bool {
		for i := n - 1; i >= 0; i-- {
			suf[i] = wizards[i] + suf[i+1]
		}
		var pref int
		for i := 0; i < n; i++ {
			if pref+1+suf[i+1] != a[i] {
				return false
			}
			if wizards[i] == 0 {
				pref++
			}
		}
		return true
	}
	var res int
	if play(0) && check() {
		res++
	}
	if play(1) && check() {
		res++
	}
	return res
}
