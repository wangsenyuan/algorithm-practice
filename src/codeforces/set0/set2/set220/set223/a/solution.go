package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	cnt, res := solve(s)
	fmt.Println(cnt)
	fmt.Println(res)
}

func solve(s string) (int, string) {
	n := len(s)
	pref := make([]int, n+1)
	stack := make([]int, n)
	ok := make([]bool, n)
	ans := make([]int, 3)
	lf := make([]int, n)

	var top int
	for i := range n {
		pref[i+1] = pref[i]
		if s[i] == '[' {
			pref[i+1]++
		}
		if s[i] == '[' || s[i] == '(' {
			stack[top] = i
			top++
		} else {
			if top == 0 || !check(s[stack[top-1]], s[i]) {
				// 全部舍弃掉
				top = 0
				continue
			}
			ok[i] = true
			j := stack[top-1]
			top--
			// 这个区间内
			lf[i] = j
			if j > 0 && ok[j-1] {
				lf[i] = lf[j-1]
			}
			tmp := pref[i+1] - pref[lf[i]]
			if tmp > ans[0] {
				ans[0] = tmp
				ans[1] = lf[i]
				ans[2] = i
			}
		}
	}
	if ans[0] == 0 {
		return 0, ""
	}
	return ans[0], s[ans[1] : ans[2]+1]
}

func check(a, b byte) bool {
	return a == '(' && b == ')' || a == '[' && b == ']'
}
