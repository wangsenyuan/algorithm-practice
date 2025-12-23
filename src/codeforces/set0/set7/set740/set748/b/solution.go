package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ok, ans := drive(reader)
	if !ok {
		fmt.Println("-1")
		return
	}
	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Println(v)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (bool, []string) {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) (bool, []string) {
	n := len(s)
	pair := make([]int, 26)
	for i := range 26 {
		pair[i] = -1
	}

	for i := range n {
		x := int(s[i] - 'a')
		y := int(t[i] - 'a')
		x, y = min(x, y), max(x, y)
		if pair[x] == -1 && pair[y] == -1 {
			pair[x] = y
			pair[y] = x
		} else if pair[x] != y || pair[y] != x {
			return false, nil
		}
	}
	var res []string
	for i := range 26 {
		if pair[i] > i {
			res = append(res, fmt.Sprintf("%c %c", byte(i+'a'), byte(pair[i]+'a')))
		}
	}
	return true, res
}
