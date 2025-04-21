package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(a []int) int {
	// n := len(a)

	var res int

	for i, v := range a {
		res = add(res, mul(v, i))
		// cnt[numLen(v)]++
	}
	cnt := make([]int, 11)

	base := make([]int, 11)
	base[0] = 1
	for i := 1; i <= 10; i++ {
		base[i] = mul(base[i-1], 10)
	}

	n := len(a)
	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= 10; j++ {
			tmp := mul(a[i], base[j])
			res = add(res, mul(tmp, cnt[j]))
		}
		cnt[numLen(a[i])]++
	}
	return res
}

func numLen(num int) int {
	var res int
	for num > 0 {
		res++
		num /= 10
	}
	return res
}
