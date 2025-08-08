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
	t := readString(reader)
	res := solve(s, t)
	fmt.Println(res[0], res[1])
}

func readString(reader *bufio.Reader) string {
	buf, _ := reader.ReadString('\n')
	return strings.TrimSpace(buf)
}

func solve(s string, t string) []int {
	fs := getFreq(s)
	ft := getFreq(t)

	res := make([]int, 2)
	for i := range 52 {
		x := min(fs[i], ft[i])
		fs[i] -= x
		ft[i] -= x
		res[0] += x
	}

	for i := range 26 {
		x := min(fs[i], ft[i+26])
		fs[i] -= x
		ft[i+26] -= x
		res[1] += x
	}

	for i := range 26 {
		x := min(fs[i+26], ft[i])
		fs[i+26] -= x
		ft[i] -= x
		res[1] += x
	}

	return res
}

func getFreq(s string) []int {
	freq := make([]int, 52)
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			freq[s[i]-'a']++
		} else {
			freq[s[i]-'A'+26]++
		}
	}
	return freq
}
