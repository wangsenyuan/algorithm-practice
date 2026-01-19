package p3811


func alternatingXOR(nums []int, target1 int, target2 int) int {
	// dp[0] = 偶数个t1, 偶数个t2    (2, 2)
	// dp[1] = 奇数个t1, 偶数个t2    (3, 2)
	// dp[2] = 偶数个t1, 奇数个t2    (4, 3)
	// dp[3] = 奇数个t1, 奇数个t2    (5, 5)
	// n := len(nums)
	const mod = 1_000_000_007
	f1 := map[int]int{}
	f2 := map[int]int{0: 1}
	preSum := 0
	for i, x := range nums {
		preSum ^= x
		last1 := f2[preSum^target1] // [0,i] 的最后一段的异或和是 target1 的方案数
		last2 := f1[preSum^target2] // [0,i] 的最后一段的异或和是 target2 的方案数
		if i == len(nums)-1 {
			return (last1 + last2) % mod
		}
		f1[preSum] = (f1[preSum] + last1) % mod
		f2[preSum] = (f2[preSum] + last2) % mod
	}

	return 0
}
