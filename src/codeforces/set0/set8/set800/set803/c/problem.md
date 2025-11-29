You are given positive integer number n. You should create such strictly increasing sequence of k positive numbers a1, a2, ..., ak, that their sum is equal to n and greatest common divisor is maximal.

Greatest common divisor of sequence is maximum of such numbers that every element of sequence is divisible by them.

If there is no possible sequence then output -1.

## Input

The first line consists of two numbers n and k (1 ≤ n, k ≤ 1010).

## Output

If the answer exists then output k numbers — resulting sequence. Otherwise output -1. If there are multiple answers, print any of them.

## Examples

### Example 1

**Input:**
```
6 3
```

**Output:**
```
1 2 3
```

### Example 2

**Input:**
```
8 2
```

**Output:**
```
2 6
```

### Example 3

**Input:**
```
5 3
```

**Output:**
```
-1
```


### ideas
1. 严格递增
2. 先得到一个序列1....k, sum = (1 + k) * k / 2 
3. 如果 sum > n => -1
4. else 肯定存在
5. 假设最大值为g
6. g, g * 2, g * 3, ... g * k
7. => g = n / sum
8. let r = n - sum * g
9. 如果 r < g 肯定不行， r >= g, 且 r % g = 0 (这样子，可以分配给最后一个)