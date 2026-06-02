package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	l, r := readTwoNums(reader)
	res := solve(l, r)
	fmt.Println(res)
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
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

func solve(l int, r int) int {
	m := int(math.Sqrt(float64(r))) + 1

	var primes []int
	lpf := make([]int, m+1)
	for x := 2; x <= m; x++ {
		if lpf[x] == 0 {
			lpf[x] = x
			primes = append(primes, x)
		}
		for _, p := range primes {
			if x*p > m {
				break
			}
			lpf[x*p] = p
			if x%p == 0 {
				break
			}
		}
	}

	diff := r - l

	marked := make([]bool, diff)

	ans := 1

	for _, p := range primes {
		x := (l/p + 1) * p
		for x <= r {
			if !marked[x-l-1] {
				marked[x-l-1] = true
				y := x
				for y%p == 0 {
					y /= p
				}
				if y == 1 {
					ans++
				}
			}

			x += p
		}
	}

	for _, v := range marked {
		if !v {
			ans++
		}
	}

	return ans
}
