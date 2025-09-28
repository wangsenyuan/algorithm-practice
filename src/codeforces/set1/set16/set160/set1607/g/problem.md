# Problem G: Chef's Balance

A known chef has prepared $n$ dishes: the $i$-th dish consists of $a_i$ grams of fish and $b_i$ grams of meat.

The banquet organizers estimate the balance of $n$ dishes as follows. The balance is equal to the absolute value of the difference between the total mass of fish and the total mass of meat.

Technically, the balance equals to $\left|\sum_{i=1}^{n}a_i - \sum_{i=1}^{n}b_i\right|$. The smaller the balance, the better.

In order to improve the balance, a taster was invited. He will eat exactly $m$ grams of food from each dish. For each dish, the taster determines separately how much fish and how much meat he will eat. The only condition is that he should eat exactly $m$ grams of each dish in total.

Determine how much of what type of food the taster should eat from each dish so that the value of the balance is as minimal as possible. If there are several correct answers, you may choose any of them.

## Input

The first line of input data contains an integer $t$ $(1 \leq t \leq 10^4)$ — the number of the test cases.

Each test case's description is preceded by a blank line. Next comes a line that contains integers $n$ and $m$ $(1 \leq n \leq 2 \cdot 10^5; 0 \leq m \leq 10^6)$. The next $n$ lines describe dishes, the $i$-th of them contains a pair of integers $a_i$ and $b_i$ $(0 \leq a_i, b_i \leq 10^6)$ — the masses of fish and meat in the $i$-th dish.

It is guaranteed that it is possible to eat $m$ grams of food from each dish. In other words, $m \leq a_i + b_i$ for all $i$ from $1$ to $n$ inclusive.

The sum of all $n$ values over all test cases in the test does not exceed $2 \cdot 10^5$.

## Output

For each test case, print on the first line the minimal balance value that can be achieved by eating exactly $m$ grams of food from each dish.

Then print $n$ lines that describe a way to do this: the $i$-th line should contain two integers $x_i$ and $y_i$ $(0 \leq x_i \leq a_i; 0 \leq y_i \leq b_i; x_i + y_i = m)$, where $x_i$ is how many grams of fish taster should eat from the $i$-th meal and $y_i$ is how many grams of meat.

If there are several ways to achieve a minimal balance, find any of them.

### ideas
1. 是让taster对于每个dish吃掉m克，使的剩余的部分，sum(a[i]) - sum(b[i])最小
2. abs(sum(a[i] - x[i]) - sum(b[i] - (m - x[i])))
3. a[i] - x[i] - (b[i] - m + x[i]) = sum(a[i] - b[i] + m) 最小
4. (3 - 4 + 5) = 4
5. -(a[i] - x[i]) + (b[i] - m + x[i]) 
6. b[i] - a[i] + 2 * x[i] - m 最小
7. 4 - 3 - 5 + 2 * 2 = 0
8. 这么奇怪的～
9. 假设每项用掉x[i]的fish, 
10. 如果最后的情况是fish 更多一些， 那么 a[i] - x[i] - (b[i] - (m - x[i])) = a[i] - b[i] + m - 2 * x[i] (x[i] <= min(a[i], x[i]))