# Problem D: Dima and Inna's Game

## Problem Description

Dima spent much time thinking about what present to give to Inna and gave her an empty sequence `w`. Now they want to fill sequence `w` with numbers zero and one. For that, they decided to play an amusing game.

### Game Rules

Before the game begins, Dima chooses `m` integers `a₁, a₂, ..., aₘ` (1 ≤ a₁ < a₂ < ... < aₘ). Then Inna and Dima start playing by adding numbers to sequence `w`. Each new number they choose is added to the end of the sequence.

At some moments of time, Dima feels that the game is going to end too soon (and he wants to play with Inna as long as possible), so he hits a table hard with his fist. At that moment:
- The a₁-th, a₂-th, a₃-th, ..., aₖ-th numbers from the beginning simultaneously fall out of the sequence
- The sequence gets `k` numbers less
- Here `k` is such maximum number that value `aₖ` doesn't exceed the current length of the sequence
- If number `a₁` is larger than the current length of `w`, then nothing falls out of the sequence

### Task

You are given the chronological sequence of events in the game. Each event is either:
- Adding a number to the end of sequence `w`, or
- Dima's hit on the table

Calculate the sequence `w` after all these events happen.

## Input

The first line contains two integers `n` and `m` (1 ≤ n, m ≤ 10⁶):
- `n`: number of events that took place
- `m`: number of numbers Dima chose

The second line contains `m` distinct integers `aᵢ` (1 ≤ aᵢ ≤ 10⁶) sorted in increasing order.

The next `n` lines describe the events in chronological order. Each line contains a single integer:
- `-1`: Dima hits the table
- `0`: Inna and Dima add number 0 to the end of the sequence
- `1`: Inna and Dima add number 1 to the end of the sequence

## Output

Print a single line containing the sequence of numbers 0 and 1 — the elements of the sequence after all events happen. Print the elements in order from the beginning to the end of the sequence.

**Special case**: If after all events the sequence ends up empty, print `"Poor stack!"`.

## Example

### Input
```
5 2
1 3
1
0
-1
1
-1
```

### Output
```
1
```

### Explanation
1. Add 1: sequence becomes `[1]`
2. Add 0: sequence becomes `[1, 0]`
3. Hit table (remove 1st element): sequence becomes `[0]`
4. Add 1: sequence becomes `[0, 1]`
5. Hit table (remove 1st and 3rd elements, but 3rd doesn't exist): sequence becomes `[1]`

### ideas
1. 使用一个segment tree似乎也可以？
2. 就是每次找到第k个，k-1个, ... 的位置，然后删除掉