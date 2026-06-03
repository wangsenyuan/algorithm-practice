# G. Buratsuta 3

[Problem link](https://codeforces.com/problemset/problem/2149/G)

time limit per test: 4.5 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

In the ruthless world of Blue Lock, Buratsuta 3 is a trio selected to overthrow the
reigning champions and lead the Japan U-20 team to glory. Sae Itoshi has already
secured his place as the first participant; the remaining two spots will be
contested in the tough Side-B selection.

To test the strategic abilities of the candidates, Buratsuta has posed the following
challenge:

You are given an array of `n` integers called "performance records" and `q` queries.
Each query specifies a subarray `[l, r]`. In this subarray, find all record values
that occur strictly more than `floor((r - l + 1) / 3)` times.

Output those values in sorted order. If there are no such values, output `-1`.

## Input

Each test consists of several test cases.

The first line contains a single integer `t` (`1 <= t <= 10^4`) — the number of test
cases. The following describes the test cases.

The first line of each test case contains two integers `n` and `q`
(`1 <= n, q <= 2 * 10^5`) — the number of records and the number of queries.

The second line of each test case contains `n` integers `a1, a2, ..., an`
(`1 <= ai <= 10^9`) — the performance records.

Each of the following `q` lines contains two integers `l` and `r`
(`1 <= l <= r <= n`) — the boundaries of the query.

It is guaranteed that the sum of `n` and the sum of `q` over all test cases does
not exceed `2 * 10^5`.

## Output

For each query, output in one line all record values (in sorted order) that occur
strictly more than `floor((r - l + 1) / 3)` times in the segment `[l, r]`. If there
are no such values, output `-1`.

Equivalently, output all values whose frequency in `[l, r]` is at least
`floor((r - l + 1) / 3) + 1`.

## Example

### Input

```text
5
1 1
5
1 1
4 2
1 1 2 3
1 4
2 3
6 3
7 7 7 8 8 9
1 6
2 5
4 6
8 2
4 4 4 5 5 5 6 6
1 8
3 6
10 5
1 2 3 3 3 4 4 4 4 5
1 10
1 5
4 9
6 9
7 10
```

### Output

```text
5
1
1 2
7
7 8
8
4 5
5
4
3
4 4 4
```

## Note

In the second test case, the array is `a = [1, 1, 2, 3]` and there are two queries:

- Query `(l, r) = (1, 4)`: The length of the segment is `len = r - l + 1 = 4`, and
  the threshold is `floor(len / 3) + 1 = 2`. Occurrences: `1 -> 2`, `2 -> 1`,
  `3 -> 1`. Only the number `1` occurs at least `2` times, so the answer is `1`.
- Query `(l, r) = (2, 3)`: The length is `len = 2`, and the threshold is
  `floor(len / 3) + 1 = 1`. Numbers `1` and `2` occur once each, so the answer is
  `1 2`.

In the fourth test case, the array is `a = [4, 4, 4, 5, 5, 5, 6, 6]` and there are
two queries:

- Query `(l, r) = (1, 8)`: The length is `len = 8`, and the threshold is
  `floor(len / 3) + 1 = 3`. Occurrences: `4 -> 3`, `5 -> 3`, `6 -> 2`. Only the
  numbers `4` and `5` occur at least `3` times, so the answer is `4 5`.
- Query `(l, r) = (3, 6)`: The length is `len = 4`, and the threshold is
  `floor(len / 3) + 1 = 2`. Occurrences: `4 -> 1`, `5 -> 3`. Only the number `5`
  occurs at least `2` times, so the answer is `5`.


### ideas
1. will sqrt(n) * q * log(n) work?
2. 还有就是，似乎只可能有2个数满足条件。
3. 所以，似乎log(n)可以不用？只用记录最多的数，和第二多的数即可
4. 只保留两个数，它被删除后，就不一定加回来了
5. 所以这个方式不大对～
6. 可以确定的是，答案只能是[0, 1, 2]
7. 2e5 * 300 * 20 差不多 1e9了。肯定不行
8. 如何快速的知道一个区间内的最多的数呢？
9. 好像有个数据结构可以知道，一个数在一个区间内的rank
10. 

## Misra-Gries Heavy Hitters Explanation

Misra-Gries is a deterministic heavy hitters algorithm. It finds a small set of
possible values whose frequency may exceed `n / (k + 1)` while keeping only `k`
counters.

For this problem, the threshold is `> len / 3`, so use `k = 2`. Any query range
can contain at most two values that occur more than one third of the range.

The streaming version keeps at most `k` candidate values:

1. If the current value is already a candidate, increment its counter.
2. Otherwise, if there is an empty slot, add it with counter `1`.
3. Otherwise, there are already `k` candidates and this is a new `(k + 1)`-th
   distinct value. Decrement all counters by `1` and remove zero counters.

For `k = 2`:

```text
stream: 1 2 3
after 1: {1:1}
after 2: {1:1, 2:1}
after 3: cancel one occurrence of 1, 2, and 3 -> {}
```

This cancellation is equivalent to deleting one occurrence of `k + 1` distinct
values. A value whose frequency is strictly greater than `n / (k + 1)` cannot be
fully deleted by these cancellations, so it must remain as a candidate.

Misra-Gries only gives candidates. It does not prove that every candidate is a
real heavy hitter. The implementation still verifies each candidate by counting
its real occurrences in `[l, r]` using the positions map and binary search.

### Segment Tree Merge

In a normal stream, every incoming item has weight `1`. In a segment tree, each
node stores a compressed Misra-Gries summary, so each candidate has a residual
weight:

```text
[{value: 4, count: 3}, {value: 7, count: 2}]
```

When merging two segment summaries, we are merging weighted candidates, not raw
single elements.

The correct merge rule is:

1. Add counters for equal values.
2. If adding a new value makes the number of candidates exceed `k`, repeatedly:
   - find the minimum counter `d` among all current candidates,
   - subtract `d` from every counter,
   - remove candidates whose counter becomes zero.

The cancellation must happen only when there are more than `k` candidates. For
`k = 2`, two distinct candidates are valid and must be kept. Cancellation starts
only when a third distinct candidate appears.

The weighted subtraction is also important. Example:

```text
left  = {1:100, 2:100}
right = {3:50}
```

After adding, there are three candidates. The correct cancellation subtracts the
minimum counter `50`:

```text
{1:100, 2:100, 3:50}
-> subtract 50 from all
-> {1:50, 2:50}
```

Subtracting only `1` would leave `{1:99, 2:99, 3:49}`, which is not yet a valid
summary with at most two candidates.

### Query Flow

The range query works as follows:

1. Build a segment tree where every node stores a Misra-Gries summary with
   `k = 2`.
2. For query `[l, r]`, merge the `O(log n)` node summaries covering the range.
3. The merged result contains at most two candidates.
4. Verify each candidate by binary searching its occurrence positions.
5. Output the verified candidates in sorted order, or `-1` if none pass.

The complexity is:

```text
Build: O(n) because k = 2 is constant
Query summary merge: O(log n)
Candidate verification: O(log n)
Total per query: O(log n)
```
