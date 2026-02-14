# Problem A

From beginning till end, this message has been waiting to be conveyed.

For a given unordered multiset of n lowercase English letters ("multi" means that a letter may appear more than once), we treat all letters as strings of length 1, and repeat the following operation n - 1 times:

- Remove any two elements s and t from the set, and add their concatenation s + t to the set.

The cost of such operation is defined to be f(s,c) * f(t,c) summed over all characters c, where f(s, c) denotes the number of times character c appears in string s.

Given a non-negative integer k, construct any valid non-empty set of no more than 100,000 letters, such that the minimum accumulative cost of the whole process is exactly k. It can be shown that a solution always exists.

## Input

The first and only line of input contains a non-negative integer k (0 ≤ k ≤ 100,000) — the required minimum cost.

## Output

Output a non-empty string of no more than 100,000 lowercase English letters — any multiset satisfying the requirements, concatenated to be a string.

Note that the printed string doesn't need to be the final concatenated string. It only needs to represent an unordered multiset of letters.

## Examples

### Example 1

**Input:**
```
12
```

**Output:**
```
abababab
```

### Example 2

**Input:**
```
3
```

**Output:**
```
codeforces
```

## Note

For the multiset {'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}, one of the ways to complete the process is as follows:

- {"ab", "a", "b", "a", "b", "a", "b"}, with a cost of 0;
- {"aba", "b", "a", "b", "a", "b"}, with a cost of 1;
- {"abab", "a", "b", "a", "b"}, with a cost of 1;
- {"abab", "ab", "a", "b"}, with a cost of 0;
- {"abab", "aba", "b"}, with a cost of 1;
- {"abab", "abab"}, with a cost of 1;
- {"abababab"}, with a cost of 8.

The total cost is 12, and it can be proved to be the minimum cost of the process.


### ideas
1. 如果都是a可行吗？
2. 最后一步把 s = w(a), t = v(a) 合并的cost  = w * v
3. dp[n] = 表示长度为n的字符串a的长度 = n时的最大 cost
4. dp[n] = dp[n/2] + dp[(n+1)/2) + (n+1) / 2 * (n/2)
5. 假设 {aa, a, a} 是先合并出 {aa, aa} 还是 {aaa, a}?
6. 1 + 4 vs 2 + 3
7. 假设目前有一个len(s) = x, 如果先合并出2，在合并x => 1 + x * 2
8. 如果先合并出 x + 1, 再合并 x + x + 1 = 1 + 2 * x (似乎没区别)？
9. (x, y, z), x * y + (x + y) * z => x * y + x * z + y * z
10.           x * z + (x + z) * y => ...
11.           y * z + (y + z) * x => ...
12. 确实是相等的，对于同一个字符，它的贡献不论什么顺序，结果是一样的
13. 假设有x个a，1 + 2 + 3 + ... <= k