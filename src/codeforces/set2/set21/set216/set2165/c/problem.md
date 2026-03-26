You are given `n` integers `a1, a2, ..., an`, each in the range `[0, 2^30)`.

You can spend `1` coin to increase any `ai` by `1`. You can perform this operation any number of times.

You need to solve `q` queries; for each query, you are given an integer `c`, also in the range `[0, 2^30)`. You would like it if there exists a sequence `b` of length `n` with the following properties:

- For every `1 <= i <= n`, `0 <= bi <= ai`.
- `b1 XOR b2 XOR ... XOR bn = c`, where XOR denotes the bitwise XOR operation.

Please calculate the minimum number of coins you will have to spend such that there exists a suitable `b`.

The queries are independent, meaning that any operations you perform on the sequence `a` will not impact future queries.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case consists of two integers `n`, `q` (`1 <= n <= 5 * 10^5`, `1 <= q <= 5 * 10^4`) — the length of sequence `a` and the number of queries.

The second line of each test case contains `n` integers `a1, a2, ..., an` (`0 <= ai < 2^30`) — the initial sequence `a`.

Each of the next `q` lines contains a single integer `c` (`0 <= c < 2^30`) — the target XOR.

It is guaranteed that the sum of `n` over all test cases does not exceed `5 * 10^5`.

It is guaranteed that the sum of `q` over all test cases does not exceed `5 * 10^4`.

## Output

For each query, output a single integer — the minimum coins you will have to spend such that there exists a suitable `b`.

## Example

**Input**

```
4
2 1
5 7
9
3 1
9 9 8
24
6 4
1 1 4 5 1 4
10
20
30
40
1 1
0
0
```

**Output**

```
1
7
3
11
16
31
0
```

## Note

In the first test case, we spend `1` coin to increase `a2` by `1`, resulting in sequence `[5, 8]`. A suitable `b` would be `[1, 8]`. It can be shown one cannot spend less than `1` coin to achieve the objective.

In the second test case, we can spend `7` coins to increase `a1` by `7`, resulting in sequence `[16, 9, 8]`. A suitable `b` would be `[16, 9, 1]`.


### ideas
1. 考虑一个数c。如果目前有一个a[i] >= c, 那么可以直接使用b[i] = c, 花费0
2. 如果a[i] < c， 那么必然是c的最高位为1，但是所有a[i]的最高位是0，那么必须使用某个a[i]，得到c的最高位
3. 那么这时候显然使用最大的a[i]是更优的。