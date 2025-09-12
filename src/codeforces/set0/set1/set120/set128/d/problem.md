# Problem Description

One day Anna got the following task at school: to arrange several numbers in a circle so that any two neighboring numbers differs exactly by 1. Anna was given several numbers and arranged them in a circle to fulfill the task. Then she wanted to check if she had arranged the numbers correctly, but at this point her younger sister Maria came and shuffled all numbers. Anna got sick with anger but what's done is done and the results of her work had been destroyed. But please tell Anna: could she have hypothetically completed the task using all those given numbers?

## Input

The first line contains an integer n — how many numbers Anna had (3 ≤ n ≤ 10^5). The next line contains those numbers, separated by a space. All numbers are integers and belong to the range from 1 to 10^9.

## Output

Print the single line "YES" (without the quotes), if Anna could have completed the task correctly using all those numbers (using all of them is necessary). If Anna couldn't have fulfilled the task, no matter how hard she would try, print "NO" (without the quotes).

## Examples

### Example 1

**Input:**
```
4
1 2 3 2
```

**Output:**
```
YES
```

### Example 2

**Input:**
```
6
1 1 2 2 2 3
```

**Output:**
```
YES
```

### Example 3

**Input:**
```
6
2 4 1 1 2 2
```

**Output:**
```
NO
```


### ideas
1. 如果答案存在，必然存在一个排列是升序，然后降序（但是中间会存在相同的数）
2. 假设有k个不同的数， 1,2 ...k
3. 先把12, 按照1212这样排列好如果， 1比2多的话 => false
4. 这样子会出现连续的11（这个在后续无法被消除掉）
5. 应该从大往下考虑k-1和k，如果k的个数超过k-1，是不行的，这样子会出现kk
6. 然后(k-1)k(k-1)k(k-1) 这样子可以把cnt[k] + 1个(k-1)连接在一起作为一个单独的k-1 (这个k-1)当作一个k-1
7. 或者把1个k去绑定2个k-1，这样子k-1就变成了一个范围
8. 假设原来有cnt[k-1]个数，现在 lo = cnt[k-1] - cnt[k], hi = cnt[k-1] - 2 * cnt[k]
9. 假设有9个k-1， 3个k，  要么当作   xyx  xyx, xyx, x, x, x hi = 6
10. 要么当作 xyxyxyx   x x x x x  hi = 6?
11. cnt[k-1] - 2 * cnt[k] = （剩下的k-1个） + cnt[k] (全部当作k-1) = cnt[k-1] - cnt[k] 
12. cnt[k-1] - (cnt[k] + 1) + 1 = cnt[k-1] - cnt[k]
13. 所以两种处理是相当的
