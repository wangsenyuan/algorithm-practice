In a very ancient country the following game was popular. Two people play the game. Initially the first player writes a string $s_1$, consisting of exactly nine digits and representing a number that does not exceed $a$. After that the second player looks at $s_1$ and writes a string $s_2$, consisting of exactly nine digits and representing a number that does not exceed $b$. Here $a$ and $b$ are some given constants, $s_1$ and $s_2$ are chosen by the players. The strings are allowed to contain leading zeroes.

If a number obtained by the concatenation (joining together) of strings $s_1$ and $s_2$ is divisible by $mod$, then the second player wins. Otherwise the first player wins. You are given numbers $a$, $b$, $mod$. Your task is to determine who wins if both players play in the optimal manner. If the first player wins, you are also required to find the lexicographically minimum winning move.

## Input

The first line contains three integers $a$, $b$, $mod$ ($0 \le a, b \le 10^9$, $1 \le mod \le 10^7$).

## Output

If the first player wins, print `1` and the lexicographically minimum string $s_1$ he has to write to win. If the second player wins, print the single number `2`.

## Examples

### Input
```
1 10 7
```

### Output
```
2
```

### Input
```
4 0 9
```

### Output
```
1 000000001
```

## Note

The lexical comparison of strings is performed by the `<` operator in modern programming languages. String $x$ is lexicographically less than string $y$ if there exists such $i$ ($1 \le i \le 9$), that $x_i < y_i$, and for any $j$ ($1 \le j < i$) $x_j = y_j$. These strings always have length $9$.



### ideas
1. 要找到一个s1, 使的， 000..0 ~ 999..9 都没法和s1组合出 + s1 * 1e9  % mod = 0
2. s2貌似不用管， 因为它肯定可以通过某种方式，得到0...mod-1的结果
3. 所以这里的关键在于 s2 <= b
4. 如果 b >= mod => 2 (无论如何， 第二个人，总能根据s1， 得到某个s2, 使的 s1 * 1e9 + s2 % mod = 0)
5. b < mod， 那么 b + 1, b + 2 .... mod - 1个数是第二个无法获得数
6. 那么此时，就需要判断 是否存在s1 <= a, 且 s1 * 1e9 + s2 ~= 0 这个时候应该是个范围，然后找出一个公共范围出来