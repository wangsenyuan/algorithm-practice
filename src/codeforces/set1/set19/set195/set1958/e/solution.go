package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	t := readNum(reader)
	for range t {
		n, k := readTwoNums(reader)
		res := solve(n, k)
		if res == nil {
			fmt.Fprintln(writer, -1)
		} else {
			for i, v := range res {
				if i > 0 {
					fmt.Fprint(writer, " ")
				}
				fmt.Fprint(writer, v)
			}
			fmt.Fprintln(writer)
		}
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

// solve finds a permutation of n that takes exactly k operations to reduce to [n].
// Returns nil if impossible.
//
// Key idea: fₖ = 2ᵏ⁻¹ + 1 is the minimum size needed for k operations.
// Construction:
// - For k=1: [1, 2]
// - For k>1: recursively insert elements between existing ones, increasing existing elements
// - If n > fₖ: add extra elements at the beginning that get deleted in first operation
func solve(n int, k int) []int {
	// Calculate fₖ = 2ᵏ⁻¹ + 1 using recurrence fₖ = 2·fₖ₋₁ − 1, f₁ = 2
	fk := 2
	for i := 1; i < k; i++ {
		fk = 2*fk - 1
		if fk > n {
			// Overflow: fₖ > n, so impossible
			return nil
		}
	}

	if n < fk {
		return nil
	}

	// Build base sequence for size fₖ
	seq := buildSequence(fk, k)

	// If n > fₖ, add extra elements at the beginning
	if n > fk {
		extra := n - fk
		// Increase all existing elements by extra
		for i := range seq {
			seq[i] += extra
		}
		// Prepend 1, 2, ..., extra (these will be deleted in first operation)
		newSeq := make([]int, n)
		for i := 0; i < extra; i++ {
			newSeq[i] = i + 1
		}
		copy(newSeq[extra:], seq)
		seq = newSeq
	}

	return seq
}

// buildSequence constructs a sequence of size fₖ that takes exactly k operations.
// Pattern: [1, 2] → [2, 1, 3] → [4, 1, 3, 2, 5] → ...
// "Even position" means 1-indexed (i.e., indices 1, 3, 5, ... in 0-indexed).
// New elements go at even positions (1-indexed), existing elements (increased) at odd positions.
func buildSequence(size int, k int) []int {
	if k == 1 {
		// Base case: [1, 2]
		return []int{1, 2}
	}

	// Recursive: build sequence for k-1
	prevSize := (size + 1) / 2 // f_{k-1} = (fₖ + 1) / 2
	prevSeq := buildSequence(prevSize, k-1)

	// Insert elements at even positions (1-indexed = odd positions 0-indexed)
	insertions := size - prevSize
	result := make([]int, size)
	
	// Fill odd positions (0-indexed) = even positions (1-indexed) with new elements
	newVal := 1
	for i := 1; i < size; i += 2 {
		result[i] = newVal
		newVal++
	}
	
	// Fill even positions (0-indexed) = odd positions (1-indexed) with existing elements (increased)
	j := 0
	for i := 0; i < size; i += 2 {
		result[i] = prevSeq[j] + insertions
		j++
	}

	return result
}
