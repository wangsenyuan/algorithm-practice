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
	var n, a, b, p, q int
	fmt.Fscan(reader, &n, &a, &b, &p, &q)
	return solve(n, a, b, p, q)
}

func solve(n int, a int, b int, p int, q int) int {
	x := n / a
	y := n / b
	z := n / lcm(a, b)
	if p >= q {
		return x*p + q*(y-z)
	}
	return (x-z)*p + y*q
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
