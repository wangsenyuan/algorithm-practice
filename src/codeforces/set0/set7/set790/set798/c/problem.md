Mike has a sequence A = [a1, a2, ..., an] of length n. He considers the sequence B = [b1, b2, ..., bn] beautiful if the gcd of all its elements is bigger than 1, i.e. gcd(B) > 1.

Mike wants to change his sequence in order to make it beautiful. In one move he can choose an index i (1 ≤ i < n), delete numbers ai, ai+1 and put numbers ai - ai+1, ai + ai+1 in their place instead, in this order. He wants perform as few operations as possible. Find the minimal number of operations to make sequence A beautiful if it's possible, or tell him that it is impossible to do so.

gcd(B) is the biggest non-negative number d such that d divides bi for every i (1 ≤ i ≤ n).

## Input

The first line contains a single integer n (2 ≤ n ≤ 100 000) — length of sequence A.

The second line contains n space-separated integers a1, a2, ..., an (1 ≤ ai ≤ 109) — elements of sequence A.

## Output

Output on the first line "YES" (without quotes) if it is possible to make sequence A beautiful by performing operations described above, and "NO" (without quotes) otherwise.

If the answer was "YES", output the minimal number of moves needed to make sequence A beautiful.

## Examples

### Example 1

**Input:**
```
2
1 1
```

**Output:**
```
YES
1
```

### Example 2

**Input:**
```
3
6 2 4
```

**Output:**
```
YES
0
```

### Example 3

**Input:**
```
2
1 3
```

**Output:**
```
YES
1
```

## Note

In the first example you can simply make one move to obtain sequence [0, 2] with gcd(0, 2) = 2 > 1.

In the second example the gcd of the sequence is already greater than 1.


### ideas
1. 咋又是个数学相关的～
2. 如果 a[i] 和 a[i+1]都是奇数，那么一次操作后，它们都变成了偶数
3. 如果 其中一个是奇数，操作后，似乎没有好处？
4. [a, b, c]
5. [a - b, a + b, c] 再进行一次操作 [a - b - (a + b), a - b + a + b, c]
6. => [-2*b, 2 * a, c]
7. 一个奇数，一个偶数，两次操作后，就变成了两个偶数
8. 那这个是最优的解吗？