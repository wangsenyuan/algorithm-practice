# Problem

You are given a matrix $A$ of size $n \times n$ where $A_{i,j} = j$ for all $1 \le i, j \le n$.

In one operation, you may choose a row and **reverse any subarray** of that row.

Find a sequence of **at most $2n$** operations so that every column of the final matrix is a **permutation** of length $n$ (i.e. contains each of $1, 2, \ldots, n$ exactly once).

It can be proven that such a construction always exists. If multiple solutions exist, output any.

**Subarray:** An array $a$ is a subarray of array $b$ if $a$ can be obtained from $b$ by deleting zero or more elements from the beginning and zero or more from the end.

**Permutation of length $n$:** An array of $n$ distinct integers from $1$ to $n$ in some order. E.g. $[2,3,1,5,4]$ is a permutation; $[1,2,2]$ is not (repeated element); $[1,3,4]$ is not a permutation of length $3$ (contains $4$).

## Input

- The first line contains an integer $t$ ($1 \le t \le 100$) — the number of test cases.
- For each test case, a single line with an integer $n$ ($3 \le n \le 5000$) — the number of rows and columns.

The sum of $n$ over all test cases does not exceed $5000$.

## Output

For each test case:

- First line: an integer $k$ ($0 \le k \le 2n$) — the number of operations.
- Next $k$ lines: each operation in the format `i l r` ($1 \le l \le r \le n$, $1 \le i \le n$), meaning reverse the subarray $A_{i,l}, A_{i,l+1}, \ldots, A_{i,r}$ in row $i$.

## Examples

### Example 1

**Input:**

```
2
3
4
```

**Output:**

```
4
2 1 3
2 2 3
3 1 2
3 2 3
5
2 1 4
3 1 3
3 2 4
4 3 4
4 1 2
```

(First test case: $n=3$, 4 operations. Second test case: $n=4$, 5 operations.)


### ideas
1. 考虑n = 4
  ```
          1234
  3214 => 3412 
  1432 => 2341
  4321 => 4123
  ```
2. 其实是旋转，就是把1往右移动一格，shift(k) = flip(:k), flip(k:), flip(:n)
3. 当k = 1 或者 n - k = 1 的时候，不flip（但是这样子的结果是 差不多 3 * n)
4. flip(:n)其实都不需要做，大家都不flip， 相当于大家都flip