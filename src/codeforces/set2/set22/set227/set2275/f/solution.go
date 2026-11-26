package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
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

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const X = 1 << 20

var lpf [X]int
var rep [X]int

func init() {
	var primes []int
	var randPrimes []int

	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			randPrimes = append(randPrimes, rand.Int())
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= X {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	for i := 2; i < X; i++ {
		var w int
		for i1 := i; i1 > 1; {
			j := lpf[i1]
			var c int
			for i1%j == 0 {
				i1 /= j
				c++
			}
			if c&1 == 1 {
				pos, _ := slices.BinarySearch(primes, j)
				w ^= randPrimes[pos]
			}
		}
		rep[i] = w
	}
}

func solve(a []int) int64 {
	// 如果 j < i, 那么 a[0]... a[j] * a[i] 是一个完全平方数,
	// 如果 i <= j, 那么就是 a[0]...a[j] / a[i] 是一个完全平方数
	freq := make(map[int]int)
	for _, v := range a {
		freq[rep[v]]++
	}

	var ans int
	var num int

	for _, v := range a {
		num ^= rep[v]
		ans += freq[num]
	}

	return int64(ans)
}
