# Problem Description

On vacations n pupils decided to go on excursion and gather all together. They need to overcome the path with the length l meters. Each of the pupils will go with the speed equal to v1. To get to the excursion quickly, it was decided to rent a bus, which has seats for k people (it means that it can't fit more than k people at the same time) and the speed equal to v2. In order to avoid seasick, each of the pupils want to get into the bus no more than once.

Determine the minimum time required for all n pupils to reach the place of excursion. Consider that the embarkation and disembarkation of passengers, as well as the reversal of the bus, take place immediately and this time can be neglected.

## Input

The first line of the input contains five positive integers n, l, v1, v2 and k (1 ≤ n ≤ 10 000, 1 ≤ l ≤ 10^9, 1 ≤ v1 < v2 ≤ 10^9, 1 ≤ k ≤ n) — the number of pupils, the distance from meeting to the place of excursion, the speed of each pupil, the speed of bus and the number of seats in the bus.

## Output

Print the real number — the minimum time in which all pupils can reach the place of excursion. Your answer will be considered correct if its absolute or relative error won't exceed 10^-6.

## Examples

### Example 1

**Input:**

```text
5 10 1 2 5
```

**Output:**

```text
5.0000000000
```

### Example 2

**Input:**

```text
3 6 1 2 1
```

**Output:**

```text
4.7142857143
```

## Note

In the first sample we should immediately put all five pupils to the bus. The speed of the bus equals 2 and the distance is equal to 10, so the pupils will reach the place of excursion in time 10 / 2 = 5.

## Ideas

1. n个人，通过l长度的距离; 如步行，l / v1
2. 或者部分人可以乘车（但貌似可以多趟）而且也不需要送到目的地
3. 假设答案为expect,
4. 第一次，车辆接送k个人，经过距离x后，花费x/v2 + (l - x) / v1 == expect
5. 它第一次接送后，运行x距离，然后返回y（此时小孩子们已经走了一段距离w, w + y = x)
6. 然后车辆再运行x距离，然后再返回。。。再运行（x距离）但是
7. 好像也能算