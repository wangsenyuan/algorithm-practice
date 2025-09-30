# Problem C - Bus Conductor

## Problem Statement

I guess there's not much point in reminding you that Nvodsk winters aren't exactly hot. That increased the popularity of the public transport dramatically. The route of bus 62 has exactly n stops (stop 1 goes first on its way and stop n goes last). The stops are positioned on a straight line and their coordinates are 0 = x₁ < x₂ < ... < xₙ.

Each day exactly m people use bus 62. For each person we know the number of the stop where he gets on the bus and the number of the stop where he gets off the bus. A ticket from stop a to stop b (a < b) costs xᵦ - xₐ rubles. However, the conductor can choose no more than one segment NOT TO SELL a ticket for. We mean that conductor should choose C and D (C ≤ D) and sell a ticket for the segments [A, C] and [D, B], or not sell the ticket at all. The conductor and the passenger divide the saved money between themselves equally. The conductor's "untaxed income" is sometimes interrupted by inspections that take place as the bus drives on some segment of the route located between two consecutive stops. The inspector fines the conductor by c rubles for each passenger who doesn't have the ticket for this route's segment.

You know the coordinates of all stops xᵢ; the numbers of stops where the i-th passenger gets on and off, aᵢ and bᵢ (aᵢ < bᵢ); the fine c; and also pᵢ — the probability of inspection on segment between the i-th and the i + 1-th stop. The conductor asked you to help him make a plan of selling tickets that maximizes the mathematical expectation of his profit.

## Input

The first line contains three integers n, m and c (2 ≤ n ≤ 150,000, 1 ≤ m ≤ 300,000, 1 ≤ c ≤ 10,000).

The next line contains n integers xᵢ (0 ≤ xᵢ ≤ 10⁹, x₁ = 0, xᵢ < xᵢ₊₁) — the coordinates of the stops on the bus's route.

The third line contains n - 1 integer pᵢ (0 ≤ pᵢ ≤ 100) — the probability of inspection in percents on the segment between stop i and stop i + 1.

Then follow m lines that describe the bus's passengers. Each line contains exactly two integers aᵢ and bᵢ (1 ≤ aᵢ < bᵢ ≤ n) — the numbers of stops where the i-th passenger gets on and off.

## Output

Print the single real number — the maximum expectation of the conductor's profit. Your answer will be considered correct if its absolute or relative error does not exceed 10⁻⁶.

Namely: let's assume that your answer is a, and the answer of the jury is b. The checker program will consider your answer correct, if |a - b| / max(1, |b|) ≤ 10⁻⁶.

## Examples

### Example 1

**Input:**
```
3 3 10
0 10 100
100 0
1 2
2 3
1 3
```

**Output:**
```
90.000000000
```

### Example 2

**Input:**
```
10 8 187
0 10 30 70 150 310 630 1270 2550 51100
13 87 65 0 100 44 67 3 4
1 10
2 9
3 8
1 5
6 10
2 7
4 10
4 5
```

**Output:**
```
76859.990000000
```

## Note

A comment to the first sample:

The first and third passengers get tickets from stop 1 to stop 2. The second passenger doesn't get a ticket. There always is inspection on the segment 1-2 but both passengers have the ticket for it. There never is an inspection on the segment 2-3, that's why the second passenger gets away with the cheating. Our total profit is (0 + 90/2 + 90/2) = 90.



### ideas
1. 有n个人，从$a_i$上车，$b_i$下车， 正常情况下，他需要购买车票，支付 $b_i - a_i$, 
2. 但是如果他不买票，那么只需要支付一半给司机；
3. 但是在途中，如果出现了检查人员，对于不买票的人，将对司机罚款c
4. 如果司机可以选择一个区间[C...D]，不卖票，而是收取一半，那么他的预期最大收入时多少？
5. 每段的检查人员出现的概率，是否有可加性？
6. dp[i][0] 表示到i为止，表示还在卖票的收益 = 0
7. 假设目前车上有w个人，其中x个人有票， w - x 个没票
8. 且下一个区间出现检察员的概率是p,那么这一段的预期收入 = (w - x) * (1 - p) * dx / 2 - c * (w - x) * p
9. 这里x没有作用， w * (1 - p) * dx / 2 - c * w * p
10. 这个值是确定的，就是选一段连续的区间，这个预期的sum是最大的
11. got