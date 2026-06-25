# E. Seek the Truth

[Problem link](https://codeforces.com/problemset/problem/2222/E)

**Contest:** [Spectral::Cup 2026 Round 1 (Codeforces Round 1094, Div. 1 + Div. 2)](https://codeforces.com/contest/2222)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

**This is an interactive problem.**

There are two hidden integers `k` and `c` such that `k ∈ {1, 2, 3}` and `1 <= c <= 2^n - 1`. Note that `c != 0`.

Before any interaction, output a non-negative integer `a` (`0 <= a <= 2^n - 1`). The grader uses it as the initial element of set `S`, so initially `S = {a}`.

You may make queries of two types:

1. `I x` (`0 <= x <= 2^n - 1`) — insert `f(x)` into `S`. The jury responds with `|S|` after the insertion.
2. `Q y` (`0 <= y <= 2^n - 1`) — ask how many `z` in `S` satisfy `z >= y`. The jury responds with that count.

Elements in `S` are not duplicated: inserting a value already present does not increase `|S|`.

When you know the hidden values, output `A k c` (`1 <= k <= 3`, `1 <= c <= 2^n - 1`).

The function `f` is defined by the hidden parameters:

```text
f(x) = x & c   if k = 1
f(x) = x | c   if k = 2
f(x) = x ^ c   if k = 3
```

You may use at most `n + 30` operations in total (the initial choice of `a`, all `I`/`Q` queries, and the final `A` command).

After each output, flush. In Go use `fmt.Println` followed by reading the response, or flush the writer explicitly.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains one line with an integer `n` (`2 <= n <= 60`). Interaction follows.

## Interaction

For each test case:

1. Read `n`.
2. Output the initial element `a`.
3. Repeat queries until you can answer:
   - output `I x` or `Q y`, then read one integer response;
   - or output `A k c` to finish the test case.

## Example

The example below shows only the values written by your program and the jury responses. Blank lines are for readability.

### Input

```text
3
2

1

2

2

1

2

2
```

### Output

```text
3
I 3
A 1 3
1
I 0
Q 2
Q 1
Q 0
I 3
A 2 2
1
I 1
I 0
A 3 1
```

### Note

- Test 1: `n = 2`, hidden `k = 1`, `c = 3`. Start with `S = {3}`. `I 3` inserts `3 & 3 = 3`, so `|S| = 1`.
- Test 2: hidden `k = 2`, `c = 2`. Start with `S = {1}`. `I 0` inserts `0 | 2 = 2`, so `|S| = 2`.
- Test 3: hidden `k = 3`, `c = 1`. Start with `S = {1}`. `I 1` inserts `1 ^ 1 = 0`, so `|S| = 2`.

## Solution

The implementation turns the interactive task into two smaller questions:

1. find the hidden value `c`;
2. decide which operation produced the set values.

The key choice is to start with `a = 0`, so initially:

```text
S = {0}
```

Then we insert `0` and look at the returned set size.

```text
I 0
```

The inserted value is `f(0)`, and this immediately separates the `AND` case from the other two:

```text
k = 1: 0 & c = 0      -> duplicate, |S| stays 1
k = 2: 0 | c = c      -> new value, |S| becomes 2
k = 3: 0 ^ c = c      -> new value, |S| becomes 2
```

So if the answer to `I 0` is `1`, then `k = 1`. Otherwise, `k` is either `2` or `3`, and in both cases
the set already contains `c`.

### How to find c

The query

```text
Q y
```

returns how many values in `S` are at least `y`.

During the algorithm we also know some values that are already in `S`, for example the initial `0`, and
later the recovered `c`. The code keeps these known values in a local map called `set`. When it asks
`Q y`, it subtracts the known values that are at least `y`. After this subtraction, the answer tells us
whether there is still some unknown value in `S` that is at least `y`.

When the only unknown value is `c`, the adjusted query has this meaning:

```text
query(y) > 0  <=>  c >= y
```

That is enough to reconstruct `c` bit by bit.

The function `guess()` builds `c` from the highest bit to the lowest bit. Suppose the already fixed
higher bits are stored in `pre`. For bit `i`, it tries setting this bit:

```text
candidate = pre | (1 << i)
```

If `c >= candidate`, then `c` must have enough value to include this bit, so we keep it:

```text
if query(candidate) > 0:
    pre = candidate
```

Otherwise this bit is `0`. Repeating this from bit `n - 1` down to `0` recovers the exact value of `c`.

### Case k = 1

If `I 0` returned `1`, then we already know:

```text
k = 1
```

But `c` is not in the set yet, because `0 & c = 0`. To make `c` appear, insert the all-ones mask:

```text
I (2^n - 1)
```

For `AND`:

```text
(2^n - 1) & c = c
```

Now the set is `{0, c}`, so `guess()` recovers `c`. The answer is `(1, c)`.

### Case k = 2 or k = 3

If `I 0` returned `2`, then the inserted value was `c`, because:

```text
0 | c = c
0 ^ c = c
```

So the set is already `{0, c}`. Run `guess()` to recover `c`, and record both `0` and `c` as known set
values.

Now only `k` remains unknown.

#### When c is all ones

Let:

```text
full = 2^n - 1
```

If `c = full`, then inserting `full` is not useful, because both operations can create already known
values. The implementation instead inserts `1`:

```text
I 1
```

For OR:

```text
1 | full = full
```

This is just `c`, already present in the set, so the set size does not change.

For XOR:

```text
1 ^ full = full - 1
```

This is a new value, so the set size increases.

Therefore:

```text
set size changed     -> k = 3
set size unchanged   -> k = 2
```

#### When c is not all ones

If `c != full`, insert `full`:

```text
I full
```

For OR:

```text
full | c = full
```

So `full` is inserted.

For XOR:

```text
full ^ c
```

This is the bitwise complement of `c` inside `n` bits. Since `c != full`, this value is not `full`.

Now ask:

```text
Q full
```

After subtracting the already known values `{0, c}`, the adjusted answer is positive exactly when `full`
was inserted. Therefore:

```text
query(full) > 0   -> k = 2
query(full) = 0   -> k = 3
```

### Why the number of operations is enough

The algorithm uses:

1. one initial output `0`;
2. one insertion `I 0`;
3. at most `n` threshold queries inside `guess()`;
4. one extra insertion to create `c` or distinguish the operation;
5. sometimes one final query `Q full`.

This is at most `n + 4` interactions before the final answer, well below the allowed `n + 30`.

### Interactive I/O detail

Every output command must be flushed immediately. Otherwise the judge may never receive the command, and
the submission can fail with `Idleness limit exceeded`.

The implementation uses one `bufio.Writer` and calls `Flush()` after:

1. the initial number `a`;
2. every `I x` insertion;
3. every `Q y` query;
4. the final answer `A k c`.

It also exits immediately if the judge returns `-1`, which means the interaction has already failed.

### Complexity

The local work is `O(n)` per test case, because `guess()` checks each bit once. The number of interactive
operations is also `O(n)`.
