## Palindrome Transformation

Nam is playing with a string on his computer. The string consists of n lowercase English letters. It is meaningless, so
Nam decided to make the string more beautiful — to make it a palindrome — using four arrow keys: left, right, up, and
down.

There is a cursor pointing at some character of the string. Suppose the cursor is at position i (1 ≤ i ≤ n, the string
uses 1-based indexing). The string is cyclic for horizontal movement:

- Pressing the left arrow moves the cursor to position i − 1 if i > 1, or to position n otherwise.
- Pressing the right arrow moves the cursor to position i + 1 if i < n, or to position 1 otherwise.

Vertical keys modify the character under the cursor (the alphabet is cyclic):

- Pressing the up arrow changes the letter to the next letter in the English alphabet (after 'z' follows 'a').
- Pressing the down arrow changes the letter to the previous letter (before 'a' goes to 'z').

Initially, the text cursor is at position p.

Because Nam has a lot of homework to do, he wants to complete this as fast as possible. Calculate the minimum number of
arrow key presses required to make the string a palindrome.

## Input

- The first line contains two space-separated integers n (1 ≤ n ≤ 105) and p (1 ≤ p ≤ n), the length of Nam's string and
  the initial position of the text cursor.
- The second line contains n lowercase English letters — Nam's string.

## Output

Print the minimum number of key presses needed to change the string into a palindrome.

## Example

Input

```
8 3
aeabcaez
```

Output

```
6
```

## Notes

- A string is a palindrome if it reads the same forward and backward.
- In the sample test, initial Nam's string is: (cursor position is shown in bold).
- In an optimal solution, Nam may do 6 steps; the result is a palindrome.

## ideas

1. 对应位置的上、下的cost是固定的, diff(s[i] - s[j])
2. 所以主要考虑移动的问题
3. 左右移动，还是会被上下移动所限制，就是假设某个位置后面，没有不match的位置，那么它就不需要移动到那里去
4. 否则的话，还是得移动
5. 