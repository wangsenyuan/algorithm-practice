package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)

	if len(res) == 0 {
		fmt.Println("No solution")
		return
	}
	output := func(arr []int) {
		s := fmt.Sprintf("%v", arr)
		fmt.Println(s[1 : len(s)-1])
	}
	output(res[0])
	output(res[1])
}

func drive(reader *bufio.Reader) (a []int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

const inf = 1 << 60

func solve(a []int) [][]int {
	n := len(a)
	if n == 2 {
		return [][]int{{a[0]}, {a[1]}}
	}

	if n <= 4 {
		h := (n + 1) / 2
		first := a[:h]
		second := a[h:]
		return [][]int{first, second}
	}

	check := func(arr []int) bool {
		if len(arr) == 0 {
			return false
		}
		if len(arr) <= 2 {
			return true
		}
		d := arr[1] - arr[0]
		for i := 2; i < len(arr); i++ {
			if arr[i]-arr[i-1] != d {
				return false
			}
		}
		return true
	}

	marked := make([]bool, n)

	play := func(f1 int, f2 int) [][]int {
		clear(marked)
		diff1 := a[f2] - a[f1]
		var v1 []int
		var v2 []int
		w := a[f1]
		last := -1
		for i, v := range a {
			if v == w {
				v1 = append(v1, v)
				w += diff1
				marked[i] = true
				last = i
			} else {
				v2 = append(v2, v)
			}
		}
		if check(v2) {
			return [][]int{v1, v2}
		}
		v1 = v1[:len(v1)-1]
		v2 = v2[:0]
		marked[last] = false
		for i := range n {
			if !marked[i] {
				v2 = append(v2, a[i])
			}
		}
		if check(v2) {
			return [][]int{v1, v2}
		}
		return nil
	}

	res := play(0, 1)
	if len(res) > 0 {
		return res
	}
	res = play(0, 2)
	if len(res) > 0 {
		return res
	}

	return play(1, 2)
}
