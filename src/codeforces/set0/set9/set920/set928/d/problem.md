# D. Autocompletion

[Problem link](https://codeforces.com/problemset/problem/928/D)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Arcady types a text in an editor. Each letter and each sign (including line
feed) costs one keyboard click.

When Arcady has a non-empty prefix of some word on the screen, the editor may
propose autocompletion: the unique already printed word that has this prefix
(if such a word is unique among printed words). With one click Arcady can
accept the proposal and replace the current prefix with that full word. No
extra symbols are printed after autocompletion.

Arcady cannot move the cursor or erase printed symbols. Find the minimum
number of clicks to print the entire text.

A word is a contiguous sequence of Latin letters bordered by spaces,
punctuation, or line/text boundaries. Only lowercase letters are used.

## Constraints

- Text consists of lowercase Latin letters, spaces, line feeds, and
  punctuation: `.` `,` `?` `!` `'` `-`
- Total length ≤ `3 * 10^5`
- All lines are non-empty

## Input

The text (possibly multiple lines).

## Output

One integer — the minimum number of clicks.

## Samples

### Sample 1

Input:

```
snow affects sports such as skiing, snowboarding, and snowmachine travel.
snowboarding is a recreational activity and olympic and paralympic sport.
```

Output:

```
141
```

### Sample 2

Input:

```
'co-co-co, codeforces?!'
```

Output:

```
25
```

### Sample 3

Input:

```
thun-thun-thunder, thunder, thunder
thunder, thun-, thunder
thun-thun-thunder, thunder
thunder, feel the thunder
lightning then the thunder
thunder, feel the thunder
lightning then the thunder
thunder, thunder
```

Output:

```
183
```
