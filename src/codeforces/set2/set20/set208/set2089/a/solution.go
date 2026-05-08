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
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func solve(n int) []int {
	h1 := n / 3
	h2 := (n*2 + 2) / 3

	p := h1
	for p < h2 && !checkPrime(p) {
		p++
	}
	res := make([]int, n)
	marked := make([]bool, n+1)
	var pos int
	for i := 0; p-i > 0 && p+i <= n && pos < n; i++ {
		res[pos] = p - i
		pos++
		marked[p-i] = true
		if pos < n && i > 0 {
			res[pos] = p + i
			pos++
			marked[p+i] = true
		}
	}
	for i := 1; pos < n; pos++ {
		for marked[i] {
			i++
		}
		res[pos] = i
		marked[i] = true
	}

	return res
}

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
