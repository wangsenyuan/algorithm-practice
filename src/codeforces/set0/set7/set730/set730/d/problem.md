Polycarp is playing a game called "Running Over The Bridges". In this game he has to run over $n$ bridges from the left to the right. Bridges are arranged one after the other, so the $i$-th bridge begins where the $(i-1)$-th bridge ends.

You have the following data about bridges: $l_i$ and $t_i$ — the length of the $i$-th bridge and the maximum allowed time which Polycarp can spend running over the $i$-th bridge. Thus, if Polycarp is at the beginning of bridge $i$ at time $T$ then he has to leave it at time $T + t_i$ or earlier. It is allowed to reach the right end of a bridge exactly at time $T + t_i$.

Polycarp can run from the left side to the right one with speed $0.5$, so he will run over a bridge with length $s$ in time $2 \cdot s$. Besides, he has several magical drinks. If he uses one drink, his speed increases twice (i.e. to value $1$) for $r$ seconds. All magical drinks are identical. Please note that Polycarp can use a drink only at integer moments of time, and he drinks it instantly and completely. Additionally, if Polycarp uses a drink at the moment $T$ he can use the next drink not earlier than at the moment $T + r$.

What is the minimal number of drinks Polycarp has to use to run over all $n$ bridges? If this number is not greater than $10^5$, then you have to find out the moments of time when Polycarp has to use each magical drink.

## Input

The first line contains two integers $n$ and $r$ ($1 \le n \le 2 \cdot 10^5$, $1 \le r \le 10^{12}$) — the number of bridges and the duration of the effect of a magical drink.

The second line contains a sequence of integers $l_1, l_2, \ldots, l_n$ ($1 \le l_i \le 5 \cdot 10^6$), where $l_i$ is equal to the length of the $i$-th bridge.

The third line contains a sequence of integers $t_1, t_2, \ldots, t_n$ ($1 \le t_i \le 10^7$), where $t_i$ is equal to the maximum allowed time which Polycarp can spend running over the $i$-th bridge.

## Output

The first line of the output should contain $k$ — the minimal number of drinks which Polycarp has to use, or `-1` if there is no solution.

If the solution exists and the value of $k$ is not greater than $10^5$ then output $k$ integers on the next line — moments of time from beginning of the game when Polycarp has to use drinks. Print the moments of time in chronological order. If there are several solutions, you can output any of them.

## Examples

### Input
```
1 3
7
10
```

### Output
```
2
0 3
```

### Input
```
3 3
3 3 3
3 3 2
```

### Output
```
-1
```

### Input
```
3 100000
5 5 5
5 7 8
```

### Output
```
1
0
```

### Input
```
4 1000
1 2 3 4
10 9 10 9
```

### Output
```
0
```

## Note

In the first case, there is only one bridge and it is clear that Polycarp cannot run over it without magical drinks. So, if he will use one magical drink at the start (moment of time $0$), and the second one three seconds later (moment of time $3$), he will be able to reach the end of the bridge in time. Please note that in this case there are several possible answers to the problem. For example, Polycarp can use the first drink at the moment of time $4$ and the second one at the moment of time $7$.

In the second case, Polycarp cannot run over all bridges even if he will use magical drinks. So, the answer in this case is `-1`.

In the fourth case, Polycarp can run over all bridges without magical drinks.


### ideas
1. if li > ti 也就是说，即使有魔法，也没法在t[i]时间内（速度为1）通过第i座桥 => -1
2. 否则肯定可以通过这座桥，并且存在一个答案
3. 需要的是最小的k, 那么岂不是，如果 li <= ti and 但是 2 * ti <= li时，就必须使用，如果要使用，就从一开始就使用（当然需要考虑上次使用的效果）
4. 似乎就模拟过去就可了～
5. 先要计算次数（有可能超过1e5)