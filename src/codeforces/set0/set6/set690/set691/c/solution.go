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
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(solve(s))
}

func solve(x string) string {
	buf := []byte(x)
	for len(buf) > 0 && buf[0] == '0' {
		buf = buf[1:]
	}
	// len(buf) > 0, must hold
	n := len(buf)
	dot := n
	for i := range n {
		if buf[i] == '.' {
			dot = i
			break
		}
	}

	if dot == 0 {
		for n > 0 && buf[n-1] == '0' {
			n--
		}
		buf = buf[:n]
		b := 0
		for buf[b+1] == '0' {
			b++
		}
		// buf[b+1] != '0'
		if b+2 == n {
			return fmt.Sprintf("%cE-%d", buf[b+1], b+1)
		}
		return fmt.Sprintf("%c.%sE-%d", buf[b+1], string(buf[b+2:]), b+1)
	}

	if dot == n {
		buf = append(buf, '.')
		n++
	}

	// dot > 0
	b := dot - 1

	if dot > 1 && dot < n {
		slices.Reverse(buf[1 : dot+1])
		slices.Reverse(buf[2 : dot+1])
	}
	// 现在dot在位置1处
	for buf[n-1] == '0' {
		n--
	}
	if buf[n-1] == '.' {
		n--
	}
	buf = buf[:n]

	if b == 0 {
		return string(buf)
	}

	if n == 1 {
		return fmt.Sprintf("%cE%d", buf[0], b)
	}
	return fmt.Sprintf("%c.%sE%d", buf[0], string(buf[2:]), b)
}
