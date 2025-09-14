# Problem Description

Daniel is organizing a football tournament. He has come up with the following tournament format:

## Tournament Rules

1. **Elimination Stages**: In the first several (possibly zero) stages, while the number of teams is even, they split in pairs and play one game for each pair. At each stage the loser of each pair is eliminated (there are no draws). Such stages are held while the number of teams is even.

2. **Final Stage**: Eventually there will be an odd number of teams remaining. If there is one team remaining, it will be declared the winner, and the tournament ends. Otherwise each of the remaining teams will play with each other remaining team once in round robin tournament (if there are x teams, there will be x(x-1)/2 games), and the tournament ends.

## Example

If there were 20 teams initially:
- They would begin by playing 10 games (20 teams → 10 teams)
- Then the remaining 10 would play 5 games (10 teams → 5 teams)  
- Then the remaining 5 teams would play 10 games in a round robin tournament (5 teams → 5×4/2 = 10 games)
- **Total**: 10 + 5 + 10 = 25 games

## Task

Daniel has already booked the stadium for n games. Help him to determine how many teams he should invite so that the tournament needs exactly n games. You should print all possible numbers of teams that will yield exactly n games in ascending order, or -1 if there are no such numbers.

## Input

The first line contains a single integer n (1 ≤ n ≤ 10^18), the number of games that should be played.

> **Note**: Do not use the %lld specifier to read or write 64-bit integers in C++. It is preferred to use the cin, cout streams or the %I64d specifier.

## Output

Print all possible numbers of invited teams in ascending order, one per line. If exactly n games cannot be played, output one number: -1.

## Examples

### Example 1
**Input:**
```
3
```

**Output:**
```
3
4
```

### Example 2
**Input:**
```
25
```

**Output:**
```
20
```

### Example 3
**Input:**
```
2
```

**Output:**
```
-1
```


### ideas
1. 假设最后剩余了x个team（x是奇数）
2. m = x * (x - 1) / 2 
3. 那么 n = m + w (w = n / 2 + n / 4 + n / 8 ===)
4. x = n / pow(2, i)
5. 这里迭代x不大行（它的量级在n的平方根，1e9)
6. 迭代i?可行吗
7. 假设进行0轮后，x = m,  n = x * (x - 1) / 2 (如果存在x) 那么就有一个值， 这个x必须是奇数
8. 如果迭代1轮后, x = m / 2, m / 2 + x * (x - 1) / 2 = n
9. 迭代2轮后， x = m / 4, m / 2 + m / 4 + x * (x - 1) / 2 = n
10. 感觉最多是60轮？ 好像是可行的