package main

import (
	"bufio"
	"fmt"
	"os"
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

func drive(reader *bufio.Reader) string {
	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)
	if res {
		return "YES"
	}
	return "NO"
}

func solve(a, b string) bool {
	x := small(a)
	y := small(b)
	return x == y
}

func small(s string) string {

	n := len(s)

	buf := []byte(s)

	var f func(l int, r int)

	back := make([]byte, n)

	f = func(l int, r int) {
		if (r-l+1)&1 == 1 {
			return
		}
		mid := (l + r) / 2
		f(l, mid)
		f(mid+1, r)

		swap := false

		for i := 0; i < mid-l+1; i++ {
			if buf[l+i] != buf[mid+1+i] {
				if buf[l+i] > buf[mid+1+i] {
					swap = true
				}
				break
			}
		}
		if swap {
			copy(back[l:mid+1], buf[l:mid+1])
			copy(buf[l:mid+1], buf[mid+1:r+1])
			copy(buf[mid+1:r+1], back[l:mid+1])
		}
	}
	f(0, n-1)
	return string(buf)
}
