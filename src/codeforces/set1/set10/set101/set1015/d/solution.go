package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k, s int
	fmt.Fscanf(reader, "%d %d %d", &n, &k, &s)
	res := solve(n, k, s)
	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func solve(n int, k int, s int) []int {
	if k > s || k*(n-1) < s {
		return nil
	}
	// k <= s
	// k * (n - 1) >= s
	pos := 1
	var res []int
	// 假设左右横跳了x次后，后面就短途移动
	// s1 = s - x * (n - 1)
	// k1 = k - x
	// k1 <= s1 and k1 * (n - 1) >= s1
	// 这个可以二分
	x := sort.Search(k, func(x int) bool {
		s1 := s - x*(n-1)
		k1 := k - x
		return k1 > s1 || k1*(n-1) < s1
	})
	x--
	k -= x
	for x > 0 {
		pos = n + 1 - pos
		res = append(res, pos)
		x--
		s -= n - 1
	}

	for k > 0 {
		d1 := pos - 1
		d2 := n - pos

		d := max(d1, d2)

		if s <= d {
			// 往一个方向运动就可以了
			for k > 1 {
				if d == d1 {
					pos--
				} else {
					pos++
				}
				res = append(res, pos)
				s--
				k--
			}
			// 最后移动s的距离
			if pos-1 >= s {
				res = append(res, pos-s)
			} else {
				res = append(res, pos+s)
			}

			break
		}

		// 假设移动到max(d1, d2)处，使用了x次
		// k - x <= s - d, 且 (k - x) * (n - 1) >= s - d
		// x >= k + d - s
		// (k - x) * (n - 1) >= s - d
		// (k - x) >= (s - d + n - 2) / (n - 1)
		// x <= k - (s - d + n - 2) / (n - 1)
		x := min(k-(s-d+n-2)/(n-1), d)
		k -= x

		for x > 1 {
			if d == d1 {
				pos--
			} else {
				pos++
			}
			res = append(res, pos)
			x--
		}
		if d == d1 {
			pos = 1
		} else {
			pos = n
		}
		res = append(res, pos)
		s -= d
	}

	return res
}
