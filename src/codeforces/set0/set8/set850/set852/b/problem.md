# Problem B: Neural Network Country

## Problem Description

Due to the recent popularity of Deep Learning, new countries are starting to look like Neural Networks. That is, the
countries are being built deep with many layers, each layer possibly having many cities. They also have one entry and
one exit point.

There are exactly **L** layers, each having **N** cities. Let us look at the two adjacent layers L1 and L2. Each city
from layer L1 is connected to each city from layer L2 with traveling cost $c_{ij}$ for $1 \leq i, j \leq N$, and each
pair of adjacent layers has the same cost between their cities as any other pair (they just stacked the same layers, as
usual). Also, the traveling costs to each city from layer L2 are the same for all cities in L1, that is $c_{ij}$ is the
same for $1 \leq i \leq N$ and fixed $j$.

Doctor G. needs to speed up his computations for this country, so he asks you to find the number of paths he can take
from entry to exit point such that his traveling cost is divisible by given number **M**.

## Input

- **First line**: Contains three integers **N**, **L**, and **M** separated by spaces
    - $1 \leq N \leq 10^6$ (number of cities in each layer)
    - $2 \leq L \leq 10^5$ (number of layers)
    - $2 \leq M \leq 100$ (divisor for traveling cost)

- **Second line**: Contains **N** integers denoting costs from entry point to the first layer
    - Each cost: $0 \leq \text{cost} \leq M$

- **Third line**: Contains **N** integers denoting costs between adjacent layers
    - Each cost: $0 \leq \text{cost} \leq M$

- **Fourth line**: Contains **N** integers denoting costs from the last layer to the exit point
    - Each cost: $0 \leq \text{cost} \leq M$

## Output

Output a single integer: the number of paths Doctor G. can take which have total cost divisible by **M**,
modulo $10^9 + 7$.

## Example

### Input

```
2 3 13
4 6
2 1
3 4
```

### Output

```
2
```

## Note

This is a country with 3 layers, each layer having 2 cities. Paths with total costs divisible by 13 are the only valid
paths. Notice that input edges for layer cities have the same cost, and they are the same for all layers.

### ideas

1. 考虑中间 $L_i$ 到 $L_{i+1}$ 中间n个城市到n个城市的转换
2. 到$city_1$的都是在之前的基础上增加$c1$, ...一般的就是在之前的基础上增加 $c_i$
3. 假设$dp_x$表示到本层之前，余数为x时的计数
4. $ndp_y$ += $dp_x$ where (x + $c1$) % m = y
5. 但是这样子，复杂性还是 $n \times L \times m$
6. dp[x][y] 表示从 x ( $w \mod m = x$) 到 y的转换次数
7. (x + ci) % m = y 的城市的个数
8. 然后 pow(a, L - 1)
9. 然后乘以前面的，然后乘以最后的
