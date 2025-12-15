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
	s := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &s[i])
	}
	return solve(s)
}

func solve(s []int) int {
	// 在某个位置，前面只能保留0，后面只能保留1
	n := len(s)
	pref := make([]int, n+1)
	for i, v := range s {
		pref[i+1] = pref[i] + v
	}
	// 删除所有的1
	ans := pref[n]
	var suf int
	for i := n - 1; i >= 0; i-- {
		// 当前位置是0, 或者1都可以
		ans = min(ans, pref[i]+suf)
		if s[i] == 0 {
			suf++
		}
	}
	ans = min(ans, suf)

	return n - ans
}
