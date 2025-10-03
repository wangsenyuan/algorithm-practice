# Problem

For the given sequence with `n` different elements find the number of increasing subsequences with `k + 1` elements. It is guaranteed that the answer is not greater than $8 \cdot 10^{18}$.

## Input

First line contain two integer values `n` and `k` ($1 \leq n \leq 10^5$, $0 \leq k \leq 10$) — the length of sequence and the number of elements in increasing subsequences.

Next `n` lines contains one integer $a_i$ ($1 \leq a_i \leq n$) each — elements of sequence. All values $a_i$ are different.

## Output

Print one integer — the answer to the problem.

## Examples

### Example 1

**Input:**
```
5 2
1
2
3
5
4
```

**Output:**
```
7
```
