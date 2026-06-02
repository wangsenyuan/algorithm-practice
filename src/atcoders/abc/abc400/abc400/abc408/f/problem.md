# Problem Statement

There are $N$ scaffolds numbered from $1$ to $N$ arranged in a line. The height of scaffold $i$ $(1 \leq i \leq N)$ is $H_i$.

Takahashi decides to play a game of moving on the scaffolds. Initially, he freely chooses an integer $i$ $(1 \leq i \leq N)$ and gets on scaffold $i$.

When he is on scaffold $i$ at some point, he can choose an integer $j$ $(1 \leq j \leq N)$ satisfying the following condition and move to scaffold $j$:

- $H_j \leq H_i - D$ and $1 \leq |i - j| \leq R$.

Find the maximum number of moves he can make when he repeats moving until he can no longer move.

## Constraints

- $1 \leq N \leq 5 \times 10^5$
- $1 \leq D, R \leq N$
- $H$ is a permutation of $(1, 2, \ldots, N)$.
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```
N D R
H_1 H_2 \ldots H_N
```

## Output

Output the answer.

## Sample Input 1

```
5 2 1
5 3 1 4 2
```

## Sample Output 1

```
2
```

Takahashi initially gets on scaffold $1$ and can move between the scaffolds as follows:

- **First move**: Since $H_2 \leq H_1 - D$ and $|2 - 1| \leq R$, he can move to scaffold $2$. Move from scaffold $1$ to scaffold $2$.
- **Second move**: Since $H_3 \leq H_2 - D$ and $|3 - 2| \leq R$, he can move to scaffold $3$. Move from scaffold $2$ to scaffold $3$.

Since the height of scaffold $3$ is $1$, he can no longer move.

As shown above, he can move $2$ times. Also, no matter how he chooses the scaffolds to move to, he cannot move $3$ or more times. Therefore, output $2$.

## Sample Input 2

```
13 3 2
13 7 10 1 9 5 4 11 12 2 8 6 3
```

## Sample Output 2

```
3


### ideas
1. range update + query
