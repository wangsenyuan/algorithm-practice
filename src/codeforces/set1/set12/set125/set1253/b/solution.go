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
	if res == nil {
		fmt.Fprintln(writer, -1)
		return
	}
	fmt.Fprintln(writer, len(res))
	for i, x := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	w := slices.Max(a)
	v := slices.Min(a)
	if w != -v {
		return nil
	}
	onboard := make([]bool, w+1)
	var todo []int
	thisTime := make([]bool, w+1)

	checkIn := func(x int) bool {
		if thisTime[x] || onboard[x] {
			return false
		}
		onboard[x] = true
		thisTime[x] = true
		todo = append(todo, x)
		return true
	}

	checkOut := func(x int) bool {
		if !onboard[x] {
			return false
		}
		onboard[x] = false
		return true
	}

	var cnt int
	var res []int
	for _, x := range a {
		if x > 0 {
			if !checkIn(x) {
				return nil
			}
			cnt++
		} else {
			x = -x
			if !checkOut(x) {
				return nil
			}
			cnt--
			if cnt == 0 {
				res = append(res, 2*len(todo))
				for _, y := range todo {
					thisTime[y] = false
				}
				todo = nil
			}
		}
	}
	if cnt != 0 {
		return nil
	}
	return res
}
