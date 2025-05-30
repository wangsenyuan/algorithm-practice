package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func getPrimes(n int) []int {
	var primes []int
	lpf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, v := range primes {
			if i*v > n {
				break
			}
			lpf[i*v] = v
			if i%v == 0 {
				break
			}
		}
	}

	for i := 4; i <= n; i++ {
		if lpf[i] != i {
			lpf[i] = lpf[i-1]
		}
	}
	return lpf
}

func solve(a []int) [][]int {
	n := len(a)
	lpf := getPrimes(n)
	var res [][]int
	marked := make([]bool, n+1)

	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pos[a[i-1]] = i
	}

	var dfs func(u int, v int)
	dfs = func(u int, v int) {
		// 目标是交换u, v
		d := v - u + 1
		// 用最大的质数去处理
		p := lpf[d]
		w := v - p + 1
		res = append(res, []int{w, v})
		if w != u {
			// 交换u, w
			dfs(u, w)
			// 交换w,v
			res = append(res, []int{w, v})
		}
	}

	for i := 1; i <= n; i++ {
		if !marked[i] {
			var tmp []int
			j := i
			for !marked[j] {
				tmp = append(tmp, j)
				marked[j] = true
				j = pos[j]
			}

			for i := 1; i < len(tmp); i++ {
				u := tmp[i-1]
				v := tmp[i]
				if u > v {
					u, v = v, u
				}
				dfs(u, v)
			}
		}
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
