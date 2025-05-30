package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) bool {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) bool {
	n := len(s)
	m := len(t)

	pos := make([][]int, 26)
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		pos[x] = append(pos[x], i)
	}

	sum := make([][]int, 26)
	for i := range 26 {
		// 都增加一个n
		pos[i] = append(pos[i], n)
		sum[i] = make([]int, len(pos[i])+1)
	}

	fp := make([]int, m+1)
	fp[m] = n
	for i, j := m-1, n-1; i >= 0; i-- {
		for j >= 0 && t[i] != s[j] {
			j--
		}
		if j < 0 {
			// s里面没有子序列能匹配t
			return false
		}
		fp[i] = j
		if j >= 0 {
			j--
		}
	}

	for i, j := 0, 0; i < m; i++ {
		if i < m && j < n {
			x := int(t[i] - 'a')
			if j < fp[i+1] {
				u := sort.SearchInts(pos[x], j)
				if u < len(pos[x]) {
					sum[x][u]++
				}
				v := sort.SearchInts(pos[x], fp[i+1])
				if v <= len(pos[x]) {
					sum[x][v]--
				}
			}
		}

		for j < n && t[i] != s[j] {
			j++
		}

		j++
	}
	ok := make([]bool, n)
	for i := range 26 {
		for j := 1; j < len(sum[i]); j++ {
			sum[i][j] += sum[i][j-1]
		}
		for j := 0; j < len(sum[i])-1; j++ {
			if sum[i][j] > 0 {
				ok[pos[i][j]] = true
			}
		}
	}

	for i := range n {
		if !ok[i] {
			return false
		}
	}
	return true
}
