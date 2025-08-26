# Problem E: John's Journey

## Problem Description

John has just bought a new car and is planning a journey around the country. The country has $N$ cities, some of which are connected by bidirectional roads. There are $N - 1$ roads and every city is reachable from any other city. Cities are labeled from 1 to $N$.

John first has to select from which city he will start his journey. After that, he spends one day in a city and then travels to a randomly chosen city which is directly connected to his current one and which he has not yet visited. He does this until he can't continue obeying these rules.

To select the starting city, he calls his friend Jack for advice. Jack is also starting a big casino business and wants to open casinos in some of the cities (max 1 per city, maybe nowhere). Jack knows John well and he knows that if he visits a city with a casino, he will gamble exactly once before continuing his journey.

He also knows that if John enters a casino in a good mood, he will leave it in a bad mood and vice versa. Since he is John's friend, he wants him to be in a good mood at the moment when he finishes his journey. John is in a good mood before starting the journey.

**Question:** In how many ways can Jack select a starting city for John and cities where he will build casinos such that no matter how John travels, he will be in a good mood at the end? Print answer modulo $10^9 + 7$.

## Input

- First line: A positive integer $N$ $(1 \leq N \leq 100000)$, the number of cities.
- Next $N - 1$ lines: Two numbers $a, b$ $(1 \leq a, b \leq N)$ separated by a single space, meaning that cities $a$ and $b$ are connected by a bidirectional road.

## Output

Output one number: the answer to the problem modulo $10^9 + 7$.

## Examples

### Example 1
**Input:**
```
2
1 2
```

**Output:**
```
4
```

**Explanation:** If Jack selects city 1 as John's starting city, he can either build 0 casinos (so John will be happy all the time), or build a casino in both cities (so John would visit a casino in city 1, become unhappy, then go to city 2, visit a casino there and become happy and his journey ends there because he can't go back to city 1). If Jack selects city 2 for start, everything is symmetrical, so the answer is 4.

### Example 2
**Input:**
```
3
1 2
2 3
```

**Output:**
```
10
```

**Explanation:** If Jack tells John to start from city 1, he can either build casinos in 0 or 2 cities (total 4 possibilities). If he tells him to start from city 2, then John's journey will either contain cities 2 and 1 or 2 and 3. Therefore, Jack will either have to build no casinos, or build them in all three cities. With other options, he risks John ending his journey unhappy. Starting from 3 is symmetric to starting from 1, so in total we have $4 + 2 + 4 = 10$ options.


### ideas
1. dp[u][0]表示在子树u中，到任何一个叶子节点的路径上，有偶数个酒店的计数
2. dp[u][0] = prod(dp[v][1]) (如果在u处开设酒店)，那么所有u的子树，都必须是奇数的
3.          + prod(dp[v][0]) 或者在u处不开设酒店 
4. dp[u][1] = prod(dp[v][0]) + prod(dp[v][1]) 
5. 咋感觉不大对呢？
6. dp[u][0] = dp[u][1]了
7. 考虑一个节点u做为root，那么它所有到叶子节点的路径的奇偶性，是否必须要一致，
8. 如果不一致呢？假设有只有两条路径，一条路径的长度是偶数，一条路径的长度是奇数
9. 如果在root处开设了酒店，似乎也没有关系呐（如果进入一条路径，这条路径上就需要使用奇数个酒店就可以了）
10. 咋感觉只和最长的路径有关系，假设一条直线，从一个端点开始，那么当前节点的可以选择有或者没有酒店
11. 假设到了最后一个节点时，如果它目前已经选择了偶数个，那么这里不能建造，否则这里就必须建造
12. 那么也就是说，一条n长度的线，它的答案 = pow(2, n - 1) (最后一个节点，被之前的节点的状态确定了)
13. 然后考虑一棵树；所有的内部节点，都可以有自己的状态，但是叶子节点的状态，被内部节点给确定了
14. 假设现在是r，有m个叶子节点 = pow(2, n - m) (如果r是叶子节点，需要处理一下)