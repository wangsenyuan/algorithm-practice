package p3728

type pair struct {
	first  int
	second int
}

func countStableSubarrays(capacity []int) int64 {
	// nums[l] and nums[r] = pref[r-1] - pref[l]
	// pref[l] = pref[r-1] - nums[r]
	freq := make(map[pair]int)
	var sum int
	n := len(capacity)
	var res int
	for i := range n {
		if i > 0 && capacity[i] == 0 && capacity[i-1] == 0 {
			res--
		}
		res += freq[pair{sum - capacity[i], capacity[i]}]
		sum += capacity[i]
		freq[pair{sum, capacity[i]}]++
	}
	return int64(res)
}
