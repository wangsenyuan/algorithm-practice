package p3583

const mod = 1e9 + 7

func specialTriplets(nums []int) int {
	freq := make(map[int]int)
	dp := make(map[int]int)
	var res int
	for _, x := range nums {
		if x%2 == 0 {
			res += dp[x/2]
			res %= mod
		}
		dp[x] += freq[x*2]
		dp[x] %= mod
		freq[x]++
	}
	return res
}
