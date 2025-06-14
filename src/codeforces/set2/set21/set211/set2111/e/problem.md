Given a string $s$ that consists only of the first three letters of the Latin alphabet, meaning each character of the string is either **a**, **b**, or **c**.

Also given are $q$ operations that need to be performed on the string. In each operation, two letters $x$ and $y$ from the set of the first three letters of the Latin alphabet are provided, and for each operation, one of the following two actions must be taken:

- Change any (one) occurrence of the letter $x$ in the string $s$ to the letter $y$ (if at least one occurrence of the letter $x$ exists);
- Do nothing.

The goal is to perform all operations in the given order in such a way that the string $s$ becomes **lexicographically minimal**.

Recall that a string $a$ is lexicographically less than a string $b$ if and only if one of the following conditions holds:

- $a$ is a prefix of $b$, but $a \neq b$;
- At the first position where $a$ and $b$ differ, the string $a$ has a letter that comes earlier in the alphabet than the corresponding letter in $b$.

---

### Input

Each test consists of several test cases. The first line contains a single integer $t$ ($1 \leq t \leq 10^3$) — the number of test cases. The description of the test cases follows.

In the first line of each test case, there are two integers $n$ and $q$ ($1 \leq n, q \leq 2 \cdot 10^5$) — the length of the string $s$ and the number of operations.

In the second line of each test case, the string $s$ is given — a string of exactly $n$ characters, each of which is **a**, **b**, or **c**.

The next $q$ lines of each test case contain the description of the operations. Each line contains two characters $x$ and $y$, each of which is **a**, **b**, or **c**.

**Additional constraints on the input:**

- The sum of $n$ across all test cases does not exceed $2 \cdot 10^5$;
- The sum of $q$ across all test cases does not exceed $2 \cdot 10^5$.

---

### Output

For each test case, output the lexicographically minimal string that can be obtained from $s$ using the given operations.


### ideas
1. 可以不进行操作
2. 考虑最后的字符串，如果可以把c变成a，那么肯定是最前面的c
3. 如果能够把b变成a，也肯定是最前面的b
4. 也就是说，始终是选择第一个不是a的字符，看能否把它变成a
5. 只要能，就应该使用这个变化
6. 但是假设存在一个b到a的操作，不管这个位置在哪里，都优先使用它
7. 如果不存在b到a，但是存在 b -> c -> a，也应该使用