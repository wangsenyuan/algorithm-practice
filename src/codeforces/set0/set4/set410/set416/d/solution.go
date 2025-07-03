package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	next := make([]int, n+1)
	next[n] = n
	for i := n - 1; i >= 0; i-- {
		if i+1 == n || a[i+1] != -1 {
			next[i] = i + 1
		} else {
			next[i] = next[i+1]
		}
	}

	check := func(l int, r int, d int) bool {
		if r == n || a[r] < 0 {
			return a[l]+(r-l)*d > 0
		}
		return a[r] == a[l]+(r-l)*d
	}

	if a[0] == -1 && next[0] == n {
		return 1
	}

	var res int

	findNextCorrectPosition := func(i int) int {
		i1 := i
		if next[i] == n || next[next[i]] == n {
			res++
			return n
		}
		i = next[i]
		j := next[i]
		// j < n
		if (a[j]-a[i])%(j-i) != 0 {
			// j前面的全部设置为a[i], 从j开始处理
			res++
			return j
		} else {
			// 检查i前面的是否也满足条件
			d := (a[j] - a[i]) / (j - i)
			// a[0] + i * d = a[i], 保证a[0] > 0
			if a[i]-(i-i1)*d <= 0 {
				// 采用a[i]策略
				res++
				return j
			}
			// else, 还是从i开始
		}
		return i
	}

	for i := 0; i < n; {
		if a[i] < 0 {
			i = findNextCorrectPosition(i)
			if i == n {
				break
			}
		}

		res++
		j := i
		i = next[i]
		if i == n {
			break
		}
		// a[i] - a[j] = (j - i) * d
		if (a[i]-a[j])%(i-j) != 0 {
			continue
		}
		d := (a[i] - a[j]) / (i - j)
		for i < n && check(j, i, d) {
			i++
		}
	}

	return res
}
