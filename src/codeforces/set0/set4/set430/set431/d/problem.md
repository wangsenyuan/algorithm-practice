One day, after a difficult lecture a diligent student Sasha saw a graffitied desk in the classroom. She came closer and read: "Find such positive integer n, that among numbers n + 1, n + 2, ..., 2·n there are exactly m numbers which binary representation contains exactly k digits one".

The girl got interested in the task and she asked you to help her solve it. Sasha knows that you are afraid of large numbers, so she guaranteed that there is an answer that doesn't exceed 10^18.

## Input

The first line contains two space-separated integers, m and k (0 ≤ m ≤ 10^18; 1 ≤ k ≤ 64).

## Output

Print the required number n (1 ≤ n ≤ 10^18). If there are multiple answers, print any of them.

## Examples

**Input:**
```
1 1
```

**Output:**
```
1
```

**Input:**
```
3 2
```

**Output:**
```
5
```


### ideas
1. 先找出 1....某个数 恰好包含m个数，有k个digit one的范围（这个可以用二分）
2. 然后找到 n = 1 << h, 2 * n 包含这些数就可以了
3. 好像有些特殊情况要判断