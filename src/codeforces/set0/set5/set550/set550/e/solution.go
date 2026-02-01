package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (a []int, res string) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

type node struct {
	left  *node
	right *node
	value int // value > 0 a leaf node, else a internal node
}

func solve(a []int) string {
	n := len(a)
	if a[n-1] == 1 {
		return ""
	}

	if n == 1 {
		return "0"
	}

	if n == 2 {
		if a[0] == 0 {
			return ""
		}
		return "1->0"
	}
	if a[n-2] == 1 {
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteByte(byte(a[i] + '0'))
			if i+1 < n {
				buf.WriteString("->")
			}
		}
		return buf.String()
	}
	// n >= 3 and a[n-2] == 0
	if a[n-3] == 0 {
		var buf bytes.Buffer
		buf.WriteByte('(')
		// 剩下就是 00000
		if len(a) >= 3 {
			// 那么把倒数两个00，合并在一起，形成一个1，又可以一直过去了
			for i := 0; i < n-3; i++ {
				buf.WriteByte(byte(a[i] + '0'))
				buf.WriteString("->")
			}
			buf.WriteString("(0->0)->0)")
			return buf.String()
		}
	}
	// a[n-3] == 1, (10) 得到的是0
	// 那么就必须找到一个0，让它们变成1
	w := n - 3
	for w >= 0 && a[w] == 1 {
		w--
	}
	if w < 0 {
		return ""
	}
	var buf bytes.Buffer
	// a[w] == 0
	for i := 0; i < w; i++ {
		buf.WriteByte(byte(a[i] + '0'))
		if i+1 < w {
			buf.WriteString("->")
		}
	}
	// (0 -> (1 -> (1 -> .... (1 -> 0))) = 1
	for i := w; i < n-1; i++ {
		if buf.Len() > 0 {
			buf.WriteString("->")
		}
		buf.WriteByte('(')
		buf.WriteByte(byte(a[i] + '0'))
	}

	for i := w; i < n-1; i++ {
		buf.WriteByte(')')
	}

	buf.WriteString("->0")

	return "(" + buf.String() + ")"
}
