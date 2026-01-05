package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	events := make([][]int, q)
	for i := range q {
		events[i] = make([]int, 2)
		fmt.Fscan(reader, &events[i][0], &events[i][1])
	}
	return solve(n, events)
}

func solve(n int, events [][]int) []int {
	// n := len(events)
	// 最多也就这么多个消息

	// 最后一个消息的id
	// n := len(events)
	pos := make([][]int, n+1)

	m := len(events)

	prev := make([]int, m)
	next := make([]int, m)
	for i := range m {
		prev[i] = i - 1
		next[i] = i + 1
	}

	ans := make([]int, m)
	var unread int

	marked := make([]bool, m)

	var head int

	read := func(i int) {
		if marked[i] {
			return
		}
		marked[i] = true
		unread--
		l, r := prev[i], next[i]
		if l >= 0 {
			next[l] = r
		} else {
			// l < 0
			head = r
		}
		if r < m {
			prev[r] = l
		}
	}

	var last int

	for i, cur := range events {
		switch cur[0] {
		case 1:
			x := cur[1]
			pos[x] = append(pos[x], last)
			last++
			unread++
		case 2:
			x := cur[1]
			// 将x的所有消息都读完
			for len(pos[x]) > 0 {
				j := pos[x][0]
				pos[x] = pos[x][1:]
				read(j)
			}
		default:
			t := cur[1]
			// 读取前t个消息，这些消息有可能被读取过了
			for head < t {
				// read 会更新head
				read(head)
				// head = tmp
			}
		}
		ans[i] = unread
	}

	return ans
}
