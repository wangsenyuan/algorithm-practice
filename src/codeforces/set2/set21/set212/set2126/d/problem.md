# Casino Problem

You are given $n$ casinos, numbered from $1$ to $n$. Each casino is described by three integers: $l_i$, $r_i$, and $real_i$ ($l_i \leq real_i \leq r_i$). You initially have $k$ coins.

You can play at casino $i$ only if the current number of coins $x$ satisfies $l_i \leq x \leq r_i$. After playing, your number of coins becomes $real_i$.

You can visit the casinos in any order and are not required to visit all of them. Each casino can be visited no more than once.

Your task is to find the maximum final number of coins you can obtain.

## Input

The first line contains a single integer $t$ ($1 \leq t \leq 10^4$) — the number of test cases.

The first line of each test case contains two integers $n$ and $k$ ($1 \leq n \leq 10^5$, $0 \leq k \leq 10^9$) — the number of casinos and the initial number of coins.

This is followed by $n$ lines. In the $i$-th line, there are three integers $l_i$, $r_i$, $real_i$ ($0 \leq l_i \leq real_i \leq r_i \leq 10^9$) — the parameters of the $i$-th casino.

It is guaranteed that the sum of all $n$ across all test cases does not exceed $10^5$.

## Output

For each test case, output a single integer — the maximum number of coins you can obtain after optimally choosing the order of visiting the casinos.

## Example

### Input
```
5
3 1
2 3 3
1 2 2
3 10 10
1 0
1 2 2
1 2
1 2 2
2 2
1 3 2
2 4 4
2 5
1 10 5
3 6 5
```

### Output
```
10
0
2
4
5
```

## Note

- **Test case 1**: You can first play at the 2nd casino. After that, you will have 2 coins. Then you can play at the 1st casino, and the number of coins will increase to 3. Finally, after playing at the 3rd casino, you will have 10 coins — this is the maximum possible amount.

- **Test case 2**: You have no money, so you cannot earn more.

- **Test case 4**: It is beneficial to play at the 2nd casino right away and earn 4 coins.