package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n, q := readTwoNums(reader)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {
	bs := int(math.Sqrt(float64(n))) + 1
	where := make([]int, n+1)
	a := make([]int, n+1)
	ps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = i
		where[a[i]] = i
		ps[i] = ps[i-1]
		if i%bs == 1 {
			ps[i]++
		}
	}

	sumb := make([][]int, bs+10)
	for i := range sumb {
		sumb[i] = make([]int, bs+10)
	}

	for i := 1; i <= n; i++ {
		sumb[ps[i]][ps[a[i]]]++
	}

	for i := 1; i <= ps[n]; i++ {
		for j := 1; j <= ps[n]; j++ {
			sumb[i][j] += sumb[i-1][j]
		}
	}

	change := func(pos int, val int, dt int) {
		x := ps[val]
		for i := ps[pos]; i <= ps[n]; i++ {
			sumb[i][x] += dt
		}
	}

	var tot int
	ask := func(x int, y int) {
		// x <= y
		mn, mx := min(a[x], a[y]), max(a[x], a[y])
		var xs int
		if a[x] < a[y] {
			xs = 2
		} else {
			xs = -2
		}
		var num int
		if ps[mn]+1 >= ps[mx] {
			for i := mn; i <= mx; i++ {
				if x < where[i] && where[i] < y {
					num++
				}
			}
			tot += xs * num
			return
		}
		if ps[x]+1 >= ps[y] {
			for i := x; i <= y; i++ {
				if mn < a[i] && a[i] < mx {
					num++
				}
			}
			tot += xs * num
			return
		}
		for i := mn; ps[i] == ps[mn]; i++ {
			if x < where[i] && where[i] < y {
				num++
			}
		}
		for i := mx; ps[i] == ps[mx]; i-- {
			if x < where[i] && where[i] < y {
				num++
			}
		}
		for i := x; ps[i] == ps[x]; i++ {
			if ps[mn] < ps[a[i]] && ps[a[i]] < ps[mx] {
				num++
			}
		}
		for i := y; ps[i] == ps[y]; i-- {
			if ps[mn] < ps[a[i]] && ps[a[i]] < ps[mx] {
				num++
			}
		}

		for i := ps[mn] + 1; i < ps[mx]; i++ {
			num += sumb[ps[y]-1][i] - sumb[ps[x]][i]
		}

		tot += xs * num
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		x, y := cur[0], cur[1]
		if x > y {
			x, y = y, x
		}
		ask(x, y)
		change(x, a[x], -1)
		change(y, a[y], -1)
		if a[x] > a[y] {
			tot--
		} else if a[x] < a[y] {
			tot++
		}
		a[x], a[y] = a[y], a[x]
		where[a[x]] = x
		where[a[y]] = y
		change(x, a[x], 1)
		change(y, a[y], 1)
		ans[i] = tot
	}
	return ans
}
