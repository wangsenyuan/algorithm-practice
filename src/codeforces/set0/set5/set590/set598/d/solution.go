package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
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

func drive(reader *bufio.Reader) []int {
	n, _, q := readThreeNums(reader)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)
	}
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []string, queries [][]int) []int {
	n := len(a)
	m := len(a[0])
	set := NewDSU(n * m)

	for i := range n {
		for j := range m {
			if a[i][j] == '.' {
				if i > 0 && a[i-1][j] == '.' {
					set.Union(i*m+j, (i-1)*m+j)
				}
				if j > 0 && a[i][j-1] == '.' {
					set.Union(i*m+j, i*m+j-1)
				}
			}
		}
	}

	see := make([]int, n*m)

	for i := range n {
		for j := range m {
			if a[i][j] == '*' {
				if i > 0 && a[i-1][j] == '.' {
					see[set.Find((i-1)*m+j)]++
				}
				if j > 0 && a[i][j-1] == '.' {
					see[set.Find(i*m+j-1)]++
				}
				if i+1 < n && a[i+1][j] == '.' {
					see[set.Find((i+1)*m+j)]++
				}
				if j+1 < m && a[i][j+1] == '.' {
					see[set.Find(i*m+j+1)]++
				}
			}
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		r, c := cur[0]-1, cur[1]-1
		ans[i] = see[set.Find(r*m+c)]
	}

	return ans
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
