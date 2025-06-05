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
		n, m, k := readThreeNums(reader)
		res := solve(n, m, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(x int, y int, k int) int {
	g := gcd(x, y)

	res := get(x/g, k)
	if res < 0 {
		return -1
	}
	res2 := get(y/g, k)
	if res2 < 0 {
		return -1
	}
	return res + res2
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func get(x int, k int) int {
	if x == 1 {
		return 0
	}
	var divs []int
	for i := 1; i <= x/i; i++ {
		if x%i == 0 {
			divs = append(divs, i)
			if i*i != x {
				divs = append(divs, x/i)
			}
		}
	}
	sort.Ints(divs)
	n := len(divs)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 100
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if divs[i]/divs[j] > k {
				break
			}
			if divs[i]%divs[j] == 0 {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	if dp[n-1] == 100 {
		return -1
	}
	return dp[n-1]
}
