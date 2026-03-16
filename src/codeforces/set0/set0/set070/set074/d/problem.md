### Problem

There is a cloakroom with a coat hanger represented by `n` hooks in a row, numbered from `1` to `n` (left to right).

At the beginning of the day, all hooks are empty. During the day:

- Employees arrive and hang their coats.
- Employees leave and take their coats.

**Arrival rule (where to hang the coat):**

1. Consider all maximal contiguous segments of **empty** hooks.
2. Choose the segment with the **maximum length**.
3. If there are multiple such segments, choose the one **closest to the right**.
4. Hang the coat at the **middle hook** of this segment:
   - If the segment length is odd, take the unique middle hook.
   - If the segment length is even, take the **right** of the two middle hooks.

**Departure rule:**

- When an employee leaves, they take their own coat, freeing that hook.
- Employee identifiers are unique; nobody ever takes someone else’s coat.

From time to time, the director asks:

- “How many coats are currently hanging on hooks from `i` to `j` (inclusive)?”

You must simulate the whole process and answer all such queries.

### Input

- First line: two integers `n`, `q`  
  - `1 ≤ n ≤ 10^9` — number of hooks  
  - `1 ≤ q ≤ 10^5` — number of requests

- Next `q` lines: requests in chronological order.

Each request is one of:

- `0 i j` (`1 ≤ i ≤ j ≤ n`) — director’s query:  
  “How many coats hang on hooks `i..j` right now?”

- `<id>` (a single positive integer, `1 ≤ id ≤ 10^9`) — employee with identifier `id`:
  - The **1st**, **3rd**, **5th**, ... occurrence of this `id` is that employee’s **arrival**.
  - The **2nd**, **4th**, **6th**, ... occurrence is that employee’s **departure**.

Additional guarantees:

- Every employee id is distinct.
- Whenever an employee arrives, there is at least one empty hook.
- There is at least one director’s query of type `0 i j` in the input.

### Output

For each director’s request `0 i j`, print a single integer on its own line —  
the number of coats hanging on hooks from `i` to `j` (inclusive) at that moment.

### Example

**Input**
```text
9 11
1
2
0 5 8
1
1
3
0 3 8
9
0 6 9
6
0 1 9
```

**Output**
```text
2
3
2
5
```

### Summary / Notes

- You must dynamically maintain:
  - The set of free segments of hooks to place new coats according to the “longest, rightmost, middle” rule.
  - The mapping from employee id → hook index they currently occupy.
  - A data structure that supports:
    - Point updates when a coat is hung or removed.
    - Range queries `count of occupied hooks in [i, j]`.
- Constraints (`n` up to `10^9`, `q` up to `10^5`) imply:
  - You cannot store all hooks explicitly.
  - You need balanced trees / priority structures for segments and something like a Fenwick tree or segment tree over compressed coordinates for query ranges.


### ideas
1. 这里n是很大的。所以还不能简单的用n去开数组