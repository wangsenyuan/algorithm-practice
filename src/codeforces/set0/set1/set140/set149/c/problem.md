Petya loves football very much, especially when his parents aren't home. Each morning he comes to the yard, gathers his friends and they play all day. From time to time they have a break to have some food or do some chores (for example, water the flowers).

The key in football is to divide into teams fairly before the game begins. There are $n$ boys playing football in the yard (including Petya), each boy's football playing skill is expressed with a non-negative characteristic $a_i$ (the larger it is, the better the boy plays).

Let's denote the number of players in the first team as $x$, the number of players in the second team as $y$, the individual numbers of boys who play for the first team as $p_i$ and the individual numbers of boys who play for the second team as $q_i$. Division of $n$ boys into two teams is considered **fair** if three conditions are fulfilled:

1. Each boy plays for exactly one team ($x + y = n$).
2. The sizes of teams differ by no more than one ($|x - y| \le 1$).
3. The total football playing skills for two teams differ by no more than the value of skill the best player in the yard has. More formally:
   \[
   \left| \sum_{i \in \text{team}_1} a_i - \sum_{i \in \text{team}_2} a_i \right| \le \max_i a_i.
   \]

Your task is to help the guys divide into two teams fairly. It is guaranteed that a fair division into two teams always exists.

## Input

The first line contains the only integer $n$ ($2 \le n \le 10^5$) which represents the number of guys in the yard. The next line contains $n$ positive space-separated integers $a_i$ ($1 \le a_i \le 10^4$), where the $i$-th number represents the $i$-th boy's playing skills.

## Output

On the first line print an integer $x$ — the number of boys playing for the first team. On the second line print $x$ integers — the individual numbers of boys playing for the first team. On the third line print an integer $y$ — the number of boys playing for the second team. On the fourth line print $y$ integers — the individual numbers of boys playing for the second team. Don't forget that you should fulfil all three conditions: $x + y = n$, $|x - y| \le 1$, and the condition that limits the total skills.

If there are multiple ways to solve the problem, print any of them.

The boys are numbered starting from one in the order in which their skills are given in the input data. You are allowed to print individual numbers of boys who belong to the same team in any order.

## Examples

### Input

```
3
1 2 1
```

### Output

```
2
1 2
1
3
```

### Input

```
5
2 3 3 1 1
```

### Output

```
3
4 1 3
2
5 2
```

### ideas
1. max(ai) 是确定的
2. 是不是相邻的交替划分就可以了？
3. 但是怎么证明呢？