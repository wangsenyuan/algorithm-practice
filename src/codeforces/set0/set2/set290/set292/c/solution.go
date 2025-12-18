package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	fmt.Fprintln(writer, len(res))
	for _, s := range res {
		fmt.Fprintln(writer, s)
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

func solve(a []int) []string {

	var res []string

	digits := make([]int, 12)

	var play func(i int, j int, m int, s string)

	play = func(i int, j int, m int, s string) {
		if i == m {
			if j == 4 {
				res = append(res, s[:len(s)-1])
			}
			return
		}
		if j == 4 {
			return
		}

		var val int
		for i1 := i; i1 < m; i1++ {
			val = val*10 + digits[i1]
			if val > 255 || i1 > i && digits[i] == 0 {
				break
			}
			play(i1+1, j+1, m, s+strconv.Itoa(val)+".")
		}
	}

	var tot int
	for _, v := range a {
		tot |= 1 << v
	}

	var fill func(i int, d int, mask int)

	fill = func(i int, d int, mask int) {
		if i == (d+1)/2 {
			if mask == tot {
				play(0, 0, d, "")
			}
			return
		}
		for _, v := range a {
			digits[i] = v
			digits[d-1-i] = v
			fill(i+1, d, mask|(1<<v))
		}
	}

	for d := 4; d <= 12; d++ {
		fill(0, d, 0)
	}

	return res
}
