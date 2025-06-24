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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const mod = 1_0000_0000

func solve(a []int) int {
	sort.Ints(a)

	n := len(a)
	sum := make([]int, n+1)
	for i, x := range a {
		sum[i+1] = sum[i] + x
	}

	var res int
	for i := 0; i < n; i++ {
		// a[i]的贡献
		j := search(i+1, n, func(j int) bool {
			return a[i]+a[j] >= mod
		})
		if i < j {
			res += (j-i-1)*a[i] + sum[j] - sum[i+1]
			if j < n {
				tmp := (n-j)*(a[i]-mod) + sum[n] - sum[j]
				res += tmp
			}
		}
	}

	return res
}

func search(l, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) >> 1
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
