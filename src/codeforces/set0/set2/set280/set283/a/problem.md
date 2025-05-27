# Problem Statement

Bessie and the cows are playing with sequences and need your help. They start with a sequence, initially containing just the number `0`, and perform `n` operations. Each operation is one of the following:

1. **Add** the integer `xi` to the first `ai` elements of the sequence.
2. **Append** an integer `ki` to the end of the sequence (the size of the sequence increases by 1).
3. **Remove** the last element of the sequence (the size of the sequence decreases by one). Note: this operation can only be done if there are at least two elements in the sequence.

After each operation, the cows would like to know the average of all the numbers in the sequence. Help them!

---

## Input

- The first line contains a single integer `n` ($1 \leq n \leq 2 \cdot 10^5$) — the number of operations.
- The next `n` lines describe the operations. Each line will start with an integer `ti` ($1 \leq ti \leq 3$), denoting the type of the operation (see above):
  - If `ti = 1`, it will be followed by two integers `ai`, `xi` ($|xi| \leq 10^3$; $1 \leq ai$).
  - If `ti = 2`, it will be followed by a single integer `ki` ($|ki| \leq 10^3$).
  - If `ti = 3`, it will not be followed by anything.

It is guaranteed that all operations are correct (don't touch nonexistent elements) and that there will always be at least one element in the sequence.

---

## Output

Output `n` lines, each containing the average of the numbers in the sequence after the corresponding operation.

The answer will be considered correct if its absolute or relative error doesn't exceed $10^{-6}$.

## ideas
1. 假设维护了sum, cnt, 操作1 = sum += xi * ai
2. 操作2， sum += ki, cnt++
3. 操作3， 比较麻烦, cnt--, sum -= 最后一个数的大小
4. 主要是操作3. 
5. 假设有这么一个操作序列， 先增加size到10， 然后操作1，增加前10个数字x
6. 然后pop掉，size变成9，再push，再pop。
7. 所以，要知道在这次pop掉的是多少，必须记录在push进来时，10处的数字，old_value
8. 然后在pop的时候，ki + new_value - old_value就是真实的变化