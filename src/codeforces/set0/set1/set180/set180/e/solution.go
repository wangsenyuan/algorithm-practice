package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(m, k, a)
}

func solve(m int, k int, a []int) int {
	n := len(a)
	pos := make([][]int, m+1)
	for i, c := range a {
		pos[c] = append(pos[c], i)
	}

	at := make([]int, m+1)

	check := func(m int) bool {
		clear(at)
		for i, c := range a {
			at[c]++
			if at[c] >= m {
				j := pos[c][at[c]-m]
				if i-j+1-m <= k {
					return true
				}
			}
		}
		return false
	}

	l, r := 1, n+1
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}
