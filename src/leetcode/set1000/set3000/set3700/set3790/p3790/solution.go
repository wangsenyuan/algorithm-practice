package p3790


func minAllOneMultiple(k int) int {
 	marked := make([]bool, k+1)
	 marked[0] = true   

	 base := 1
	 val := base % k
	 n := 1
	 for val != 0 && !marked[val]{
		n++
		marked[val] = true
		base = base * 10 % k
		val = (val + base) % k
	 }

	 if val == 0 {
		return n
	 }
	 
	 return -1
}