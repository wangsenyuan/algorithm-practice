package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, d := readTwoNums(reader)
		res := solve(n, d)
		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
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

func solve(n int, d int) []int {
	if d > n*(n-1)/2 {
		// 肯定没有结果
		return nil
	}

	dep := make([]int16, n)
	cnt := make([]int16, n)
	var sum int

	fa := make([]int, n)
	for i := 0; i < n; i++ {
		dep[i] = int16(i)
		fa[i] = i - 1
		if i > 0 {
			cnt[i-1]++
		}
		sum += i
	}

	bad := make([]bool, n)

	for sum > d {
		v := -1
		for i := 0; i < n; i++ {
			if !bad[i] && cnt[i] == 0 && (v < 0 || dep[v] > dep[i]) {
				v = i
			}
		}
		if v < 0 {
			return nil
		}
		p := -1
		for i := 0; i < n; i++ {
			if cnt[i] < 2 && dep[i] == dep[v]-2 {
				p = i
				break
			}
		}
		if p < 0 {
			bad[v] = true
			continue
		}
		sum--
		cnt[fa[v]]--
		cnt[p]++
		dep[v]--
		fa[v] = p
	}

	for i := 0; i < n; i++ {
		fa[i]++
	}

	return fa[1:]
}
