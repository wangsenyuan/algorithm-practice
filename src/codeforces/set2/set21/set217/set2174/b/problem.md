It is said that writing wishing cards on New Year's Eve can bring good luck. You have written wishing cards and plan to give them to Little A.

You have invited $n$ friends to help you deliver the wishing cards. You plan to give the $i$-th friend $b_i$ wishing cards. The friends will visit Little A's house in the order of $1$ to $n$, and each will hand over all his $b_i$ wishing cards to Little A. Little A always remembers the friend who gives her the most wishing cards, so after the $j$-th friend's visit, she gains happiness equal to the largest number of cards received so far, that is, $\max(b_1, b_2, \ldots, b_j)$.

The $i$-th friend can carry at most $a_i$ cards, and the total number of cards cannot exceed $k$. Your task is to optimally choose values of $b_i$ so that the total happiness of Little A after $n$ visits is maximized under these constraints.

Note that some friends may carry no cards at all, Little A won't get mad.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^3$). The description of the test cases follows.

The first line of each test case contains two integers $n$ and $k$ ($1 \le n \le 10^5$, $1 \le k \le 360$), the number of friends and the maximum total number of cards.

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($0 \le a_i \le k$), the maximum number of cards each friend can carry.

It is guaranteed that the sum of $n$ over all test cases does not exceed $5 \cdot 10^5$ and the sum of $k$ over all test cases does not exceed $1800$.

## Output

For each test case, output the maximum happiness Little A can get.

## Example

### Input
```
4
3 4
0 0 1
6 8
1 2 0 5 1 8
3 4
1 0 4
5 8
2 4 5 4 3
```

### Output
```
1
20
5
19
```

## Note

In the first test case, the only way to distribute wishing cards is $b = [0, 0, 1]$.

In the third test case, $b = [1, 0, 4]$ is invalid because you only have $k = 4$ wishing cards.

In the fourth test case, one of the optimal ways to distribute wishing cards is to set $b = [2, 0, 5, 0, 0]$, in which Little A will gain $2 + 2 + 5 + 5 + 5 = 19$ happiness.


