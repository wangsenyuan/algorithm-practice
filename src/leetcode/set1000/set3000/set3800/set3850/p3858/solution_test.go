package p3858

import "testing"

func TestMinimumOR(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 5}, {2, 4}}, 3},
		{[][]int{{3, 5}, {6, 4}}, 5},
	}

	for _, test := range tests {
		got := minimumOR(test.grid)
		if got != test.want {
			t.Errorf("minimumOR(%v) = %d, want %d", test.grid, got, test.want)
		}
	}
}
