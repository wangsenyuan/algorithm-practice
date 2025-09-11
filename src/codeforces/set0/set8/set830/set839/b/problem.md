# Problem: Army Placement

## Description

Daenerys Targaryen has an army consisting of k groups of soldiers, the i-th group contains ai soldiers. She wants to bring her army to the other side of the sea to get the Iron Throne. She has recently bought an airplane to carry her army through the sea. The airplane has n rows, each of them has 8 seats. We call two seats neighbor, if they are in the same row and in seats {1, 2}, {3, 4}, {4, 5}, {5, 6} or {7, 8}.

**A row in the airplane:**
```
[1] [2] [3] [4] [5] [6] [7] [8]
```

Daenerys Targaryen wants to place her army in the plane so that there are no two soldiers from different groups sitting on neighboring seats.

Your task is to determine if there is a possible arranging of her army in the airplane such that the condition above is satisfied.

## Input

The first line contains two integers n and k (1 ≤ n ≤ 10000, 1 ≤ k ≤ 100) — the number of rows and the number of groups of soldiers, respectively.

The second line contains k integers a1, a2, a3, ..., ak (1 ≤ ai ≤ 10000), where ai denotes the number of soldiers in the i-th group.

It is guaranteed that a1 + a2 + ... + ak ≤ 8·n.

## Output

If we can place the soldiers in the airplane print "YES" (without quotes). Otherwise print "NO" (without quotes).

You can choose the case (lower or upper) for each letter arbitrary.

## Examples

### Example 1
**Input:**
```
2 2
5 8
```
**Output:**
```
YES
```

### Example 2
**Input:**
```
1 2
7 1
```
**Output:**
```
NO
```

### Example 3
**Input:**
```
1 2
4 4
```
**Output:**
```
YES
```

### Example 4
**Input:**
```
1 4
2 2 1 2
```
**Output:**
```
YES
```

## Note

In the first sample, Daenerys can place the soldiers like in the figure below:

In the second sample, there is no way to place the soldiers in the plane since the second group soldier will always have a seat neighboring to someone from the first group.

In the third example Daenerys can place the first group on seats (1, 2, 7, 8), and the second group on all the remaining seats.

In the fourth example she can place the first two groups on seats (1, 2) and (7, 8), the third group on seats (3), and the fourth group on seats (5, 6).


### ideas
1. sum(a[i]) + waste <= 8 * n
2. 按照a[i] % 4 进行分组，所有 = 0的部分，可以放在最中间，如果有多出来的，就分到两边去
3. rem = 1 的部分， 这样子的肯定会造成一个浪费，rem = 3 的也必然造成一个浪费
4. rem = 2 的部分， 如果少于 2 * n, 那么也不会造成浪费，如果超过了， 就会占用中间的部分，
5. 但是它们可以和1的部分去匹配