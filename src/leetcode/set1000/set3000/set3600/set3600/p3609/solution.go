package p3609

func minMoves(sx int, sy int, tx int, ty int) int {
	if sx == tx && sy == ty {
		return 0
	}
	if sx > tx || sy > ty || sx+sy == 0 {
		return -1
	}

	if tx == ty {
		if sx == 0 {
			sx, sy = sy, sx
			tx, ty = ty, tx
		}
		ans := minMoves(sx, sy, tx, 0)
		if ans == -1 {
			return -1
		}
		return ans + 1
	}

	if tx < ty {
		sx, sy = sy, sx
		tx, ty = ty, tx
	}
	if tx >= 2*ty {
		if tx%2 == 1 {
			return -1
		}
		ans := minMoves(sx, sy, tx/2, ty)
		if ans < 0 {
			return -1
		}
		return ans + 1
	}
	ans := minMoves(sx, sy, tx-ty, ty)
	if ans < 0 {
		return -1
	}
	return ans + 1
}
