package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	for _, cur := range res {
		for _, v := range cur {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
		buf.WriteByte('\n')
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

func process(reader *bufio.Reader) (n int, likes [][]int, res [][]int) {
	n, m := readTwoNums(reader)
	likes = make([][]int, m)
	for i := range m {
		likes[i] = readNNums(reader, 2)
	}
	res = solve(n, likes)
	return
}

func solve(n int, likes [][]int) [][]int {
	set := NewDSU(n)

	for _, cur := range likes {
		u, v := cur[0], cur[1]
		set.Union(u-1, v-1)
	}
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		p := set.Find(i)
		arr[p] = append(arr[p], i+1)
		if len(arr[p]) > 3 {
			// 不能超过3个人
			return nil
		}
	}

	parts := make([][]int, 4)
	for i, cur := range arr {
		parts[len(cur)] = append(parts[len(cur)], i)
	}

	if len(parts[2]) > len(parts[1]) {
		return nil
	}

	var res [][]int

	for _, i := range parts[3] {
		res = append(res, arr[i])
	}

	for i := 0; i < len(parts[2]); i++ {
		x := arr[parts[1][i]]
		y := arr[parts[2][i]]
		res = append(res, append(y, x...))
	}

	for i := len(parts[2]); i < len(parts[1]); i += 3 {
		tmp := []int{
			arr[parts[1][i]][0],
			arr[parts[1][i+1]][0],
			arr[parts[1][i+2]][0],
		}
		res = append(res, tmp)
	}

	return res
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
