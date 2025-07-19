package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Printf("%c\n", res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) byte {
	var pos1 []int
	var pos2 []int
	n := len(s)
	for i := range n {
		if s[i] == 'R' {
			pos1 = append(pos1, i)
		} else {
			pos2 = append(pos2, i)
		}
	}

	// 最多n次，因为每次至少有一个位置被移除掉
	for len(pos1) > 0 && len(pos2) > 0 {
		x := pos1[0]
		y := pos2[0]
		if x < y {
			pos2 = pos2[1:]
			pos1 = pos1[1:]
			pos1 = append(pos1, x+n)
		} else {
			pos1 = pos1[1:]
			pos2 = pos2[1:]
			pos2 = append(pos2, y+n)
		}
	}

	if len(pos1) > 0 {
		return 'R'
	}

	return 'D'
}
