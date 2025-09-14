package p3683

func earliestTime(tasks [][]int) int {
	res := 1 << 30
	for _, task := range tasks {
		res = min(res, task[0]+task[1])
	}
	return res
}
