package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	flats := make([][]int, m)
	for i := 0; i < m; i++ {
		flats[i] = readNNums(reader, 2)
	}
	return solve(n, flats)
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

const inf = 1 << 30

func solve(n int, flats [][]int) int {
	if n == 1 {
		return 1
	}
	if len(flats) == 0 {
		return -1
	}
	slices.SortFunc(flats, func(a, b []int) int {
		return a[0] - b[0]
	})
	var arr []int
	for k := 1; k <= 100; k++ {
		ok := true
		for _, cur := range flats {
			if (cur[0]+k-1)/k != cur[1] {
				ok = false
				break
			}
		}
		if ok {
			arr = append(arr, (n+k-1)/k)
		}
	}
	sort.Ints(arr)
	arr = slices.Compact(arr)
	if len(arr) != 1 {
		return -1
	}
	return arr[0]
}
