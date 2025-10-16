# Problem

Let's define a function \(f(a)\) for an array \(a\) of length \(n\) as:

$$f(a) = \text{cost} + (a_1 - a_2 + a_3 - a_4 \cdots a_n)$$

where \(\text{cost}\) is zero initially.

Now consider the scenario where Alice and Bob are given an array \(a\) of length \(n\). They play a game taking at most \(10^{100}\) turns alternately with Alice going first.

In each turn, they must perform any one (only one) of the following operations:

1. End the game for both Alice and Bob.
2. Choose two indices \(l, r\) with \(1 \le l \le r \le n\) and swap \(a_l\) and \(a_r\); this adds \((r - l)\) to the cost.

Assume that Alice tries to maximize \(f(a)\) and Bob tries to minimize it.

Your task is to determine the final value of \(f(a)\) assuming both players play optimally.

## Input

Each test contains multiple test cases. The first line contains the number of test cases \(t\) (\(1 \le t \le 10^4\)). The description of the test cases follows.

The first line of each test case contains a single integer \(n\) (\(1 \le n \le 2 \cdot 10^5\)) — the length of the array \(a\).

The second line contains \(n\) integers \(a_1, a_2, a_3, \ldots, a_n\) (\(1 \le a_i \le 10^9\)) — the elements of the array \(a\).

It is guaranteed that the sum of \(n\) over all test cases does not exceed \(2 \cdot 10^5\).

## Output

For each testcase, output a single integer — the final value of \(f(a)\) assuming both players play optimally.

## Example

### Input
```
5
2
1000 1
5
9 9 9 9 9
4
7 1 8 4
6
1 14 1 14 1 15
9
31 12 14 22 89 6 78 25 91
```

### Output
```
999
13
12
-7
265
```

## Note

For the first testcase, it is optimal for Alice to end the game on her first turn.

So the final value of \(\text{cost} = 0\) and \(f(a) = 0 + 1000 - 1 = 999\).

For the fourth testcase, it is optimal for Alice to swap \(a_1\) and \(a_6\), and then it is optimal for Bob to end the game on his first turn.

So the final value of \(\text{cost} = 5\) and \(f(a) = 5 + 15 - 14 + 1 - 14 + 1 - 1 = -7\).