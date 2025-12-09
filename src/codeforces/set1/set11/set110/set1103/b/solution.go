package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(x int, y int) string {
		if x == y {
			return "x"
		}
		fmt.Printf("? %d %d\n", x, y)
		return readString(reader)
	}

	for {
		message := readString(reader)
		if message == "end" {
			break
		}
		// message = "start"
		res := solve(ask)
		fmt.Printf("! %d\n", res)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(ask func(int, int) string) int {
	l0, r0 := 0, 1

	for {
		res := ask(l0, r0)
		if res == "x" {
			break
		}
		l0, r0 = r0, r0*2
	}

	// l0 < a <= r0, and r0 < 2 * a

	l, r := l0+1, r0
	for l < r {
		mid := (l + r) / 2
		res := ask(l0, mid)
		if res == "x" {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
