package p3755

type pair struct {
	first  int
	second int
}

func maxBalancedSubarray(nums []int) int {
	pos := make(map[pair]int)
	pos[pair{0, 0}] = -1

	var best int
	var xor int
	var cnt int
	for i, num := range nums {
		xor ^= num
		if num&1 == 0 {
			cnt++
		} else {
			cnt--
		}
		key := pair{xor, cnt}
		if j, ok := pos[key]; ok {
			best = max(best, i-j)
		} else {
			pos[key] = i
		}
	}
	return best
}
