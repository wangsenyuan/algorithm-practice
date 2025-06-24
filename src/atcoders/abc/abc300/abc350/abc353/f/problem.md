# Problem Statement

Tiles are laid out on a coordinate plane. There are two types of tiles: small tiles of size $1 \times 1$ and large tiles of size $K \times K$, laid out according to the following rules:

- For each pair of integers $(i, j)$, the square $([i, i+1) \times [j, j+1))$ is contained within either one small tile or one large tile.
    - If $\left\lfloor \frac{i}{K} \right\rfloor + \left\lfloor \frac{j}{K} \right\rfloor$ is even, it is contained within a small tile.
    - Otherwise, it is contained within a large tile.

Tiles include their boundaries, and no two different tiles have a positive area of intersection.

For example, when $K = 3$, the tiles are laid out as follows:

Example tile layout for $K=3$:
```

| S | S | S | L | L | L | S | S | S |
| S | S | S | L | L | L | S | S | S |
| S | S | S | L | L | L | S | S | S |
| L | L | L | S | S | S | L | L | L |
| L | L | L | S | S | S | L | L | L |
| L | L | L | S | S | S | L | L | L |
| S | S | S | L | L | L | S | S | S |
| S | S | S | L | L | L | S | S | S |
| S | S | S | L | L | L | S | S | S |

```
Takahashi starts at the point $(S_x + 0.5, S_y + 0.5)$ on the coordinate plane.

He can impart the following movement any number of times:
- Choose a direction (left, up, down, right) and a positive integer $n$. Move $n$ units in that direction.

Each time he crosses from one tile to another, he must pay a toll of 1.

Determine the minimum toll Takahashi must pay to reach the point $(T_x + 0.5, T_y + 0.5)$.

## Constraints
- $1 \leq K \leq 10^{10}$
- $0 \leq S_x, S_y \leq 2 \times 10^5$
- $0 \leq T_x, T_y \leq 2 \times 10^5$
- $0 \leq S_x, T_x, S_y, T_y$
- All input values are integers.

## Input
The input is given from Standard Input in the following format:

```
K
S_x S_y
T_x T_y
```

## Output
Print the minimum toll Takahashi must pay.

## Sample Input 1
```
3
7 2
1 6
```

## Sample Output 1
```
5
```
For example, he can move as follows, paying a toll of 5.

- Move up by 3. Pay a toll of 1.
- Move left by 2. Pay a toll of 1.
- Move up by 1. Pay a toll of 1.
- Move left by 4. Pay a toll of 2.

The toll paid cannot be 4 or less, so print 5.

## Sample Input 2
```
1
41 42
18 86
```

## Sample Output 2
```
42
```
When he moves the shortest distance, he will always pay a toll of 42.
The toll paid cannot be 41 or less, so print 42.

## Sample Input 3
```
100
100 99
199 1
```

## Sample Output 3
```
0
```
There are cases where no toll needs to be paid.

## Sample Input 4
```
502964321
5182164213865191 18842645763346525
143271389130967 14412423443308528
```

## Sample Output 4
```
79314849
```


## ideas
1. 显然要尽可能得经过K-tile
2. 任意两个k-tile，可以通过一个1-tile联通
3. 找到离S/T的最近的K-tile（或者都试一下，就4个）然后看这两个K-tile的cost）
4. 