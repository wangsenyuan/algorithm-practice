package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
	for _, x := range res {
		fmt.Println(x)
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

func solve(n int) []string {
	if n == 2 {
		return []string{"-1"}
	}
	// n >= 3
	nums := []int{15, 10, 6}
	for i := 3; i < n; i++ {
		nums = append(nums, 6*(i-1))
	}

	res := make([]string, n)
	for i := range n {
		res[i] = fmt.Sprintf("%d", nums[i])
	}
	return res
}

func solve1(n int) []string {
	if n == 2 {
		return []string{"-1"}
	}
	// n >= 3
	// first n  prime numbers
	// 2, 3, 5
	// 3 * 5, 2 * 5, 3 * 2
	var primes []int
	lpf := make([]int, 1000)
	for i := 2; i < 1000 && len(primes) < n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= 1000 {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}

	get := func(i int) string {
		res := big.NewInt(1)

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			res = res.Mul(res, big.NewInt(int64(primes[j])))
		}

		return res.String()
	}

	ans := make([]string, n)
	for i := range n {
		ans[i] = get(i)
	}
	return ans
}
