# Problem E: Taxi Drivers

Andrew prefers taxi to other means of transport, but recently most taxi drivers have been acting inappropriately. In order to earn more money, taxi drivers started to drive in circles. Roads in Andrew's city are one-way, and people are not necessary able to travel from one part to another, but it pales in comparison to insidious taxi drivers.

The mayor of the city decided to change the direction of certain roads so that the taxi drivers wouldn't be able to increase the cost of the trip endlessly. More formally, if the taxi driver is on a certain crossroads, they wouldn't be able to reach it again if he performs a nonzero trip.

Traffic controllers are needed in order to change the direction the road goes. For every road it is known how many traffic controllers are needed to change the direction of the road to the opposite one. It is allowed to change the directions of roads one by one, meaning that each traffic controller can participate in reversing two or more roads.

You need to calculate the minimum number of traffic controllers that you need to hire to perform the task and the list of the roads that need to be reversed.

## Input

The first line contains two integers $n$ and $m$ ($2 \leq n \leq 100000$, $1 \leq m \leq 100000$) — the number of crossroads and the number of roads in the city, respectively.

Each of the following $m$ lines contain three integers $u_i$, $v_i$ and $c_i$ ($1 \leq u_i, v_i \leq n$, $1 \leq c_i \leq 10^9$, $u_i \neq v_i$) — the crossroads the road starts at, the crossroads the road ends at and the number of traffic controllers required to reverse this road.

## Output

In the first line output two integers the minimal amount of traffic controllers required to complete the task and amount of roads $k$ which should be reversed. $k$ should not be minimized.

In the next line output $k$ integers separated by spaces — numbers of roads, the directions of which should be reversed. The roads are numerated from $1$ in the order they are written in the input. If there are many solutions, print any of them.

## Examples

### Example 1

**Input:**
```
5 6
2 1 1
5 2 6
2 3 2
3 4 3
4 5 5
1 5 4
```

**Output:**
```
2 2
1 3
```

### Example 2

**Input:**
```
5 7
2 1 5
3 2 3
1 3 3
2 4 1
4 3 5
5 4 1
1 5 3
```

**Output:**
```
3 3
3 4 7
```

## Note

There are two simple cycles in the first example: $1 \to 5 \to 2 \to 1$ and $2 \to 3 \to 4 \to 5 \to 2$. One traffic controller can only reverse the road $2 \to 1$ and he can't destroy the second cycle by himself. Two traffic controllers can reverse roads $2 \to 1$ and $2 \to 3$ which would satisfy the condition.

In the second example one traffic controller can't destroy the cycle $1 \to 3 \to 2 \to 1$. With the help of three controllers we can, for example, reverse roads $1 \to 3$, $2 \to 4$, $1 \to 5$.


### ideas
1. 要把图改成成dag
2. 找出所有的强连通分量，二分check需要的人数
3. 所有超过x的边保留，如果这些边，仍然组成了环，那么就false，否则就是ok的（即使断开的情况下，也可以在那些 <= x的边中，选择一部分，让它们重新连接起来）