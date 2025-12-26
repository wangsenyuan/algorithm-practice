package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("Impossible")
	} else {
		fmt.Println(res[0])
		fmt.Println(res[1])
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (string, []string) {
	s := readString(reader)
	res := solve(s)
	return s, res
}

func solve(s string) []string {
	// len(s) = 27
	pos := make([][]int, 26)
	var x int
	for i := range 27 {
		w := int(s[i] - 'A')
		pos[w] = append(pos[w], i)
		if len(pos[w]) == 2 {
			x = w
			break
		}
	}

	d := pos[x][1] - pos[x][0]
	if d == 1 {
		// impossible
		return nil
	}
	// ABA   ABCA
	buf := make([][]byte, 2)
	for i := range 2 {
		buf[i] = make([]byte, 13)
		for j := range 13 {
			buf[i][j] = '.'
		}
	}
	// ABABA
	// 比如 ABA 需要2个字符位置, ABCA 需要3个字符位置
	need := pos[x][1] - pos[x][0]
	// 这些位置都放到最后去
	top := (need + 1) / 2
	bot := need / 2

	w := pos[x][0]
	for i := 13 - top; i < 13; i++ {
		buf[0][i] = s[w]
		w++
	}

	for i := 12; i >= 13-bot; i-- {
		buf[1][i] = s[w]
		w++
	}

	prev := pos[x][0]
	if prev <= 13-top {
		w := 13 - top - prev
		for i := range prev {
			buf[0][w+i] = s[i]
		}
	} else {
		prev -= 13 - top
		var w int
		for i := prev - 1; i >= 0; i-- {
			buf[1][i] = s[w]
			w++
		}
		for i := 0; i < 13-bot; i++ {
			buf[0][i] = s[w]
			w++
		}
	}
	v := pos[x][1] + 1
	// 把剩余的放置进去
	r, c := 1, 13-bot-1
	for v < 27 {
		buf[r][c] = s[v]
		v++
		if r > 0 {
			c--
			if c < 0 {
				r--
				c = 0
			}
		} else {
			c++
		}
	}

	ans := make([]string, 2)
	ans[0] = string(buf[0])
	ans[1] = string(buf[1])
	return ans
}
