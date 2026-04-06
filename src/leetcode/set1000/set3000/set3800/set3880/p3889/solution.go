package p3889

func mirrorFrequency(s string) int {
	freq := make(map[byte]int)
	for i := range s {
		freq[s[i]]++
	}
	var res int
	const letters = "abcdefghijklmnopqrstuvwxyz"
	for i := range len(letters) / 2 {
		res += abs(freq[letters[i]] - freq[letters[len(letters)-i-1]])
	}
	const nums = "0123456789"
	for i := range len(nums) / 2 {
		res += abs(freq[nums[i]] - freq[nums[len(nums)-i-1]])
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
