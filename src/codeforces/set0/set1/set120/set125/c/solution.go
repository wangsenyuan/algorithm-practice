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
	res := solve(n)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, v := range res {
		for _, x := range v {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
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

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func solve(n int) [][]int {
	var res [][]int
	res = append(res, []int{1, 2}, []int{1, 3}, []int{2, 3})
	if n <= 5 {
		return res
	}
	// 假设m天，每天所有人都参加
	// n = 6
	// 1 2 4 7
	// 1 3 5 8
	// 2 3 6 9
	// 4 5 6 10
	// 7 8 9 10
	for i := 4; i < n; {
		m := len(res)
		if i+m-1 > n {
			break
		}
		var tmp []int
		for j := 0; j < m; j++ {
			tmp = append(tmp, i+j)
			if j < len(res) {
				res[j] = append(res[j], i+j)
			}
		}
		res = append(res, tmp)
		i += m
	}
	return res
}
