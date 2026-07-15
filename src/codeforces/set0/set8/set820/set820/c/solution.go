package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var a, b, l, r int
	fmt.Fscan(reader, &a, &b, &l, &r)
	return solve(a, b, l, r)
}

func solve(a, b, l, r int) int {
	ans := 26
	for x := range a {
		ans = min(ans, play(a, b, l, r, x))
	}
	return ans
}

func play(a int, b int, l int, r int, x int) int {
	var buf []byte
	var pref [][]int
	vis := make(map[int]int)

	add := func(v int) {
		cur := make([]int, 26)
		if len(pref) > 0 {
			copy(cur, pref[len(pref)-1])
		}
		buf = append(buf, byte(v+'a'))
		cur[v]++
		pref = append(pref, cur)
	}

	// Place the lex-smallest letters indicated by bitmask v (exactly a bits).
	computer := func(v int) {
		vis[v] = len(buf)
		for i := 0; v > 0; i++ {
			if v&1 == 1 {
				add(i)
			}
			v >>= 1
		}
	}

	player := func() {
		last := x
		if len(buf) > a {
			last = int(buf[len(buf)-1] - 'a')
		}

		for range b {
			add(last)
		}
	}

	computer(1<<a - 1)
	player()

	var cycleLen int
	for len(buf) < r {
		var flag int
		n := len(buf)
		for i := n - a; i < n; i++ {
			flag |= 1 << (buf[i] - 'a')
		}
		var next, cnt int
		for i := range 26 {
			if flag>>i&1 == 0 {
				next |= 1 << i
				cnt++
				if cnt == a {
					break
				}
			}
		}
		if j, ok := vis[next]; ok {
			cycleLen = len(buf) - j
			break
		}
		computer(next)
		player()
	}

	get := func(pos, i int) int {
		if pos < 0 {
			return 0
		}
		return pref[pos][i]
	}

	// Counts of each letter in s[1..pos] (1-indexed length pos).
	countAt := func(pos int) [26]int {
		var res [26]int
		if pos <= 0 {
			return res
		}
		if cycleLen == 0 || pos <= len(buf) {
			copy(res[:], pref[pos-1])
			return res
		}
		pre := len(buf) - cycleLen
		after := pos - pre
		full := after / cycleLen
		rem := after % cycleLen
		for i := range 26 {
			base := get(pre-1, i)
			cycleCnt := get(len(buf)-1, i) - base
			res[i] = base + full*cycleCnt
			if rem > 0 {
				res[i] += get(pre+rem-1, i) - base
			}
		}
		return res
	}

	cr := countAt(r)
	cl := countAt(l - 1)
	var ans int
	for i := range 26 {
		if cr[i]-cl[i] > 0 {
			ans++
		}
	}
	return ans
}
