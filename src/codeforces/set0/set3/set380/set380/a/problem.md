# Sereja and Number Sequences

Sereja loves number sequences very much. That's why he decided to make himself a new one following a certain algorithm.

## Problem Description

Sereja takes a blank piece of paper. Then he starts writing out the sequence in m stages. Each time he either:

1. **Adds a new number** to the end of the sequence, or
2. **Takes the first l elements** of the current sequence and adds them c times to the end

More formally, if we represent the current sequence as $a_1, a_2, ..., a_n$, then after applying the described operation, the sequence transforms into:

$a_1, a_2, ..., a_n, [a_1, a_2, ..., a_l]$ (the block in square brackets must be repeated c times)

A day has passed and Sereja has completed the sequence. He wonders what are the values of some of its elements. Help Sereja.

## Input

- **Line 1**: Integer $m$ $(1 \leq m \leq 10^5)$ — the number of stages to build a sequence.

- **Lines 2 to m+1**: Description of the stages in order:
  - **Type 1**: Add one number to the end
    - Format: `1 x_i`
    - $x_i$ $(1 \leq x_i \leq 10^5)$ — the number to add
  - **Type 2**: Copy a prefix of length $l_i$ to the end $c_i$ times
    - Format: `2 l_i c_i`
    - $l_i$ $(1 \leq l_i \leq 10^5)$ — length of the prefix
    - $c_i$ $(1 \leq c_i \leq 10^4)$ — number of copyings
    - It is guaranteed that $l_i$ is never larger than the current length of the sequence

- **Line m+2**: Integer $n$ $(1 \leq n \leq 10^5)$ — the number of elements Sereja is interested in

- **Line m+3**: $n$ numbers representing the positions of elements in the final sequence that Sereja is interested in
  - Numbers are given in strictly increasing order
  - All numbers are strictly larger than zero and do not exceed the length of the resulting sequence
  - Elements are numbered starting from 1

**Note**: Do not use the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use `cin`, `cout` streams or the `%I64d` specifier.

## Output

Print the elements that Sereja is interested in, in the order in which their numbers occur in the input.

## Examples

### Example 1

**Input:**
```
6
1 1
1 2
2 2 1
1 3
2 5 2
1 4
16
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
```

**Output:**
```
1 2 1 2 3 1 2 1 2 3 1 2 1 2 3 4
```

**Explanation:**
- Stage 1: Add 1 → [1]
- Stage 2: Add 2 → [1, 2]
- Stage 3: Copy first 2 elements 1 time → [1, 2, 1, 2]
- Stage 4: Add 3 → [1, 2, 1, 2, 3]
- Stage 5: Copy first 5 elements 2 times → [1, 2, 1, 2, 3, 1, 2, 1, 2, 3, 1, 2, 1, 2, 3]
- Stage 6: Add 4 → [1, 2, 1, 2, 3, 1, 2, 1, 2, 3, 1, 2, 1, 2, 3, 4]


## ideas
1. 每个查询的结果，只可能是Add的数
2. f(i) = f((i - m) % l) (m是操作2前的长度)
3. 这样子，只需要生成前1e4个就可以了