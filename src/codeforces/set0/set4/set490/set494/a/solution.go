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
	res := solve(s)
	if res == nil {
		fmt.Println("-1")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func solve(s string) []int {
	// bal >= 0
	// 好像最后一个特殊处理下
	var res []int
	var level int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			level++
		case ')':
			level--
		default:
			res = append(res, 1)
			level--
		}
		if level < 0 {
			return nil
		}
	}
	if level > 0 {
		m := len(res)
		if m == 0 {
			return nil
		}
		res[m-1] += level
	}

	level = 0

	for i, j := 0, 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			level++
		case ')':
			level--
		default:
			level -= res[j]
			j++
		}
		if level < 0 {
			return nil
		}
	}

	return res
}
