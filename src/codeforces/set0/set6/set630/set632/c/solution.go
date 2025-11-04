package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func drive(reader *bufio.Reader) string {
	n := readNum(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	return solve(words)
}

func cmp(a string, b string) int {
	n := len(a)
	m := len(b)
	// s1 = a + b
	// s2 = b + a
	for i := 0; i < n+m; i++ {
		var x, y byte
		if i < n {
			x = a[i]
		} else {
			x = b[i-n]
		}
		if i < m {
			y = b[i]
		} else {
			y = a[i-m]
		}
		if x != y {
			return int(x-'a') - int(y-'a')
		}
	}
	return 0
}

func solve(words []string) string {
	slices.SortFunc(words, cmp)
	var buf strings.Builder
	for _, word := range words {
		buf.WriteString(word)
	}
	return buf.String()
}
