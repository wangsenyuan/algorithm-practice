# Problem C: Bulmart

## Problem Description

A new trade empire is rising in Berland. Bulmart, an emerging trade giant, decided to dominate the market of ... shovels! And now almost every city in Berland has a Bulmart store, and some cities even have several of them! The only problem is, at the moment sales are ... let's say a little below estimates. Some people even say that shovels retail market is too small for such a big company to make a profit. But the company management believes in the future of that market and seeks new ways to increase income.

There are $n$ cities in Berland connected with $m$ bi-directional roads. All roads have equal lengths. It can happen that it is impossible to reach a city from another city using only roads. There is no road which connects a city to itself. Any pair of cities can be connected by at most one road.

There are $w$ Bulmart stores in Berland. Each of them is described by three numbers:

- $c_i$ — the number of city where the $i$-th store is located (a city can have no stores at all or have several of them)
- $k_i$ — the number of shovels in the $i$-th store
- $p_i$ — the price of a single shovel in the $i$-th store (in burles)

The latest idea of the Bulmart management is to create a program which will help customers get shovels as fast as possible for affordable budget. Formally, the program has to find the minimum amount of time needed to deliver $r_j$ shovels to the customer in the city $g_j$ for the total cost of no more than $a_j$ burles. The delivery time between any two adjacent cities is equal to 1. If shovels are delivered from several cities, the delivery time is equal to the arrival time of the last package. The delivery itself is free of charge.

The program needs to find answers to $q$ such queries. Each query has to be processed independently from others, i.e. a query does not change number of shovels in stores for the next queries.

## Input

The first line contains two integers $n, m$ ($1 \leq n \leq 5000$, $0 \leq m \leq \min(5000, n \cdot (n - 1) / 2)$).

Each of the next $m$ lines contains two integers $x_e$ and $y_e$, meaning that the $e$-th road connects cities $x_e$ and $y_e$ ($1 \leq x_e, y_e \leq n$).

The next line contains a single integer $w$ ($1 \leq w \leq 5000$) — the total number of Bulmart stores in Berland.

Each of the next $w$ lines contains three integers describing the $i$-th store: $c_i, k_i, p_i$ ($1 \leq c_i \leq n$, $1 \leq k_i, p_i \leq 2 \cdot 10^5$).

The next line contains a single integer $q$ ($1 \leq q \leq 1000$) — the number of queries.

Each of the next $q$ lines contains three integers describing the $j$-th query: $g_j, r_j$ and $a_j$ ($1 \leq g_j \leq n$, $1 \leq r_j, a_j \leq 10^9$)

## Output

Output $q$ lines. On the $j$-th line, print an answer for the $j$-th query — the minimum amount of time needed to deliver $r_j$ shovels to the customer in city $g_j$ spending no more than $a_j$ burles. Print `-1` if there is no solution for the $j$-th query.

## Example

### Input
```
6 4
4 2
5 4
1 2
3 2
2
4 1 2
3 2 3
6
1 2 6
2 3 7
3 1 2
4 3 8
5 2 5
6 1 10
```

### Output
```
2
-1
2
2
3
-1
```

### ideas
1. 假设给定时间w，是否能够在w时间内，在给定budget a[i], 从整个地图获得r[i]的货物？
2. 那么就是找出和g[i]相聚不超过w的城市中，按照架构从低往高排列，看能前r[i]个货物，价值是否 <= a[i]
3. 