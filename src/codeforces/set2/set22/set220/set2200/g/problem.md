# G. Operation Permutation

[Problem link](https://codeforces.com/problemset/problem/2200/G)

**Contest:** [Codeforces Round 1084 (Div. 3)](https://codeforces.com/contest/2200)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

AksLolCoding has an integer `x` and a list of `n` operations. Each operation is a string starting with one of the symbols `+`, `-`, `x`, or `/` (representing addition, subtraction, multiplication, and real number division respectively), followed immediately by a positive integer `y` (`1 ≤ y ≤ 10^9`). For example, the operation `x3` represents multiplying `x` by `3`.

AksLolCoding will randomly permute the operations and then apply all operations sequentially to `x` in the permuted order. Help AksLolCoding compute the expected final value of `x` modulo `10^9 + 7`.

Formally, let `M = 10^9 + 7`. It can be shown that the answer can be expressed as an irreducible fraction `p / q`, where `p` and `q` are integers and `q ≢ 0 (mod M)`. Output the integer equal to `p · q^{-1} (mod M)`. In other words, output such an integer `a` that `0 ≤ a < M` and `a · q ≡ p (mod M)`.

The expected final value of `x` is the average of the final value of `x` over all `n!` permutations.

Note: `x` is used to represent multiplication, not `*`.

## Input

The first line contains a single integer `t` (`1 ≤ t ≤ 1000`), the number of test cases.

For each test case, the first line contains two integers `n` and `x` (`1 ≤ n ≤ 3000`, `1 ≤ x ≤ 10^9`).

The second line of each test case contains `n` strings, each representing an operation in the format described above.

The sum of `n^2` over all test cases does not exceed `3000^2`.

## Output

For each test case, output a single integer: the expected final value of `x` modulo `10^9 + 7`.

## Example

### Input

```text
4
2 10
x2 -10
4 2
+6 +7 /1 -13
8 1
+1 x2 x3 +4 +5 +6 -7 -8
9 864209753
-918273645 x564738291 /365107362 x734582911 -654321789 x998244353 +172519103 /482193765 /482091376
```

### Output

```text
5
2
166666677
601980218
```

### Note

In the first test case, `x` can either be `(10 · 2) − 10 = 10` or `(10 − 10) · 2 = 0`, resulting in an expected value of `5`.

In the second test case, all possible permutations result in `x = 2`.

In the third test case, the expected value of `x` is `55 / 6`.


## ideas
1. -y可以修改为 +(mod - y), /y可以修改为x(inv(y))
2. 所以只需要考虑+/*; 对于一个数+y来说, 需要知道它的贡献
3. 如果没有乘法, 那么+y的贡献 = y * 1
4. 如果有 *2, *3, 那么+y的贡献 = y * 2 + y * 3 + y * (2 * 3)
5. 看乘法的位置

## Solution

First normalize every operation in the field modulo `M = 1e9 + 7`:

- `-y` becomes `+(M-y)`;
- `/y` becomes multiplication by `y^(M-2)`.

After this transformation, there are only additions and multiplications. Let
`a` be the sum of all addition values, and let the multiplicative values be
`q[0..m-1]`.

### Contribution of the initial value

The initial `x` is affected by every multiplication regardless of the
permutation, so its contribution is fixed:

```text
x * product(q[i])
```

It must **not** be divided by the number of permutations.

### Contribution of one addition

Consider a fixed addition. Its value is multiplied only by the multiplication
operations placed after it. Suppose that exactly a particular subset `S` of
the `m` multiplications is after this addition, where `|S| = k`.

Among this addition and the `m` multiplications, the number of relative orders
having exactly this subset after the addition is:

```text
k! * (m-k)!
```

The total number of relative orders is `(m+1)!`, hence the probability is:

```text
k! * (m-k)! / (m+1)!
= 1 / ((m+1) * C(m, k))
```

The other additions do not change this probability or the multiplier. Thus all
additions have the same expected trailing multiplier, and we can multiply it
by their total `a`.

### DP over subset sizes

Let `dp[k]` be the sum of products of all size-`k` subsets of `q`:

```text
dp[k] = sum(product(q[i]) for every subset of size k)
```

Initialize `dp[0] = 1`. For each multiplier `q`, update in descending order:

```text
dp[k] += dp[k-1] * q
```

Then the expected multiplier after an addition is:

```text
E = sum(dp[k] * k! * (m-k)! / (m+1)!) for k = 0..m
```

Therefore the answer is:

```text
x * product(q[i]) + a * E
```

All divisions above are modular inverses modulo `M`.

### Complexity

The DP uses `O(m)` memory and `O(m^2)` time. Since `m <= n` and the sum of
`n^2` over test cases is bounded, this fits the limit.
