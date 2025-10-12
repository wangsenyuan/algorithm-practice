package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) int {
	n := len(s)
	fp := make([]int, 2)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			fp[0] = max(fp[0], fp[1]+1)
		} else {
			fp[1] = max(fp[1], fp[0]+1)
		}
	}
	res := max(fp[0], fp[1])

	// 在s中，有x个为止 s[i] != s[i+1], 那么flip后，这个数字是不会变的
	// 区别在于两头的数字
	// 假设反转[l...r]
	// pref(r) 表示在r处的01或者10的个数，那么反转的位置l，必然是00或者11的中间
	// 只要在r的前面存在一个00,或者11,就+1即可
	var add int
	for i := 0; i+1 < n; i++ {
		if s[i] == s[i+1] {
			add++
		}
	}

	return res + min(add, 2)
}
