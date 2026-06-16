# A1. Collective Mindsets (easy)

[Problem link](https://codeforces.com/problemset/problem/690/A1)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Tonight is brain dinner night: `N` guests attend — `N - 1` zombies and Heidi in
disguise. Each real zombie has a distinct rank from `1` to `N - 1`; Heidi has rank
`N` (highest).

A chest contains some indivisible brains. They are distributed as follows:

1. The highest-ranked living guest proposes how to split the brains.
2. Everyone votes. If at least half accept, the split is used.
3. Otherwise the proposer is killed and the next highest rank proposes.
4. On a tie, the highest-ranked **living** voter counts twice and always votes for
   their own proposal to survive.

Zombies are perfectly rational and greedy. Each guest's priorities are:

1. Survive the night.
2. Maximize their own number of brains.

A zombie rejects an offer unless it is **strictly better** for them than what they
could get from later rounds.

Heidi proposes first. Find the **minimum** number of brains in the chest so that
Heidi can make an offer that:

- at least half of the attendees accept, and
- gives Heidi at least one brain.

## Input

One integer `N` (`1 <= N <= 10^9`) — the number of guests.

## Output

One integer — the minimum chest size that lets Heidi take at least one brain home.

## Examples

### Input

```text
1
```

### Output

```text
1
```

### Input

```text
4
```

### Output

```text
2
```


### ideas
1. 海盗分金的问题, 考虑最后两个海盗, 那么带头的那个,可以得到全部,而最后一个海盗啥都得不到
2. 考虑3个海盗. 头领只需要给最后一个海盗1个, 就可以得到最后一个海盗的投票, 从而获胜
3. 不需要给第二个海盗,因为轮到它的时候, 最后一个海盗啥也得不到. 
4. 也就是倒数第3个海盗的状态是 [x-1, 0, 1]
5. 如果是4个海盗, [x-1, 0, 1, 0] (他只需要争取倒数第二个人的支持) (那这里4的时候,x >= 2)
6. 如果是5个海盗, [x-1, 0, 1, 0, 0] (正好3票 vs 3票)
7. 如果是6个海盗, [x-2, 0, 1, 0, 0, 1]
8. 如果n是奇数, 比如 n = 5, 那么他要获得 (n - 1) // 2 个人的支持
9. 如果n是偶数, 比如 n = 6, 那么他要获得 (n-1) // 2 个人的支持