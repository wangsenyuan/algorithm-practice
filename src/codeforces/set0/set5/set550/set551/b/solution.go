package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) string {
	a := readString(reader)
	b := readString(reader)
	c := readString(reader)
	return solve(a, b, c)
}

func solve(a string, b string, c string) string {
	x := getFreq(a)
	y := getFreq(b)
	z := getFreq(c)

	check := func(n int) int {
		u := len(a)
		for i := range 26 {
			if y[i]*n > x[i] {
				return -1
			}
			if z[i] > 0 {
				u = min(u, (x[i]-y[i]*n)/z[i])
			}
		}
		return u
	}

	best := []int{0, 0}
	for u := 0; ; u++ {
		n := check(u)
		if n < 0 {
			break
		}
		if best[0]+best[1] < n+u {
			best[0] = u
			best[1] = n
		}
	}

	var buf strings.Builder

	for range best[0] {
		buf.WriteString(b)
	}
	for i := range 26 {
		x[i] -= y[i] * best[0]
	}
	for range best[1] {
		buf.WriteString(c)
	}
	for i := range 26 {
		x[i] -= z[i] * best[1]
	}
	for i := range 26 {
		for range x[i] {
			buf.WriteByte(byte(i + 'a'))
		}
	}
	return buf.String()
}

func getFreq(s string) []int {
	freq := make([]int, 26)
	for i := range s {
		freq[int(s[i]-'a')]++
	}
	return freq
}
