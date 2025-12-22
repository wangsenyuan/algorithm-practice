package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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
	b := readNNums(reader, n)
	return solve(a, b) 
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int) int {
	n := len(b)
	buf := make([]pair, n)
	for i := 0; i < n; i++ {
		buf[i] = pair{b[i], i}
	}
	slices.SortFunc(buf, func(x, y pair) int {
		return x.first - y.first
	})

	arr := make([]pair, n)
	for i := 0; i < n; i++ {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return x.first - y.first
	})

	tmp := make([]int, n)

	for i := 0; i < n; i++ {
		tmp[arr[i].second] = buf[i].second
	}

	return countInversions(tmp) % (1e8 - 3)
}

func countInversions(arr []int) int {
	n := len(arr)

	set := make(BIT, n+10)

	var res int

	for i, v := range arr {
		res += i - set.Query(v)
		set.Update(v, 1)
	}

	return res
}

type BIT []int

func (bit BIT) Update(pos int, v int) {
	pos++
	for pos < len(bit) {
		bit[pos] += v
		pos += pos & -pos
	}
}

func (bit BIT) Query(pos int) int {
	pos++
	var res int
	for pos > 0 {
		res += bit[pos]
		pos -= pos & -pos
	}
	return res
}
