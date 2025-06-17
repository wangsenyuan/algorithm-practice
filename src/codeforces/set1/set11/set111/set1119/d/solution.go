package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	s := readNNums(reader, n)
	m := readNum(reader)
	qs := make([][]int, m)
	for i := range m {
		qs[i] = readNNums(reader, 2)
	}
	return solve(s, qs)
}

func solve(s []int, queries [][]int) []int {
	sort.Ints(s)
	s = slices.Compact(s)
	n := len(s)
	// 1, 5, 6, 10, 100
	// 当q = 3 时
	// 1, 2, 3, 5, 6, 7, 8, 10, 11, 12, 100, 101, 102

	type pair struct {
		first  int
		second int
	}

	qs := make([]pair, len(queries))

	for i, q := range queries {
		l, r := q[0], q[1]
		qs[i] = pair{r - l + 1, i}
	}

	ans := make([]int, len(queries))
	if n == 1 {
		for i := range qs {
			ans[i] = qs[i].first
		}
		return ans
	}

	slices.SortFunc(qs, func(a, b pair) int {
		return a.first - b.first
	})

	arr := make([]pair, n-1)
	for i := 0; i < n-1; i++ {
		arr[i] = pair{s[i+1] - s[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})
	var i int
	var sum int
	for _, q := range qs {
		sz := q.first
		for i < n-1 && arr[i].first <= sz {
			sum += arr[i].first
			i++
		}
		ans[q.second] = sum + (n-i)*sz
	}

	return ans
}
