package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, res := drive(reader)
		if len(res) == 0 {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, len(res))
			for _, v := range res {
				fmt.Fprintln(writer, v[0], v[1])
			}
		}
	}
}

func drive(reader *bufio.Reader) (h []int, m int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n, &m)
	h = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	res = solve(h, m)
	return
}

type elf struct {
	id     int
	health int
}

func solve(h []int, m int) [][]int {
	n := len(h)
	if m > n/2 {
		return nil
	}

	m2 := 2 * m
	// 使用策略1，剩余m2个人

	arr := make([]elf, n)
	for i := range n {
		arr[i] = elf{id: i, health: h[i]}
	}
	slices.SortFunc(arr, func(x elf, y elf) int {
		return x.health - y.health
	})
	if m == 0 {
		return solve0(arr)
	}
	var res [][]int
	for i := 0; i+1 <= n-m2; i++ {
		// x是还没有攻击过别人的
		res = append(res, []int{arr[i+1].id + 1, arr[i].id + 1})
	}
	// 剩下的m2个是偶数
	for i := n - m2; i < n; i += 2 {
		res = append(res, []int{arr[i+1].id + 1, arr[i].id + 1})
	}
	return res
}

func solve0(arr []elf) [][]int {
	n := len(arr)
	var sum int
	w := n - 2
	for w >= 0 && sum < arr[n-1].health {
		sum += arr[w].health
		w--
	}
	if sum < arr[n-1].health {
		return nil
	}
	var res [][]int
	// 虽然第w+1个人的health少了，但是它的伤害还在
	for i := 0; i <= w; i++ {
		res = append(res, []int{arr[i].id + 1, arr[i+1].id + 1})
	}
	for i := w + 1; i < n-1; i++ {
		res = append(res, []int{arr[i].id + 1, arr[n-1].id + 1})
	}
	return res
}
