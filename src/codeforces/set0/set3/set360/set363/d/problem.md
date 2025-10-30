## Bike Renting

A group of n schoolboys decided to ride bikes. As nobody of them has a bike, the boys need to rent them.

The renting site offered them m bikes. The renting price is different for different bikes, renting the j-th bike costs pj rubles.

In total, the boys' shared budget is a rubles. Besides, each of them has his own personal money, the i-th boy has bi personal rubles. The shared budget can be spent on any schoolchildren arbitrarily, but each boy's personal money can be spent on renting only this boy's bike.

Each boy can rent at most one bike, one cannot give his bike to somebody else.

What maximum number of schoolboys will be able to ride bikes? What minimum sum of personal money will they have to spend in total to let as many schoolchildren ride bikes as possible?

### Input
The first line of the input contains three integers `n`, `m` and `a` (`1 ≤ n, m ≤ 10^5; 0 ≤ a ≤ 10^9`). The second line contains the sequence of integers `b1, b2, ..., bn` (`1 ≤ bi ≤ 10^4`), where `bi` is the amount of the i-th boy's personal money. The third line contains the sequence of integers `p1, p2, ..., pm` (`1 ≤ pj ≤ 10^9`), where `pj` is the price for renting the j-th bike.

### Output
Print two integers `r` and `s`, where `r` is the maximum number of schoolboys that can rent a bike and `s` is the minimum total personal money needed to rent `r` bikes. If the schoolchildren cannot rent any bikes, then `r = s = 0`.

### Examples
Input
```text
2 2 10
5 5
7 6
```
Output
```text
2 3
```

Input
```text
4 5 2
8 1 1 2
6 3 7 5 2
```
Output
```text
3 8
```

### Note
In the first sample both schoolchildren can rent a bike. For instance, they can split the shared budget in half (5 rubles each). In this case one of them will have to pay 1 ruble from the personal money and the other one will have to pay 2 rubles from the personal money. In total, they spend 3 rubles of their personal money. This way of distribution of money minimizes the amount of spent personal money.


### ideas
1. 先考虑第一个问题的答案
2. 假设要检查能否有x个人可以做到
3. 那么选择b[?]最大的x个人，然后和x个最小的p[?]， 进行配对
4. 看 sum(max(p[i] - b[i]), 0) <= a
5. 这样就能得到x。
6. 第二个要怎么处理呢？
7. 总的花费是确定的 = sum(p), 然后分配a即可