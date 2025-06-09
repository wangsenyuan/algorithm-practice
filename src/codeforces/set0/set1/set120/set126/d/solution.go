package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	for range tc {
		n := readNum(reader)
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func solve(num int) int {
	if num <= 2 {
		return 1
	}
	f := []int{1, 1}
	for f[len(f)-1] < num {
		f = append(f, f[len(f)-1]+f[len(f)-2])
	}
	n := len(f)

	var arr []int
	for i := n - 1; i >= 0 && num > 0; i-- {
		if num >= f[i] {
			arr = append(arr, i)
			num -= f[i]
		}
	}
	reverse(arr)
	dp1 := 1
	dp2 := (arr[0] - 1) / 2

	for i := 1; i < len(arr); i++ {
		// 不展开当前位置
		ndp1 := dp1 + dp2
		// 为啥是 (两个位置的距离) / 2呢？
		// 10001
		// 01101
		// 分解后，看头部的1在哪里，它后面的1(分解出来的)不能继续分解
		// 所以这里就看头部的1有多少个位置可以到达
		// 如果张开当前位的1，那么可选的位置，只有(arr[i] - 1 - arr[i-1]) / 2
		// 如果前一个位置也展开了，那么多一个位置可用
		ndp2 := (arr[i]-arr[i-1]-1)/2*dp1 + (arr[i]-arr[i-1])/2*dp2
		dp1, dp2 = ndp1, ndp2
	}
	return dp1 + dp2
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
