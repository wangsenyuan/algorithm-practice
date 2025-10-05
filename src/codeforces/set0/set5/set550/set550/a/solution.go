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
	var first []int
	var second []int
	for i := 0; i+1 < len(s); i++ {
		if s[i] == 'A' && s[i+1] == 'B' {
			first = append(first, i)
		}
		if s[i] == 'B' && s[i+1] == 'A' {
			second = append(second, i)
		}
	}
	if len(first) == 0 || len(second) == 0 {
		return false
	}
	n := len(first)
	m := len(second)
	if first[0]+1 < second[m-1] || second[0]+1 < first[n-1] {
		return true
	}
	return false
}
