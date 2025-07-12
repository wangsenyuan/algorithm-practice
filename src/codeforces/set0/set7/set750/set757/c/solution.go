package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)
	gyms := make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		gyms[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &gyms[i][j])
		}
	}
	return solve(m, gyms)
}

func solve(m int, gyms [][]int) int {
	n := len(gyms)
	mul := make([]int, n)

	r := rand.New(rand.NewSource(99))

	for i := range n {
		mul[i] = r.Int() + 1
	}

	vals := make([]int, m)

	freq := make([]int, m)
	pos := make([]int, m)
	for j, cur := range gyms {
		for i, v := range cur {
			v--
			freq[v]++
			pos[v] = i + 1
		}

		for i, v := range cur {
			v--
			if pos[v] == i+1 {
				vals[v] += mul[j] * freq[v]
				pos[v] = 0
				freq[v] = 0
			}
		}
	}

	sort.Ints(vals)

	P := make([]int, m+1)
	P[0] = 1
	for i := 1; i < len(P); i++ {
		P[i] = i * P[i-1]
		P[i] %= MOD
	}

	ans := 1

	for i := 0; i < m; {
		j := i
		for i < m && vals[i] == vals[j] {
			i++
		}
		ans *= P[i-j]
		ans %= MOD
	}

	return ans
}

const MOD = 1000000007
