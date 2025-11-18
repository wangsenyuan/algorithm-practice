# Problem E

Notice that the memory limit is non-standard.

Recently Arthur and Sasha have studied correct bracket sequences. Arthur understood this topic perfectly and become so amazed about correct bracket sequences, so he even got himself a favorite correct bracket sequence of length 2n. Unlike Arthur, Sasha understood the topic very badly, and broke Arthur's favorite correct bracket sequence just to spite him.

All Arthur remembers about his favorite sequence is for each opening parenthesis ('(') the approximate distance to the corresponding closing one (')'). For the i-th opening bracket he remembers the segment [li, ri], containing the distance to the corresponding closing bracket.

Formally speaking, for the i-th opening bracket (in order from left to right) we know that the difference of its position and the position of the corresponding closing bracket belongs to the segment [li, ri].

Help Arthur restore his favorite correct bracket sequence!

## Input

The first line contains integer n (1 ≤ n ≤ 600), the number of opening brackets in Arthur's favorite correct bracket sequence.

Next n lines contain numbers li and ri (1 ≤ li ≤ ri < 2n), representing the segment where lies the distance from the i-th opening bracket and the corresponding closing one.

The descriptions of the segments are given in the order in which the opening brackets occur in Arthur's favorite sequence if we list them from left to right.

## Output

If it is possible to restore the correct bracket sequence by the given data, print any possible choice.

If Arthur got something wrong, and there are no sequences corresponding to the given information, print a single line "IMPOSSIBLE" (without the quotes).

## Examples

### Example 1

**Input:**

```text
4
1 1
1 1
1 1
1 1
```

**Output:**

```text
()()()()
```

### Example 2

**Input:**

```text
3
5 5
3 3
1 1
```

**Output:**

```text
((()))
```

### Example 3

**Input:**

```text
3
5 5
3 3
2 2
```

**Output:**

```text
IMPOSSIBLE
```

### Example 4

**Input:**

```text
3
2 3
1 4
1 4
```

**Output:**

```text
(())()
```


### ideas
1. 左括号的位置也不确定～
2. 完全没有想法～～～
3. 前后两个括号i, j, 有两种关系
4. 要么i包含了j，要么i在j的前面
5. 如果 i....j 是一个完整的一段(....) 那么这段的长度 k = 2 * (j - i + 1)
6. 那么 li + 1 <= k <= ri + 1
7. 还必须保证内部的条件
8. 还有一种情况是 dp[i][l] and dp[l+1][j] 也就是分别满足条件
9. 但是～前后两段必然有限制
10. 