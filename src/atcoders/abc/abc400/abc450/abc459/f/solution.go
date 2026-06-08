package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	for _, x := range res {
		fmt.Println(x)
	}
}

func drive(reader *bufio.Reader) []int {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([]int, tc)
	for i := range tc {
		res[i] = readCase(reader)
	}
	return res
}

func readCase(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	for i := range n {
		a[i] -= i
	}
	var d []pair

	for _, v := range a {
		now := pair{1, v}
		for len(d) > 0 {
			tail := d[len(d)-1]
			if divCeil(tail.second, tail.first) <= divFloor(now.second, now.first) {
				break
			}
			now.first += tail.first
			now.second += tail.second
			d = d[:len(d)-1]
		}
		d = append(d, now)
	}

	var b []int
	for _, v := range d {
		l, s := v.first, v.second
		for i := range l {
			b = append(b, divFloor(s+i, l))
		}
	}
	var res int
	for i := range n {
		res += i * (b[i] - a[i])
	}
	return res
}

type pair struct {
	first  int
	second int
}

func divFloor(a int, b int) int {
	res := a / b
	if a < 0 && a%b != 0 {
		res--
	}
	return res
}

func divCeil(a int, b int) int {
	res := a / b
	if a > 0 && a%b != 0 {
		res++
	}
	return res
}
