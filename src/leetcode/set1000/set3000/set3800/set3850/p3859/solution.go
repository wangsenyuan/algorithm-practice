package p3859

import "slices"

func countSubarrays(nums []int, k int, m int) int64 {
	// mx := slices.Max(nums)

	play := func(limit int) int {
		freq := make(map[int]int)
		var cnt int
		var l int
		var ans int
		for _, v := range nums {
			freq[v]++
			if freq[v] == m {
				cnt++
			}

			for len(freq) >= limit && cnt >= k {
				out := nums[l]
				if freq[out] == m {
					cnt--
				}
				freq[out]--
				if freq[out] == 0 {
					delete(freq, out)
				}
				l++
			}
			ans += l
		}
		return ans
	}
	return int64(play(k) - play(k+1))
}

func countSubarrays1(nums []int, k int, m int) int64 {
	n := len(nums)
	mx := slices.Max(nums)

	fq := NewFreqSet(mx)

	pos := make([][]int, mx+1)
	tr := NewTree(n)

	update := func(w int, i int) {
		pos[w] = append(pos[w], i)
		if len(pos[w]) >= m {
			if len(pos[w]) > m {
				j := pos[w][len(pos[w])-m-1]
				tr.Update(j, -1)
			}
			j := pos[w][len(pos[w])-m]
			tr.Update(j, j)
		}
	}

	var res int

	var l int

	play := func(l int, r int) {
		cnt := tr.Count(l, r)
		if cnt == k {
			v := tr.Get(l, r)
			res += v - l + 1
		}
	}

	for i := range n {
		w := nums[i]
		fq.Add(w)
		update(w, i)
		for l < i && fq.cnt > k {
			fq.Rem(nums[l])
			if fq.cnt == k {
				fq.Add(nums[l])
				break
			}
			l++
		}

		if fq.cnt == k {
			play(0, i)
		} else if fq.cnt > k {
			play(l+1, i)
		}
	}

	return int64(res)
}

type FreqSet struct {
	freq []int
	cnt  int
}

func NewFreqSet(x int) *FreqSet {
	return &FreqSet{make([]int, x+1), 0}
}

func (set *FreqSet) Add(x int) {
	set.freq[x]++
	if set.freq[x] == 1 {
		set.cnt++
	}
}

func (set *FreqSet) Rem(x int) {
	set.freq[x]--
	if set.freq[x] == 0 {
		set.cnt--
	}
}

const inf = 1 << 60

type Tree struct {
	val []int
	cnt []int
	sz  int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	cnt := make([]int, 4*n)
	for i := range val {
		val[i] = inf
		cnt[i] = 0
	}
	return &Tree{
		val: val,
		cnt: cnt,
		sz:  n,
	}
}

func (t *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			if v >= 0 {
				t.val[i] = v
				t.cnt[i] = 1
			} else {
				t.val[i] = inf
				t.cnt[i] = 0
			}
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		t.val[i] = min(t.val[i*2+1], t.val[i*2+2])
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) Count(L int, R int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return t.cnt[i]
		}
		mid := (l + r) >> 1
		return f(i*2+1, l, mid) + f(i*2+2, mid+1, r)
	}
	return f(0, 0, t.sz-1)
}

func (t *Tree) Get(L int, R int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if R < l || r < L {
			return inf
		}

		if L <= l && r <= R {
			return t.val[i]
		}
		mid := (l + r) >> 1
		a := f(i*2+1, l, mid)
		b := f(i*2+2, mid+1, r)
		return min(a, b)
	}
	return f(0, 0, t.sz-1)
}
