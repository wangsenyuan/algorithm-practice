# Problem D

Recently Maxim has found an array of n integers, needed by no one. He immediately come up with idea of changing it: he invented positive integer x and decided to add or subtract it from arbitrary array elements. Formally, by applying single operation Maxim chooses integer i (1 ≤ i ≤ n) and replaces the i-th element of array ai either with ai + x or with ai - x. Please note that the operation may be applied more than once to the same position.

Maxim is a curious minimalis, thus he wants to know what is the minimum value that the product of all array elements (i.e. ∏ai) can reach, if Maxim would apply no more than k operations to it. Please help him in that.

## Input

The first line of the input contains three integers n, k and x (1 ≤ n, k ≤ 200 000, 1 ≤ x ≤ 10^9) — the number of elements in the array, the maximum number of operations and the number invented by Maxim, respectively.

The second line contains n integers a1, a2, ..., an (|ai| ≤ 10^9) — the elements of the array found by Maxim.

## Output

Print n integers b1, b2, ..., bn in the only line — the array elements after applying no more than k operations to the array. In particular, |bi - ai| should be divisible by x for every 1 ≤ i ≤ n, but the product of all array elements should be minimum possible.

If there are multiple answers, print any of them.

## Examples

### Example 1

**Input:**
```
5 3 1
5 4 3 5 2
```

**Output:**
```
5 4 3 5 -1
```

### Example 2

**Input:**
```
5 3 1
5 4 3 5 5
```

**Output:**
```
5 4 0 5 5
```

### Example 3

**Input:**
```
5 3 1
5 4 4 5 5
```

**Output:**
```
5 1 4 5 5
```

### Example 4

**Input:**
```
3 2 7
5 4 2
```

**Output:**
```
5 11 -5
```

### ideas
1. 如果最后的结果能变成负数，那么就应该优先变成负数，然后尽量的让结果更大（绝对值）
2. 如果不能变成负数，那么就应该尽可能的把最大值（绝对值）变小