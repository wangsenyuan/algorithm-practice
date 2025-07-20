package p3618

func splitArray(nums []int) int64 {
	n := len(nums)
	var primes []int
	lpf := make([]int, n)
	for i := 2; i < n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= n {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	var sum int
	for i := 0; i < n; i++ {
		if i > 0 && lpf[i] == i {
			sum += nums[i]
		} else {
			sum -= nums[i]
		}
	}
	if sum < 0 {
		sum *= -1
	}
	return int64(sum)
}
