# Problem D

Steve made the foolish decision to mine at night, and came across a monstrous creature: the chicken jockey𝑛!

A chicken jockey𝑛 consists of 𝑛 mobs stacked in order on top of each other, with mob 1 at the bottom and mob 𝑛 at the top. Mob 𝑖 initially has ℎ𝑖 health.

In one attack, Steve can deal 1 damage to any mob. If any mob reaches 0 or less health, it dies, and all the mobs on top of it fall down and form a new stack. The bottom mob in the new stack takes 1 fall damage for every mob it was on top of before falling (i.e. the number of mobs below it in the previous stack, including the one that just died). This may kill it as well, in which case all mobs on top of it fall down again and the process repeats.

For example, consider a chicken jockey6 with initial mob healths [1,2,1,3,5,2]. If Steve damages the third mob in the stack, it dies and the mobs with health [3,5,2] fall down in a new stack. The bottom mob takes 3 units of fall damage so it also dies, and the mobs with health [5,2] fall down in a new stack. The bottom mob takes 1 unit of fall damage. As a result, after Steve's first attack, there will be two stacks with healths [1,2] and [4,2].

Steve's sword's durability is low, so he wishes to know the minimum attacks required to kill all the mobs.

## Input

Each test contains multiple test cases. The first line contains the number of test cases 𝑡 (1≤𝑡≤10⁴). The description of the test cases follows.

The first line of each test case contains a single integer 𝑛 (2≤𝑛≤2⋅10⁵) — the number of mobs.

The second line of each test case contains 𝑛 integers ℎ₁,ℎ₂,…,ℎₙ (1≤ℎᵢ≤10⁹) — the initial health of each mob.

It is guaranteed that the sum of 𝑛 over all test cases does not exceed 2⋅10⁵.

## Output

For each test case, output a single integer — the minimum attacks required to kill all the mobs.

## Example

**Input:**
```
5
5
3 1 4 1 2
4
1 1 1 1
6
1 2 1 3 5 2
6
3 1 1 3 2 1
3
1000000000 1000000000 1000000000
```

**Output:**
```
7
1
7
5
2999999998
```

## Note

In the first test case, the initial stack has mobs with health [3,1,4,1,2]. Steve can use one attack to damage the second mob in the stack and kill it. The third mob takes 2 units of fall damage. There are now two stacks: [3] and [2,1,2]. Now Steve can kill the second mob in the second stack. The third mob in the stack takes 2 units of fall damage, killing it. There are now two stacks: [3] and [2]. Finally, Steve can use five attacks to kill the remaining mobs.

In the second test case, Steve can deal 1 damage to the bottom mob in the stack. When it dies, the second mob will take 1 unit of fall damage and die; then the third mob will take 1 unit of fall damage and die; finally the fourth mob will take 1 unit of fall damage and die.


### ideas
1. 每次产生一个新的stack的时候，最底部的那个受到的伤害 = 它之前所在的位置
2. 假设操作的位置是 x1, x2, x3... xk （操作了k次）
3. 那么这个序列本身不变的情况下，从高往低操作是更优的
4. dp[i] = 如果 i >= a[i] (不管前面怎么弄的，让i形成了一个新的stack) = dp[i+1]
5. 如果 i >= a[i], dp[i] = dp[i+1]
6. dp[i] = a[i] - i + dp[i+1]
7. 