# Problem D - Tavas and Malekas

## Description

Today Tavas fell asleep in Malekas' place. While he was sleeping, Malekas did a little process on string `s`. 

Malekas has a favorite string `p`. He determined all positions $x_1 < x_2 < ... < x_k$ where `p` matches `s`. More formally, for each $x_i$ $(1 \leq i \leq k)$, the condition $s_{x_i}s_{x_i + 1}... s_{x_i + |p| - 1} = p$ is fulfilled.

Then Malekas wrote down one of subsequences of $x_1, x_2, ... x_k$ (possibly, he didn't write anything) on a piece of paper. Here a sequence `b` is a subsequence of sequence `a` if and only if we can turn `a` into `b` by removing some of its elements (maybe none of them or all).

After Tavas woke up, Malekas told him everything. He couldn't remember string `s`, but he knew that both `p` and `s` only contain lowercase English letters and also he had the subsequence he had written on that piece of paper.

Tavas wonders, what is the number of possible values of `s`? He asked SaDDas, but he wasn't smart enough to solve this. So, Tavas asked you to calculate this number for him.

Answer can be very large, so Tavas wants you to print the answer modulo $10^9 + 7$.

## Input

- The first line contains two integers `n` and `m`, the length of `s` and the length of the subsequence Malekas wrote down $(1 \leq n \leq 10^6$ and $0 \leq m \leq n - |p| + 1)$.
- The second line contains string `p` $(1 \leq |p| \leq n)$.
- The next line contains `m` space separated integers $y_1, y_2, ..., y_m$, Malekas' subsequence $(1 \leq y_1 < y_2 < ... < y_m \leq n - |p| + 1)$.

## Output

In a single line print the answer modulo $1000000007$.

## Examples

### Example 1

**Input:**
```
6 2
ioi
1 3
```

**Output:**
```
26
```

### Example 2

**Input:**
```
5 2
ioi
1 2
```

**Output:**
```
0
```

## Note

In the first sample test all strings of form "ioioi?" where the question mark replaces arbitrary English letter satisfy.

Here $|x|$ denotes the length of string `x`.

Please note that it's possible that there is no such string (answer is 0).


### ideas
1. s的长度为n，其中有m个位置，yi, yi+1... yi+|p|-1 = p
2. 剩余位置可以随便填
3. 就是检查是否满足条件，如果不满足 => 0
4. 满足的情况下，看还空多少个位置，pow(26, k)
5. 