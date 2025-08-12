# Problem A: Hay Blocks Theft

## Problem Statement

Once upon a time in the Kingdom of Far Far Away lived Sam the Farmer. Sam had a cow named Dawn and he was deeply attached to her. Sam would spend the whole summer stocking hay to feed Dawn in winter. Sam scythed hay and put it into haystack. As Sam was a bright farmer, he tried to make the process of storing hay simpler and more convenient to use. He collected the hay into cubical hay blocks of the same size. Then he stored the blocks in his barn. After a summer spent in hard toil Sam stored A·B·C hay blocks and stored them in a barn as a rectangular parallelepiped A layers high. Each layer had B rows and each row had C blocks.

At the end of the autumn Sam came into the barn to admire one more time the hay he'd been stacking during this hard summer. Unfortunately, Sam was horrified to see that the hay blocks had been carelessly scattered around the barn. The place was a complete mess. As it turned out, thieves had sneaked into the barn. They completely dissembled and took away a layer of blocks from the parallelepiped's front, back, top and sides. As a result, the barn only had a parallelepiped containing (A - 1) × (B - 2) × (C - 2) hay blocks. To hide the evidence of the crime, the thieves had dissembled the parallelepiped into single 1 × 1 × 1 blocks and scattered them around the barn. After the theft Sam counted n hay blocks in his barn but he forgot numbers A, B and C.

Given number n, find the minimally possible and maximally possible number of stolen hay blocks.

## Input

The only line contains integer n from the problem's statement (1 ≤ n ≤ 10^9).

## Output

Print space-separated minimum and maximum number of hay blocks that could have been stolen by the thieves.

**Note:** The answer to the problem can be large enough, so you must use the 64-bit integer type for calculations. Please, do not use the %lld specificator to read or write 64-bit integers in С++. It is preferred to use cin, cout streams or the %I64d specificator.

## Examples

### Example 1
**Input:**
```
4
```

**Output:**
```
28 41
```

### Example 2
**Input:**
```
7
```

**Output:**
```
47 65
```

### Example 3
**Input:**
```
12
```

**Output:**
```
48 105
```

## Explanation

Let's consider the first sample test. If initially Sam has a parallelepiped consisting of 32 = 2 × 4 × 4 hay blocks in his barn, then after the theft the barn has 4 = (2 - 1) × (4 - 2) × (4 - 2) hay blocks left. Thus, the thieves could have stolen 32 - 4 = 28 hay blocks. 

If Sam initially had a parallelepiped consisting of 45 = 5 × 3 × 3 hay blocks in his barn, then after the theft the barn has 4 = (5 - 1) × (3 - 2) × (3 - 2) hay blocks left. Thus, the thieves could have stolen 45 - 4 = 41 hay blocks. 

No other variants of the blocks' initial arrangement (that leave Sam with exactly 4 blocks after the theft) can permit the thieves to steal less than 28 or more than 41 blocks.


### ideas
1. 最多被偷走的，比较容易计算 = A = 2, 那么最后只剩一层
2. 那么这个时候，超过一半被偷走了, B * C = n, B = 1, 那么 C = n
3. 那么原来有 2 * (1 + 2) * (n + 2) - n
4. 那么最小被偷走多少呢？这里一共5个面，有3个值 x(前后), y(左右), z (顶部)
5. x = a * b, y = a * c, z = b * c
6. 在满足 (a - 1) * (b - 2) * (c - 2) = n 的情况下
7. 满足 x + y + z 最小
8. 感觉应该是 b 和 c越接近, s 越小, 似乎
9. 假设在确定a的情况下， 那么 (b - 2) * (c - 2) = n / (a - 1)
10. 如果只看 b * c (这个结果基本上被 n/a 限定了)
11. b * c - 2 * (b + c) + 4 = n / (a - 1)
12. b * c - 2 * (b + c) + (a + 2) * (b + c)  = n / (a - 1) - 4 + (a + 2) * (b + c)
13. => b * c + a * (b + c) = s
14. s = n / (a - 1) - 4 + (a + 2) * (b + c)
15. 那么就是 b + c 最小
16. 面积固定的情况下, n是体积， a是高度， (b * c)是面积
17. 圆的周长最短？比如面积为 5 * 5 = 25时， 周长为20, 2 * 12.5 周长为 29
18. 所以，确实是 b ~~ c 最好