package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	n := readNum(reader)
	cmds := make([]string, 2*n)
	for i := 0; i < 2*n; i++ {
		cmds[i] = readString(reader)
	}
	return solve(n, cmds)
}

func solve(n int, cmds []string) int {
	stack := make([]int, n)
	var top int

	when := make([]int, n+1)
	var res int
	next := 1
	lastReorder := -1

	for i, cur := range cmds {
		if cur == "remove" {
			if stack[top-1] == next {
				top--
			} else if when[stack[top-1]] > lastReorder {
				res++
				// 现在stack中的，要从大往小排列
				lastReorder = i
			}

			next++
		} else {
			var num int
			readInt([]byte(cur[4:]), 0, &num)
			stack[top] = num
			top++
			when[num] = i
		}
	}

	return res
}
