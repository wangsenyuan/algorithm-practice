# Elevator Simulation

## Problem Description

The m-floor (m > 1) office of international corporation CodeForces has the advanced elevator control system established. It works as follows.

All office floors are sequentially numbered with integers from 1 to m. At time t = 0, the elevator is on the first floor, the elevator is empty and nobody is waiting for the elevator on other floors. Next, at times ti (ti > 0) people come to the elevator. For simplicity, we assume that one person uses the elevator only once during the reported interval. For every person we know three parameters: the time at which the person comes to the elevator, the floor on which the person is initially, and the floor to which he wants to go.

### Elevator Movement Rules

The movement of the elevator between the floors is as follows. At time t (t ≥ 0, t is an integer) the elevator is always at some floor:

1. **Release passengers**: The elevator releases all people who are in the elevator and want to get to the current floor
2. **Board passengers**: Then it lets in all the people waiting for the elevator on this floor
3. **Move**: After that the elevator decides which way to move and at time (t + 1) the elevator gets to the selected floor

We can assume that all of these actions (going in or out from the elevator) are made instantly. If a person comes to the elevator exactly at time t, then he has enough time to get into it.

### Direction Selection Algorithm

The elevator selects the direction of moving by the following algorithm:

- **If the elevator is empty and at the current time no one is waiting for the elevator on any floor**, then the elevator remains at the current floor
- **Otherwise**, let's assume that the elevator is on the floor number x (1 ≤ x ≤ m). Then elevator calculates the directions' "priorities" pup and pdown:
  - **pup** is the sum of the number of people waiting for the elevator on the floors with numbers greater than x, and the number of people in the elevator, who want to get to the floors with the numbers greater than x
  - **pdown** is the sum of the number of people waiting for the elevator on the floors with numbers less than x, and the number of people in the elevator, who want to get to the floors with the numbers less than x
  - If pup ≥ pdown, then the elevator goes one floor above the current one (that is, from floor x to floor x + 1)
  - Otherwise the elevator goes one floor below the current one (that is, from floor x to floor x - 1)

Your task is to simulate the work of the elevator and for each person to tell the time when the elevator will get to the floor this person needs. Please note that the elevator is large enough to accommodate all the people at once.

## Input

- The first line contains two space-separated integers: n, m (1 ≤ n ≤ 10^5, 2 ≤ m ≤ 10^5) — the number of people and floors in the building, correspondingly
- Next n lines each contain three space-separated integers: ti, si, fi (1 ≤ ti ≤ 10^9, 1 ≤ si, fi ≤ m, si ≠ fi) — the time when the i-th person begins waiting for the elevator, the floor number where the i-th person was initially located, and the number of the floor where he wants to go

## Output

Print n lines. In the i-th line print a single number — the moment of time when the i-th person gets to the floor he needs. The people are numbered in the order in which they are given in the input.

**Note:** Please don't use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout streams or the %I64d specifier.

## Examples

### Example 1
**Input:**
```
3 10
1 2 7
3 6 5
3 4 8
```

**Output:**
```
7
11
8
```

### Example 2
**Input:**
```
2 10
1 2 5
7 4 5
```

**Output:**
```
5
9
```

## Detailed Simulation (Example 1)

In the first sample the elevator worked as follows:

- **t = 1**: The elevator is on floor 1. Floor 2 has one person waiting. pup = 1 + 0 = 1, pdown = 0 + 0 = 0, pup ≥ pdown. So the elevator goes to floor 2
- **t = 2**: The elevator is on floor 2. One person enters the elevator, he wants to go to floor 7. pup = 0 + 1 = 1, pdown = 0 + 0 = 0, pup ≥ pdown. So the elevator goes to floor 3
- **t = 3**: The elevator is on floor 3. There is one person in the elevator, he wants to go to floor 7. Floors 4 and 6 have two people waiting. pup = 2 + 1 = 3, pdown = 0 + 0 = 0, pup ≥ pdown. So the elevator goes to floor 4
- **t = 4**: The elevator is on floor 4. There is one person in the elevator who wants to go to floor 7. One person goes into the elevator, he wants to get to floor 8. Floor 6 has one man waiting. pup = 1 + 2 = 3, pdown = 0 + 0 = 0, pup ≥ pdown. So the elevator goes to floor 5
- **t = 5**: The elevator is on floor 5. There are two people in the elevator, they want to get to floors 7 and 8. There is one person waiting on floor 6. pup = 1 + 2 = 3, pdown = 0 + 0 = 0, pup ≥ pdown. So the elevator goes to floor 6
- **t = 6**: The elevator is on floor 6. There are two people in the elevator, they want to get to floors 7 and 8. One man enters the elevator, he wants to get to floor 5. pup = 0 + 2 = 2, pdown = 0 + 1 = 1, pup ≥ pdown. So the elevator goes to floor 7
- **t = 7**: The elevator is on floor 7. One person leaves the elevator (wanted floor 7). There are two people in the elevator, they want to get to floors 8 and 5. pup = 0 + 1 = 1, pdown = 0 + 1 = 1, pup ≥ pdown. So the elevator goes to floor 8
- **t = 8**: The elevator is on floor 8. One person leaves the elevator (wanted floor 8). There is one person in the elevator, he wants to go to floor 5. pup = 0 + 0 = 0, pdown = 0 + 1 = 1, pup < pdown. So the elevator goes to floor 7
- **t = 9**: The elevator is on floor 7. There is one person in the elevator, this person wants to get to floor 5. pup = 0 + 0 = 0, pdown = 0 + 1 = 1, pup < pdown. So the elevator goes to floor 6
- **t = 10**: The elevator is on floor 6. There is one person in the elevator, he wants to get to floor 5. pup = 0 + 0 = 0, pdown = 0 + 1 = 1, pup < pdown. So the elevator goes to floor 5
- **t = 11**: The elevator is on floor 5. One person leaves the elevator (wanted floor 5). The elevator is empty and nobody needs it, so the elevator remains at floor 5

### ideas
1. 模拟，但是要避免出现在空楼层停靠的情况
2. 那么每个人有一次上梯的事件、一次离开的事件，所以是O(m)的。再加上每次判断，可能需要log(n)
3. 假设当前(时刻t）在x层，且通过某种方式上移，那么要尽快的找到下一个目标楼层
4. 下一个目标楼层 = min(y > x)当前电梯中，需要离开的人的最低目标楼层
5. 或者是，在时刻 <= t + y - x 时，有人会在y层等待乘梯的人
6. 但是人是按照时间到达的，所以还是有问题
7. 比如可能在向上的过程中，下部出现了新的等待乘梯的人，那么又要往下去移动
8. 换个角度，共n个人，也就是共n个时间点
9. 假设在这些时间点t，知道目前电梯的情况，在x楼层，其中有up个想上，down个人想下
10. 然后在下个时间点前（下一个人出现前）
11. 是可以根据现有的状态，计算出型的 t_next, x_next, x_up, x_down
12. 其中t -> t_next （因为没有新的人加入进来，所以移动的方向是可以计算的）
13. 这个就是上面的那个情况
14. 然后在时刻t_next + 1, 把人们加入进来，重新计算一遍
15. 然后就是在不加入新人时，对当前状态的处理
16. 使用一个segment tree，可以快速的计算，在x两头等待的人的数量，还有位置
17. 然后使用一个segment tree，来计算目前乘梯中的，要离开的事件。同样包括位置、数量
18. 