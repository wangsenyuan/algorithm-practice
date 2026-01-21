package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var a, b int
	fmt.Fscan(reader, &a, &b)
	return solve(a, b)
}

func solve(a int, b int) int {
	if a == b {
		return 0
	}

	// Ensure a <= b for easier handling
	if a > b {
		a, b = b, a
	}

	d := b - a

	// Get all divisors of d
	divisors := getDivisors(d)

	bestK := 0
	bestLcm := lcm(a, b)

	// For each divisor g of d, find k such that g divides (a+k)
	// We want to maximize gcd(a+k, d) to minimize lcm
	for _, g := range divisors {
		// Find smallest k >= 0 such that g | (a+k)
		// k ≡ -a (mod g), so k = (g - a%g) % g
		rem := a % g
		var k int
		if rem == 0 {
			k = 0
		} else {
			k = g - rem
		}

		ak := a + k
		bk := b + k
		currentLcm := lcm(ak, bk)

		if currentLcm < bestLcm || (currentLcm == bestLcm && k < bestK) {
			bestLcm = currentLcm
			bestK = k
		}
	}

	return bestK
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func getDivisors(n int) []int {
	var divisors []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i*i != n {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}
