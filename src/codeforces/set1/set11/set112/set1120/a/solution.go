package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, ok, res := process(reader)
	if !ok {
		fmt.Println(-1)
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (n int, k int, a []int, b []int, ok bool, res []int) {
	nums := readNNums(reader, 4)
	m, k, n, s := nums[0], nums[1], nums[2], nums[3]
	a = readNNums(reader, m)
	b = readNNums(reader, s)
	ok, res = solve(n, k, a, b)
	return
}

const X = 500000 + 10

func solve(n int, k int, a []int, b []int) (bool, []int) {
	m := len(a)

	f1 := make([]int, X)
	for _, v := range b {
		f1[v]++
	}

	var cnt1 int
	for i := range X {
		if f1[i] > 0 {
			cnt1++
		}
	}

	f2 := make([]int, X)
	var cnt2 int

	add := func(x int) {
		f2[x]++
		if f2[x] == f1[x] {
			cnt2++
		}
	}

	rem := func(x int) {
		if f1[x] > 0 && f2[x] == f1[x] {
			cnt2--
		}
		f2[x]--
	}

	check := func(l int, r int) bool {
		// cnt2 == cnt1
		u := l / k
		u++
		u += (m - r - 1) / k
		return u >= n
	}

	play := func(l int, r int) []int {
		var res []int
		// 删除掉w个数
		u := min(n-1, l/k)
		w := l - u*k
		for i := range w {
			// 把前w个数删除掉
			res = append(res, i+1)
		}

		// 当前区间保留k个数，且包含序列b
		clear(f2)
		marked := make([]bool, m)
		// 先把b标记出来
		var v int
		for i := l; i <= r; i++ {
			x := a[i]
			if f1[x] > 0 {
				f2[x]++
				if f2[x] <= f1[x] {
					marked[i] = true
					v++
				}
			}
		}

		for i := l; i <= r; i++ {
			if marked[i] || v < k {
				if !marked[i] {
					v++
				}
				continue
			}
			// i是需要被删除的
			res = append(res, i+1)
		}

		u++
		// 后面的不用删除
		return res
	}

	for l, r := 0, 0; r < m; r++ {
		add(a[r])
		// 至少需要k个数 => r - l + 1 >= k => r >= l + k - 1
		for l+k-1 <= r && cnt2 == cnt1 {
			if check(l, r) {
				return true, play(l, r)
			}
			rem(a[l])
			if cnt2 < cnt1 {
				add(a[l])
				break
			}
			l++
		}
	}

	return false, nil
}
