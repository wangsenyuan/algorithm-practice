package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) string {
	m := readNum(reader)
	s := readString(reader)
	return solve(m, s)
}

func solve(m int, s string) string {
	pos := make([][]int, 26)
	n := len(s)
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		pos[x] = append(pos[x], i)
	}
	// 如果存在b，那么最好选择所有的a
	// 如果存在c，那么最好选择所有的b
	// 依次类推

	marked := make([]bool, n)

	// 如果把a的全部放上，还无法满足条件，那么就需要啊
	for x := range 26 {
		for _, i := range pos[x] {
			marked[i] = true
		}
		flag := false
		prev := -1
		for i := range n {
			if !marked[i] && i-prev >= m {
				flag = true
				break
			}
			if marked[i] {
				prev = i
			}
		}
		if !flag {
			// 全部放x太多了
			// 先取消
			for _, i := range pos[x] {
				marked[i] = false
			}
			prev := -1
			for i := range n {
				if marked[i] {
					prev = i
				}

				if i-prev == m {
					for len(pos[x]) > 1 && pos[x][1]-prev <= m {
						pos[x] = pos[x][1:]
					}
					marked[pos[x][0]] = true
					prev = max(prev, pos[x][0])
					pos[x] = pos[x][1:]
				}
			}
			break
		}
	}

	var buf []byte

	for i := range n {
		if marked[i] {
			buf = append(buf, s[i])
		}
	}
	slices.Sort(buf)

	return string(buf)
}
