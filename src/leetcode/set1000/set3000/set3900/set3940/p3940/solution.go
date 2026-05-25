package p3940

func limitOccurrences(nums []int, k int) []int {
	freq := make([]int, 101)
	var res []int
	for _, num := range nums {
		if freq[num] < k {
			res = append(res, num)
			freq[num]++
		}
	}
	return res
}
