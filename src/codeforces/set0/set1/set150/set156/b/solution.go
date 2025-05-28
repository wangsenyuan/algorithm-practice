package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) []string {
	n, m := readTwoNums(reader)
	answers := make([]string, n)
	for i := 0; i < n; i++ {
		answers[i] = readString(reader)
	}
	return solve(n, m, answers)
}

func solve(n int, m int, answers []string) []string {

	pos := make([]int, n+1)
	neg := make([]int, n+1)
	var sum int
	a := make([]int, n)
	for j, cur := range answers {
		var i int
		readInt([]byte(cur), 1, &i)
		if cur[0] == '+' {
			a[j] = i
			pos[i]++
		} else {
			a[j] = -i
			neg[i]++
			sum++
		}
	}

	var arr []int

	for i := 1; i <= n; i++ {
		tmp := pos[i] + sum - neg[i]
		if tmp == m {
			arr = append(arr, i)
		}
	}
	ans := make([]string, n)
	// len(arr) >= 1
	if len(arr) == 1 {
		x := arr[0]
		for j := range n {
			if a[j] == x {
				ans[j] = "Truth"
			} else if a[j] == -x {
				// 说x不是嫌犯的，在撒谎
				ans[j] = "Lie"
			} else if a[j] > 0 {
				// 说y(!= x)是嫌犯的，在撒谎
				ans[j] = "Lie"
			} else {
				// 说不是y(y!=x)的，在说真话
				ans[j] = "Truth"
			}
		}
		return ans
	}
	// len(arr) > 1
	suspect := make([]bool, n+1)
	for _, x := range arr {
		suspect[x] = true
	}
	for j := range n {
		x := a[j]
		if x > 0 {
			if !suspect[x] {
				ans[j] = "Lie"
			} else {
				ans[j] = "Not defined"
			}
		} else {
			x = -x
			if !suspect[x] {
				ans[j] = "Truth"
			} else {
				ans[j] = "Not defined"
			}
		}
	}
	return ans
}
