# D. Make a Palindrome

[Problem link](https://codeforces.com/problemset/problem/2124/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

You are given an array `a` of size `n` and an integer `k`. You may perform the
following operation any number of times:

1. Choose two integers `l` and `r` (`1 <= l <= r <= |a|`) such that `r - l + 1 >= k`.
2. Choose an index `i` with `l <= i <= r`, where `a_i` is the `k`-th smallest value
   in the subarray `a[l], a[l+1], ..., a[r]`. If several indices are valid, choose
   any one of them.
   - Example: for `a = [1, 2, 2, 1, 3]`, `l = 1`, `r = 5`, `k = 3`, valid indices
     are `2` and `3`.
3. Delete `a_i` from `a` and concatenate the remaining parts.

Determine whether it is possible to obtain a palindrome after any number of
operations. An empty array is considered a palindrome.

An array `b = [b_1, b_2, ..., b_m]` is a palindrome if `b_i = b_{m+1-i}` for every
`1 <= i <= m`.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test
cases.

The first line of each test case contains two integers `n` and `k`
(`1 <= k <= n <= 2 * 10^5`).

The second line contains `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= n`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print `YES` if a palindrome can be created, and `NO`
otherwise. You may use any letter case.

## Example

### Input

```text
8
5 3
5 4 3 4 5
4 1
1 1 2 1
6 6
2 3 4 5 3 2
5 4
5 2 4 3 1
8 5
4 7 1 2 3 1 3 4
5 4
1 2 1 2 2
3 3
1 2 2
4 4
2 1 2 2
```

### Output

```text
YES
YES
YES
NO
NO
YES
NO
YES
```

## Note

In the first test case, `a` is already a palindrome.

In the second test case, one sequence of operations is:

```text
[1, 1, 2, 1] -> [1, 2, 1] -> [1, 1]
```

In the third test case, one operation suffices:

```text
[2, 3, 4, 5, 3, 2] -> [2, 3, 4, 3, 2]
```

In the fourth and fifth test cases, no sequence of operations can produce a
palindrome.

## Solution

The operation always deletes the `k`-th smallest value inside a chosen subarray.
So the key is which values can never be removed, and how many copies of the
threshold value `w` we can afford to delete while pairing the rest into a
palindrome.

### Values that cannot be deleted

Let `w` be the `k`-th smallest value in the whole array (after sorting).

- Every value **strictly less than `w`** is always among the `k - 1` smallest
  elements inside any long enough subarray that contains it, so it can never be
  deleted.
- Every value **greater than `w`** can be deleted (take the whole current array;
  those values are never the `k`-th smallest).
- Values **equal to `w`** may be deleted when we choose a subarray where `w` is
  exactly the `k`-th smallest.

So only elements `<= w` matter for the final shape. Build `a1` by keeping the
original-order subsequence of values `<= w`.

### Two-pointer pairing on `a1`

Try to form a palindrome from `a1` using only deletions of value `w`:

- Use two pointers `i` (left) and `j` (right).
- If `a1[i] == a1[j]`, they can mirror each other. If the common value is `w`,
  count how many `w` are used in such pairs (`cnt`).
- If `a1[i] != a1[j]`:
  - If both are `< w`, answer is `NO` (fixed elements cannot be removed or
    rearranged).
  - Otherwise delete one `w` by moving the pointer on the `w` side inward.

This greedy scan checks whether the undeletable part can be matched into a
palindrome once enough flexible `w` copies are removed.

### Final check

Let `l` be the number of elements strictly less than `w` in the sorted array.
At least `k - l` copies of `w` exist globally, and the pairing process needs
enough matched `w` to cover the deletions.

Answer is `YES` iff the two-pointer scan does not fail early and:

```text
cnt >= k - l - 1
```

### How the code maps to this

1. Sort a copy of `a` to get `w = sorted[k-1]` and `l = count of values < w`.
2. Filter `a` into `a1` keeping only values `<= w`.
3. Run the two-pointer procedure, updating `cnt` for matched `w` pairs.
4. Return whether the scan succeeds and `cnt >= k - l - 1`.

### Complexity

- Sorting: `O(n log n)`
- Filtering and two pointers: `O(n)`
- Total per test case: `O(n log n)`