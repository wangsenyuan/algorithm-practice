# Problem

This is an interactive problem.

There are `n` words in a text editor. The `i`-th word has length `li` (`1 <= li <= 2000`).
The array `l` is hidden and known only by the grader.

The editor displays words in lines, placing at least one space between adjacent words in the same line. A line does not need to end with a space.
For a fixed width, words are arranged to minimize the number of used lines.

More formally, for width `w`, let `a` be an array of length `k + 1` such that:

`1 = a1 < a2 < ... < a(k+1) = n + 1`

Array `a` is valid if for each `1 <= i <= k`, words `a_i ... a_(i+1)-1` fit into one line:

`l[a_i] + l[a_i+1] + ... + l[a_(i+1)-1] + ((a_(i+1)-a_i)-1) <= w`

The editor height is the minimum possible `k` over all valid arrays.

If `w < max(li)`, the editor crashes and the height is defined as `0`.

You can ask at most `n + 30` queries. In one query, you provide width `w`, and the grader returns `h_w` (the resulting height).

Find the minimum editor area:

`min(w * h_w)` over all `w` such that `h_w != 0`.

The hidden lengths are fixed in advance (the interactor is not adaptive).

## Input

The first and only line contains integer `n` (`1 <= n <= 2000`) — the number of words.

It is guaranteed that hidden `li` satisfy `1 <= li <= 2000`.

## Interaction

Begin by reading `n`.

To make a query, print:

`? w`

where `1 <= w <= 10^9`.

Then read the response `h_w`.

If your program makes an invalid query or exceeds the query limit, the interactor terminates immediately and your program gets `Wrong answer`.

To output the final answer, print:

`! area`

This final output is not counted toward the `n + 30` query limit.

After each query, print a newline and flush output; otherwise, you may get `Idleness limit exceeded`.

Use, for example:

- `fflush(stdout)` or `cout.flush()` in C++
- `System.out.flush()` in Java
- `flush(output)` in Pascal
- `stdout.flush()` in Python

## Hacks

In hack format (non-interactive):

- First line contains `n` (`1 <= n <= 2000`)
- Second line contains `n` space-separated integers `l1, l2, ..., ln` (`1 <= li <= 2000`)

## Example

**Input**

```text
6

0

4

2
```

**Output**

```text
? 1

? 9

? 16

! 32
```

## Note

In this sample, words are `{glory, to, ukraine, and, anton, trygub}`, so:

`l = {5, 2, 7, 3, 5, 6}`.

If `w = 1`, the editor crashes, so `h1 = 0`.

If `w = 9`, one possible layout is:

```text
glory  to
ukraine
and anton
  trygub
```

So `h9 = 4`.

If `w = 16`, one possible layout is:

```text
glory to ukraine
and anton trygub
```

So `h16 = 2`.

The minimum area is `32`.


## ideas
1. 最小的w可以通过最多20次，二分查询出来（也是最大的l[i])
2. 那么如何找到最优的w * h[w] 呢？
3. 对于给定的w, 肯定存在一列 l[a[1]] + l[a[2]] + ... + l[a[k]] + k - 1 = w
4. 假设对于(w1, h1), (w2, h2), 且 w1 < w2, h1 > h2
5. 是不是计算出w1(最小的h > 0)， w2(最小的宽度, h = 1)
6. 那么 sum(l) = w2 - (n - 1)
7. 那么 avg = sum(l) / n
8. 是不是用avg递进查询？

## key insights

1. The solution first finds `S`, the minimum width such that `h(S) = 1`, using binary search:
   - search range: `[1, n * 2000 + n - 1]`
   - predicate: `ask(w) == 1`
   - interpretation: `S = sum(li) + (n - 1)` (all words in one line).

2. `S` gives an immediate valid area candidate (`S * 1 = S`), so initialize:
   - `best = S`.

3. Then iterate possible target line counts by probing widths near `S / h`:
   - for `h = 2..n`
   - query `w = S / h`
   - read `tmp = ask(w)`
   - if `tmp == 0`, widths are already too small (editor crashes), so stop.
   - otherwise update `best = min(best, w * tmp)`.

4. Why this is efficient:
   - binary search costs about `log2(4e6) <= 22` queries,
   - loop costs at most `n - 1` queries,
   - total `<= n + 22`, which is within the limit `n + 30`.

5. Final answer is the minimum observed area `best`.
