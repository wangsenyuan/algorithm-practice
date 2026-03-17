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
		if len(res) == 0 {
			fmt.Fprintln(writer, "-1")
		} else {
			s := fmt.Sprintf("%v", res)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

func solve(n int) []int {
	if n%2 == 0 {
		res := make([]int, n)
		for i := 0; i < n; i += 2 {
			res[i] = i/2 + 1
			res[i+1] = i/2 + 1
		}
		return res
	}
	if n <= 26 {
		return nil
	}
	// n >= 26
	res := make([]int, n)
	res[0] = 1e6
	res[9] = 1e6
	res[25] = 1e6
	res[10] = 1e6 - 1
	res[26] = res[10]
	num := 1
	var cnt int
	for i := range n {
		if res[i] > 0 {
			continue
		}
		res[i] = num
		cnt++
		if cnt == 2 {
			num++
			cnt = 0
		}
	}
	return res
}
