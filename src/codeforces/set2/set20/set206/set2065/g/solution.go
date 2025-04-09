package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const N = 2e5 + 10

var lpf [N + 1]int
var sem_prime [N]bool

func init() {
	var primes []int
	for i := 2; i <= N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > N {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	for i := 2; i <= N; i++ {
		if lpf[i] == i {
			for j := i * 2; j <= N; j += i {
				if lpf[j/i] == j/i {
					sem_prime[j] = true
				}
			}
		}
	}
}

func solve(a []int) int {
	n := len(a)
	cnt := make([]int, n+1)
	var res int
	sort.Ints(a)

	var prime_cnt int

	for i := n - 1; i >= 0; i-- {
		v := a[i]
		if lpf[v] != v {
			if sem_prime[v] {
				cnt[v]++
				res += cnt[v]
			}
			continue
		}
		j := i
		for i >= 0 && a[i] == a[j] {
			i--
		}
		cur := j - i
		i++
		res += cur * prime_cnt
		prime_cnt += cur

		for u := 2 * v; u <= n; u += v {
			if sem_prime[u] {
				res += cnt[u] * cur
			}
		}
	}

	return res
}

func bruteForce(a []int) int {
	sort.Ints(a)
	var res int
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			v := lcm(a[i], a[j])
			if v <= N && sem_prime[v] {
				res++
			}
		}
	}
	return res
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
