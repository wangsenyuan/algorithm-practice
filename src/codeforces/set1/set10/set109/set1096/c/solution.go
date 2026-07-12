package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var tc int
	fmt.Fscan(reader, &tc)
	qs := make([]int, tc)
	for i := range tc {
		fmt.Fscan(reader, &qs[i])
	}
	return solve(qs)
}

func solve(qs []int) []int {
	res := make([]int, len(qs))
	for i, ang := range qs {
		res[i] = minPolygon(ang)
	}
	return res
}

func minPolygon(ang int) int {
	g := gcd(ang, 180)
	n := 180 / g
	k := ang / g
	if k+1 == n {
		n *= 2
	}
	return n
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
