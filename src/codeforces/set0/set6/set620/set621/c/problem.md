# Problem: Wet Shark and Flowers

## Problem Statement

There are `n` sharks who grow flowers for Wet Shark. They are all sitting around the table, such that sharks `i` and `i + 1` are neighbours for all `i` from `1` to `n - 1`. Sharks `n` and `1` are neighbours too.

Each shark will grow some number of flowers `s[i]`. For i-th shark, value `s[i]` is a random integer equiprobably chosen in range from `l[i]` to `r[i]`. 

Wet Shark has its favourite prime number `p`, and he really likes it! If for any pair of neighbouring sharks `i` and `j` the product `s[i] * s[j]` is divisible by `p`, then Wet Shark becomes happy and gives **1000 dollars** to each of these sharks.

At the end of the day, sharks sum all the money Wet Shark granted to them. Find the expectation of this value.

## Input

The first line of the input contains two space-separated integers `n` and `p` (`3 ≤ n ≤ 100,000`, `2 ≤ p ≤ 10^9`) — the number of sharks and Wet Shark's favourite prime number. It is guaranteed that `p` is prime.

The i-th of the following `n` lines contains information about i-th shark — two space-separated integers `l[i]` and `r[i]` (`1 ≤ l[i] ≤ r[i] ≤ 10^9`), the range of flowers shark `i` can produce. Remember that `s[i]` is chosen equiprobably among all integers from `l[i]` to `r[i]`, inclusive.

## Output

Print a single real number — the expected number of dollars that the sharks receive in total. Your answer will be considered correct if its absolute or relative error does not exceed `10^-6`.

## Examples

### Example 1

**Input:**
```
3 2
1 2
420 421
420420 420421
```

**Output:**
```
4500.0
```

### Example 2

**Input:**
```
3 5
1 4
2 3
11 14
```

**Output:**
```
0.0
```

## Note

A prime number is a positive integer number that is divisible only by 1 and itself. 1 is not considered to be prime.

### Example 1 Explanation

First shark grows some number of flowers from 1 to 2, second shark grows from 420 to 421 flowers, and third from 420420 to 420421. There are eight cases for the quantities of flowers `(s[0], s[1], s[2])` each shark grows:

1. **(1, 420, 420420)**: Note that `s[0] * s[1] = 420`, `s[1] * s[2] = 176576400`, and `s[2] * s[0] = 420420`. For each pair, 1000 dollars will be awarded to each shark. Therefore, each shark will be awarded 2000 dollars, for a total of **6000** dollars.

2. **(1, 420, 420421)**: Now, the product `s[2] * s[0]` is not divisible by 2. Therefore, sharks `s[0]` and `s[2]` will receive 1000 dollars, while shark `s[1]` will receive 2000. The total is **4000**.

3. **(1, 421, 420420)**: Total is **4000**.

4. **(1, 421, 420421)**: Total is **0**.

5. **(2, 420, 420420)**: Total is **6000**.

6. **(2, 420, 420421)**: Total is **6000**.

7. **(2, 421, 420420)**: Total is **6000**.

8. **(2, 421, 420421)**: Total is **4000**.

The expected value is `(6000 + 4000 + 4000 + 0 + 6000 + 6000 + 6000 + 4000) / 8 = 36000 / 8 = 4500.0`.

### Example 2 Explanation

In the second sample, no combination of quantities will garner the sharks any money.


### ideas
1. p是质数，所以， s[i] * s[j] % p = 0, 要么两个都可以整除p, 要么只有一个可以整除p
2. dp[i][0][?] 到i为止，前一个可以整除p, 且第一个是?的收益，
3. dp[i][1][?] 。。。 不能整除p
4. l[i]...r[i]中间共有 r - l + 1 个数，其中p的倍数 = r / p - (l - 1) / p 个数
5. 如果当前选到了一个p的倍数, 那么收益 = dp[i][0][?] + 2000 * w
6. 这个思路没有问题，但是概率的作用没想清楚
7. 期望值是线性可加的～
8. 