# Dog Show Problem

A new dog show on TV is starting next week. On the show dogs are required to demonstrate bottomless stomach, strategic thinking and self-preservation instinct. You and your dog are invited to compete with other participants and naturally you want to win!

## Problem Description

On the show a dog needs to eat as many bowls of dog food as possible (bottomless stomach helps here). Dogs compete separately of each other and the rules are as follows:

At the start of the show the dog and the bowls are located on a line. The dog starts at position $x = 0$ and $n$ bowls are located at positions $x = 1, x = 2, ..., x = n$. The bowls are numbered from 1 to n from left to right. After the show starts the dog immediately begins to run to the right to the first bowl.

The food inside bowls is not ready for eating at the start because it is too hot (dog's self-preservation instinct prevents eating). More formally, the dog can eat from the i-th bowl after $t_i$ seconds from the start of the show or later.

It takes dog 1 second to move from the position $x$ to the position $x + 1$. The dog is not allowed to move to the left, the dog runs only to the right with the constant speed 1 distance unit per second. When the dog reaches a bowl (say, the bowl i), the following cases are possible:

- **the food had cooled down** (i.e. it passed at least $t_i$ seconds from the show start): the dog immediately eats the food and runs to the right without any stop,
- **the food is hot** (i.e. it passed less than $t_i$ seconds from the show start): the dog has two options: to wait for the i-th bowl, eat the food and continue to run at the moment $t_i$ or to skip the i-th bowl and continue to run to the right without any stop.

After $T$ seconds from the start the show ends. If the dog reaches a bowl of food at moment $T$ the dog can not eat it. The show stops before $T$ seconds if the dog had run to the right of the last bowl.

You need to help your dog create a strategy with which the maximum possible number of bowls of food will be eaten in $T$ seconds.

## Input

Two integer numbers are given in the first line - $n$ and $T$ ($1 \leq n \leq 200,000$, $1 \leq T \leq 2 \cdot 10^9$) — the number of bowls of food and the time when the dog is stopped.

On the next line numbers $t_1, t_2, ..., t_n$ ($1 \leq t_i \leq 10^9$) are given, where $t_i$ is the moment of time when the i-th bowl of food is ready for eating.

## Output

Output a single integer — the maximum number of bowls of food the dog will be able to eat in $T$ seconds.

## Examples

### Example 1
**Input:**
```
3 5
1 5 3
```
**Output:**
```
2
```

### Example 2
**Input:**
```
1 2
1
```
**Output:**
```
1
```

### Example 3
**Input:**
```
1 1
1
```
**Output:**
```
0
```

## Note

In the first example the dog should skip the second bowl to eat from the two bowls (the first and the third).


### ideas
1. dp[i]表示在i处开始，且正好bowl i刚好cool down下来
2. dp[i] = 1 + ? + dp[j] where j 是下一个可以刚好开始的位置 
3.       而中间的，必须是到达的时候，已经能够冷却到能吃，或者直接pass（也就是不等待）
4. ? = t[k] - t[i] <= k - i =》 t[k] - k <= t[i] - i
5. 在i确定的情况下，（还有j确定的情况下）k的数量是确定的
6. dp[i] = max(dp[j] + count(k))
7. f(i, j) = count of k, t[k] - k <= t[i] - i, where k < j
8. 是不是符合中间大两头小的情况呢？
9. 