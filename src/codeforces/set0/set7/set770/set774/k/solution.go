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
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) string {
	var res []byte
	for i := 0; i < len(s); {
		if !isVowel(s[i]) {
			res = append(res, s[i])
			i++
		} else {
			j := i
			for i < len(s) && s[i] == s[j] {
				i++
			}
			if i-j == 2 && (s[j] == 'e' || s[j] == 'o') {
				res = append(res, s[j], s[j])
				continue
			}
			res = append(res, s[j])
		}
	}
	return string(res)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'y'
}
