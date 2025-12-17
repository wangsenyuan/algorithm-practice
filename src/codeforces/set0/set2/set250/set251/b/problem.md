# Problem Description

Little Petya likes permutations a lot. Recently his mom has presented him permutation $q_1, q_2, \ldots, q_n$ of length $n$.

A permutation $a$ of length $n$ is a sequence of integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq n$), all integers there are distinct.

There is only one thing Petya likes more than permutations: playing with little Masha. As it turns out, Masha also has a permutation of length $n$. Petya decided to get the same permutation, whatever the cost may be. For that, he devised a game with the following rules:

Before the beginning of the game Petya writes permutation $1, 2, \ldots, n$ on the blackboard. After that Petya makes exactly $k$ moves, which are described below.

During a move Petya tosses a coin. If the coin shows heads, he performs point 1, if the coin shows tails, he performs point 2.

1. Let's assume that the board contains permutation $p_1, p_2, \ldots, p_n$ at the given moment. Then Petya removes the written permutation $p$ from the board and writes another one instead: $p_{q_1}, p_{q_2}, \ldots, p_{q_n}$. In other words, Petya applies permutation $q$ (which he has got from his mother) to permutation $p$.
2. All actions are similar to point 1, except that Petya writes permutation $t$ on the board, such that: $t_{q_i} = p_i$ for all $i$ from $1$ to $n$. In other words, Petya applies a permutation that is inverse to $q$ to permutation $p$.

We know that after the $k$-th move the board contained Masha's permutation $s_1, s_2, \ldots, s_n$. Besides, we know that throughout the game process Masha's permutation never occurred on the board before the $k$-th move. Note that the game has exactly $k$ moves, that is, throughout the game the coin was tossed exactly $k$ times.

Your task is to determine whether the described situation is possible or else state that Petya was mistaken somewhere. See samples and notes to them for a better understanding.

## Input

The first line contains two integers $n$ and $k$ ($1 \leq n, k \leq 100$). The second line contains $n$ space-separated integers $q_1, q_2, \ldots, q_n$ ($1 \leq q_i \leq n$) — the permutation that Petya's got as a present. The third line contains Masha's permutation $s$, in the similar format.

It is guaranteed that the given sequences $q$ and $s$ are correct permutations.

## Output

If the situation that is described in the statement is possible, print "YES" (without the quotes), otherwise print "NO" (without the quotes).

## Examples

### Example 1

**Input:**

```text
4 1
2 3 4 1
1 2 3 4
```

**Output:**

```text
NO
```

### Example 2

**Input:**

```text
4 1
4 3 1 2
3 4 2 1
```

**Output:**

```text
YES
```

### Example 3

**Input:**

```text
4 3
4 3 1 2
3 4 2 1
```

**Output:**

```text
YES
```

### Example 4

**Input:**

```text
4 2
4 3 1 2
2 1 4 3
```

**Output:**

```text
YES
```

### Example 5

**Input:**

```text
4 1
4 3 1 2
2 1 4 3
```

**Output:**

```text
NO
```

## Note

In the first sample Masha's permutation coincides with the permutation that was written on the board before the beginning of the game. Consequently, that violates the condition that Masha's permutation never occurred on the board before $k$ moves were performed.

In the second sample the described situation is possible, in case if after we toss a coin, we get tails.

In the third sample the possible coin tossing sequence is: heads-tails-tails.

In the fourth sample the possible coin tossing sequence is: heads-heads.


### ideas
1. 这个题目像谜语一样
2. q里面应该有一些环，如果是操作1，环顺时针转动，如果是操作2，环逆时针转动
3. 如果一开始 = p相同 => false
4. [4 3 1 2] q[1] = 4, q[2] = 3, q[3] = 1, q[4] = 2
5.  要保证在k次操作后，能够得到s，且在这之前不能获得s
6.  操作1 [p[q[1]], p[q[2]], ... p[q[4]]] => [4, 3, 1, 2] => [2, 1, 4, 3] => [3, 4, 2, 1]
7.  操作2，【1， 2， 3， 4】 =》 t[q[i]] = p[i]
8.  t[q[1]] = p[1] => t[4] = 1
9.  t[q[2]] = p[2] => t[3] = 2
10. t[q[3]] = p[3] => t[1] = 3
11. t[q[4]] = p[4] => t[2] = 4
12. [3, 4, 2, 1]
13. 使用操作1，到s的距离，如果是 <= k, 且和k的diff 是偶数 => true