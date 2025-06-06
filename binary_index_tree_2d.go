package main

// BinaryIndexedTree2D represents a 2D Binary Indexed Tree
type BinaryIndexedTree2D struct {
	tree [][]int
	rows int
	cols int
}

// NewBinaryIndexedTree2D creates a new 2D Binary Indexed Tree with given dimensions
func NewBinaryIndexedTree2D(rows, cols int) *BinaryIndexedTree2D {
	tree := make([][]int, rows+1)
	for i := range tree {
		tree[i] = make([]int, cols+1)
	}
	return &BinaryIndexedTree2D{
		tree: tree,
		rows: rows,
		cols: cols,
	}
}

// Update updates the value at position (row, col) by adding delta
func (bit *BinaryIndexedTree2D) Update(row, col, delta int) {
	for i := row + 1; i <= bit.rows; i += i & -i {
		for j := col + 1; j <= bit.cols; j += j & -j {
			bit.tree[i][j] += delta
		}
	}
}

// Query returns the sum of all elements in the rectangle from (0,0) to (row,col)
func (bit *BinaryIndexedTree2D) Query(row, col int) int {
	sum := 0
	for i := row + 1; i > 0; i -= i & -i {
		for j := col + 1; j > 0; j -= j & -j {
			sum += bit.tree[i][j]
		}
	}
	return sum
}

// QueryRange returns the sum of all elements in the rectangle from (row1,col1) to (row2,col2)
func (bit *BinaryIndexedTree2D) QueryRange(row1, col1, row2, col2 int) int {
	return bit.Query(row2, col2) - bit.Query(row2, col1-1) - bit.Query(row1-1, col2) + bit.Query(row1-1, col1-1)
}

// Set sets the value at position (row, col) to val
func (bit *BinaryIndexedTree2D) Set(row, col, val int) {
	current := bit.QueryRange(row, col, row, col)
	delta := val - current
	bit.Update(row, col, delta)
} 