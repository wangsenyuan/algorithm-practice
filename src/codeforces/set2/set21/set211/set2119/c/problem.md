# Problem

You are given four positive integers $n, l, r, k$. You need to find the lexicographically smallest\* array $a$ of length $n$, consisting of integers, such that:

- For every $1 \le i \le n$, $l \le a_i \le r$.
- $a_1 \mathbin{\&} a_2 \mathbin{\&} \ldots \mathbin{\&} a_n = a_1 \oplus a_2 \oplus \ldots \oplus a_n$, where $\&$ denotes the bitwise AND operation and $\oplus$ denotes the bitwise XOR operation.

If no such array exists, output $-1$. Otherwise, since the entire array might be too large to output, output $a_k$ only.

\* An array $a$ is lexicographically smaller than an array $b$ if and only if one of the following holds:

- $a$ is a prefix of $b$, but $a \ne b$; or
- in the first position where $a$ and $b$ differ, the array $a$ has a smaller element than the corresponding element in $b$.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

Each test case contains four positive integers $n, l, r, k$ ($1 \le k \le n \le 10^{18}$, $1 \le l \le r \le 10^{18}$).

## Output

For each test case, output $a_k$ or $-1$ if no array meets the conditions.

## Example

### Input

```
9
1 4 4 1
3 1 3 3
4 6 9 2
4 6 9 3
4 6 7 4
2 5 5 1
2 3 6 2
999999999999999999 1000000000000000000 1000000000000000000 999999999999999999
1000000000000000000 1 999999999999999999 1000000000000000000
```

### Output

```
4
1
6
8
-1
-1
-1
1000000000000000000
2
```

## Note

In the first test case, the array $a = [4]$. It can be proven that there is no array that meets the above requirements and has a smaller lexicographic order.

In the second test case, the array $a = [1, 1, 1]$. It can be proven that there is no array that meets the above requirements and has a smaller lexicographic order.

In the third test case and the fourth test case, the array $a = [6, 6, 8, 8]$. It can be proven that there is no array that meets the above requirements and has a smaller lexicographic order.

In the fifth test case and the sixth test case, it can be proven that there is no array that meets the above requirements.

### ideas
1. 如果n=1, 这个很简单， => l
2. 如果n=2, a[1] & a[2] = a[1] ^ a[2] => a[1]和a[2]不能有相同的位, 1 & 1 != 1 ^ 1
3. 且a[1] & a[2] = 0, a[1] ^ a[2] = 0
4. 所以，就是让a[1] ^ .... ^ a[n] = 0 
5. 前面尽量是l，如果只留下一位，是不行的
6. 因为如果n是偶数，前面xor = l的，那么最后的那位必须是l (and就是l)
7. 如果n是奇数，那么xor = 0, 那么最后一个必须是0
8. 考虑剩下两位 l, .... l, a, b
9. l & a & b = 0， 如果n是偶数, a ^ b = 0, 如果n是奇数, l ^ a ^ b = 0
10, 如果n是偶数, a ^ b = 0 => a = b, 进而找到最小的a, a & l = 0即可(这个很好处理)
11. 如果n是奇数, l ^ a ^ b = 0, => a ^ b = l (a和b在l=1的位置，要刚好相反，在l=0的地方，要同时为0或者同时为1)
12. let h = len(l), a[h] = 1, b[h] = 1 => a和b > l (这里必须为1)， 假设不是，那么没法保证a > l
13. 然后把剩下l = 1 的地方，全部分配给b，l = 0的地方，全部分配1个a, a = 1 << h
14. 可以
