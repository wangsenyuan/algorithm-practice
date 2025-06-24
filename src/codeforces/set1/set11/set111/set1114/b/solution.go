package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, sum, res := process(reader)
	fmt.Println(sum)
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

func process(reader *bufio.Reader) (a []int, m int, k int, sum int, res []int) {
	n, m, k := readThreeNums(reader)
	a = readNNums(reader, n)
	sum, res = solve(a, m, k)
	return
}

func solve(a []int, m int, k int) (int, []int) {
	n := len(a)
	// 要舍弃掉这么多的数
	u := n - m*k
	type pair struct {
		first  int
		second int
	}
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}
	slices.SortFunc(arr, func(x, y pair) int {
		return cmp.Or(x.first-y.first, x.second-y.second)
	})
	marked := make([]bool, n)
	for i := range u {
		marked[arr[i].second] = true
	}
	var sum int
	var res []int
	var cnt int
	for i := range n {
		if !marked[i] {
			sum += a[i]
			cnt++
		}
		if cnt == m {
			res = append(res, i+1)
			cnt = 0
		}
	}
	// len(res) = k
	return sum, res[:k-1]
}
