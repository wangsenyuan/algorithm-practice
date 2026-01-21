# Problem Statement

We had a string $s$ consisting of $n$ lowercase Latin letters. We made $k$ copies of this string, thus obtaining $k$ identical strings $s_1, s_2, \ldots, s_k$. After that, in each of these strings we swapped exactly two characters (the characters we swapped could be identical, but they had different indices in the string).

You are given $k$ strings $s_1, s_2, \ldots, s_k$, and you have to restore any string $s$ so that it is possible to obtain these strings by performing the aforementioned operations. Note that the total length of the strings you are given doesn't exceed $5000$ (that is, $k \cdot n \leq 5000$).

## Input

The first line contains two integers $k$ and $n$ $(1 \leq k \leq 2500$, $2 \leq n \leq 5000$, $k \cdot n \leq 5000)$ — the number of strings we obtained, and the length of each of these strings.

Next $k$ lines contain the strings $s_1, s_2, \ldots, s_k$, each consisting of exactly $n$ lowercase Latin letters.

## Output

Print any suitable string $s$, or `-1` if such string doesn't exist.

## Examples

### Input 1

```
3 4
abac
caab
acba
```

### Output 1

```
acab
```

### Input 2

```
3 4
kbbu
kbub
ubkb
```

### Output 2

```
kbub
```

### Input 3

```
5 4
abcd
dcba
acbd
dbca
zzzz
```

### Output 3

```
-1
```

## Note

In the first example, $s_1$ is obtained by swapping the second and the fourth character in `acab`, $s_2$ is obtained by swapping the first and the second character, and to get $s_3$, we swap the third and the fourth character.

In the second example, $s_1$ is obtained by swapping the third and the fourth character in `kbub`, $s_2$ — by swapping the second and the fourth, and $s_3$ — by swapping the first and the third.

In the third example, it's impossible to obtain the given strings by the aforementioned operations.


### ideas
1. let s = abcde
2. 任意两个 s[i], s[j], 它们之间存在几种情况，完全相同，有n-4个位置相同, 剩下的4个位置不相同，但频次相同
3. s[1] = bacde, s[2] = acbde 有n-3个位置相同, 剩下3个位置不相同，但是频次相同
4. 好像就不存在其他的情况了（如果都选择了相同的位置，l，r)， 那么 s[i] = s[j]
5. 以s[1]为基准，根据情况1和情况2，可以确定它的交换范围(如果完全相同的，可以跳过)（也就确定了s的剩余部分）
6. 假设有个s[i]和s[1]确定的部分重合了，在s[1]中变化的(i, j), 在s[i]中变化的是(l, r)