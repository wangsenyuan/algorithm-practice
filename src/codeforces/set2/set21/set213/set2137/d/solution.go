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
		res := drive(reader)
		if res == nil {
			fmt.Fprintln(writer, -1)
			continue
		}
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(b)
}

func solve(b []int) []int {
	n := len(b)
	todo := make([][]int, n+1)
	for i, v := range b {
		if v > n {
			return nil
		}
		todo[v] = append(todo[v], i)
	}
	ans := make([]int, n)

	id := 1
	for i := 1; i <= n; i++ {
		w := len(todo[i])
		if w%i != 0 {
			return nil
		}

		for j := 0; j < w; j += i {
			for k := j; k < j+i; k++ {
				ans[todo[i][k]] = id
			}
			id++
		}
	}

	return ans
}
