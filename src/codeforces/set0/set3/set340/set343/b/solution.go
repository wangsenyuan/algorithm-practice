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
	if solve(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) bool {
	buf := []byte(s)
	var top int
	n := len(s)
	for i := range n {
		if top > 0 && buf[top-1] == s[i] {
			top--
		} else {
			buf[top] = s[i]
			top++
		}
	}
	return top == 0
}
