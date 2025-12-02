package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(s string) bool {
	const target = "heidi"

	var j int
	for i := range s {
		if s[i] == target[j] {
			j++
		}
		if j == len(target) {
			return true
		}
	}
	return false
}
