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

### Key Insights

**Step 1 — Recurrence for occurrence count**

Let `cnt[n]` = number of times pattern `s` appears in `f_n`. Because `f_n = f_{n-1} · f_{n-2}`, every occurrence falls into exactly one of three buckets:

- wholly inside `f_{n-1}` → contributes `cnt[n-1]`
- wholly inside `f_{n-2}` → contributes `cnt[n-2]`
- straddles the join point → contributes `cross[n]`

So: `cnt[n] = cnt[n-1] + cnt[n-2] + cross[n]`

**Step 2 — Computing `cross[n]` cheaply**

An occurrence that crosses the boundary uses some suffix of `f_{n-1}` followed by some prefix of `f_{n-2}`. The straddling window is at most `|s|-1` characters on each side, so we only need to keep:

- `suff[n]` = last `|s|-1` characters of `f_n`
- `pref[n]` = first `|s|-1` characters of `f_n`

`cross[n]` = KMP count of `s` in `suff[n-1] + pref[n-2]` (a string of length ≤ `2(|s|-1)`)

These short prefixes and suffixes are updated cheaply each step by appending/trimming.

**Step 3 — Stabilisation of `cross[n]`**

Once `|f_n| ≥ |s|-1`, both `pref[n]` and `suff[n]` are fully determined by just the first/last `|s|-1` characters, which no longer grow. Since `f_n = f_{n-1} · f_{n-2}`, the parity of `n` determines which prior pair of (suff, pref) is used at the join. Therefore, after a small bootstrap phase (roughly `O(log |s|)` steps), `cross[n]` is periodic with period 2:

```
cross[n] = g[n & 1]   for all sufficiently large n
```

**Step 4 — Matrix exponentiation for large `k`**

Once `cross[n]` is constant per parity, `cnt[n]` satisfies:

```
[cnt[n]  ]   [1 1 g[n&1]] [cnt[n-1]  ]
[cnt[n-1]] = [1 0    0  ] [cnt[n-2]  ]
[   1    ]   [0 0    1  ] [    1     ]
```

Two consecutive steps (even then odd, or vice versa) can be combined into a single 3×3 matrix. Raising that matrix to a power via fast exponentiation handles `k` up to `10^18` in `O(log k)` multiplications.

**Step 5 — Per-query algorithm**

1. Build KMP failure function for `s`.
2. Simulate the recurrence for small `n` (until `cross` stabilises), tracking `cnt`, `pref`, `suff`.
3. If `k` is already reached, return `cnt[k]`.
4. Otherwise apply matrix exponentiation to jump from the stable base to `k`, all arithmetic mod `10^9+7`.