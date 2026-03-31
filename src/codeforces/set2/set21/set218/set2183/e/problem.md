You are given a sequence `a` of length `n` and a positive integer `m`. Each element of `a` is an integer in the range `[0, m]`.

A sequence `a` is **good** if and only if both of the following hold:

1. `a_1 < a_2 < a_3 < … < a_n` (strictly increasing).
2.  
   `1/lcm(a_1, a_2) + 1/lcm(a_2, a_3) + … + 1/lcm(a_{n-1}, a_n) + 1/lcm(a_n, a_1) >= 1`.

You must replace every `0` in `a` with an integer from `[1, m]`. Count the number of ways to do this so that the resulting sequence is good. Print the answer modulo `998244353`.

\* The least common multiple `lcm(x, y)` of two positive integers is the smallest positive integer divisible by both. For example, `lcm(2, 3) = 6`, `lcm(4, 6) = 12`.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 1000`). The description of the test cases follows.

The first line of each test case contains two integers `n` and `m` (`2 <= n <= m <= 3000`).

The second line contains `n` integers `a_1, a_2, …, a_n` (`0 <= a_i <= m`).

It is guaranteed that the sum of `m` over all test cases does not exceed `3000`.

## Output

For each test case, output a single integer — the number of valid completions modulo `998244353`.

## Example

**Input**

```
5
4 6
1 0 0 6
2 2
2 1
5 24
0 0 4 0 0
5 6
0 0 6 0 0
20 2000
1 0 0 0 0 14 0 0 0 0 0 0 0 0 0 514 0 0 0 0
```

**Output**

```
2
0
10
0
973702700
```

## Note

In the first test case, there are `2` good completions:

- `[1, 2, 3, 6]`:  
  `1/lcm(1,2) + 1/lcm(2,3) + 1/lcm(3,6) + 1/lcm(6,1) = 1/2 + 1/6 + 1/6 + 1/6 = 1`.
- `[1, 2, 4, 6]`:  
  `1/lcm(1,2) + 1/lcm(2,4) + 1/lcm(4,6) + 1/lcm(6,1) = 1/2 + 1/4 + 1/12 + 1/6 = 1`.

In the second test case, the sequence is `[2, 1]`. Since `2 < 1` is false, the strictly increasing condition fails and the answer is `0`.

In the fourth test case, the fixed nonzero entry is `6` at position `3` with `m = 6`. For a strictly increasing sequence with values in `[1, 6]`, we would need `6 < a_4 < a_5 <= 6`, which is impossible, so the answer is `0`.


### ideas
1. 第一个条件比较好处理。主要是第二个条件咋搞呢？
2. lcm(a[i], a[i+1]) = a[i] * a[i+1] / gcd(a[i], a[i+1])
3. 1/lcm(a[i], a[i+1]) = gcd(a[i], a[i+1]) / (a[i] * a[i+1])
4. >=  1/a[i] - 1 / a[i+1] 是否成立呢？
5. (a[i+1] - a[i]) / (a[i] * a[i+1])
6. a[i+1] - a[i] >= gcd(a[i], a[i+1]) 这个成立