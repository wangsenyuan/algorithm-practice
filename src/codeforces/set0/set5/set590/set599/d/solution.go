package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	fmt.Fprintln(writer, len(res))
	for _, p := range res {
		fmt.Fprintln(writer, p[0], p[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var x int
	fmt.Fscan(reader, &x)
	return solve(x)
}

func solve(x int) [][]int {

	// sum[w] =

	calc := func(w int, h int) int {
		var sum int
		// w <= h
		for i := 1; i <= w && i <= h; i++ {
			sum += (w - i + 1) * (h - i + 1)
			if sum > x {
				return x + 1
			}
		}

		return sum
	}

	// 先找出最大的w
	lo, hi := 1, x
	for lo < hi {
		mid := (lo + hi) >> 1
		if calc(mid, mid) >= x {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	find := func(w int) int {
		h := (6*x - w*(w+1)*(2*w+1) + 3*w*w*(w+1)) / (3 * w * (w + 1))
		sum := w*w*h - w*w*(w-1)/2 - w*h*(w-1)/2 + w*(w-1)*(2*w-1)/6
		if sum == x {
			return h
		}
		return -1
	}

	var res [][]int
	for w := 1; w <= hi; w++ {
		h := find(w)
		if h != -1 {
			if h < w {
				// already found
				break
			}
			// w <= h
			res = append(res, []int{w, h})
			if h != w {
				res = append(res, []int{h, w})
			}
		}
	}
	slices.SortFunc(res, func(i, j []int) int {
		return i[0] - j[0]
	})

	return res
}
