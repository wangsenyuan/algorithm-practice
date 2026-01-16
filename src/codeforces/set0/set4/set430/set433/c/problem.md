Though Ryouko is forgetful, she is also born with superb analyzing abilities. However, analyzing depends greatly on gathered information, in other words, memory. So she has to shuffle through her notebook whenever she needs to analyze, which is tough work.

Ryouko's notebook consists of n pages, numbered from 1 to n. To make life (and this problem) easier, we consider that to turn from page x to page y, \|x - y\| pages should be turned. During analyzing, Ryouko needs m pieces of information, the i-th piece of information is on page ai. Information must be read from the notebook in order, so the total number of pages that Ryouko needs to turn is the sum of absolute differences between consecutive pages.

Ryouko wants to decrease the number of pages that need to be turned. In order to achieve this, she can merge two pages of her notebook. If Ryouko merges page x to page y, she would copy all the information on page x to y (1 ≤ x, y ≤ n), and consequently, all elements in sequence a that were x would become y. Note that x can be equal to y, in which case no changes take place.

Please tell Ryouko the minimum number of pages that she needs to turn. Note she can apply the described operation at most once before the reading. Note that the answer can exceed 32-bit integers.

## Input

The first line of input contains two integers n and m (1 ≤ n, m ≤ 10^5).

The next line contains m integers separated by spaces: a1, a2, ..., am (1 ≤ ai ≤ n).

## Output

Print a single integer — the minimum number of pages Ryouko needs to turn.

## Examples

**Input:**
```
4 6
1 2 3 4 3 2
```

**Output:**
```
3
```

**Input:**
```
10 5
9 4 3 8 8
```

**Output:**
```
6
```

## Note

In the first sample, the optimal solution is to merge page 4 to 3, after merging sequence a becomes {1, 2, 3, 3, 3, 2}, so the number of pages Ryouko needs to turn is \|1 - 2\| + \|2 - 3\| + \|3 - 3\| + \|3 - 3\| + \|3 - 2\| = 3.

In the second sample, optimal solution is achieved by merging page 9 to 4.


### ideas
1. 操作执行一次，不会得到更差的结果，所以都需要执行一次
2. 假设有一个序列 a[i-1], a[i], a[i+1]
3. 在改变前 abs(a[i-1] - a[i]) + abs(a[i] - a[i+1])
4. 如果 a[i-1] <= a[i] <= a[i+1] (改变a[i], 没啥好处)
5. 所以，a[i], 是个峰谷，那就好办了啊
6. 