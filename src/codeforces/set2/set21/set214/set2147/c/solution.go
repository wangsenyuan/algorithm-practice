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
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		res := solve(s)
		if res {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func solve(s string) bool {
	n := len(s)
	// s[i] = '1 的地方不能放置rabbit，但是必须保证没有兔子jump进去
	var state int
	// state = 1, 表示，最近的空的草地可以被mark
	var arr []int

	check := func(j int, i int) bool {
		if s[j] == '1' {
			if (i-j > 1 || i == n) && state == 1 {
				return false
			}

			return true
		}
		// 兔子
		if i-j > 1 {
			// 有多只兔子，可以*推迟*选择最后一只兔子的朝向
			state = 2
			return true
		}
		// 只有一只兔子
		if len(arr) > 0 {
			grass := arr[len(arr)-1]
			if grass > 1 {
				// 只能朝向右边
				state = 1
				return true
			}
			// grass == 1
			switch state {
			case 1:
				// 只能朝向左边
				state = 0
			case 2:
				// 前面的可以朝向左边，那么它也可以后续选择朝向
				state = 2
			default:
				// state == 0, 前面的朝向左边，那么这里必须朝向右边
				state = 1
			}

		} else {
			state = 2
		}

		return true
	}

	for i := 0; i < n; {
		j := i
		for i < n && s[i] == s[j] {
			i++
		}

		if !check(j, i) {
			return false
		}

		arr = append(arr, i-j)
	}

	return true
}
