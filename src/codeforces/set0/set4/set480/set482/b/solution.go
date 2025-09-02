package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, ans := drive(reader)
	if len(ans) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, queries [][]int, ans []int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	queries = make([][]int, m)
	for i := range m {
		var l, r, q int
		fmt.Fscan(reader, &l, &r, &q)
		queries[i] = []int{l, r, q}
	}
	ans = solve(n, queries)
	return
}

const H = 30

const inf = 1 << 60

func solve(n int, queries [][]int) []int {
	res := make([]int, n)

	cnt := make([]int, n+1)
	sum := make([]int, n+1)
	work := func(d int) bool {
		// 处理d位
		// 对于1的，必须全部时1
		clear(cnt)
		clear(sum)
		for _, cur := range queries {
			l, r, q := cur[0]-1, cur[1]-1, cur[2]
			if (q>>d)&1 == 1 {
				cnt[l]++
				cnt[r+1]--
			}
		}
		for i := 0; i < n; i++ {
			if i > 0 {
				sum[i] = sum[i-1]
				cnt[i] += cnt[i-1]
			}
			if cnt[i] > 0 {
				// 这个地方必须是1,
				res[i] |= 1 << d
				sum[i]++
			}
		}
		// 然后就是检查剩余的区间内，是否满足条件, 必须存在一个q[d] = 0的位置
		for _, cur := range queries {
			l, r, q := cur[0]-1, cur[1]-1, cur[2]
			if (q>>d)&1 == 0 {
				tmp := sum[r]
				if l > 0 {
					tmp -= sum[l-1]
				}
				if tmp == (r - l + 1) {
					// 这个区间里面全部是1的话，是不行的
					return false
				}
			}
		}
		return true
	}

	for d := range H {
		if !work(d) {
			return nil
		}
	}

	return res
}
