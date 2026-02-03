package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, k, res := drive(reader)
	fmt.Println(k)
	if k <= 0 {
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (r int, l []int, t []int, k int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &r)
	l = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &l[i])
	}
	t = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &t[i])
	}
	k, res = solve(r, l, t)

	return
}

func solve(r int, l []int, t []int) (int, []int) {
	for i := range l {
		if l[i] > t[i] {
			// no way to pass this bridge
			return -1, nil
		}
	}

	var res []int
	var k int
	expired := -1
	var cur int
	for i := range l {
		w := l[i]
		// 剩余的时间
		y := t[i]
		if cur < expired {
			// 魔法效果还在, 速度为1
			v := min(w, expired-cur)
			w -= v
			cur += v
			y -= v
		}

		if w*2 <= y {
			// 不需要魔法，也可以通过
			cur += w * 2
			continue
		}
		// 假设用速度0.5经过了v米
		// w -  v = y - v * 2
		// v = y - w
		v := y - w
		cur += v * 2
		w -= v
		y -= 2 * v
		// 剩余的必须使用速度1通过
		x := (w + r - 1) / r
		k += x
		expired = cur + x*r
		if k <= 1e5 {
			tmp := cur
			for range x {
				res = append(res, tmp)
				tmp += r
			}
		} else {
			res = nil
		}
		cur += w
	}

	return k, res
}
