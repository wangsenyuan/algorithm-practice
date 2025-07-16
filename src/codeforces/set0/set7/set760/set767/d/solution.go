package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, ok, res := process(reader)
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

func process(reader *bufio.Reader) (k int, a []int, b []int, ok bool, res []int) {
	n, m, k := readThreeNums(reader)
	a = readNNums(reader, n)
	b = readNNums(reader, m)
	ok, res = solve(k, a, b)
	return
}

type milk struct {
	id  int
	exp int
}

func solve(k int, a []int, b []int) (bool, []int) {
	n := len(a)
	m := len(b)

	arr := make([]milk, n+m)
	for i := range n {
		arr[i] = milk{-(i + 1), a[i]}
	}
	for i := range m {
		arr[n+i] = milk{i + 1, b[i]}
	}

	slices.SortFunc(arr, func(x, y milk) int {
		return cmp.Or(x.exp-y.exp, x.id-y.id)
	})

	var res []int
	var free int
	prev := -1
	for i := 0; i < n+m; {
		j := i
		free += (arr[j].exp - prev) * k
		for i < n+m && arr[i].exp == arr[j].exp && arr[i].id < 0 {
			i++
		}
		if i-j > free {
			return false, nil
		}
		free -= i - j

		for i < n+m && arr[i].exp == arr[j].exp {
			if free > 0 {
				res = append(res, arr[i].id)
				free--
			}
			i++
		}
		prev = arr[j].exp
	}

	return true, res
}
