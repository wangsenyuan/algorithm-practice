package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) int {
	p := mersennePrimeExponents[n-1]
	return (pow(2, p-1) - 1 + mod) % mod
}

const mod = 1000000007

var mersennePrimeExponents = []int{
	2, 3, 5, 7, 13, 17, 19, 31, 61, 89,
	107, 127, 521, 607, 1279, 2203, 2281, 3217, 4253, 4423,
	9689, 9941, 11213, 19937, 21701, 23209, 44497, 86243, 110503, 132049,
	216091, 756839, 859433, 1257787, 1398269, 2976221, 3021377,
	6972593, 13466917, 20996011,
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = int(int64(res) * int64(a) % mod)
		}
		a = int(int64(a) * int64(a) % mod)
		b >>= 1
	}
	return res
}
