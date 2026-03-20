package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, cur := range res {
		s := fmt.Sprintf("%v", cur)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i := range nums {
		nums[i], _ = strconv.Atoi(ss[i])
	}
	return nums
}

func drive(reader *bufio.Reader) [][]int {
	nums := readNums(reader)
	s := readString(reader)
	queries := make([][]int, nums[1])

	for i := range queries {
		queries[i] = readNums(reader)
	}
	return solve(s, queries)
}

func solve(s string, queries [][]int) [][]int {
	next := make([]int, len(s))
	prev := make([]int, len(s))

	play := func(l int, r int) []int {
		res := make([]int, 10)

		buf := []byte(s[l : r+1])
		n := len(buf)
		for i := range n {
			next[i] = i + 1
			prev[i] = i - 1
		}

		earse := func(i int) {
			l, r := prev[i], next[i]
			if l >= 0 {
				next[l] = r
			}
			if r < n {
				prev[r] = l
			}
		}

		var cp int
		dp := 1
		for cp >= 0 && cp < n {
			if buf[cp] >= '0' && buf[cp] <= '9' {
				res[buf[cp]-'0']++
				ncp := next[cp]
				if dp == -1 {
					ncp = prev[cp]
				}
				if buf[cp] == '0' {
					earse(cp)
				} else {
					buf[cp]--
				}
				cp = ncp
			} else {
				var ncp int
				if buf[cp] == '<' {
					dp = -1
					ncp = prev[cp]
				} else {
					dp = 1
					ncp = next[cp]
				}
				if ncp >= 0 && ncp < n && (buf[ncp] == '<' || buf[ncp] == '>') {
					earse(cp)
				}
				cp = ncp
			}
		}

		return res
	}
	ans := make([][]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		ans[i] = play(l-1, r-1)
	}
	return ans
}
