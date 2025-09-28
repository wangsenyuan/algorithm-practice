package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	tc := readNum(reader)
	for range tc {
		readString(reader)
		_, best, ans := drive(reader)
		fmt.Fprintln(writer, best)
		for i := range ans {
			fmt.Fprintln(writer, ans[i][0], ans[i][1])
		}
	}
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

func drive(reader *bufio.Reader) (dishes [][]int, best int, ans [][]int) {
	n := readNum(reader)
	dishes = make([][]int, n)
	for i := range n {
		dishes[i] = readNNums(reader, 3)
	}
	best, ans = solve(dishes)
	return
}

type dish struct {
	id int
	a  int
	b  int
	m  int
}

func (d dish) getUpLimit() int {
	// x >= max(0, m - b)
	// x <= min(a, m)
	// a - x <= a - max(0, m - b)
	// a - x >= a - min(a, m)
	return d.a - max(0, d.m-d.b)
}

func (d dish) getDownLimit() int {
	return d.a - min(d.m, d.a)
}

func solve(dishes [][]int) (best int, ans [][]int) {
	// 根据 a[i] + b[i]- m[i]分组
	n := len(dishes)
	arr := make([]dish, n)
	for i := range n {
		arr[i] = dish{i, dishes[i][0], dishes[i][1], dishes[i][2]}
	}
	slices.SortFunc(arr, func(x, y dish) int {
		return (x.a + x.b - x.m) - (y.a + y.b - y.m)
	})

	ans = make([][]int, n)

	play := func(buf []dish) {
		slices.SortFunc(buf, func(x, y dish) int {
			return x.getDownLimit() - y.getDownLimit()
		})

		upLimit := inf

		var todo []dish

		for _, cur := range buf {
			if upLimit < cur.getDownLimit() {
				best++
				for _, cur := range todo {
					id := cur.id
					ans[id] = []int{cur.a - upLimit, cur.m - (cur.a - upLimit)}
				}
				todo = todo[:0]
				upLimit = inf
			}

			todo = append(todo, cur)
			upLimit = min(upLimit, cur.getUpLimit())
		}
		best++
		for _, cur := range todo {
			id := cur.id
			ans[id] = []int{cur.a - upLimit, cur.m - (cur.a - upLimit)}
		}
	}

	for i := 0; i < n; {
		var todo []dish
		j := i
		for i < n && arr[i].a+arr[i].b-arr[i].m == arr[j].a+arr[j].b-arr[j].m {
			todo = append(todo, arr[i])
			i++
		}
		play(todo)
	}

	return best, ans
}

const inf = 1 << 60
