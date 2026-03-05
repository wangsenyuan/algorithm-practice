package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	buf := make([]string, len(res))
	for i, v := range res {
		buf[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(buf, " "))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) []int {
	s1 := readString(reader)
	s2 := readString(reader)
	s3 := readString(reader)
	return solve(s1, s2, s3)
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b, c int) int {
	return a % mod * (b % mod) % mod * (c % mod) % mod
}

func solve(s1, s2, s3 string) []int {
	arr := []string{s1, s2, s3}

	slices.SortFunc(arr, func(a, b string) int {
		return len(a) - len(b)
	})

	n := len(arr[0])
	totalLen := len(arr[0]) + len(arr[1]) + len(arr[2])

	sam := NewSAM(totalLen + 3)

	for _, c := range arr[0] {
		sam.extend(int(c - 'a'))
	}
	sam.extend(26)
	for _, c := range arr[1] {
		sam.extend(int(c - 'a'))
	}
	sam.extend(27)
	for _, c := range arr[2] {
		sam.extend(int(c - 'a'))
	}
	sam.extend(28)

	sz := sam.size

	// Walk each string through the SAM and count occurrences
	for d := 0; d < 3; d++ {
		cur := 0
		for _, c := range arr[d] {
			x := int(c - 'a')
			cur = sam.next[cur][x]
			sam.cnt[d][cur]++
		}
	}

	// Topological sort by length (counting sort)
	order := make([]int, sz)
	maxLen := totalLen + 3
	buckets := make([]int, maxLen+1)
	for i := 0; i < sz; i++ {
		buckets[sam.length[i]]++
	}
	for i := 1; i <= maxLen; i++ {
		buckets[i] += buckets[i-1]
	}
	for i := sz - 1; i >= 0; i-- {
		l := sam.length[i]
		buckets[l]--
		order[buckets[l]] = i
	}

	// Propagate counts up suffix links (longest first)
	for i := sz - 1; i >= 1; i-- {
		v := order[i]
		p := sam.link[v]
		if p >= 0 {
			for d := 0; d < 3; d++ {
				sam.cnt[d][p] += sam.cnt[d][v]
			}
		}
	}

	// Difference array for answer
	ans := make([]int, n+2)

	for i := 1; i < sz; i++ {
		c0, c1, c2 := sam.cnt[0][i], sam.cnt[1][i], sam.cnt[2][i]
		if c0 == 0 || c1 == 0 || c2 == 0 {
			continue
		}
		tmp := mul(c0, c1, c2)
		maxL := sam.length[i]
		minL := 0
		if sam.link[i] >= 0 {
			minL = sam.length[sam.link[i]]
		}
		lo := minL + 1
		hi := maxL
		if lo > n {
			continue
		}
		if hi > n {
			hi = n
		}
		ans[lo] = add(ans[lo], tmp)
		if hi+1 <= n+1 {
			ans[hi+1] = sub(ans[hi+1], tmp)
		}
	}

	res := make([]int, n)
	cur := 0
	for i := 1; i <= n; i++ {
		cur = add(cur, ans[i])
		res[i-1] = cur
	}

	return res
}

type SAM struct {
	length []int
	link   []int
	next   []map[int]int
	cnt    [3][]int
	last   int
	size   int
}

func NewSAM(maxN int) *SAM {
	cap := 2*maxN + 2
	sam := &SAM{
		length: make([]int, 1, cap),
		link:   make([]int, 1, cap),
		next:   make([]map[int]int, 1, cap),
		last:   0,
		size:   1,
	}
	sam.link[0] = -1
	sam.next[0] = make(map[int]int)
	for d := 0; d < 3; d++ {
		sam.cnt[d] = make([]int, 1, cap)
	}
	return sam
}

func (sam *SAM) addState() int {
	id := sam.size
	sam.size++
	sam.length = append(sam.length, 0)
	sam.link = append(sam.link, -1)
	sam.next = append(sam.next, make(map[int]int))
	for d := 0; d < 3; d++ {
		sam.cnt[d] = append(sam.cnt[d], 0)
	}
	return id
}

func (sam *SAM) cloneState(src int) int {
	id := sam.size
	sam.size++
	sam.length = append(sam.length, 0)
	sam.link = append(sam.link, sam.link[src])
	m := make(map[int]int, len(sam.next[src]))
	for k, v := range sam.next[src] {
		m[k] = v
	}
	sam.next = append(sam.next, m)
	for d := 0; d < 3; d++ {
		sam.cnt[d] = append(sam.cnt[d], 0)
	}
	return id
}

func (sam *SAM) extend(c int) {
	cur := sam.addState()
	sam.length[cur] = sam.length[sam.last] + 1

	p := sam.last
	for p >= 0 {
		if _, ok := sam.next[p][c]; ok {
			break
		}
		sam.next[p][c] = cur
		p = sam.link[p]
	}

	if p < 0 {
		sam.link[cur] = 0
	} else {
		q := sam.next[p][c]
		if sam.length[p]+1 == sam.length[q] {
			sam.link[cur] = q
		} else {
			clone := sam.cloneState(q)
			sam.length[clone] = sam.length[p] + 1
			for p != -1 && sam.next[p][c] == q {
				sam.next[p][c] = clone
				p = sam.link[p]
			}
			sam.link[q] = clone
			sam.link[cur] = clone
		}
	}
	sam.last = cur
}
