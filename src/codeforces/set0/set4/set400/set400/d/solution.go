package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if len(res) == 0 {
		fmt.Println("No")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("Yes\n")
	for _, cur := range res {
		for _, v := range cur {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
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

func process(reader *bufio.Reader) [][]int {
	n, m, k := readThreeNums(reader)
	c := readNNums(reader, k)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, k, c, edges)

}

const inf = 1 << 60

func solve(n int, k int, c []int, edges [][]int) [][]int {

	ds := NewDSU(n)

	for _, e := range edges {
		u, v, x := e[0], e[1], e[2]
		if x == 0 {
			ds.Union(u-1, v-1)
		}
	}

	sum := make([]int, k+1)
	for i, x := range c {
		sum[i+1] = sum[i] + x
		for j := sum[i] + 1; j < sum[i+1]; j++ {
			if ds.Find(j) != ds.Find(sum[i]) {
				return nil
			}
		}
	}

	d := make([][]int, k)
	for i := range k {
		d[i] = make([]int, k)
		for j := range k {
			d[i][j] = inf
		}
		d[i][i] = 0
	}

	findType := func(s int) int {
		i := sort.Search(k+1, func(i int) bool {
			return sum[i] >= s
		})
		return i - 1
	}

	for _, e := range edges {
		u, v, x := e[0], e[1], e[2]
		i := findType(u)
		j := findType(v)
		d[i][j] = min(d[i][j], x)
		d[j][i] = min(d[j][i], x)
	}

	for u := range k {
		for i := range k {
			for j := range k {
				d[i][j] = min(d[i][j], d[i][u]+d[u][j])
			}
		}
	}

	for i := range k {
		for j := range k {
			if d[i][j] == inf {
				d[i][j] = -1
			}
		}
	}

	return d
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
