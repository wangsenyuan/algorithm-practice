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
	var sum int

	var w int

	var res []int

	for i := 0; i < n; i++ {
		sum += i + 1
		for w*w < sum {
			w++
		}
		if w*w != sum {
			res = append(res, i+1)
			continue
		}
		// w * w == sum
		if i == n-1 {
			return nil
		}
		res = append(res, i+2)
		res = append(res, i+1)
		sum += i + 2
		i++
	}

	return res
}
