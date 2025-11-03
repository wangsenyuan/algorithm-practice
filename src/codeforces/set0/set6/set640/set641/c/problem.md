## Little Artem and rueda

Artem is fond of dancing. Most of all, Artem likes rueda — a Cuban dance performed by pairs of boys and girls forming a circle and dancing together.

There are $n$ pairs of boys and girls standing in a circle. Initially, boy $1$ dances with girl $1$, boy $2$ with girl $2$, and so on. Girls are numbered in clockwise order. During the dance, different moves are announced and all pairs perform these moves. While performing moves, boys move along the circle, while girls always stay at their initial positions. There are two types of moves:

- **Type 1 (shift by $x$)**: A value $x$ and a direction are announced, and all boys move $x$ positions in the corresponding direction.
- **Type 2 (pairwise swap)**: Boys dancing with even-indexed girls swap positions with boys who are dancing with odd-indexed girls. For example, the one who was dancing with girl $1$ swaps with the one who was dancing with girl $2$, the one with girl $3$ swaps with the one with girl $4$, and so on. It is guaranteed that $n$ is even.

Your task is to determine the final position of each boy.

## Input

- The first line contains two integers $n$ and $q$ ($2 \le n \le 10^6$, $1 \le q \le 2\cdot 10^6$) — the number of couples in the rueda and the number of commands, respectively. It is guaranteed that $n$ is even.
- The next $q$ lines describe the commands. Each command starts with an integer type ($1$ or $2$).
  - If the type is $1$, it is followed by an integer $x$ ($-n \le x \le n$). Here, $0 \le x \le n$ means all boys move $x$ girls clockwise, while negative $x$ means all boys move $|x|$ positions counter-clockwise.
  - If the type is $2$, there is no additional input.

## Output

Output $n$ integers: the $i$-th should be the index of the boy that girl $i$ is dancing with after performing all $q$ moves.

## Examples

### Example 1
Input
```
6 3
1 2
2
1 2
```
Output
```
4 3 6 5 2 1
```

### Example 2
Input
```
2 3
1 1
2
1 -2
```
Output
```
1 2
```

### Example 3
Input
```
4 2
2
1 3
```
Output
```
1 4 3 2
```


### ideas
1. [1, 2, 3, 4, 5, 6]
2. move 1
3. [6, 1, 2, 3, 4, 5]
4. swap
5. [1, 6, 3, 2, 5, 4]
6. move 1
7. [4, 1, 6, 3, 2, 5]
8. swap 
9. [1, 4, 3, 6, 5, 2]
10. 看起来偶数位不会改变顺序（奇数位也一样）
11. 假设偶数位移动了w次，奇数位移动了v次，
12. 没次swap，相当于奇数位+1，偶数位-1