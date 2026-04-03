# E. XOR Expectation Squared

You are given an array of \(n\) integers \(a_1, a_2, \ldots, a_n\). You are also given an array \(p_1, p_2, \ldots, p_n\).

Let \(S\) denote the **random multiset** (it may contain equal elements) constructed as follows:

1. Initially, \(S\) is empty.
2. For each \(i\) from \(1\) to \(n\), insert \(a_i\) into \(S\) with probability \(\frac{p_i}{10^4}\). Each insertion is independent of the others.

Denote \(f(S)\) as the **bitwise XOR** of all elements of \(S\) (XOR of an empty multiset is \(0\)).

Compute the **expected value** of \((f(S))^2\). Output the answer **modulo** \(10^9 + 7\).

### Modular fraction output

Let \(M = 10^9 + 7\). The answer can be written as an irreducible fraction \(\frac{p}{q}\) with integers \(p, q\) and \(q \not\equiv 0 \pmod{M}\).

Output the integer \(p \cdot q^{-1} \bmod M\) — that is, the unique \(x\) with \(0 \le x < M\) and \(x \cdot q \equiv p \pmod{M}\).

## Input

Each test contains multiple test cases. The first line contains the number of test cases \(t\) (\(1 \le t \le 10^4\)). The description of the test cases follows.

The first line of each test case contains a single integer \(n\) (\(1 \le n \le 2 \cdot 10^5\)).

The second line of each test case contains \(n\) integers \(a_1, a_2, \ldots, a_n\) (\(1 \le a_i \le 1023\)).

The third line of each test case contains \(n\) integers \(p_1, p_2, \ldots, p_n\) (\(1 \le p_i \le 10^4\)).

It is guaranteed that the sum of \(n\) over all test cases does not exceed \(2 \cdot 10^5\).

## Output

For each test case, output the expected value of \((f(S))^2\), modulo \(10^9 + 7\).

## Example

**Input**

```
4
2
1 2
5000 5000
2
1 1
1000 2000
6
343 624 675 451 902 820
6536 5326 7648 2165 9430 5428
1
1
10000
```

**Output**

```
500000007
820000006
280120536
1
```

## Note

**First test case:** \(a = [1, 2]\) and each element is inserted with probability \(\frac{1}{2}\), since \(p_1 = p_2 = 5000\) and \(\frac{p_i}{10^4} = \frac{1}{2}\). There are \(4\) outcomes for \(S\), each with probability \(\frac{1}{4}\):

| \(S\) | \(f(S)\) | \((f(S))^2\) |
| :--- | :---: | :---: |
| \(\emptyset\) | \(0\) | \(0\) |
| \(\{1\}\) | \(1\) | \(1\) |
| \(\{2\}\) | \(2\) | \(4\) |
| \(\{1, 2\}\) | \(1 \oplus 2 = 3\) | \(9\) |

Hence the expectation is \(0 \cdot \frac{1}{4} + 1 \cdot \frac{1}{4} + 4 \cdot \frac{1}{4} + 9 \cdot \frac{1}{4} = \frac{14}{4} = \frac{7}{2} \equiv 500000007 \pmod{10^9+7}\).

**Second test case:** \(a = [1, 1]\). Element \(a_1\) is inserted with probability \(0.1\), and \(a_2\) with probability \(0.2\). The possible multisets:

| \(S\) | \(f(S)\) | \((f(S))^2\) | Probability |
| :--- | :---: | :---: | :--- |
| \(\emptyset\) | \(0\) | \(0\) | \((1-0.1)(1-0.2) = 0.72\) |
| \(\{1\}\) (one copy) | \(1\) | \(1\) | \((1-0.1)\cdot 0.2 + 0.1\cdot(1-0.2) = 0.26\) |
| \(\{1, 1\}\) | \(0\) | \(0\) | \(0.1 \cdot 0.2 = 0.02\) |

So the answer is \(0 \cdot 0.72 + 1 \cdot 0.26 + 0 \cdot 0.02 = 0.26 = \frac{26}{100} \equiv 820000006 \pmod{10^9+7}\).


### ideas
1. dp[x] x >= 0, x <= 1023 出现的期望概率
2. f(S) = sum(x * dp[x]) ?

## insights

1. `dp[x]` is the probability that the final xor value equals `x`.
   - It is not "the contribution of `x` to the answer".
   - After processing all elements, `dp` is exactly the distribution of the random variable `X = f(S)`.

2. The transition is a standard probability DP.
   - If the current xor is `x`, then after seeing `a[i]`:
   - with probability `1 - p[i]`, it stays `x`
   - with probability `p[i]`, it becomes `x ^ a[i]`
   - So `ndp` stores the probabilities of each xor result after considering one more element.

3. The problem asks for `E[(f(S))^2]`.
   - Let `X = f(S)`.
   - If we know the full distribution of `X`, then for any function `g`,
     `E[g(X)] = sum g(x) * P(X = x)`.
   - Here `g(x) = x^2`, so
     `E[X^2] = sum x^2 * P(X = x) = sum x^2 * dp[x]`.

4. This is different from `(E[X])^2`.
   - `E[X] = sum x * dp[x]`
   - so `(E[X])^2 = (sum x * dp[x])^2`
   - but the answer we need is `E[X^2] = sum x^2 * dp[x]`
   - in general they are not equal.

5. A small example:
   - If `X = 0` with probability `1/2`, and `X = 2` with probability `1/2`,
   - then `E[X] = 1`, so `(E[X])^2 = 1`
   - but `E[X^2] = (0^2 + 2^2) / 2 = 2`
   - this is exactly why the final formula must be `sum x^2 * dp[x]`.
