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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var r, g, b int
	fmt.Fscan(reader, &r, &g, &b)
	return solve(r, g, b)
}

func solve(r int, g int, b int) string {
	left := [3]int{r, g, b}
	colors := [3]byte{'R', 'G', 'B'}
	res := make([]byte, 0, r+g+b)

	for {
		best := -1
		for color := range left {
			if left[color] == 0 || (len(res) > 0 && colors[color] == res[len(res)-1]) ||
				(len(res) >= 3 && colors[color] == res[len(res)-3]) {
				continue
			}
			if best < 0 || left[color] > left[best] ||
				(left[color] == left[best] && len(res) >= 2 && colors[color] == res[len(res)-2]) {
				best = color
			}
		}
		if best < 0 {
			break
		}
		res = append(res, colors[best])
		left[best]--
	}

	return string(res)
}
