package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd, _ := os.Open("input.txt")
	defer rd.Close()
	wr, _ := os.Create("output.txt")
	defer wr.Close()

	reader := bufio.NewReader(rd)
	res := drive(reader)

	writer := bufio.NewWriter(wr)
	defer writer.Flush()

	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var m, k int
	fmt.Fscan(reader, &m, &k)
	a := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &a[i])
	}
	var n int
	fmt.Fscan(reader, &n)
	casts := make([][]int, n)
	for i := range n {
		var s string
		fmt.Fscan(reader, &s)
		var d int
		fmt.Fscan(reader, &d)
		casts[i] = make([]int, d)
		for j := range d {
			fmt.Fscan(reader, &casts[i][j])
		}
	}

	return solve(m, a, casts)
}

func solve(m int, favouriates []int, casts [][]int) []int {
	n := len(casts)
	if n == 1 {
		return []int{0}
	}

	marked := make([]bool, m+1)
	for _, v := range favouriates {
		marked[v] = true
	}

	// 如果一个电影,+ 它不确定的那些人,都是不喜欢的演员, 仍然比其他电影(不确定的都是喜欢的演员), 还要多
	// 那么它肯定是最爱电影
	// 如果应该电影,+... 还要少, 那么它就是最不爱电影

	vis := make([]bool, m+1)

	play := func(cur []int) []int {
		clear(vis)
		// 那些是0的,全部算作不喜欢的演员(但有可能不够)
		var zero int
		var stars int
		for _, v := range cur {
			if v == 0 {
				zero++
			} else if marked[v] {
				stars++
			}
			vis[v] = true
		}
		// 有zero个人可以分配, 尽量的分配不喜欢的演员
		var c1 int
		for i := 1; i <= m && c1 < zero; i++ {
			if !vis[i] && !marked[i] {
				c1++
			}
		}
		s1 := stars + zero - c1
		// 这些zero个人都是star
		var c2 int
		for i := 1; i <= m && c2 < zero; i++ {
			if !vis[i] && marked[i] {
				c2++
			}
		}
		s2 := stars + c2

		return []int{s1, s2}
	}

	type player struct {
		id int
		s1 int
		s2 int
	}

	arr := make([]player, n)

	var best []player
	var best2 []player

	for i, cur := range casts {
		tmp := play(cur)
		arr[i] = player{i, tmp[0], tmp[1]}
		p := arr[i]
		for i := range best {
			if p.s2 >= best[i].s2 {
				p, best[i] = best[i], p
			}
		}
		if len(best) < 2 {
			best = append(best, p)
		}

		p = arr[i]
		for i := range best2 {
			if p.s1 >= best2[i].s1 {
				p, best2[i] = best2[i], p
			}
		}
		if len(best2) < 2 {
			best2 = append(best2, p)
		}
	}

	ans := make([]int, n)

	for i, cur := range arr {
		ans[i] = 2
		if cur.id == best[0].id {
			if cur.s1 >= best[1].s2 {
				// 这个是favariable
				ans[i] = 0
			}
		} else if cur.s1 >= best[0].s2 {
			ans[i] = 0
		}

		if ans[i] == 0 {
			continue
		}

		if cur.id == best2[0].id {
			if cur.s2 < best2[1].s1 {
				ans[i] = 1
			}
		} else if cur.s2 < best2[0].s1 {
			ans[i] = 1
		}
	}

	return ans
}
