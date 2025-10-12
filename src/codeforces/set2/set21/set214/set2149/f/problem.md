Nezuko suddenly woke up on the number line at point 0 and has ℎ health points. She wants to reach point 𝑑. In one turn, she can do exactly one of two things:

- rest in the shade and increase her current health by 1;
- move from her current position 𝑥 to 𝑥+1.

Each movement decreases Nezuko's health; if the movement is the 𝑗-th consecutive movement, her health will decrease by 𝑗 points. If as a result of a move Nezuko's health drops to 0 or below, she cannot make that move.

For example, if Nezuko initially had 7 health points and 𝑑=4, her moves could look like this:

- Move from 0 to 1 and decrease health by 1. Now she is at point 1 with 6 health points.
- Move from 1 to 2 and decrease health by 2. Now she is at point 2 with 4 health points.
- Move from 2 to 3 and decrease health by 3. Now she is at point 3 with 1 health point.
- Rest and restore 1 health point. Now she is at point 3 with 2 health points.
- Move from 3 to 4 and decrease health by 1. Now she is at point 4 with 1 health point.

Find the minimum number of turns required for her to reach point 𝑑.

## Input

Each test consists of several test cases.

The first line contains one integer 𝑡 (1≤𝑡≤10⁴) — the number of test cases. The following describes the test cases.

The first line of each test case contains two integers ℎ and 𝑑 (1≤ℎ,𝑑≤10⁹) — the number of health points and the destination point, respectively.

## Output

For each test case, output one number — the minimum number of turns required for Nezuko to reach point 𝑑.

## Example

### Input
```
5
3 2
1 1
5 3
2 4
10 7
```

### Output
```
3
2
4
7
10
```

## Note

In the first test case, ℎ=3, 𝑑=2 the actions could be as follows:

- Move from 0 to 1 and decrease health by 1. Now she is at point 1 with 2 health points.
- Rest and restore 1 health point. Now she is at point 1 with 3 health points.
- Move from 1 to 2 and decrease health by 1. Now she is at point 2 with 2 health points.

In total, 3 turns.

In the fourth test case, ℎ=2, 𝑑=4 the actions could be as follows:

- Move from 0 to 1 and decrease health by 1. Now she is at point 1 with 1 health point.
- Rest and restore 1 health point. Now she is at point 1 with 2 health points.
- Move from 1 to 2 and decrease health by 1. Now she is at point 2 with 1 health point.
- Rest and restore 1 health point. Now she is at point 2 with 2 health points.
- Move from 2 to 3 and decrease health by 1. Now she is at point 3 with 1 health point.
- Rest and restore 1 health point. Now she is at point 3 with 2 health points.
- Move from 3 to 4 and decrease health by 1. Now she is at point 4 with 1 health point.

In total, 7 turns.


### ideas
1. 这个应该是具备二分的性质的。但是策略是什么呢？ 休息一秒，然后move，这样交替进行吗？
2. 还是尽量的move？
3. 如果当前位置x, 能够在到达d之前，hp不减到0， hp > dist * (dist + 1) / 2， 那么就move
4. 移动的时间是定的，就是d. 所以，这里要最小化等待的时间w
5. 假设分配的位置为 x1, x2, ... x2
6. h0 = hp, h1 = h0 - d1 * (d1 + 1) / 2 + 1 (等了一秒), h2 = h1 - d2 * (d2 + 1) / 2 + 1 ...
7. hi = h(i-1) - di * (di + 1) / 2 + 1
8. hw = h(w-1) - dw * (dw + 1) / 2 > 0
9. hw = hp - d1 * (d1 + 1) / 2 + 1 - d2 * (d2 + 1) / 2 + 1 - ... - dw * (dw + 1) / 2
10.   = hp + w (停了w次) - (d1 * (d1 + 1) / 2 + d2 * (d2 + 1) / 2 + ... + dw * (dw + 1) / 2)
11.   => 也就是后面的 sum = d1 * (d1 + 1) / 2 + d2 * (d2 + 1) / 2 + ... + dw * (dw + 1) / 2 最小
12. 还有个条件时 hi > 1, hw > 0
13. d1 + d2 + ... + dw = d
14. 这里基本可以判断出 d1 >= d2 >= ... >= d(w-1) (因为前面的hp高，可以跑的远一些，后面的hp低，要多休息)
15. 感觉就是要让d1 .... dw 差不多相同的时候， sum能取最小值 这个貌似又可以二分一次（尽量的让dw大）
16. 但是这个不大行，因为w会很大，没法迭代～