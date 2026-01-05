package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, a := drive(reader)
	if len(a) == 0 {
		fmt.Println("-1")
		return
	}
	s := fmt.Sprintf("%v", a)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (l []int, a []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	l = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &l[i])
	}
	return l, solve(n, l)
}

func solve(n int, l []int) []int {
	a := make([]int, n+1)
	m := len(l)

	marked := make([]bool, n+1)

	for i := range m - 1 {
		v := l[i]
		w := l[i+1] - v
		if w <= 0 {
			w += n
		}
		if a[v] == 0 {
			if marked[w] {
				return nil
			}
			marked[w] = true
			a[v] = w
		} else if a[v] != w {
			return nil
		}
	}

	var arr []int
	for i := 1; i <= n; i++ {
		if !marked[i] {
			arr = append(arr, i)
		}
	}

	for i := 1; i <= n; i++ {
		if a[i] == 0 {
			a[i] = arr[0]
			arr = arr[1:]
		}
	}

	return a[1:]
}
