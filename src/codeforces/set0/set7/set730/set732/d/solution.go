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
	d := readNNums(reader, n)
	a := readNNums(reader, m)
	return solve(d, a)
}

func solve(d []int, a []int) int {
	n := len(d)
	m := len(a)

	take := make([]int, m+1)

	check := func(mid int) bool {
		if mid > n {
			return false
		}
		clear(take)
		for i := 0; i < mid; i++ {
			if d[i] > 0 {
				take[d[i]] = i + 1
			}
		}
		for i := 1; i <= m; i++ {
			if take[i] == 0 {
				return false
			}
		}
		var prepare int
		for i := 0; i < mid; i++ {
			if d[i] == 0 {
				prepare++
				continue
			}
			if take[d[i]] == i+1 {
				// 这是必须参加这次考试的时间
				if a[d[i]-1] > prepare {
					return false
				}
				prepare -= a[d[i]-1]
			} else {
				prepare++
			}
		}
		return true
	}

	if !check(n) {
		return -1
	}

	return sort.Search(n, check)
}
