# Text Editor Correction

## Problem Statement
You wanted to write a text `t` consisting of `m` lowercase Latin letters. But instead, you have written a text `s` consisting of `n` lowercase Latin letters, and now you want to fix it by obtaining the text `t` from the text `s`.

Initially, the cursor of your text editor is at the end of the text `s` (after its last character). In one move, you can do one of the following actions:

- Press the **left** button: move the cursor left by one position (or do nothing if it is pointing at the beginning of the text, i.e., before its first character).
- Press the **right** button: move the cursor right by one position (or do nothing if it is pointing at the end of the text, i.e., after its last character).
- Press the **home** button: move the cursor to the beginning of the text (before the first character).
- Press the **end** button: move the cursor to the end of the text (after the last character).
- Press the **backspace** button: remove the character before the cursor (if there is no such character, nothing happens).

Your task is to calculate the minimum number of moves required to obtain the text `t` from the text `s` using the given set of actions, or determine if it is impossible to obtain the text `t` from the text `s`.

## Input
- The first line contains one integer `T` (`1 ≤ T ≤ 5000`) — the number of test cases.
- For each test case:
  - The first line contains two integers `n` and `m` (`1 ≤ m ≤ n ≤ 5000`) — the length of `s` and the length of `t`, respectively.
  - The second line contains the string `s` consisting of `n` lowercase Latin letters.
  - The third line contains the string `t` consisting of `m` lowercase Latin letters.
- It is guaranteed that the sum of `n` over all test cases does not exceed 5000 (`∑n ≤ 5000`).

## Output
For each test case, print one integer — the minimum number of moves required to obtain the text `t` from the text `s` using the given set of actions, or `-1` if it is impossible to obtain the text `t` from the text `s` in the given test case.