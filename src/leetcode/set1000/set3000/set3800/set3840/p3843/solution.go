package p3843

func firstUniqueFreq(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	f2 := make(map[int]int)
	for _, v := range freq {
		f2[v]++
	}

	for _, num := range nums {
		v := freq[num]
		if f2[v] == 1 {
			return num
		}
	}
	return -1
}
