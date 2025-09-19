# Problem B: Vova and Tests

## Problem Statement

Little Vova studies programming in an elite school. Vova and his classmates are supposed to write $n$ progress tests, for each test they will get a mark from $1$ to $p$. Vova is very smart and he can write every test for any mark, but he doesn't want to stand out from the crowd too much.

- If the sum of his marks for all tests exceeds value $x$, then his classmates notice how smart he is and start distracting him asking to let them copy his homework.
- If the median of his marks will be lower than $y$ points (the definition of a median is given in the notes), then his mom will decide that he gets too many bad marks and forbid him to play computer games.

Vova has already wrote $k$ tests and got marks $a_1, \ldots, a_k$. He doesn't want to get into the first or the second situation described above and now he needs to determine which marks he needs to get for the remaining tests. Help him do that.

## Input

The first line contains $5$ space-separated integers: $n$, $k$, $p$, $x$ and $y$ ($1 \leq n \leq 999$, $n$ is odd, $0 \leq k < n$, $1 \leq p \leq 1000$, $n \leq x \leq n \cdot p$, $1 \leq y \leq p$). Here:

- $n$ is the number of tests that Vova is planned to write
- $k$ is the number of tests he has already written
- $p$ is the maximum possible mark for a test
- $x$ is the maximum total number of points so that the classmates don't yet disturb Vova
- $y$ is the minimum median point so that mom still lets him play computer games

The second line contains $k$ space-separated integers: $a_1, \ldots, a_k$ ($1 \leq a_i \leq p$) — the marks that Vova got for the tests he has already written.

## Output

If Vova cannot achieve the desired result, print `-1`.

Otherwise, print $n - k$ space-separated integers — the marks that Vova should get for the remaining tests. If there are multiple possible solutions, print any of them.

## Examples

### Example 1

**Input:**
```
5 3 5 18 4
3 5 4
```

**Output:**
```
4 1
```

### Example 2

**Input:**
```
5 3 5 16 4
5 5 5
```

**Output:**
```
-1
```

## Note

The median of sequence $a_1, \ldots, a_n$ where $n$ is odd (in this problem $n$ is always odd) is the element staying on $(n + 1) / 2$ position in the sorted list of $a_i$.

In the first sample:

- The sum of marks equals $3 + 5 + 4 + 4 + 1 = 17$, which doesn't exceed $18$, so Vova won't be disturbed by his classmates
- The median point of the sequence $\{1, 3, 4, 4, 5\}$ equals to $4$, which isn't less than $4$, so his mom lets him play computer games

Please note that you do not have to maximize the sum of marks or the median mark. Any of the answers: `"4 2"`, `"2 4"`, `"5 1"`, `"1 5"`, `"4 1"`, `"1 4"` for the first test is correct.

In the second sample Vova got three `'5'` marks, so even if he gets two `'1'` marks, the sum of marks will be $17$, that is more than the required value of $16$. So, the answer to this test is `-1`.

## Ideas

1. 在保证median >= y 的情况下，让总分不超过x