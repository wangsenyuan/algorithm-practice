package p3699

func zigZagArrays(n int, l int, r int) int {
	// 如果不考虑n太大的情况， dp[i][x][0] 表示到i为止，最后一个数是x，且x的前面的数比x大
	d := r - l + 1
	// dp[x][0] 表示当前是x，且它是最小数时的最大值
	d2 := d * 2
	mat := NewMatrix(d2, d2)

	for i := 0; i < d; i++ {
		for j := range i {
			mat[i*2+1][j*2+0] = 1
		}
		for j := i + 1; j < d; j++ {
			mat[i*2+0][j*2+1] = 1
		}
	}

	res := pow(mat, n-2)

	first := NewMatrix(1, d2)
	for i := range d {
		first[0][i*2+0] = d - i - 1
		first[0][i*2+1] = i
	}

	res = first.Mul(res)

	var ans int
	for _, v := range res[0] {
		ans = modAdd(ans, v)
	}

	return ans
}

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return a * b % MOD
}

type Mat [][]int

func NewMatrix(m int, n int) Mat {
	v := make([][]int, m)
	for i := 0; i < m; i++ {
		v[i] = make([]int, n)
	}
	return Mat(v)
}

func (this Mat) Mul(that Mat) Mat {
	m := len(this)
	w := len(that)
	n := len(that[0])
	v := make([][]int, m)
	for i := 0; i < m; i++ {
		v[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var sum int
			for k := 0; k < w; k++ {
				sum = modAdd(sum, modMul(this[i][k], that[k][j]))
			}
			v[i][j] = sum
		}
	}

	return Mat(v)
}

func identity(n int) Mat {
	v := make([][]int, n)
	for i := 0; i < n; i++ {
		v[i] = make([]int, n)
		v[i][i] = 1
	}
	return Mat(v)
}

func pow(m Mat, n int) Mat {
	res := identity(len(m))
	for n > 0 {
		if n&1 == 1 {
			res = res.Mul(m)
		}
		m = m.Mul(m)
		n >>= 1
	}
	return res
}
