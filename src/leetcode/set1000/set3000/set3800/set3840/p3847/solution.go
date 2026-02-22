package p3847

func scoreDifference(nums []int) int {
	score := []int{0, 0}
	var player int
	for i, num := range nums {
		if num%2 == 1 {
			player ^= 1
		}
		if (i+1)%6 == 0 {
			player ^= 1
		}
		score[player] += num
	}
	return score[0] - score[1]
}
