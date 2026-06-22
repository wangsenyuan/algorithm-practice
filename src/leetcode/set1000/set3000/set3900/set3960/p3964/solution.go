package p3964

func minLights(lights []int) int {
	n := len(lights)
	// 如果安装额外的灯泡,它的v相当于1
	// 看长度
	diff := make([]int, n+1)
	for i, v := range lights {
		if v > 0 {
			l := max(0, i-v)
			r := min(n-1, i+v)
			diff[l]++
			diff[r+1]--
		}
	}
	var res int
	var cnt int
	for i := 0; i < n; i++ {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		if diff[i] == 0 {
			cnt++
		} else {
			res += (cnt + 2) / 3
			cnt = 0
		}
	}
	res += (cnt + 2) / 3
	return res
}
