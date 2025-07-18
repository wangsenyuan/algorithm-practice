package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, l, r := readThreeNums(reader)
	res := solve(n, l, r)
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

func solve(n int, l int, r int) int {

	mem := make(map[int]int)

	var f func(x int) int
	f = func(x int) int {
		if v, ok := mem[x]; ok {
			return v
		}
		if x <= 1 {
			return 1
		}
		res := 1 + 2*f(x/2)
		mem[x] = res
		return res
	}

	var g func(x int, l int, r int) int

	g = func(x int, l int, r int) int {
		if x <= 1 {
			return x
		}
		w := f(x / 2)
		// 左边有w个，中间有1个，右边有w个
		var res int
		if l <= w {
			res += g(x/2, l, min(r, w))
		}
		if l <= w+1 && w+1 <= r {
			res += x & 1
		}
		// 右边也有w个
		if w+1 < r {
			res += g(x/2, max(1, l-(w+1)), r-(w+1))
		}

		return res
	}

	return g(n, l, r)
}
