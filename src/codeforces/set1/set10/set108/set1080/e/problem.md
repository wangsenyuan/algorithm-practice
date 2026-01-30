Sonya had a birthday recently. She was presented with a matrix of size $n \times m$ consisting of lowercase Latin letters. We assume that the rows are numbered by integers from $1$ to $n$ from bottom to top, and the columns are numbered from $1$ to $m$ from left to right.

Let's call a submatrix $(i_1, j_1, i_2, j_2)$ ($1 \le i_1 \le i_2 \le n$; $1 \le j_1 \le j_2 \le m$) the elements $a_{ij}$ of this matrix such that $i_1 \le i \le i_2$ and $j_1 \le j \le j_2$. Sonya states that a submatrix is **beautiful** if we can independently reorder the characters in each row (not in column) so that all rows and columns of this submatrix form palindromes.

Let's recall that a string is called a **palindrome** if it reads the same from left to right and from right to left. For example, strings $aba$, $caac$, $a$ are palindromes while strings $abca$, $acbba$, $ab$ are not.

Help Sonya to find the number of beautiful submatrices. Submatrices are different if there is an element that belongs to only one submatrix.

## Input

The first line contains two integers $n$ and $m$ ($1 \le n, m \le 250$) — the matrix dimensions.

Each of the next $n$ lines contains $m$ lowercase Latin letters.

## Output

Print one integer — the number of beautiful submatrices.

## Examples

### Input
```
1 3
aba
```

### Output
```
4
```

### Input
```
2 3
aca
aac
```

### Output
```
11
```

### Input
```
3 5
accac
aaaba
cccaa
```

### Output
```
43
```

## Note

In the first example, the following submatrices are beautiful: $((1,1),(1,1))$; $((1,2),(1,2))$; $((1,3),(1,3))$; $((1,1),(1,3))$.

In the second example, all submatrices that consist of one element and the following are beautiful: $((1,1),(2,1))$; $((1,1),(1,3))$; $((2,1),(2,3))$; $((1,1),(2,3))$; $((2,1),(2,2))$.

In the third example, some of the beautiful submatrices are: $((1,1),(1,5))$; $((1,2),(3,4))$; $((1,1),(3,5))$.

The submatrix $((1,1),(3,5))$ is beautiful since it can be reordered as:

```
accca
aabaa
accca
```

In such a matrix every row and every column form palindromes.


### ideas
1. 在一个子矩阵里面，对每行进行重拍后，保证每行、每列（子矩阵）都是回文，的子矩阵的数量
2. 假设w * h, w是偶数，那么每行每个字符都必须出现偶数次，
3. 考虑列的，那么第一行、最后一行，必须有相同的字符组成，字频要一致
4. 固定一行（或者相邻两行）做为中心，依次移动检查。
5. 固定两列，c1, c2， 然后根据字频对所有行进行分组（字频一样的，就时能够匹配的）
6. 然后固定一行（奇数行的情况）往两边扩张
7. 