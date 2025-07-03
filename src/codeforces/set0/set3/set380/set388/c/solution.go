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
	fmt.Println(res[0], res[1])
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
	piles := make([][]int, n)
	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k)
		piles[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos+1, &piles[i][j])
		}
	}

	return solve(piles)
}

func solve(piles [][]int) []int {
	// n := len(piles)
	res := make([]int, 2)
	var arr []int
	for _, cur := range piles {
		l, r := 0, len(cur)-1
		for l < r {
			res[0] += cur[l]
			res[1] += cur[r]
			l++
			r--
		}
		if l == r {
			arr = append(arr, cur[l])
		}
	}

	sort.Ints(arr)

	slices.Reverse(arr)

	for i := 0; i < len(arr); i++ {
		res[i%2] += arr[i]
	}

	return res
}
