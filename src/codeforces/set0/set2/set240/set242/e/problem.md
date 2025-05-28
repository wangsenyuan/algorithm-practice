# Problem Statement

You've got an array $a$, consisting of $n$ integers $a_1, a_2, \ldots, a_n$. You are allowed to perform two operations on this array:

1. **Calculate the sum** of current array elements on the segment $[l, r]$, that is, count value $a_l + a_{l+1} + \ldots + a_r$.
2. **Apply the xor operation** with a given number $x$ to each array element on the segment $[l, r]$, that is, execute $a_i = a_i \oplus x$ for all $i$ in $[l, r]$. This operation changes exactly $r - l + 1$ array elements.

Expression $x \oplus y$ means applying bitwise xor operation to numbers $x$ and $y$. The given operation exists in all modern programming languages, for example in C++ and Java it is marked as `^`, in Pascal — as `xor`.

You've got a list of $m$ operations of the indicated type. Your task is to perform all given operations, for each sum query you should print the result you get.

---

## Input

- The first line contains integer $n$ ($1 \leq n \leq 10^5$) — the size of the array.
- The second line contains space-separated integers $a_1, a_2, \ldots, a_n$ ($0 \leq a_i \leq 10^6$) — the original array.
- The third line contains integer $m$ ($1 \leq m \leq 5 \cdot 10^4$) — the number of operations with the array.
- The $i$-th of the following $m$ lines first contains an integer $t_i$ ($1 \leq t_i \leq 2$) — the type of the $i$-th query.
    - If $t_i = 1$, then this is the query of the sum, and next follow two integers $l_i, r_i$ ($1 \leq l_i \leq r_i \leq n$).
    - If $t_i = 2$, then this is the query to change array elements, and next follow three integers $l_i, r_i, x_i$ ($1 \leq l_i \leq r_i \leq n, 1 \leq x_i \leq 10^6$).
- The numbers on the lines are separated by single spaces.

---

## Output

For each query of type 1 print in a single line the sum of numbers on the given segment. Print the answers to the queries in the order in which the queries go in the input.

---

*Please, do not use the %lld specifier to read or write 64-bit integers in C++. It is preferred to use the cin, cout streams, or the %I64d specifier.*

### ideas
1. 拆成30位，每一位，在xor的时候，只需要知道这个区域那set的数量，反转就可以了
2. 就可以同时操作xor和sum了