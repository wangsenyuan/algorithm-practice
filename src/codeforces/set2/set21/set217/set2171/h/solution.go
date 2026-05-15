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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

const X = 2e5 + 10

var lpf [X]int

func init() {
	var primes []int
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if j*i >= X {
				break
			}
			lpf[j*i] = j
			if i%j == 0 {
				break
			}
		}
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	return solve(n, m)
}

func solve(n int, m int) int {

	calc := func(i int, a int) int {
		var res int
		for a%i == 0 {
			res++
			a /= i
		}
		return res
	}

	bit := make(BIT, m-n+2)

	for i := 2; i <= n; i++ {
		for j := (m - n + i) / i * i; j >= i; j -= i {
			v := bit.query(j-i) + calc(i, j)
			bit.update(j-i, v)
		}
	}

	return bit.query(m - n)
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] = max(bit[i], v)
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	i++
	var res int
	for i > 0 {
		res = max(res, bit[i])
		i -= i & -i
	}
	return res
}
