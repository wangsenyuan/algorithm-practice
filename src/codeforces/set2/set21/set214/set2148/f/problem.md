Farmer John has `n` arrays `a1, a2, ..., an` that may have different lengths. He will stack the arrays on top of each other, resulting in a grid with `n` rows. The arrays are left-aligned and can be stacked in any order he desires.

Then, gravity takes place. Any cell that is not on the bottom row and does not have an element beneath it will fall down one row. This process repeats until no cell satisfies that condition.

Over all possible ways FJ can stack the arrays, output the lexicographically minimum bottom row after gravity finishes.

## Input

The first line contains `t` (`1 <= t <= 1000`) — the number of test cases.

The first line of each test case contains `n` (`1 <= n <= 2 * 10^5`).

For the following `n` lines, the first integer `ki` (`1 <= ki <= 2 * 10^5`) denotes the length of `ai`.

Then `ki` space-separated integers `ai1, ai2, ..., aiki` follow (`1 <= aij <= 2 * 10^5`).

It is guaranteed that the sum of `n` and the sum of all `ki` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output the lexicographically minimum bottom row on a new line.

## Example

### Input

```text
4
1
3 5 2 7
2
2 2 9
3 3 1 4
3
1 5
2 5 1
2 5 2
3
3 4 4 9
7 7 6 5 4 3 2 1
4 2 4 5 1
```

### Output

```text
5 2 7
2 9 4
5 1
2 4 5 1 3 2 1
```


### ideas
1. 如果a[i]是第0列最小的，那么就把它放在最小面，如果有多个最小的，那么就需要在它们中间挑选
2. 挑选的逻辑，也是一直找下去。但是假设某个a[i]比较短，这个时候怎么处理？
3. 这时候，那些长度 > 它的，都可以往下掉。所以，要把它们也全部算进来，继续处理
4. 

## Solution explanation

Think about the final bottom row column by column.

After gravity finishes, for a fixed column `d`, only arrays with length `> d` have an element in that column. Among those arrays, the bottom cell of column `d` comes from the lowest stacked array that still has an element in column `d`.

So choosing a stacking order is equivalent to choosing a priority order of arrays from bottom to top:

- for column `d`, scan the priority order from bottom to top;
- the first array whose length is greater than `d` supplies the bottom-row value at column `d`.

Thus we need to choose this priority order so that the produced sequence is lexicographically minimum.

## Greedy choice

Suppose we are currently deciding answer position `d`.

All arrays with length `<= d` are already irrelevant, because they do not have a value at this column. Among all remaining arrays, the first answer value should clearly be the minimum possible `a[i][d]`.

If only one array has this minimum value, then it must be placed below all other currently active arrays, because using any larger value would make the answer lexicographically worse immediately.

If several arrays tie on `a[i][d]`, then this column cannot distinguish them. We compare their next values:

`a[i][d+1]`, then `a[i][d+2]`, and so on.

This is exactly a lexicographic comparison of the suffixes starting at column `d`, except for one important detail: arrays can have different lengths.

## What happens when one tied array ends

Assume several candidate arrays are equal from column `d` up to column `e-1`, and one of them ends at `e`.

If we put that shorter array at the bottom among these tied candidates, then it contributes the equal prefix `d..e-1`. Starting at column `e`, it has no cell, so gravity exposes the next lowest active array with length `> e`.

This is never worse than choosing a longer tied array first, because the already written prefix is identical, and ending earlier gives us the chance to choose the best continuation from all longer active arrays.

That is why in the code, if during comparison `len(a[i]) == d`, `find` immediately returns this row and this stopping position:

```go
if len(a[i]) == d {
	return []int{i, d}
}
```

The caller then copies the selected row only up to that ending position, removes all arrays that are no longer active, and continues.

## Matching the code

First, rows are sorted by length:

```go
slices.SortFunc(rows, func(i int, j int) int {
	return cmp.Or(len(a[i])-len(a[j]), i-j)
})
```

This lets `pos` skip all rows that are too short for the current column:

```go
for pos < n && len(a[rows[pos]]) <= d {
	pos++
}
```

Now `rows[pos:]` contains exactly the arrays that still have a value at column `d`.

The helper:

```go
find(rows[pos:], d)
```

chooses which row should contribute the next segment of the answer.

Inside `find`:

1. Among all candidate rows, find the minimum value in the current column `d`.
2. Keep only rows with that minimum value.
3. Move to column `d + 1` and repeat.
4. Stop when either only one row remains, or some tied row ends.

If one row remains, it is lexicographically best among the active candidates, so it can safely provide values until it ends:

```go
return []int{rows[0], len(a[rows[0]])}
```

The main loop receives `(r, d1)` and copies:

```go
for i := d; i < d1; i++ {
	ans[i] = a[r][i]
}
```

Then it sets:

```go
d = d1
```

and repeats the process for the next still-empty column.

## Why this is correct

At each current position `d`, lexicographic order says the smallest possible value at `d` is the only thing that matters first. Therefore all arrays with non-minimum `a[i][d]` can be ignored for the current bottom choice.

If there is a tie, the same argument applies to the next column among the tied rows. This continues until:

- a unique best row is found; or
- a tied row ends.

In the first case, the unique row has the lexicographically smallest suffix, so choosing it is optimal.

In the second case, all tied rows produce the same segment up to the shorter row's end. Choosing the shorter row first cannot make the written segment worse, and after it ends, all longer rows are still available for the remaining columns. So it is optimal to stop there and re-run the same greedy choice on the remaining active rows.

Since every step fixes the lexicographically smallest possible next segment, the whole answer is lexicographically minimum.

## Complexity

Each call to `find` advances column comparisons while filtering candidate rows by the current minimum value. The total amount of row data processed over the test case is bounded by the total input size up to sorting/filtering overhead.

The implementation also sorts rows by length, so the complexity is:

`O(total length + n log n)` per test case.
