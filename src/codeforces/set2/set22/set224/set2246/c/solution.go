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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func add(a, b int) int {
	return (a + b) % mod
}

const N = 2e5 + 10

var PW [N]int

func init() {
	PW[0] = 1
	for i := 1; i < N; i++ {
		PW[i] = mul(PW[i-1], 2)
	}
}

type data struct {
	val  int
	len  int
	ways int
}

func solve(a []int) int {
	res1 := 1
	n := len(a)
	var arr []data
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		// 使用偶数个a[j], 包括0个
		res1 = mul(res1, PW[i-j-1])
		if a[j] > 0 {
			if len(arr) == 0 {
				arr = append(arr, data{val: a[j], len: i - j, ways: PW[i-j-1]})
			} else {
				last := arr[len(arr)-1]
				cur := data{val: a[j], len: i - j, ways: mul(last.ways, PW[i-j-1])}
				arr = append(arr, cur)
			}
		}
	}
	if a[0] == -1 {
		var cnt int
		for cnt < n && a[cnt] == -1 {
			cnt++
		}
		res2 := PW[cnt-1]
		suf := 1
		var res3 int
		// 后面要构造一个 a - b = -1的结果出来
		for i := len(arr) - 1; i > 0; i-- {
			if arr[i].val == arr[i-1].val+1 {
				// arr[i] 和 arr[i-1]必须选择奇数个, 也正好是pw[len-1]
				tmp := mul(PW[arr[i].len-1], PW[arr[i-1].len-1])
				tmp = mul(tmp, suf)
				if i > 1 {
					tmp = mul(tmp, arr[i-2].ways)
				}
				res3 = add(res3, tmp)
			}
			suf = mul(suf, PW[arr[i].len-1])
		}

		res1 = add(res1, mul(res2, res3))
	}

	return res1
}
