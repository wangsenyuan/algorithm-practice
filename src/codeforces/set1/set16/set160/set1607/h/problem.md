The chef has cooked $n$ dishes yet again: the $i$-th dish consists of $a_i$ grams of fish and $b_i$ grams of meat.

Banquet organizers consider two dishes $i$ and $j$ equal if $a_i = a_j$ and $b_i = b_j$ at the same time.

The banquet organizers estimate the variety of $n$ dishes as follows. The variety of a set of dishes is equal to the number of different dishes in it. The less variety is, the better.

In order to reduce the variety, a taster was invited. He will eat exactly $m_i$ grams of food from each dish. For each dish, the taster determines separately how much fish and how much meat he will eat. The only condition is that he will eat exactly $m_i$ grams of the $i$-th dish in total.

Determine how much of what type of food the taster should eat from each dish so that the value of variety is the minimum possible. If there are several correct answers, you may output any of them.

## Input

The first line of input data contains an integer $t$ ($1 \leq t \leq 10^4$) — the number of test cases.

Each test case's description is preceded by a blank line. Next comes a line that contains an integer $n$ ($1 \leq n \leq 2 \cdot 10^5$) — the number of dishes. Then follows $n$ lines, $i$-th of which contains three integers $a_i$, $b_i$ and $m_i$ ($0 \leq a_i, b_i \leq 10^6$; $0 \leq m_i \leq a_i + b_i$) — the mass of fish in $i$-th dish, the mass of meat in $i$-th dish and how many grams in total the taster should eat in $i$-th dish.

The sum of all $n$ values for all input data sets in the test does not exceed $2 \cdot 10^5$.

## Output

For each test case, print on the first line the minimum value of variety that can be achieved by eating exactly $m_i$ grams of food (for all $i$ from 1 to $n$) from a dish $i$.

Then print $n$ lines that describe a way to do this: the $i$-th line should contain two integers $x_i$ and $y_i$ ($0 \leq x_i \leq a_i$; $0 \leq y_i \leq b_i$; $x_i + y_i = m_i$), where $x_i$ is how many grams of fish the taster should eat from $i$-th dish, and $y_i$ is how many grams of meat.

If there are several ways to achieve a minimum balance, print any of them.

## Example

### Input
```
5

3
10 10 2
9 9 0
10 9 1

2
3 4 1
5 1 2

3
7 2 5
6 5 4
5 5 6

1
13 42 50

5
5 7 12
3 1 4
7 3 7
0 0 0
4 1 5
```

### Output
```
1
1 1
0 0
1 0
2
0 1
1 1
2
3 2
0 4
1 5
1
8 42
2
5 7
3 1
4 3
0 0
4 1
```


### ideas
1. 考虑i, j使的它们相同, a[i] - x[i]  = a[j] - x[j], b[i] - (m[i] - x[i]) = b[j] - (m[j] - x[j])
2. a[i] - a[j] = x[i] - x[j]
3. b[i] - b[j] = (m[i] - x[i]) - (m[j] - x[j]) = (m[i] - m[j]) - (x[i] - x[j])
4. (x[i] - x[j]) = (a[i] - a[j], (m[i] - m[j]) - (b[i] - b[j]))
5. 如果不考虑x >= 0, x <= min(a[i], m[i])的条件， 简单的分组就可以了
6. 怎么处理x呢？假设已经处理了一部分，现在加入x[j], 但是x[j] 最大到(a[j], m[j])
7. 也就是说，虽然j满足分组的条件，但是无法加入里面，那么这个分组还得有个要求
8. 对于分组(a[i] - a[j], (...) - ..) 按照x的下限进行降序处理。然后对于当前的，根据它的上限，加入到对应的分组里面去
9. 