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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, x, s int
	fmt.Fscan(reader, &n, &x, &s)
	var u string
	fmt.Fscan(reader, &u)
	return solve(n, x, s, u)
}

func solve(n, x, s int, u string) int {
	// 那些可以移动的A的人
	if s == 1 {
		return solve1(x, u)
	}
	var free []int
	pos := make([]int, n)
	for i := range n {
		pos[i] = -1
	}
	var res int
	cnt := make([]int, x)
	var que []int
	var lastTable int

	for i := range n {
		if u[i] == 'I' {
			if lastTable == x {
				continue
			}
			res++
			pos[i] = lastTable
			cnt[lastTable]++
			que = append(que, lastTable)
			lastTable++
		} else if u[i] == 'A' {
			// 这个要找到第一个>0的table
			for len(que) > 0 && cnt[que[0]] == s {
				que = que[1:]
			}
			if len(que) > 0 {
				res++
				pos[i] = que[0]
				cnt[que[0]]++
				free = append(free, i)
			} else if lastTable < x {
				res++
				pos[i] = lastTable
				cnt[lastTable]++
				que = append(que, lastTable)
				lastTable++
				// 这个不是free的
			}
		} else {
			// E
			for len(que) > 0 && cnt[que[0]] == s {
				que = que[1:]
			}
			if len(que) > 0 {
				res++
				pos[i] = que[0]
				cnt[que[0]]++
				continue
			}
			// 就需要先安排一个free的人过来
			if len(free) > 0 && lastTable < x {
				res++
				first := free[0]
				free = free[1:]
				// 这个有可能减到0吗? 好像是有可能的, 但是似乎不能影响到后面I的情况
				cnt[pos[first]]--
				que = append(que, pos[first])
				pos[first] = lastTable
				cnt[lastTable]++
				que = append(que, lastTable)
				pos[i] = lastTable
				cnt[lastTable]++
				lastTable++
			}
		}
	}

	return res
}

func solve1(x int, u string) int {
	var cnt int
	for i := range u {
		if u[i] != 'E' {
			cnt++
		}
	}
	return min(cnt, x)
}
