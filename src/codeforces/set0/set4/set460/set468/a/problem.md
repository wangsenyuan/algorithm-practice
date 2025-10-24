# Problem

Little X used to play a card game called "24 Game", but recently he has found it too easy. So he invented a new game.

Initially you have a sequence of $n$ integers: $1, 2, ..., n$. In a single step, you can pick two of them, let's denote them $a$ and $b$, erase them from the sequence, and append to the sequence either $a + b$, or $a - b$, or $a \times b$.

After $n - 1$ steps there is only one number left. Can you make this number equal to $24$?

## Input

The first line contains a single integer $n$ ($1 \leq n \leq 10^5$).

## Output

If it's possible, print "YES" in the first line. Otherwise, print "NO" (without the quotes).

If there is a way to obtain $24$ as the result number, in the following $n - 1$ lines print the required operations an operation per line. Each operation should be in form: "a op b = c". Where $a$ and $b$ are the numbers you've picked at this operation; $op$ is either "+", or "-", or "*"; $c$ is the result of corresponding operation. Note, that the absolute value of $c$ mustn't be greater than $10^{18}$. The result of the last operation must be equal to $24$. Separate operator sign and equality sign from numbers with spaces.

If there are multiple valid answers, you may print any of them.

## Examples

### Example 1

**Input**
```
1
```

**Output**
```
NO
```

### Example 2

**Input**
```
8
```

**Output**
```
YES
8 * 7 = 56
6 * 5 = 30
3 - 4 = -1
1 - 2 = -1
30 - -1 = 31
56 - 31 = 25
25 + -1 = 24
```


### ideas
1. 如果有 n >= 4, 似乎一定有答案
2. 1 * 2 * 3 * 4 = 24
3. n > 4的时候， 如果能让后边的变成0，就可以了
4. 如果后面有可以整除4， 那么4个一组，就可以++--就可以变成0
5. 如果 n % 4 = 0, 
6. 如果 n % 4 = 1, 1, 2, 3, 4, 5 (5 * 4 + 3 + 2 - 1)
7. 如果 n % 4 = 2, [1 ~ 6], [4 * 6 + 1 * 2 + 3 - 5]
8. 如果 n % 4 = 3, [1 ~7], [3 * 7 + 6 - 5 + 4 - 2] * 1