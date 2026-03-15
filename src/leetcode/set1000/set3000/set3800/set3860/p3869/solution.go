package p3869

import (
	"slices"
	"sort"
	"strconv"
)

// 数位和不是好数的好数（只有 139 个）
var goodNums []int

func init() {
	for mask := 1; mask < 1<<10; mask++ {
		// 构造严格递减好数
		x := 0
		sum := 0
		for i := 9; i >= 0; i-- {
			if mask>>i&1 > 0 {
				x = x*10 + i
				sum += i
			}
		}
		if !isGood(sum) {
			goodNums = append(goodNums, x)
		}

		// 构造严格递增好数
		if mask&1 > 0 { // 不能包含 0
			continue
		}
		x = 0
		sum = 0
		for i := 1; i < 10; i++ {
			if mask>>i&1 > 0 {
				x = x*10 + i
				sum += i
			}
		}
		if !isGood(sum) {
			goodNums = append(goodNums, x)
		}
	}

	slices.Sort(goodNums) // 方便二分求个数
}

// 判断数位和 s 是否为好数
func isGood(s int) bool {
	if s < 100 { // s 是个位数或者两位数
		return s/10 != s%10 // 十位和个位不相等即为好数
	}
	// s 是三位数，其百位一定是 1
	return 1 < s/10%10 && s/10%10 < s%10 // 只能严格递增
}

func countFancy(l, r int64) int64 {
	// 计算 [l, r] 内的数位和不是好数的好数的个数
	ans := int64(sort.SearchInts(goodNums, int(r+1)) - sort.SearchInts(goodNums, int(l)))

	// 计算 [l, r] 内的数位和是好数的数的个数（这个数是不是好数都可以）
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][]int64, n)
	for i := range memo {
		memo[i] = make([]int64, n*9+1) // 数位和最大 9n
	}

	var dfs func(int, int, bool, bool) int64
	dfs = func(i, digitSum int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			if isGood(digitSum) {
				return 1 // 合法
			}
			return 0 // 不合法
		}
		if !limitLow && !limitHigh {
			dv := &memo[i][digitSum]
			if *dv > 0 {
				return *dv - 1
			}
			defer func() { *dv = res + 1 }() // 这样写无需初始化 memo 为 -1
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		// 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, true, false)
			d = 1 // 下面填数字，从 1 开始填
		}

		for ; d <= hi; d++ {
			res += dfs(i+1, digitSum+d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}

	return ans + dfs(0, 0, true, true)
}
