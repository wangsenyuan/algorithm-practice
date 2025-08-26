# Problem D - Ralph's Tours in Binary Country

## Problem Description

Ralph is in the Binary Country. The Binary Country consists of $n$ cities and $(n-1)$ bidirectional roads connecting the cities. The roads are numbered from 1 to $(n-1)$, where the $i$-th road connects the city labeled $\lfloor \frac{i}{2} \rfloor$ (here $\lfloor x \rfloor$ denotes $x$ rounded down to the nearest integer) and the city labeled $(i+1)$, and the length of the $i$-th road is $L_i$.

Now Ralph gives you $m$ queries. In each query, he tells you some city $A_i$ and an integer $H_i$. He wants to make some tours starting from this city. He can choose any city in the Binary Country (including $A_i$) as the terminal city for a tour. He gains happiness $(H_i - L)$ during a tour, where $L$ is the distance between the city $A_i$ and the terminal city.

Ralph is interested in tours from $A_i$ in which he can gain positive happiness. For each query, compute the sum of happiness gains for all such tours.

**Note:** Ralph will never take the same tour twice or more (in one query), and he will never pass the same city twice or more in one tour.

## Input

- The first line contains two integers $n$ and $m$ $(1 \leq n \leq 10^6, 1 \leq m \leq 10^5)$.
- $(n-1)$ lines follow, each containing one integer $L_i$ $(1 \leq L_i \leq 10^5)$, which denotes the length of the $i$-th road.
- $m$ lines follow, each containing two integers $A_i$ and $H_i$ $(1 \leq A_i \leq n, 0 \leq H_i \leq 10^7)$.

## Output

Print $m$ lines, on the $i$-th line print one integer — the answer for the $i$-th query.

## Examples

### Example 1

**Input:**
```
2 2
5
1 8
2 4
```

**Output:**
```
11
4
```

### Example 2

**Input:**
```
6 4
2
1
1
3
2
2 4
1 3
3 2
1 7
```

**Output:**
```
11
6
3
28
```

## Explanation for Example 2

Ralph's first query is to start tours from city 2 with $H_i = 4$. Here are the options:

- **City 5 as terminal:** Distance = 3, happiness = $4 - 3 = 1$
- **City 4 as terminal:** Distance = 1, happiness = $4 - 1 = 3$
- **City 1 as terminal:** Distance = 2, happiness = $4 - 2 = 2$
- **City 3 as terminal:** Distance = 3, happiness = $4 - 3 = 1$
- **City 2 as terminal:** Distance = 0, happiness = $4 - 0 = 4$

**Note:** Ralph won't choose city 6 as his terminal city because the distance between city 6 and city 2 is 5, which leads to negative happiness.

So the answer for the first query is $1 + 3 + 2 + 1 + 4 = 11$.

### ideas
1. (1, 2), (1, 3), (2, 4), (2, 5), ...