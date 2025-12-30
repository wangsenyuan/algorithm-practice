package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ask := func(x string) []int {
		fmt.Println(x)
		res := make([]int, 2)
		fmt.Fscan(reader, &res[0], &res[1])
		return res
	}

	solve(ask)
}

func solve(ask func(string) []int) string {
	// Strategy based on Python solution:
	// 1. Query "0123" and "4567" to get group counts
	// 2. Generate all candidate strings based on group combinations
	// 3. Filter by matching bulls from first two queries
	// 4. Use minimax strategy to select optimal guesses
	// 5. Query and filter until solution found
	
	// Query 1: Test first group {0,1,2,3}
	r1 := ask("0123")
	if r1[0] == 4 {
		return "0123"
	}
	b1, c1 := r1[0], r1[1]
	total1 := b1 + c1
	
	// Query 2: Test second group {4,5,6,7}
	r2 := ask("4567")
	if r2[0] == 4 {
		return "4567"
	}
	b2, c2 := r2[0], r2[1]
	total2 := b2 + c2
	
	// Calculate count from {8,9}
	total3 := 4 - total1 - total2
	if total3 < 0 {
		total3 = 0
	} else if total3 > 2 {
		total3 = 2
	}
	
	// Generate all candidate strings
	candidates := generateCandidates(total1, total2, total3, b1, b2)
	if len(candidates) == 0 {
		return ""
	}
	
	// Main query loop with minimax strategy
	for round := 0; round < 5; round++ {
		if len(candidates) == 0 {
			break
		}
		
		if len(candidates) == 1 {
			guess := candidates[0]
			res := ask(guess)
			if res[0] == 4 {
				return guess
			}
			// Filter candidates based on response
			candidates = filterCandidates(candidates, guess, res[0], res[1])
			continue
		}
		
		// Minimax: find guess that minimizes worst-case partition size
		bestGuess := ""
		bestWorst := len(candidates) + 1
		
		for _, g := range candidates {
			partition := make(map[[2]int]int)
			
			for _, s := range candidates {
				bulls, cows := computeBullsAndCowsString(s, g)
				key := [2]int{bulls, cows}
				partition[key]++
			}
			
			maxSize := 0
			for _, count := range partition {
				if count > maxSize {
					maxSize = count
				}
			}
			
			if maxSize < bestWorst {
				bestWorst = maxSize
				bestGuess = g
			}
		}
		
		if bestGuess == "" {
			bestGuess = candidates[0]
		}
		
		res := ask(bestGuess)
		if res[0] == 4 {
			return bestGuess
		}
		
		// Filter candidates based on response
		candidates = filterCandidates(candidates, bestGuess, res[0], res[1])
	}
	
	return ""
}

func generateCandidates(total1, total2, total3, b1, b2 int) []string {
	group1 := []byte{'0', '1', '2', '3'}
	group2 := []byte{'4', '5', '6', '7'}
	group3 := []byte{'8', '9'}
	
	candidates := make([]string, 0)
	
	// Generate combinations from each group
	comb1 := generateCombinations(group1, total1)
	comb2 := generateCombinations(group2, total2)
	comb3 := generateCombinations(group3, total3)
	
	// Combine sets from each group
	for _, set1 := range comb1 {
		for _, set2 := range comb2 {
			for _, set3 := range comb3 {
				// Combine all digits
				digitSet := make(map[byte]bool)
				for _, d := range set1 {
					digitSet[d] = true
				}
				for _, d := range set2 {
					digitSet[d] = true
				}
				for _, d := range set3 {
					digitSet[d] = true
				}
				
				// Must have exactly 4 distinct digits
				if len(digitSet) != 4 {
					continue
				}
				
				// Convert to slice for permutation
				digits := make([]byte, 0, 4)
				for d := range digitSet {
					digits = append(digits, d)
				}
				
				// Generate all permutations
				perms := generateStringPermutations(digits)
				for _, perm := range perms {
					s := string(perm)
					// Filter by matching bulls from first query
					bulls1 := 0
					for i := 0; i < 4; i++ {
						if s[i] == "0123"[i] {
							bulls1++
						}
					}
					if bulls1 != b1 {
						continue
					}
					
					// Filter by matching bulls from second query
					bulls2 := 0
					for i := 0; i < 4; i++ {
						if s[i] == "4567"[i] {
							bulls2++
						}
					}
					if bulls2 != b2 {
						continue
					}
					
					candidates = append(candidates, s)
				}
			}
		}
	}
	
	return candidates
}

func generateCombinations(arr []byte, k int) [][]byte {
	if k == 0 {
		return [][]byte{{}}
	}
	if k > len(arr) {
		return [][]byte{}
	}
	if k == len(arr) {
		result := make([]byte, len(arr))
		copy(result, arr)
		return [][]byte{result}
	}
	
	var result [][]byte
	// Include first element
	withFirst := generateCombinations(arr[1:], k-1)
	for _, comb := range withFirst {
		newComb := make([]byte, 0, k)
		newComb = append(newComb, arr[0])
		newComb = append(newComb, comb...)
		result = append(result, newComb)
	}
	
	// Exclude first element
	withoutFirst := generateCombinations(arr[1:], k)
	result = append(result, withoutFirst...)
	
	return result
}

func generateStringPermutations(arr []byte) [][]byte {
	if len(arr) == 0 {
		return [][]byte{{}}
	}
	
	var result [][]byte
	for i := 0; i < len(arr); i++ {
		rest := make([]byte, 0, len(arr)-1)
		rest = append(rest, arr[:i]...)
		rest = append(rest, arr[i+1:]...)
		
		subPerms := generateStringPermutations(rest)
		for _, subPerm := range subPerms {
			newPerm := make([]byte, 0, len(arr))
			newPerm = append(newPerm, arr[i])
			newPerm = append(newPerm, subPerm...)
			result = append(result, newPerm)
		}
	}
	return result
}

func computeBullsAndCowsString(s, query string) (int, int) {
	bulls := 0
	bullPos := make([]bool, 4)
	
	// Count bulls
	for i := 0; i < 4; i++ {
		if s[i] == query[i] {
			bulls++
			bullPos[i] = true
		}
	}
	
	// Count cows
	sDigits := make(map[byte]int)
	queryDigits := make(map[byte]int)
	
	for i := 0; i < 4; i++ {
		if !bullPos[i] {
			sDigits[s[i]]++
			queryDigits[query[i]]++
		}
	}
	
	cows := 0
	for digit, qCount := range queryDigits {
		if sCount, exists := sDigits[digit]; exists {
			if qCount < sCount {
				cows += qCount
			} else {
				cows += sCount
			}
		}
	}
	
	return bulls, cows
}

func filterCandidates(candidates []string, guess string, bulls, cows int) []string {
	filtered := make([]string, 0)
	
	for _, s := range candidates {
		compBulls, compCows := computeBullsAndCowsString(s, guess)
		if compBulls == bulls && compCows == cows {
			filtered = append(filtered, s)
		}
	}
	
	return filtered
}

