In Ann's favorite book shop there are \(n\) books on math and economics. Books are numbered from 1 to \(n\). Each of them contains a non-negative number of problems.

Today there is a sale: any subsegment of a segment from \(l\) to \(r\) can be bought at a fixed price.

Ann decided that she wants to buy such a non-empty subsegment that the sale operates on it and the number of math problems is greater than the number of economics problems **exactly by \(k\)**. Note that \(k\) may be positive, negative or zero.

Unfortunately, Ann is not sure on which segment the sale operates, but she has \(q\) assumptions. For each of them she wants to know the number of options to buy a subsegment satisfying the condition (because the time she spends on choosing depends on that).

Currently Ann is too busy solving other problems, so she asks you for help. For each of her assumptions determine the number of subsegments of the given segment such that the number of math problems is greater than the number of economics problems on that subsegment exactly by \(k\).

## Input

The first line contains two integers \(n\) and \(k\) \((1 \le n \le 100\,000,\ -10^9 \le k \le 10^9)\) — the number of books and the needed difference between the number of math problems and the number of economics problems.

The second line contains \(n\) integers \(t_1, t_2, \dots, t_n\) \((1 \le t_i \le 2)\), where \(t_i = 1\) if the \(i\)-th book is on math and \(t_i = 2\) if the \(i\)-th book is on economics.

The third line contains \(n\) integers \(a_1, a_2, \dots, a_n\) \((0 \le a_i \le 10^9)\), where \(a_i\) is the number of problems in the \(i\)-th book.

The fourth line contains a single integer \(q\) \((1 \le q \le 100\,000)\) — the number of assumptions.

Each of the next \(q\) lines contains two integers \(l_i\) and \(r_i\) \((1 \le l_i \le r_i \le n)\) describing the \(i\)-th Ann's assumption.

## Output

Print \(q\) lines, in the \(i\)-th of them print the number of subsegments for the \(i\)-th Ann's assumption.

## Examples

### Input

```
4 1
1 1 1 2
1 1 1 1
4
1 2
1 3
1 4
3 4
```

### Output

```
2
3
4
1
```

### Input

```
4 0
1 2 1 2
0 0 0 0
1
1 4
```

### Output

```
10
```

## Note

In the first sample Ann can buy subsegments \([1;1], [2;2], [3;3], [2;4]\) if they fall into the sales segment, because the number of math problems is greater by 1 on them than the number of economics problems. So we should count for each assumption the number of these subsegments that are subsegments of the given segment.

Segments \([1;1]\) and \([2;2]\) are subsegments of \([1;2]\).

Segments \([1;1], [2;2]\) and \([3;3]\) are subsegments of \([1;3]\).

Segments \([1;1], [2;2], [3;3], [2;4]\) are subsegments of \([1;4]\).

Segment \([3;3]\) is a subsegment of \([3;4]\).


### ideas
1. 如果[i, j] 满足条件, 也就是 math[i...j] - ec[i...j] = k
2. 那么所有区间[l...r], l <= i 且 j <= r 的都需要加1
3. 假设迭代j, 找到i(这个有可能是多个值)， 对于每个i，现在active的query(r >= j)的部分，对[0...i]进行+1
4. 主要是i的问题，有可能会变成O(n)
5. 最大的i比较容易找出来，其他的i1 < i 满足这个区间内math = ec
6. 好像有个技巧可以弄～
7. 假设是 i1, i2, i3, ... ik
8. 且在ik处加1， 那么在i1处 +k
9. 所以对于i1来说，只要知道有多少个数在它的后面增加了,且它们的sum多少， 那么 i1的贡献 = sum - cnt * 1 （不是i1, ik)