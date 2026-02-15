package p3845

func maxXor(nums []int, k int) int {
	var s1 []int
	var s2 []int

	tr := new(Trie)
	tr.next()

	var best int

	var l int
	n := len(nums)
	a := make([]int, n+1)
	tr.update(0, 1)

	for r, num := range nums {
		a[r+1] = a[r] ^ num
		tr.update(a[r+1], 1)

		for len(s1) > 0 && nums[s1[len(s1)-1]] < num {
			s1 = s1[:len(s1)-1]
		}

		s1 = append(s1, r)

		for len(s2) > 0 && nums[s2[len(s2)-1]] > num {
			s2 = s2[:len(s2)-1]
		}
		s2 = append(s2, r)

		for l < r && len(s1) > 0 && len(s2) > 0 && nums[s1[0]]-nums[s2[0]] > k {
			if s1[0] == l {
				s1 = s1[1:]
			}
			if s2[0] == l {
				s2 = s2[1:]
			}
			tr.update(a[l], -1)
			l++
		}
		tmp := tr.get(a[r+1])

		best = max(best, tmp)
	}

	return best
}

const H = 15

type Trie struct {
	children [][2]int
	cnt      []int
}

func (tr *Trie) next() int {
	tr.children = append(tr.children, [2]int{})
	tr.cnt = append(tr.cnt, 0)
	return len(tr.children) - 1
}

func (tr *Trie) update(num int, v int) {
	var cur int
	for i := H - 1; i >= 0; i-- {
		w := (num >> i) & 1
		if tr.children[cur][w] == 0 {
			tr.children[cur][w] = tr.next()
		}
		cur = tr.children[cur][w]
		tr.cnt[cur] += v
	}
}

func (tr *Trie) get(num int) int {
	var cur int
	var res int
	for i := H - 1; i >= 0; i-- {
		w := (num >> i) & 1
		j := tr.children[cur][w^1]
		if j != 0 && tr.cnt[j] > 0 {
			res ^= 1 << i
			cur = j
		} else {
			cur = tr.children[cur][w]
		}
	}
	return res
}
