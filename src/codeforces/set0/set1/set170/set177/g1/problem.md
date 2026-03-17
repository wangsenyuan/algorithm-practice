### Problem

Fibonacci strings are defined as follows:

- `f1 = "a"`
- `f2 = "b"`
- `fn = f_{n-1} f_{n-2}` for `n > 2` (concatenation)

Thus, the first five Fibonacci strings are: `"a"`, `"b"`, `"ba"`, `"bab"`, `"babba"`.

You are given `k` and `m`, and then `m` query strings `s_i`. For each `s_i`, you must find how many times it occurs **as a substring** in the Fibonacci string `f_k`.

### Input

- First line: two space-separated integers `k` and `m` — the index of the Fibonacci string and the number of queries.
- Next `m` lines: each contains a non-empty string `s_i`, consisting only of characters `'a'` and `'b'`.

**Subtasks:**

- For 30 points:\n
  - `1 ≤ k ≤ 3000`\n
  - `1 ≤ m ≤ 3000`\n
  - Total length of all `s_i` ≤ 3000
- For 100 points:\n
  - `1 ≤ k ≤ 10^18`\n
  - `1 ≤ m ≤ 10^4`\n
  - Total length of all `s_i` ≤ 10^5

### Output

For each string `s_i`, print the number of times it occurs in `f_k` as a substring.  
Since the answers can be large, print each result **modulo** `1000000007` (`10^9 + 7`).  
Print the answers in the same order as the queries appear in the input.

### Examples

**Input**
```text
6 5
a
b
ab
ba
aba
```

**Output**
```text
3
5
3
3
1
```


### ideas
1. a, b, ba, bab, babba, babbabab
2. len(s) <= 1e5
3. 可以先找出来s[i] 被包含在f[i]中， 如果找不到 = 0
4. 这个i, 不会很大（为了加快）可以使用pattern match？
5. 找到i以后呢？看看它出现了几次，然后计算下一个位置，出现了几次。
6. 然后是一个新的fib序列