# Problem

Well, the series which Stepan watched for a very long time, ended. In total, the series had $n$ episodes. For each of them, Stepan remembers either that he definitely has watched it, or that he definitely hasn't watched it, or he is unsure whether he has watched this episode or not.

Stepan's dissatisfaction is the maximum number of consecutive series that Stepan did not watch.

Your task is to determine according to Stepan's memories if his dissatisfaction could be exactly equal to $k$.

## Input

The first line contains two integers $n$ and $k$ ($1 \le n \le 100$, $0 \le k \le n$) — the number of episodes in the series and the dissatisfaction which should be checked.

The second line contains a sequence of $n$ symbols `"Y"`, `"N"` and `"?"`. If the $i$-th symbol equals `"Y"`, Stepan remembers that he has watched episode $i$. If the $i$-th symbol equals `"N"`, Stepan remembers that he hasn't watched episode $i$. If the $i$-th symbol equals `"?"`, Stepan doesn't exactly remember whether he has watched episode $i$ or not.

## Output

If Stepan's dissatisfaction can be exactly equal to $k$, then print `YES` (without quotes). Otherwise print `NO` (without quotes).

## Examples

### Example 1

**Input**

```
5 2
NYNNY
```

**Output**

```
YES
```

### Example 2

**Input**

```
6 1
????NN
```

**Output**

```
NO
```

## Note

In the first test Stepan remembers about all the episodes whether he has watched them or not. His dissatisfaction is $2$, because he hasn't watched two episodes in a row — episode number $3$ and episode number $4$. The answer is `YES`, because $k = 2$.

In the second test $k = 1$. Stepan's dissatisfaction is greater than or equal to $2$ (because he remembers that he hasn't watched at least two episodes in a row — numbers $5$ and $6$), even if he has watched the episodes from the first to the fourth, inclusive.
