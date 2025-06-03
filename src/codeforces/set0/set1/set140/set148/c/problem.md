# Princess and Grooms

## Problem Statement

*«Next please»*, — the princess called and cast an estimating glance at the next groom.

The princess intends to choose the most worthy groom, this is, the richest one. Whenever she sees a groom who is more rich than each of the previous ones, she says a measured *«Oh...»*. Whenever the groom is richer than all previous ones added together, she exclaims *«Wow!»* (no *«Oh...»* in this case). At the sight of the first groom the princess stays calm and says nothing.

The fortune of each groom is described with an integer between 1 and 50000. You know that during the day the princess saw `n` grooms, said *«Oh...»* exactly `a` times and exclaimed *«Wow!»* exactly `b` times. 

Your task is to output a sequence of `n` integers `t1, t2, ..., tn`, where `ti` describes the fortune of i-th groom. If several sequences are possible, output any of them. If no sequence exists that would satisfy all the requirements, output a single number `-1`.

## Input

The only line of input data contains three integer numbers `n`, `a` and `b` (`1 ≤ n ≤ 100`, `0 ≤ a, b ≤ 15`, `n > a + b`), separated with single spaces.

## Output

Output any sequence of integers `t1, t2, ..., tn`, where `ti` (`1 ≤ ti ≤ 50000`) is the fortune of i-th groom, that satisfies the given constraints. If no sequence exists that would satisfy all the requirements, output a single number `-1`.

## Constraints

- `1 ≤ n ≤ 100`
- `0 ≤ a, b ≤ 15`
- `n > a + b`
- `1 ≤ ti ≤ 50000` for each fortune value

## ideas
1. Wow表示 t[i] > sum for now
2. Oh表示t[i] > max(pref) for now
3. 那么先把Oh制作出来？ 1, 2, 3, 4, 5, a + 1 (a + 1)， 然后处理b-1个（问题在于后面的增长太快，不一定能在50000的范围内完成）
4. 1, 2, 4, 8, ... (b+1个数字) + a个递增的数字就可以了