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
	fmt.Printf("%.10f\n", res)
}

func solve(s string) float64 {
	n := len(s)
	f := make([]float64, n+1)
	for i := range n {
		f[i+1] = f[i] + 1.0/float64(i+1)
	}

	var res float64
	for i := range n {
		if isVowel(s[i]) {
			j := n - i - 1
			cur := float64(n+1)*f[n] - float64(n-i)*f[j] - float64(i+1)*f[i] - 1
			res += cur
		}
	}

	return res
}

func isVowel(c byte) bool {
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'Y'
}

func bruteForce(s string) float64 {
	n := len(s)
	var res float64

	for i := range n {
		var cnt int
		for j := i; j >= 0; j-- {
			if isVowel(s[j]) {
				cnt++
			}
			res += float64(cnt) / float64(i-j+1)
		}
	}
	return res
}
