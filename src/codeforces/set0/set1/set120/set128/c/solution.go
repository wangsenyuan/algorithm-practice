package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	n, m, k := readThreeNums(bufio.NewReader(os.Stdin))
	fmt.Println(solve(n, m, k))
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

const mod = 1000000007

func add(a int, b int) int {
	return (a + b) % mod
}

func mul(a int, b int) int {
	return a * b % mod
}

func pow(a int, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, m int, k int) int {
	k *= 2
	if k > min(n-1, m-1) {
		return 0
	}
	f := make([]int, max(n, m))
	f[0] = 1
	for i := 1; i < len(f); i++ {
		f[i] = mul(f[i-1], i)
	}
	inv := make([]int, len(f))
	inv[len(f)-1] = pow(f[len(f)-1], mod-2)
	for i := len(f) - 1; i > 0; i-- {
		inv[i-1] = mul(inv[i], i)
	}

	nCr := func(n int, r int) int {
		return mul(f[n], mul(inv[r], inv[n-r]))
	}

	return mul(nCr(n-1, k), nCr(m-1, k))
}
