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
	res := process(reader)
	fmt.Println(len(res))
	if len(res) == 1 {
		fmt.Println(res[0])
	} else {
		fmt.Println(res[0], res[1])
	}
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
	stones := make([][]int, n)
	for i := 0; i < n; i++ {
		stones[i] = readNNums(reader, 3)
	}
	return solve(stones)
}

type pair struct {
	first  int
	second int
}

func solve(stones [][]int) []int {

	mem := make(map[pair][]pair)

	add := func(a, b, c int, i int) {
		a, b = min(a, b), max(a, b)
		mem[pair{a, b}] = append(mem[pair{a, b}], pair{c, i})
	}

	for i, stone := range stones {
		a, b, c := stone[0], stone[1], stone[2]
		add(a, b, c, i)
		add(a, c, b, i)
		add(b, c, a, i)
	}

	for k, v := range mem {
		slices.SortFunc(v, func(a, b pair) int {
			return cmp.Or(b.first-a.first, a.second-b.second)
		})
		v = slices.Compact(v)
		mem[k] = v
	}

	var ans []int
	var best int

	for k, v := range mem {
		a, b := k.first, k.second
		var c int
		if len(v) == 1 {
			c = v[0].first
		} else {
			c = v[0].first + v[1].first
		}
		if min(a, b, c) > best {
			best = min(a, b, c)
			if len(v) == 1 {
				ans = []int{v[0].second + 1}
			} else {
				ans = []int{v[0].second + 1, v[1].second + 1}
			}
		}
	}

	return ans
}
