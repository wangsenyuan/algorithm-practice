You are given a set Y of n distinct positive integers y1, y2, ..., yn.

Set X of n distinct positive integers x1, x2, ..., xn is said to generate set Y if one can transform X to Y by applying some number of the following two operation to integers in X:

1. Take any integer xi and multiply it by two, i.e. replace xi with 2·xi.
2. Take any integer xi, multiply it by two and add one, i.e. replace xi with 2·xi + 1.

Note that integers in X are not required to be distinct after each operation.

Two sets of distinct integers X and Y are equal if they are equal as sets. In other words, if we write elements of the sets in the array in the increasing order, these arrays would be equal.

Note, that any set of integers (or its permutation) generates itself.

You are given a set Y and have to find a set X that generates Y and the maximum element of X is mininum possible.

## Input

The first line of the input contains a single integer n (1 ≤ n ≤ 50 000) — the number of elements in Y.

The second line contains n integers y1, ..., yn (1 ≤ yi ≤ 10^9), that are guaranteed to be distinct.

## Output

Print n integers — set of distinct integers that generate Y and the maximum element of which is minimum possible. If there are several such sets, print any of them.

## Examples

### Example 1

**Input:**
```
5
1 2 3 4 5
```

**Output:**
```
4 5 2 3 1
```

### Example 2

**Input:**
```
6
15 14 3 13 1 12
```

**Output:**
```
12 13 14 7 3 1
```

### Example 3

**Input:**
```
6
9 7 13 17 5 11
```

**Output:**
```
4 5 2 6 3 1
```


### ideas
1. 感觉和树很像，
2. 把y作为叶子节点，y[i] / 2 是它的父节点
3. x是这棵树上的中间节点，从x[i]...y[i]之间不能有分叉