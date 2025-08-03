package p3635

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	// 找到最早结束的land项目
	get := func(a []int, b []int, c []int, d []int) int {
		best := inf
		for i, t := range a {
			best = min(best, t+b[i])
		}

		ans := inf

		for i, t := range c {
			if t >= best {
				ans = min(ans, t+d[i])
			} else {
				ans = min(ans, best+d[i])
			}
		}
		return ans
	}

	x := get(landStartTime, landDuration, waterStartTime, waterDuration)
	y := get(waterStartTime, waterDuration, landStartTime, landDuration)

	return min(x, y)
}

const inf = 1 << 60
