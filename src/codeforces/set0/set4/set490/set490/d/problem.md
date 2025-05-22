# Equal Chocolate Bars

## Problem Description
Polycarpus likes giving presents to Paraskevi. He has bought two chocolate bars, each of them has the shape of a segmented rectangle. The first bar is a1 × b1 segments large and the second one is a2 × b2 segments large.

Polycarpus wants to give Paraskevi one of the bars at the lunch break and eat the other one himself. Besides, he wants to show that Polycarpus's mind and Paraskevi's beauty are equally matched, so the two bars must have the same number of squares.

## Rules
To make the bars have the same number of squares, Polycarpus eats a little piece of chocolate each minute. Each minute he does one of the following:
1. Either breaks one bar exactly in half (vertically or horizontally) and eats exactly a half of the bar
2. Or he chips off exactly one third of a bar (vertically or horizontally) and eats exactly a third of the bar

### Important Notes
- In the first case, he is left with a half of the bar
- In the second case, he is left with two thirds of the bar
- Both variants aren't always possible, and sometimes Polycarpus cannot chip off a half nor a third

### Examples
- Bar of 16 × 23: Can chip off a half, but not a third
- Bar of 20 × 18: Can chip off both a half and a third
- Bar of 5 × 7: Cannot chip off a half nor a third

## Task
Find the minimum number of minutes Polycarpus needs to make two bars consist of the same number of squares, and determine the possible sizes of the bars after the process.

## Input
- First line: integers a1, b1 (1 ≤ a1, b1 ≤ 10^9) — initial sizes of the first chocolate bar
- Second line: integers a2, b2 (1 ≤ a2, b2 ≤ 10^9) — initial sizes of the second bar

Note: You can use int64 (Pascal), long long (C++), or long (Java) to process large integers.

## Output
- First line: m — the minimum number of minutes required
- Second line: sizes of the first bar after m minutes
- Third line: sizes of the second bar after m minutes

Format the sizes using the same format as the input. Print the sizes (numbers in pairs) in any order.
If there are multiple solutions, print any of them.
If there is no solution, print a single line with integer -1.


## ideas
1. 先单独考虑一个维度，进行2分或者3分
2. 对于给定x，对它进行分解得到 $x = 2^{p2} \times 3^{p3} \times x_0$
3. w * h 可以合并在一起
4. $x \times y$ = $2^{p1+p2} \times 3^{q1+q2} \times other$
5. 那么如果两个other不相同，无解（因为肯定不可能相同）
6. 