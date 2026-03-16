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
	fmt.Println(solve(s))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) int {
	var stack []int

	for i := 0; i < len(s); i++ {
		if s[i] == '*' || s[i] == '+' {
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			stack = stack[:len(stack)-1]
			if s[i] == '*' {
				first *= second
			} else {
				first += second
			}
			stack[len(stack)-1] = first
		} else {
			d := int(s[i] - '0')
			stack = append(stack, d)
		}
	}

	return stack[len(stack)-1]
}
