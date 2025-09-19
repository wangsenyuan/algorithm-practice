# Problem C - Scales

## Problem Statement

Vanya has a scales for weighing loads and weights of masses $w^0, w^1, w^2, ..., w^{100}$ grams where $w$ is some integer not less than 2 (exactly one weight of each nominal value). Vanya wonders whether he can weight an item with mass $m$ using the given weights, if the weights can be put on both pans of the scales. 

Formally speaking, your task is to determine whether it is possible to place an item of mass $m$ and some weights on the left pan of the scales, and some weights on the right pan of the scales so that the pans of the scales were in balance.

## Input

The first line contains two integers $w, m$ ($2 \leq w \leq 10^9, 1 \leq m \leq 10^9$) — the number defining the masses of the weights and the mass of the item.

## Output

Print word 'YES' if the item can be weighted and 'NO' if it cannot.

## Examples

### Example 1
**Input:**
```
3 7
```
**Output:**
```
YES
```

### Example 2
**Input:**
```
100 99
```
**Output:**
```
YES
```

### Example 3
**Input:**
```
100 50
```
**Output:**
```
NO
```

## Notes

- **Note to the first sample test:** One pan can have an item of mass 7 and a weight of mass 3, and the second pan can have two weights of masses 9 and 1, correspondingly. Then $7 + 3 = 9 + 1$.

- **Note to the second sample test:** One pan of the scales can have an item of mass 99 and the weight of mass 1, and the second pan can have the weight of mass 100.

- **Note to the third sample test:** It is impossible to measure the weight of the item in the manner described in the input.

### ideas
1. w = 2 答案 = yes
2. w > 2, 有个复杂的地方是，砝码是可以放置到两边去的，从而得到一个差值m
3. 这类数学类的问题，好麻烦，不是我的菜～
4. 