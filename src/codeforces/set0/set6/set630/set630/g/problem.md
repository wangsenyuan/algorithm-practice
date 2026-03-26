Because of budget cuts, one IT company established a new non-financial reward system instead of bonuses.

Two kinds of actions are rewarded: fixing critical bugs and suggesting new interesting features. A person who fixed a critical bug gets an "I fixed a critical bug" pennant on their table. A person who suggested a new interesting feature gets an "I suggested a new feature" pennant on their table.

Because of the limited budget of the new reward system, only `5` "I fixed a critical bug" pennants and `3` "I suggested a new feature" pennants were bought.

In order to use these pennants for a long time, they were made challenge ones. When someone fixes a new critical bug, one of the earlier awarded "I fixed a critical bug" pennants is passed on to their table. When someone suggests a new interesting feature, one of the earlier awarded "I suggested a new feature" pennants is passed on to their table.

One person can have several pennants of one type, and they can have pennants of both types on their table. There are `n` tables in the IT company. Find the number of ways to place the pennants on these tables, given that each pennant is situated on one of the tables and each table is big enough to contain any number of pennants.

## Input

The only line of the input contains one integer `n` (`1 <= n <= 500`) — the number of tables in the IT company.

## Output

Output one integer — the number of ways to place the pennants on `n` tables.

## Example

**Input**

```
2
```

**Output**

```
24
```


### ideas
1. dp[i][j] 表示到目前为止，分别剩余i,j个pennats的ways