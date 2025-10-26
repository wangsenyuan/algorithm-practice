package p3729

func numGoodSubarrays(nums []int, k int) int64 {
	freq := make(map[int]int)
	freq[0] = 1
	var res int
	n := len(nums)
	var sum int
	for i := 0; i < n; {
		j := i
		var cur int
		for i < n && nums[i] == nums[j] {
			cur += nums[i]
			res += freq[(sum+cur)%k]
			i++
		}
		for j < i {
			sum += nums[j]
			freq[sum%k]++
			j++
		}
	}
	return int64(res)
}
