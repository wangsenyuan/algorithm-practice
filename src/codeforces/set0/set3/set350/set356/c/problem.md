# Problem

A team of students from the city S is sent to the All-Berland Olympiad in Informatics. They travel by train in a single carriage with $n$ compartments, each compartment having exactly 4 seats.

**Compartment Classifications:**
- **Bored**: a compartment with 1 or 2 students
- **Fun**: a compartment with 3 or 4 students

**Rules:**
- Students may swap seats with other passengers
- A sympathetic conductor can persuade any passenger to switch seats with a student
- Each persuaded passenger counts as 1

**Objective:**
Find the minimum number of people the conductor must persuade so that, after all swaps, every compartment with students has either 3 or 4 students (or 0 students).

If it is impossible, print -1.

## Input

- The first line contains an integer $n$ ($1 \le n \le 10^6$): the number of compartments
- The second line contains $n$ integers $a_1, a_2, \dots, a_n$ ($0 \le a_i \le 4$), where $a_i$ is the number of students in the $i$-th compartment
- It is guaranteed that at least one student is on the train

## Output

Print a single integer: the minimum number of passengers to persuade, or -1 if it is impossible.

## Examples

### Example 1

**Input:**
```
5
1 2 2 4 3
```

**Output:**
```
2
```

### Example 2

**Input:**
```
3
4 1 1
```

**Output:**
```
2
```

### Example 3

**Input:**
```
4
0 3 0 4
```

**Output:**
```
0
```

### ideas
1. 所有 a[i] = 4 的可以直接排除掉
2. a[i] = 1 肯定要移动， a[i] = 2可以不移动， 
3. 如果sum <= 2, 或者 sum = 5 没有答案
4. a[i] = 4还不能被排除，考虑 x = 4 + 2, 6个人是可以被分配对的，但是如果是2个人，就无法分配了
5. a[i] = 2 的优先和 a[i] = 1的匹配（这样只需要移动1个，解决两个问题）
6. 然后a[i] = 2和 a[i] = 4的进行匹配（这样只需要移动1个人，解决一个问题）
7. 然后如果还有多的 a[i] = 2 的部分 （这个时候应该没有 a[i] = 1 和 a[i] = 4)
8. freq[a[i] = 2] / 3 个一组，可以分解一个2, 得到两个3， 花费是2
9. 如果最后剩于两个，合并它们，花费2；如果剩于1个，看是否有两个3（不管怎么产生的，花费2）
10. 如果剩于1个2，但是没有3（这个不可能）
11. 如果通过前两步的操作后，没有剩余2，但是剩余了很多1，那么这些1都必须移动，所以就简单了