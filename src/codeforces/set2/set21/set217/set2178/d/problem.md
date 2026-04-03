# D. Mass Hysteria

There are \(n\) elves numbered \(1\) to \(n\), where elf \(i\) has a health value of \(h_i\) and an attack value of \(a_i\). Initially, \(h_i = a_i\) for all \(i\), and all \(a_i\) are distinct. Elf \(i\) is alive if and only if its health is positive (i.e., \(h_i > 0\)).

When Franklin casts **Mass Hysteria**, the following process is repeated:

1. Choose a pair of distinct living elves \(x\) and \(y\) (\(h_x, h_y > 0\)) such that elf \(x\) has not attacked before. If no such pair exists, terminate the process.
2. Then, elf \(x\) attacks elf \(y\), decreasing \(h_y\) by \(a_x\). Additionally, due to recoil, \(h_x\) is decreased by \(a_y\). Note that \(a_x\) and \(a_y\) remain unchanged.

The process repeats until it is impossible to choose a valid pair of elves. It can be shown that Mass Hysteria terminates after at most \(n\) iterations.

Given an integer \(m\), construct a valid sequence of attacks such that exactly \(m\) elves are alive when the process ends; or determine that no such sequence exists.

## Input

Each test contains multiple test cases. The first line contains the number of test cases \(t\) (\(1 \le t \le 10^4\)). The description of the test cases follows.

The first line of each test case contains two integers \(n\) and \(m\) (\(2 \le n \le 2 \cdot 10^5\), \(0 \le m \le n\)) — the number of elves in the village and the number of elves to be left alive.

The second line contains \(n\) integers \(a_1, a_2, \ldots, a_n\) (\(1 \le a_i \le 10^9\)) — the initial attack and health values of the elves.

It is guaranteed that all \(a_i\) are distinct.

It is guaranteed that the sum of \(n\) over all test cases does not exceed \(2 \cdot 10^5\).

## Output

For each test case, if it is impossible for exactly \(m\) elves to remain alive, print \(-1\).

Otherwise, output a valid sequence of attacks as follows:

- On the first line, output an integer \(k\) (\(0 \le k \le n\)) — the number of iterations in Mass Hysteria.
- Then \(k\) lines follow; the \(i\)-th line contains two integers \(x_i\) and \(y_i\) (\(1 \le x_i, y_i \le n\), \(x_i \ne y_i\)), indicating that elf \(x_i\) attacks elf \(y_i\) in the \(i\)-th iteration.

The sequence must satisfy all of the following conditions:

- Immediately before the \(i\)-th iteration, both elves \(x_i\) and \(y_i\) are alive and elf \(x_i\) has not attacked in any previous iteration.
- After the \(k\)-th iteration, exactly \(m\) elves are alive and there does not exist a pair of distinct living elves \(x\) and \(y\) such that elf \(x\) has not attacked before.

If there are multiple answers, you may output any of them. Any valid sequence satisfying the above conditions will be accepted.

## Example

**Input**

```
7
4 2
1 4 2 3
2 2
6 7
3 0
1 2 3
3 1
1 2 3
3 2
1 2 3
4 1
2 3 4 5
6 0
998244353 1000000000 314159265 676767677 999999999 987654321
```

**Output**

```
2
3 1
2 4
-1
2
3 2
1 3
2
1 2
3 2
-1
2
1 4
4 2
4
3 1
2 5
6 1
4 2
```

## Note

In the first test case, one possible sequence of attacks is shown below:

| Step | \(x\) | \(y\) | Health values after attack | Elves that have attacked |
| ---: | :---: | :---: | :------------------------- | :----------------------- |
|    0 |   —   |   —   | \([1,4,2,3]\)              | \([]\)                   |
|    1 |   3   |   1   | \([−1,4,1,3]\)             | \([3]\)                  |
|    2 |   2   |   4   | \([−1,1,1,−1]\)            | \([2,3]\)                |

After \(2\) iterations, only elves \(2\) and \(3\) are alive. Since both of them have already attacked, no further valid attack is possible, and Mass Hysteria terminates.

In the second test case, the only possible choices for \((x, y)\) in the first iteration are \((1,2)\) or \((2,1)\). In either case, elf \(1\) ends up with \(-1\) health, so it is impossible for both elves to remain alive at the end. Note that Mass Hysteria will last at least one iteration as there exists a valid \((x, y)\) for the first iteration.

In the sixth test case, only elf \(3\) is alive after all attacks. Even though elf \(3\) has not attacked before, Mass Hysteria terminates since there is no other elf it can attack.


### ideas
1. 这个题目很奇怪，必须选到，所有活着的人，都攻击过别人了
2. 所以，h[i]最大的人，肯定是活着的（它攻击了别人，别人直接挂了）（好像也不一定，它可以被其他人攻击）
3. 那么它可以攻击一个（攻击过其他人的）人吗？可以的，对y没有要求
4. 如果m=1, 那么就可以按照h升序，依次攻击
5. 如果m=n/2, 那么以两人一组，相互攻击（剩下一半的人）；如果多出来一个人，让它去攻击某个人
6. 正常情况下肯定会再挂一个人，但是也可能多挂一个人（选择一个和剩下人相同h的人）
7. 不可能> n / 2
8. 那么 1...n/2一定可以达到吗？是可以的
9. 假设 m = 1...，先用策略1，达到某个数n1 = 2 * m 即可
10. 然后剩余的人，按照策略2进行