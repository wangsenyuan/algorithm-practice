package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	q := readNum(reader)
	for range q {
		a, b, m := readThreeNums(reader)
		res := solve(a, b, m)
		if res == nil {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprint(writer, len(res))
			for _, v := range res {
				fmt.Fprint(writer, " ", v)
			}
			fmt.Fprintln(writer)
		}
	}
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

// solve finds an m-cute sequence starting at a and ending at b.
// Returns nil if impossible.
//
// Formula: xₙ = 2^(n-2)·a + 2^(n-3)·r₂ + ... + rₙ₋₁ + rₙ
// Bounds:  2^(n-2)·(a+1) ≤ b ≤ 2^(n-2)·(a+m)
//
// Construction: write b = 2^(n-2)·(a+k) + r where 1 ≤ k ≤ m, 0 ≤ r < 2^(n-2).
// Then rᵢ = k + d_{n-1-i}, where d_j is the j-th bit of r (d_{-1} = 0).
func solve(a int, b int, m int) []int {
	if a == b {
		return []int{a}
	}

	// Try each sequence length n from 2 to 50
	for n := 2; n <= 50; n++ {
		pow := 1 << (n - 2) // 2^(n-2)

		// Check lo = pow*(a+1) <= b, i.e., pow <= b/(a+1)
		if pow > b/(a+1) {
			break // pow only grows, no more valid n
		}

		// Check hi = pow*(a+m) >= b, i.e., pow >= ceil(b/(a+m))
		if pow < (b+a+m-1)/(a+m) {
			continue
		}

		// b is in [pow*(a+1), pow*(a+m)]. Construct the sequence.
		// Decompose: b = pow*(a+k) + r
		t := b - pow*a
		k := t / pow
		r := t - k*pow

		seq := make([]int, n)
		seq[0] = a
		s := a // running prefix sum
		for j := 1; j < n-1; j++ {
			bit := (r >> (n - 2 - j)) & 1
			ri := k + bit
			seq[j] = s + ri
			s += seq[j]
		}
		seq[n-1] = b
		return seq
	}

	return nil
}
