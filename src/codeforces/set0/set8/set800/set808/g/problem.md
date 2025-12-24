# Problem

Berland has a long and glorious history. To increase awareness about it among younger citizens, King of Berland decided to compose an anthem.

Though there are lots and lots of victories in history of Berland, there is the one that stand out the most. King wants to mention it in the anthem as many times as possible.

He has already composed major part of the anthem and now just needs to fill in some letters. King asked you to help him with this work.

The anthem is the string $s$ of no more than $10^5$ small Latin letters and question marks. The most glorious victory is the string $t$ of no more than $10^5$ small Latin letters. You should replace all the question marks with small Latin letters in such a way that the number of occurrences of string $t$ in string $s$ is maximal.

Note that the occurrences of string $t$ in $s$ can overlap. Check the third example for clarification.

## Input

The first line contains string of small Latin letters and question marks $s$ ($1 \leq |s| \leq 10^5$).

The second line contains string of small Latin letters $t$ ($1 \leq |t| \leq 10^5$).

Product of lengths of strings $|s| \cdot |t|$ won't exceed $10^7$.

## Output

Output the maximum number of occurrences of string $t$ you can achieve by replacing all the question marks in string $s$ with small Latin letters.

## Examples

### Example 1

```text
winlose???winl???w??
win
```

```text
5
```

### Example 2

```text
glo?yto?e??an?
or
```

```text
3
```

### Example 3

```text
??c?????
abcab
```

```text
2
```

## Note

In the first example the resulting string $s$ is "winlosewinwinlwinwin".

In the second example the resulting string $s$ is "glorytoreorand". The last letter of the string can be arbitrary.

In the third example occurrences of string $t$ are overlapping. String $s$ with maximal number of occurrences of $t$ is "abcabcab".


### ideas
1. 如果 s[i] != ‘?'， for s[i] != t[j], j = p[j-1]
2. dp[i][j] = dp[i-j][0] + j
3. 如果 s[i] = '?', 