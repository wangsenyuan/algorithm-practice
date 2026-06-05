# D. String Transformation

[Problem link](https://codeforces.com/problemset/problem/119/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

Let `s` be a string of length `n`. Its characters are numbered from `0` to `n - 1`.
Let `i` and `j` be integers with `0 <= i < j < n`. Define:

```text
f(s, i, j) = s[i+1 ... j-1] + r(s[j ... n-1]) + r(s[0 ... i])
```

Here:

- `s[p ... q]` is the substring of `s` from index `p` to `q` (inclusive)
- `+` is string concatenation
- `r(x)` is the reverse of string `x`

If `j = i + 1`, then `s[i+1 ... j-1]` is the empty string.

You are given two non-empty strings `a` and `b`. Find integers `i` and `j` such that
`f(a, i, j) = b`.

Among all valid pairs, choose the **maximum possible** `i`. If several valid `j`
exist for that `i`, choose the **minimum** `j`.


### ideas.
1. r(s[:i]) 时s的一个后缀
2. 比如 a = 12345678, b = 45687321
3. 然后a[i:..]是b的一个前缀 
4. i成立的条件是, a[:i]的反转是b的一个后缀 (这个很好处理)
5. 且 a[i:?]是b的一个前缀。
6. 确定了j以后， a[i:j]是前缀， reverse(a[j+1:n]) 是b[j-i+1:n-i]
7. 在b中找到a的后缀的那些位置(这个有可能有很多个)
8. 对于这些位置，i似乎是计算出来的。比如b[4] = 8, b[5] = 87,
9. 那么i就确定了 = n - 4, 然后检查i是否满足条件即可

## Solution

The target form is:

```text
b = a[i+1 ... j-1] + reverse(a[j ... n-1]) + reverse(a[0 ... i])
```

Let:

```text
l = i + 1
r = j
```

Then we need:

```text
b = a[l ... r-1] + reverse(a[r ... n-1]) + reverse(a[0 ... l-1])
```

The answer printed is `(l - 1, r)`.

There are three parts to verify:

1. the middle prefix `a[l ... r-1]`,
2. the reversed suffix `reverse(a[r ... n-1])`,
3. the reversed prefix `reverse(a[0 ... l-1])`.

### Precomputing Prefix And Suffix Checks

First compute:

```go
dp1 := zFunction(b + "#" + a)[n+1:]
```

Here:

```text
dp1[x] = longest length such that a[x ...] matches prefix of b
```

So `dp1[l] >= r-l` means:

```text
a[l ... r-1] == b[0 ... r-l-1]
```

Next, define:

```go
br := reverse(b)
```

For a valid `l`, the final part must satisfy:

```text
reverse(a[0 ... l-1]) == suffix of b
```

Equivalently:

```text
a[0 ... l-1] == prefix of reverse(b)
```

The code only keeps such valid `l`:

```go
for i := 0; i+1 < n && a[i] == br[i]; i++ {
	dp2[i+1] = dp1[i+1]
}
```

So:

```text
dp2[l] = dp1[l] if reverse(a[0 ... l-1]) matches the suffix of b
dp2[l] = -inf otherwise
```

Then the middle part check becomes:

```go
l + dp2[l] >= r
```

If this is true, both the final reversed-prefix part and the first middle part
are valid.

### Finding The Reversed Suffix Part

The remaining part is:

```text
reverse(a[r ... n-1])
```

Let:

```go
ar := reverse(a)
```

Then `reverse(a[r ... n-1])` is a prefix of `ar` with length `n-r`.

The code scans `b` with KMP over pattern `ar`. At each position `i` in `b`, let
`j` be the current matched prefix length of `ar`. This means:

```text
b[i-j+1 ... i] == ar[0 ... j-1]
```

So this matched substring can be the second part:

```text
reverse(a[r ... n-1])
```

where:

```go
r := n - j
l := n - i - 1
```

Once `l` and `r` are known, the candidate answer is:

```go
update(l-1, r)
```

but only if:

```go
l + dp2[l] >= r
```

### Why No Need To Try KMP Fallbacks

At a fixed scan position `i`, KMP gives the longest possible `j`.

If we fall back with:

```go
j = p[j-1]
```

then `j` becomes smaller. Since:

```go
r = n - j
```

a smaller `j` gives a larger `r`.

But `l` is still the same, because it depends only on `i`:

```go
l = n - i - 1
```

Therefore, a fallback candidate requires matching an even longer middle part
`a[l ... r-1]`.

If the current longest `j` fails:

```text
l + dp2[l] < r
```

then every fallback has `r2 > r`, so it also fails:

```text
l + dp2[l] < r < r2
```

If the current longest `j` succeeds, fallback is unnecessary too, because it
would produce a larger `r`, and for the same `l` the required answer uses the
minimum possible `j`.

This is why the code checks only the current KMP `j` at each position.

### Selecting The Best Answer

The problem wants:

- maximum output `i`,
- and among those, minimum output `j`.

Since output `i = l - 1`, the update rule is:

```go
if res[0] < i || res[0] == i && res[1] > j {
	res = (i, j)
}
```

### Complexity

The solution uses one Z-function and one KMP scan, both linear. The total time is:

```text
O(n)
```

with `O(n)` memory.
