package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, d int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscan(reader, &n, &k, &d)
	res := solve(n, k, d)
	if res == nil {
		writer.WriteString("-1\n")
	} else {
		for _, row := range res {
			s := fmt.Sprintf("%v", row)
			writer.WriteString(s[1:len(s)-1] + "\n")
		}
	}
}

func solve(n int, k int, d int) [][]int {
	ok := false
	tmp := 1
	for range d {
		tmp *= k
		if tmp >= n {
			ok = true
			break
		}
	}
	if !ok {
		return nil
	}

	res := make([][]int, d)
	for i := range d {
		res[i] = make([]int, n)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < d; j++ {
			res[j][i] = res[j][i-1]
		}
		for j := d - 1; j >= 0; j-- {
			res[j][i] = (res[j][i] + 1) % k
			if res[j][i] > 0 {
				break
			}
		}
	}
	for i := range d {
		for j := range n {
			res[i][j]++
		}
	}
	return res
}
