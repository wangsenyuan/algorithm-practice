# Problem

There is a long string `s` of `n` digits. Iahub may delete some digits (possibly none, but not all) so that the remaining digits form a non-empty number divisible by `5`. The result may have leading zeros.

Count the number of different ways to do this, **modulo** `10^9 + 7`. Two ways differ if the set of deleted positions differs.

The string `s` is not given directly: see **Input**.

## Input

- First line: string `a` (`1 <= |a| <= 10^5`), digits only.
- Second line: integer `k` (`1 <= k <= 10^9`).

The plate `s` is `a` repeated `k` times (concatenated). So `n = |a| * k`.

## Output

Print one integer — the number of valid ways, modulo `10^9 + 7`.

## Examples

### Example 1

Input

```text
1256
1
```

Output

```text
4
```

### Example 2

Input

```text
13990
2
```

Output

```text
528
```

### Example 3

Input

```text
555
2
```

Output

```text
63
```

## Note

- In Example 1, the valid endings divisible by `5` give four ways: `5`, `15`, `25`, `125`.
- In Example 2, `s` is `1399013990` (two copies of `a`).
- In Example 3, any non-empty subset of deletions works except deleting all digits, so `2^6 - 1 = 63` ways.


### ideas
1. 结尾要0/5的数，可以被5整除
2. 假设最后一个数是w[i] = 0/5, 那么它后面的数都必须删除掉, 但是它前面的数，可以不删或者删除
3. pow(2, i)  (i 从0开始)
4. 对于每个i，它的真实的位置 = i, i + m, i + 2 * m, ... i + (k - 1) * m
5. pow(2, i) + pow(i + m) + pow(i + 2 * m)