package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(a [][]int) []int {
	n := len(a)
	m := len(a[0])
	// 所有的都必须连接在一起
	set := NewDSU(n * m)

	first := []int{-1, -1}

	var odd int
	var cnt int
	for i := range n {
		for j := range m {
			if a[i][j] == 1 {
				cnt++
				if first[0] == -1 {
					first[0] = i
					first[1] = j
				}
				var deg int
				for k := range 4 {
					r, c := i+dd[k], j+dd[k+1]
					if r >= 0 && r < n && c >= 0 && c < m && a[r][c] == 1 {
						deg++
						set.Union(i*m+j, r*m+c)
					}
				}
				if deg%2 == 1 {
					odd++
				}
			}
		}
	}
	r0, c0 := first[0], first[1]

	if cnt == 1 || odd != 0 && odd != 2 || cnt != set.cnt[set.Find(r0*m+c0)] {
		return nil
	}

	var w int

	for i := range n {
		for j := 0; j < m; {
			if a[i][j] == 0 {
				j++
				continue
			}
			l := j

			for j < m && a[i][j] == 1 {
				j++
			}

			if j-l > 1 {
				w = gcd(w, j-l-1)
			}
		}
	}

	for j := range m {
		for i := 0; i < n; {
			if a[i][j] == 0 {
				i++
				continue
			}
			l := i
			for i < n && a[i][j] == 1 {
				i++
			}
			if i-l > 1 {
				w = gcd(w, i-l-1)
			}
		}
	}
	var res []int
	for k := 2; k <= w; k++ {
		if w%k == 0 {
			res = append(res, k)
		}
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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
