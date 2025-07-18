# Problem D: Mishka and Interesting Sum

## Problem Description

Little Mishka enjoys programming. Since her birthday has just passed, her friends decided to present her with an array of non-negative integers $a_1, a_2, \ldots, a_n$ of $n$ elements!

Mishka loved the array and she instantly decided to determine its beauty value, but she is too little and can't process large arrays. Right because of that she invited you to visit her and asked you to process $m$ queries.

Each query is processed in the following way:

1. Two integers $l$ and $r$ ($1 \leq l \leq r \leq n$) are specified — bounds of query segment.
2. Integers, presented in array segment $[l, r]$ (in sequence of integers $a_l, a_{l+1}, \ldots, a_r$) an **even number of times**, are written down.
3. XOR-sum of written down integers is calculated, and this value is the answer for a query. Formally, if integers written down in point 2 are $x_1, x_2, \ldots, x_k$, then Mishka wants to know the value $x_1 \oplus x_2 \oplus \ldots \oplus x_k$, where $\oplus$ is the operator of exclusive bitwise OR.

Since only the little bears know the definition of array beauty, all you are to do is to answer each of the queries presented.

## Input

- The first line of the input contains a single integer $n$ ($1 \leq n \leq 1,000,000$) — the number of elements in the array.
- The second line of the input contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^9$) — array elements.
- The third line of the input contains a single integer $m$ ($1 \leq m \leq 1,000,000$) — the number of queries.
- Each of the next $m$ lines describes the corresponding query by a pair of integers $l$ and $r$ ($1 \leq l \leq r \leq n$) — the bounds of the query segment.

## Output

Print $m$ non-negative integers — the answers for the queries in the order they appear in the input.

## Examples

### Example 1

**Input:**
```
3
3 7 8
1
1 3
```

**Output:**
```
0
```

### Example 2

**Input:**
```
7
1 2 1 3 3 2 3
5
4 7
4 5
1 3
1 7
1 5
```

**Output:**
```
0
3
1
3
2
```

## Explanation

In the second sample:

- **Query 1 (4-7)**: There are no integers in the segment presented an even number of times — the answer is 0.
- **Query 2 (4-5)**: Only integer 3 is presented an even number of times — the answer is 3.
- **Query 3 (1-3)**: Only integer 1 is written down — the answer is 1.
- **Query 4 (1-7)**: All array elements are considered. Only 1 and 2 are presented an even number of times. The answer is $1 \oplus 2 = 3$.
- **Query 5 (1-5)**: 1 and 3 are written down. The answer is $1 \oplus 3 = 2$.

### ideas
1. sum[r] ^ sum[l-1] = 这个区间内，出现了奇数次的数
2. 如果知道 sum2[r] ^ sum2[l-1] 知道这个区间内，至少出现过一次的数，就能找到答案了？
3. 假设a[i], 如果有任何一个查询(l.r)覆盖了区间i，那么它就有a[i]
4. 固定位置r，那么对于l来说，就是那些a[i]出现了，但是i >= l 的数 
