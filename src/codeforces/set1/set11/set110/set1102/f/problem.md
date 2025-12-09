# Problem F

You are given a matrix $a$, consisting of $n$ rows and $m$ columns. Each cell contains an integer in it.

You can change the order of rows arbitrarily (including leaving the initial order), but you can't change the order of cells in a row. After you pick some order of rows, you traverse the whole matrix the following way: firstly visit all cells of the first column from the top row to the bottom one, then the same for the second column and so on. During the traversal you write down the sequence of the numbers on the cells in the same order you visited them. Let that sequence be $s_1, s_2, \ldots, s_{nm}$.

The traversal is $k$-acceptable if for all $i$ $(1 \leq i \leq nm-1)$ $|s_i - s_{i+1}| \geq k$.

Find the maximum integer $k$ such that there exists some order of rows of matrix $a$ that it produces a $k$-acceptable traversal.

## Input

The first line contains two integers $n$ and $m$ $(1 \leq n \leq 16, 1 \leq m \leq 10^4, 2 \leq nm)$ — the number of rows and the number of columns, respectively.

Each of the next $n$ lines contains $m$ integers $(1 \leq a_{i,j} \leq 10^9)$ — the description of the matrix.

## Output

Print a single integer $k$ — the maximum number such that there exists some order of rows of matrix $a$ that it produces an $k$-acceptable traversal.

## Examples

### Example 1

**Input:**
```
4 2
9 9
10 8
5 3
4 3
```

**Output:**
```
5
```

### Example 2

**Input:**
```
2 4
1 2 3 4
10 3 7 3
```

**Output:**
```
0
```

### Example 3

**Input:**
```
6 1
3
6
2
5
1
4
```

**Output:**
```
3
```

## Note

In the first example you can rearrange rows as following to get the 5-acceptable traversal:


### ideas
1. F[16] 太大了
2. dp[state][i][j] 表示由state表示的行，第一个是i，最后一个是j 时的最优解
3. TSP