package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

type pair struct {
	first  int
	second int
}

// x.first / x.second < y.first / y.second
func comparePair(x, y pair) int {
	return x.first*y.second - y.first*x.second
}

const mod = 998244353

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return (a * b) % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func solve(a, b []int) int {
	var arr1 []pair

	// a[i] / a[j]
	for j := range len(a) {
		for i := range j {
			x, y := a[i], a[j]
			z := gcd(x, y)
			arr1 = append(arr1, pair{x / z, y / z})
		}
	}

	slices.SortFunc(arr1, comparePair)

	var arr2 []pair

	for i := range len(b) {
		for j := range len(b) {
			if i != j {
				x, y := b[i], b[j]
				z := gcd(x, y)
				arr2 = append(arr2, pair{x / z, y / z})
			}
		}
	}

	slices.SortFunc(arr2, comparePair)

	var sum int
	for i, j := 0, 0; i < len(arr1); i++ {
		for j < len(arr2) && comparePair(arr2[j], arr1[i]) < 0 {
			j++
		}
		sum += j
	}

	w := mul(len(a), len(a)-1)

	w = pow(w, mod-2)

	return mul(sum%mod, w)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
