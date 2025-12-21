package p3785

func minSwaps(nums []int, forbidden []int) int {
	n := len(nums)

	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	f2 := make(map[int]int)
	var k int
	var mx int
	for i, v := range forbidden {
		freq[v]++
		if freq[v] > n {
			return -1
		}

		if v == nums[i] {
			f2[v]++
			k++
			mx = max(mx, f2[v])
		}
	}
	return max(mx, (k+1)/2)
}
