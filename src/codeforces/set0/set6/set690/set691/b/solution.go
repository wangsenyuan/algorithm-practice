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
	fmt.Println(solve(s))
}

func solve(s string) string {

	n := len(s)

	for l, r := 0, n-1; l <= r; l, r = l+1, r-1 {
		if !checkMirror(s[l], s[r]) {
			return "NIE"
		}
	}
	return "TAK"
}

var mirrors = []string{"AA", "bd", "db", "pq", "qp", "oo", "OO", "xx", "XX",
	"YY", "MM", "HH", "II",  "VV", "vv", "TT", "WW", "ww", "UU"}

func checkMirror(a, b byte) bool {
	for _, cur := range mirrors {
		if cur[0] == a && cur[1] == b {
			return true
		}
	}
	return false
}
