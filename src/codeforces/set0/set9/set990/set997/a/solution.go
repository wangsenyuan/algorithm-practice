package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	first := readString(reader)
	ss := strings.Split(first, " ")
	x, _ := strconv.Atoi(ss[1])
	y, _ := strconv.Atoi(ss[2])
	s := readString(reader)
	return solve(x, y, s)
}

func solve(x int, y int, s string) int {
	buf := []byte(s)
	buf = slices.Compact(buf)
	n := len(buf)
	if n == 1 {
		if buf[0] == '1' {
			return 0
		}
		return y
	}

	p := n / 2
	if buf[0] == '0' {
		p = (n + 1) / 2
	}

	return (p-1)*min(x, y) + y
}

const inf = 1 << 60
