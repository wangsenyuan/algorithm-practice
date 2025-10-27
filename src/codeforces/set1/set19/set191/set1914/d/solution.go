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
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, 3)
	for i := range 3 {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int) int {
	x := getLargestThree(a[0])
	y := getLargestThree(a[1])
	z := getLargestThree(a[2])
	var best int
	for _, u := range x {
		for _, v := range y {
			for _, w := range z {
				if u.second != v.second && u.second != w.second && v.second != w.second {
					best = max(best, u.first+v.first+w.first)
				}
			}
		}
	}
	return best
}

type pair struct {
	first  int
	second int
}

func getLargestThree(a []int) []pair {
	res := make([]pair, 3)
	for i, v := range a {
		if v >= res[0].first {
			res[2], res[1] = res[1], res[0]
			res[0] = pair{v, i}
		} else if v >= res[1].first {
			res[2] = res[1]
			res[1] = pair{v, i}
		} else if v >= res[2].first {
			res[2] = pair{v, i}
		}
	}
	return res
}
