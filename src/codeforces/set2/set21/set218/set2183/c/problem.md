You are given a battle. As the country's top general, you must decide where to place your troops.

There are `n` bases in a line, and the `k`-th base is the home base for your army. Initially, there is exactly one soldier at base `k`. Each day, the following happens in order:

1. You choose a base `i` (`1 <= i <= n`) and also choose any number of soldiers currently located at base `i` (it can be `0` or all soldiers in that base). Then, you command all chosen soldiers to move either to base `i-1` or to base `i+1`.
2. All chosen soldiers must move in the same direction, and no soldier is allowed to move to the left of base `1` or to the right of base `n`.
3. After that, a new soldier moves onto base `k`. This newly added soldier cannot be ordered by that day's commands.

However, time is tight: there are only `m` days until the enemy attacks. A base is called **fortified** if at least one soldier resides in it. Your job is to find the maximum number of fortified bases (including the home base) you can have by the end of the `m`-th day.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 1e4`). The description of the test cases follows.

The first line of each test case contains three integers `n`, `m`, `k` (`1 <= k <= n <= 1e5`, `1 <= m <= 1e9`) — denoting the number of bases, the number of days until you need to fortify, and the index of the home base.

It is guaranteed that the sum of `n` across all test cases does not exceed `2e5`.

## Output

For each test case, print the maximum number of bases you can fortify at the end of the `m`-th day.

## Example

**Input**

```
7
3 1 3
3 3 2
4 2 2
3 2 1
4 3 3
7 7 4
100000 1000000000 100000
```

**Output**

```
2
3
3
2
3
6
100000
```

## Note

In the second test case, one way to fortify 3 bases is:

- On the first day, order `0` soldiers in base `3` to move to base `2`. At the end of the day, a new soldier moves to base `2` (now there are `2` soldiers in base `2` and `0` on any other base).
- On the second day, order `1` soldier in base `2` to move to base `1`. At the end of the day, a new soldier moves to base `2`. Now there are `2` soldiers in base `2` and `1` soldier on base `1`.
- On the third day, order `2` soldiers in base `2` to move to base `3`. At the end of the day, a new soldier moves to base `2`. Now there is `1` soldier in base `1` and `2` soldiers in base `3`.

There is now at least one soldier in each of bases `1`, `2`, and `3`. Therefore, the answer is `3`.

In the third test case, one way you can achieve 3 bases being fortified is:

- On the first day, order the existing soldier to move from base `2` to base `3`. At the end of the day, a new soldier moves to base `2`.
- On the second day, order the soldier in base `2` to move to base `1`. At the end of the day, a new soldier moves to base `2`.

There is now a soldier at each of bases `1`, `2`, and `3`. Therefore, the answer is `3`. It can be shown we cannot have more than `3` fortified bases by the end of day `2`.


### ideas
1. 先考虑一边的情况，假设一共守卫了w长度的bases，那么这些bases肯定都被守卫住了，没有空的
2. 那么这时候，最好的策略时等到有w个人以后，再开始行动？
3. 考虑覆盖区域3，先等待2 seconds，这时候3个人，然后让2个人移动，再让1个人移动
4. 这时候一共花费2 + 1 + 1 = 4 秒
5. 如果一开始让第一个人移动到位置，花费 2秒，这时候在base里有2个人，那么再移动1个，也就是3秒
6. 所以等待是不对的（因为在移动的时候，也可以有士兵到达）
7. 假设第一个士兵分配到了w处，那么此时正好有w个士兵出现了，再分配w-1次，就可以了
8. 就可以覆盖整个区域