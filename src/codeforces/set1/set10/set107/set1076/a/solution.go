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
	fmt.Println(solve(s))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) string {
	n := len(s)
	buf := []byte(s)
	for i := 0; i+1 < n; i++ {
		if buf[i] > buf[i+1] {
			copy(buf[i:], buf[i+1:])
			break
		}
	}
	return string(buf[:n-1])
}
