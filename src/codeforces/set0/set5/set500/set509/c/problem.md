# Problem Description

Vasya had a strictly increasing sequence of positive integers a₁, ..., aₙ. Vasya used it to build a new sequence b₁, ..., bₙ, where bᵢ is the sum of digits of aᵢ's decimal representation. Then sequence aᵢ got lost and all that remained is sequence bᵢ.

Vasya wonders what the numbers aᵢ could be like. Of all the possible options he likes the one sequence with the minimum possible last number aₙ. Help Vasya restore the initial sequence.

It is guaranteed that such a sequence always exists.

## Input

The first line contains a single integer number n (1 ≤ n ≤ 300).

Next n lines contain integer numbers b₁, ..., bₙ — the required sums of digits. All bᵢ belong to the range 1 ≤ bᵢ ≤ 300.

## Output

Print n integer numbers, one per line — the correct option for numbers aᵢ, in order of following in sequence. The sequence should be strictly increasing. The sum of digits of the i-th number should be equal to bᵢ.

If there are multiple sequences with least possible number aₙ, print any of them. Print the numbers without leading zeroes.

## Examples

### Example 1

**Input:**
```
3
1
2
3
```

**Output:**
```
1
2
3
```

### Example 2

**Input:**
```
3
3
2
1
```

**Output:**
```
3
11
100
```


### ideas
1. 序列a估计不能用int表示，因为b[i] = 300 => a[i] >= 300 / 9 = 33个9组成, 远超1e18了
2. 假设dp[i]是目前满足b[...i]为止最小的a[i] ?
3. 然后继续目前的b[i+1], 这样的dp[i]是不是最优的结果？
4. 也就是说，每一步都得到满足条件的最小的a[i], 能否得到最优的a[n]?
5. 好像是的，因为a[i]越小，对a[i+1]的限制越少，越能得到更小的a[i+1]
6. 那么剩下就是，如何在a[i], b[i+1]的基础上得到a[i+1]?
7. 首先是相同长度，如果可以的话，就使用相同长度，否则就增加1位，2位，这样子