# Problem

There is a variable `sum`, initially `0`.

There is also a data structure that supports:

- `pushback x` — append an element with value `x`;
- `pushfront x` — prepend an element with value `x`;
- `popback` — remove the last element;
- `popfront` — remove the first element;
- `min` — add the current minimum element in the structure to `sum`.

Operations `popback`, `popfront`, and `min` must not be used on an empty structure.

Using this structure, you want to produce a sequence of commands so that after all operations, `sum` equals the sum of minima over **all non-empty subarrays** of an array `a` of `n` elements.

More formally: find at most `n * (n + 2)` commands such that for **any** array `a` of length `n`, after executing them,

`sum = sum_{0 <= l <= r < n} min(a[l], ..., a[r])`.

## Input

The first line contains a single integer `n` (`1 <= n <= 500`) — the length of the array.

## Output

Output `k` (`1 <= k <= n * (n + 2)`) lines — each command is one of:

- `pushback a[i]`, where `i` is in `0 .. n-1`;
- `pushfront a[i]`, where `i` is in `0 .. n-1`;
- `popback`;
- `popfront`;
- `min`.

If there are multiple valid answers, output any.

## Example

### Example 1

**Input**

```text
1
```

**Output**

```text
3
pushback a[0]
min
popfront
```

### Example 2

**Input**

```text
2
```

**Output**

```text
6
pushfront a[1]
min
pushback a[0]
min
popfront
min
```


### ideas
1. 也就是要在不超过n*(n+2)个操作里， 得到sum(min(l..r)) (l <= r)
2. n * (n + 2) = n * n + 2 * n
3. 