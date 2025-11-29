# Problem C

Petya recieved a gift of a string s with length up to 105 characters for his birthday. He took two more empty strings t and u and decided to play a game. This game has two possible moves:

1. Extract the first character of s and append t with this character.
2. Extract the last character of t and append u with this character.

Petya wants to get strings s and t empty and string u lexicographically minimal.

You should write a program that will help Petya win the game.

## Input

First line contains non-empty string s (1 ≤ |s| ≤ 105), consisting of lowercase English letters.

## Output

Print resulting string u.

## Examples

### Example 1

**Input:**

```text
cab
```

**Output:**

```text
abc
```

### Example 2

**Input:**

```text
acdb
```

**Output:**

```text
abdc
```


### ideas
1. t有点像一个stack， 假设目前stack的栈顶是w, 如果后面有任何一个字符比w小，那么至少要把这个加入进来
2. 