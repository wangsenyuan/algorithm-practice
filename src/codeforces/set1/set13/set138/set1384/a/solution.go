package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		for _, s := range drive(reader) {
			fmt.Fprintln(writer, s)
		}
	}
}

func drive(reader *bufio.Reader) []string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const letters = "abc"

func solve(a []int) []string {
	n := len(a)
	res := make([]string, n+1)
	res[0] = string(createBuf(max(a[0], 1), 'a'))
	for i := range n {
		w := a[i]
		if i+1 < n {
			w = max(w, a[i+1])
		}
		w = max(1, w)
		// a, ab, abc
		c := letters[(i+1)%len(letters)]
		buf := createBuf(w, c)
		copy(buf, res[i])
		if a[i] < len(buf) && a[i] < len(res[i]) {
			buf[a[i]] = 'a'
			if res[i][a[i]] == 'a' {
				buf[a[i]] = 'b'
			}
		}
		res[i+1] = string(buf)
	}
	return res
}

func createBuf(n int, c byte) []byte {
	res := make([]byte, n)
	for i := range n {
		res[i] = c
	}
	return res
}
