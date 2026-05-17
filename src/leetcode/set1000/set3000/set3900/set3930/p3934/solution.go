package p3934

import (
	"index/suffixarray"
	"sort"
	"unsafe"
)

func smallestUniqueSubarray1(nums []int) int {
	n := len(nums)
	bases := make([]Hash, n+1)
	bases[0] = NewHash(1)

	keys := make([]Hash, n+1)
	for i, v := range nums {
		bases[i+1] = bases[i].MulInt(10)
		keys[i+1] = keys[i].MulInt(10).AddInt(v)
	}

	check := func(mid int) bool {
		freq := make(map[Hash]int)

		for i := 0; i+mid <= n; i++ {
			w := keys[i+mid]
			v := keys[i].Mul(bases[mid])
			w = w.Sub(v)
			freq[w]++
		}
		for _, v := range freq {
			if v == 1 {
				return true
			}
		}
		return false
	}
	return sort.Search(n, check)
}

var MOD = [...]uint64{1000000007, 1000000009}

type Hash struct {
	h [2]uint64
}

func NewHash(x uint64) Hash {
	h := [2]uint64{x % MOD[0], x % MOD[1]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(x int) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + uint64(x)%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) MulInt(x int) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * uint64(x)) % MOD[i]
	}
	return Hash{h}
}

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	// 把 1 个 int 拆成 4 个 byte，从而可以调用库函数计算 SA
	tmp := make([]byte, 0, n*4)
	for _, x := range nums {
		tmp = append(tmp, byte(x>>24), byte(x>>16), byte(x>>8), byte(x))
	}

	type _tp struct {
		_  []byte
		sa []int32
	}
	_sa := (*_tp)(unsafe.Pointer(suffixarray.New(tmp))).sa

	sa := make([]int32, 0, n)
	for _, p := range _sa {
		if p&3 == 0 { // 是 4 的倍数的 _sa[i] 就对应着 nums 的 sa[i]
			sa = append(sa, p>>2)
		}
	}

	// 计算后缀名次数组
	// 后缀 nums[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 nums 在后缀字典序中的排名，rank[n-1] 即 nums[n-1:] 在字典序中的排名
	// 相当于 sa 的反函数，即 rank[sa[i]] = i
	rank := make([]int, n)
	for i, p := range sa {
		rank[p] = i
	}

	// 计算高度数组（也叫 LCP 数组）
	// height[0] = 0（哨兵）
	// height[i] = LCP(nums[sa[i]:], nums[sa[i-1]:])  (i > 0)
	// 获取 nums[i] 所在位置的高度：height[rank[i]]
	height := make([]int, n)
	h := 0
	// 计算 nums 与 nums[sa[rank[0]-1]:] 的 LCP（记作 LCP0）
	// 计算 nums[1:] 与 nums[sa[rank[1]-1]:] 的 LCP（记作 LCP1）
	// 计算 nums[2:] 与 nums[sa[rank[2]-1]:] 的 LCP
	// ...
	// 计算 nums[n-1:] 与 nums[sa[rank[n-1]-1]:] 的 LCP
	// 从 LCP0 到 LCP1，我们只去掉了 nums[0] 和 nums[sa[rank[0]-1]] 这两个数
	// 所以 LCP1 >= LCP0 - 1
	// 这样就能加快 LCP 的计算了（类似滑动窗口）
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && nums[i+h] == nums[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	ans := n
	for i, h := range height {
		// 对于后缀 nums[sa[i]:]，其长为 uniqueLength 的前缀是唯一的
		uniqueLength := h + 1
		if i < n-1 {
			uniqueLength = max(h, height[i+1]) + 1
		}
		// 注意 uniqueLength 不能超过后缀 nums[sa[i]:] 的长度
		if uniqueLength <= n-int(sa[i]) {
			ans = min(ans, uniqueLength)
		}
	}
	return ans
}
