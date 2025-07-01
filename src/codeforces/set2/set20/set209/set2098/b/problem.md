# Problem B: Apartment Selection

## Problem Description

Sasha wants to buy an apartment on a street where the houses are numbered from 1 to $10^9$ from left to right.

There are $n$ bars on this street, located in houses with numbers $a_1, a_2, \ldots, a_n$. Note that there might be multiple bars in the same house, and in this case, these bars are considered distinct.

Sasha is afraid that by the time he buys the apartment, some bars may close, but no more than $k$ bars can close.

For any house with number $x$, define $f(x)$ as the sum of $|x - y|$ over all open bars $y$ (that is, after closing some bars).

Sasha can potentially buy an apartment in a house with number $x$ (where $1 \leq x \leq 10^9$) if and only if it is possible to close at most $k$ bars so that after that $f(x)$ becomes minimal among all houses.

Determine how many different houses Sasha can potentially buy an apartment in.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each test case contains two integers $n$ and $k$ ($1 \leq n \leq 10^5$, $0 \leq k < n$) — the number of bars and the maximum number of bars that can close.

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^9$) — the house numbers where the bars are located.

It is guaranteed that the sum of $n$ over all test cases does not exceed $10^5$.

## Output

For each test case, output a single integer — the number of houses where Sasha can buy an apartment.

