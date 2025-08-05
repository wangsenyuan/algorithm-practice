# Problem G: Little John's Milk Satisfaction

## Problem Description

Little John is as little as night is day — he was known to be a giant, at possibly 2.1 metres tall. It has everything to do with his love for milk.

His dairy diary has $n$ entries, showing that he acquired $a_i$ pints of fresh milk on day $d_i$. Milk declines in freshness with time and stays drinkable for a maximum of $k$ days. In other words, fresh milk acquired on day $d_i$ will be drinkable between days $d_i$ and $d_i + k - 1$ inclusive.

Every day, Little John drinks drinkable milk, up to a maximum of $m$ pints. In other words:
- If there are less than $m$ pints of milk, he will drink them all and not be satisfied
- If there are at least $m$ pints of milk, he will drink exactly $m$ pints and be satisfied, and it's a **milk satisfaction day**

Little John always drinks the freshest drinkable milk first.

**Determine the number of milk satisfaction days for Little John.**

## Input

The first line of the input contains a single integer $t$ ($1 \leq t \leq 10^4$), the number of test cases.

The first line of each test case consists of three integers $n$, $m$, $k$ ($1 \leq n, m, k \leq 10^5$):
- $n$: the number of diary entries
- $m$: the maximum pints needed for a milk satisfaction day
- $k$: the duration of milk's freshness

Then follow $n$ lines of each test case, each with two integers $d_i$ and $a_i$ ($1 \leq d_i, a_i \leq 10^6$):
- $d_i$: the day on which the milk was acquired
- $a_i$: the number of pints acquired

**Note:** They are sorted in increasing values of $d_i$, and all values of $d_i$ are distinct.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output a single integer, the number of milk satisfaction days.

## Example

### Input
```
6
1 1 3
1 5
2 3 3
1 5
2 7
4 5 2
1 9
2 6
4 9
5 6
5 2 4
4 7
5 3
7 1
11 2
12 1
4 1 3
5 10
9 4
14 8
15 3
5 5 5
8 9
10 7
16 10
21 5
28 9
```

### Output
```
3
3
4
5
10
6
```

## Explanation

### Test Case 1
- 5 pints of milk are good for 3 days before spoiling.

### Test Case 2
The following will happen:

1. **Day 1**: He will receive 5 pints of milk and drink 3 of them (leaving 2 pints from day 1)
2. **Day 2**: He will receive 7 pints of milk and drink 3 of them (leaving 2 pints from day 1 and 4 pints from day 2)
3. **Day 3**: He will drink 3 pints from day 2 (leaving 2 pints from day 1 and 1 pint from day 2)
4. **Day 4**: The milk acquired on day 1 will spoil, and he will drink 1 pint from day 2 (no more milk is left)


### ideas
1. 如果每天能喝到m单位的牛奶，那么这天是good的，计算有多少天good
2. 最多到1e6天
3. 还有个关键是，牛奶是按照新鲜程度喝掉的。
4. 用segment tree，表示每天剩余的牛奶的数量
5. 然后只查询i-k+1...i的部分