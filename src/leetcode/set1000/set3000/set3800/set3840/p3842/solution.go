package p3842

func toggleLightBulbs(bulbs []int) []int {
	status := make([]int, 101)
	for _, i := range bulbs {
		status[i] ^= 1
	}
	var res []int
	for i := 1; i <= 100; i++ {
		if status[i] == 1 {
			res = append(res, i)
		}
	}
	return res
}
