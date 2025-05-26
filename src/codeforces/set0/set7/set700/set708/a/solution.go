package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) string {
	buf := []byte(s)

	ok := false
	for i := 0; i < len(buf); i++ {
		if buf[i] > 'a' {
			for j := i; j < len(buf) && buf[j] != 'a'; j++ {
				buf[j]--
			}
			ok = true
			break
		}
	}

	if !ok {
		n := len(buf)
		buf[n-1] = 'z'
	}

	return string(buf)
}
