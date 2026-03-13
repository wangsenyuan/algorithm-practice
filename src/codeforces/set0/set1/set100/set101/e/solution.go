package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, best, res := drive(reader)
	fmt.Println(best)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (x []int, y []int, p int, best int, res string) {
	var n, m int
	fmt.Fscan(reader, &n, &m, &p)
	x = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	y = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &y[i])
	}
	best, res = solve(x, y, p)
	return
}

var (
	gx, gy       []int
	gp, gn, gm   int
	garr1, garr2 []int
	garr3        []int
	gpath        []byte
)

func solve(x []int, y []int, p int) (best int, res string) {
	gn, gm, gp = len(x), len(y), p
	gx, gy = x, y
	garr1 = make([]int, gn)
	garr2 = make([]int, gn)
	garr3 = make([]int, gn)
	gpath = make([]byte, gn+gm-1)

	best = f(0, 0) + build(0, 0, gn-1, gm-1)
	res = string(gpath[1:])
	return
}

func f(i int, j int) int {
	return (gx[i] + gy[j]) % gp
}

func dp(r1 int, c1 int, r2 int, c2 int, k int) []int {
	const neg = -1 << 60
	for i := r1; i <= r2; i++ {
		garr1[i] = neg
		garr2[i] = neg
	}
	garr1[r1] = f(r1, c1)
	prevLo, prevHi := r1, r1

	for x := r1 + c1 + 1; x <= k; x++ {
		lo := max(r1, x-c2)
		hi := min(r2, x-c1)

		for i := lo; i <= hi; i++ {
			j := x - i
			garr2[i] = f(i, j)
			bestHere := neg
			if i > r1 && prevLo <= i-1 && i-1 <= prevHi {
				bestHere = max(bestHere, garr1[i-1])
			}
			if j > c1 && prevLo <= i && i <= prevHi {
				bestHere = max(bestHere, garr1[i])
			}
			garr2[i] += bestHere
		}

		for i := prevLo; i <= prevHi; i++ {
			garr1[i] = neg
		}
		for i := lo; i <= hi; i++ {
			garr1[i] = garr2[i]
			garr2[i] = neg
		}
		prevLo, prevHi = lo, hi
	}
	return garr1
}

func fp(r1 int, c1 int, r2 int, c2 int, k int) []int {
	const neg = -1 << 60
	for i := r1; i <= r2; i++ {
		garr3[i] = neg
		garr2[i] = neg
	}
	garr3[r2] = f(r2, c2)
	prevLo, prevHi := r2, r2

	for x := r2 + c2 - 1; x >= k; x-- {
		lo := max(r1, x-c2)
		hi := min(r2, x-c1)
		for i := lo; i <= hi; i++ {
			j := x - i
			garr2[i] = f(i, j)
			bestHere := neg
			if i < r2 && prevLo <= i+1 && i+1 <= prevHi {
				bestHere = max(bestHere, garr3[i+1])
			}
			if j < c2 && prevLo <= i && i <= prevHi {
				bestHere = max(bestHere, garr3[i])
			}
			garr2[i] += bestHere
		}

		for i := prevLo; i <= prevHi; i++ {
			garr3[i] = neg
		}
		for i := lo; i <= hi; i++ {
			garr3[i] = garr2[i]
			garr2[i] = neg
		}
		prevLo, prevHi = lo, hi
	}
	return garr3
}

func build(r1 int, c1 int, r2 int, c2 int) int {
	var sum int
	if r1 == r2 {
		for j := c1; j < c2; j++ {
			gpath[r1+j+1] = 'S'
			sum += f(r1, j+1)
		}
		return sum
	}
	if c1 == c2 {
		for i := r1; i < r2; i++ {
			gpath[i+c1+1] = 'C'
			sum += f(i+1, c1)
		}
		return sum
	}

	k := (r1 + c1 + r2 + c2) / 2
	d1 := dp(r1, c1, r2, c2, k)
	d2 := fp(r1, c1, r2, c2, k)

	lo := max(r1, k-c2)
	hi := min(r2, k-c1)

	bestPair := []int{-1 << 60, r1}
	for i := lo; i <= hi; i++ {
		j := k - i
		tmp := d1[i] + d2[i] - f(i, j)
		if tmp > bestPair[0] {
			bestPair[0] = tmp
			bestPair[1] = i
		}
	}

	i := bestPair[1]
	j := k - bestPair[1]
	sum += build(r1, c1, i, j)
	sum += build(i, j, r2, c2)
	return sum
}
