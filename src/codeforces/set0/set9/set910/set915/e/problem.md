# Problem E: Physical Education Lessons

## Problem Description

This year Alex has finished school, and now he is a first-year student of Berland State University. For him it was a total surprise that even though he studies programming, he still has to attend physical education lessons. The end of the term is very soon, but, unfortunately, Alex still hasn't attended a single lesson!

Since Alex doesn't want to get expelled, he wants to know the number of working days left until the end of the term, so he can attend physical education lessons during these days. But in BSU calculating the number of working days is a complicated matter:

There are $n$ days left before the end of the term (numbered from 1 to $n$), and initially all of them are working days. Then the university staff sequentially publishes $q$ orders, one after another. Each order is characterised by three numbers $l$, $r$ and $k$:

- If $k = 1$, then all days from $l$ to $r$ (inclusive) become non-working days. If some of these days are made working days by some previous order, then these days still become non-working days.
- If $k = 2$, then all days from $l$ to $r$ (inclusive) become working days. If some of these days are made non-working days by some previous order, then these days still become working days.

Help Alex to determine the number of working days left after each order!

## Input

- The first line contains one integer $n$
- The second line contains one integer $q$ 
- Then $q$ lines follow, where the $i$-th line contains three integers $l_i$, $r_i$ and $k_i$ representing the $i$-th order

**Constraints:**
- $1 \leq n \leq 10^9$
- $1 \leq q \leq 3 \cdot 10^5$
- $1 \leq l_i \leq r_i \leq n$
- $1 \leq k_i \leq 2$

## Output

Print $q$ integers. The $i$-th of them must be equal to the number of working days left until the end of the term after the first $i$ orders are published.

## Example

### Input
```
4
6
1 2 1
3 4 1
2 3 2
1 3 2
2 4 1
1 4 2
```

### Output
```
2
0
2
3
1
4
```

### Explanation

Let's trace through the orders:
1. **Order 1**: Days 1-2 become non-working → 2 working days remain (days 3-4)
2. **Order 2**: Days 3-4 become non-working → 0 working days remain
3. **Order 3**: Days 2-3 become working → 2 working days remain (days 2-3)
4. **Order 4**: Days 1-3 become working → 3 working days remain (days 1-3)
5. **Order 5**: Days 2-4 become non-working → 1 working day remains (day 1)
6. **Order 6**: Days 1-4 become working → 4 working days remain (days 1-4)



### ideas
1. 用一个按需创建节点的，带lazy tag的segment tree，似乎是可以处理的