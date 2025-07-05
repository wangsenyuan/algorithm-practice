package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	h := readNNums(reader, n)
	p := readNNums(reader, m)
	return solve(h, p)
}

func solve(h []int, p []int) int {

	calc := func(a int, l int, r int) int {
		if a <= l {
			return r - a
		}
		if r <= a {
			return a - l
		}
		d1 := a - l
		d2 := r - a
		return min(d1, d2) + r - l
	}

	check := func(want int) bool {
		var j int
		for i := 0; i < len(h) && j < len(p); i++ {
			if h[i]-p[j] > want {
				// too far
				return false
			}

			if p[j]-h[i] > want {
				continue
			}

			k := j
			for k < len(p) && calc(h[i], p[j], p[k]) <= want {
				k++
			}

			j = k
		}
		return j == len(p)
	}
	return sort.Search(1<<60, check)
}
