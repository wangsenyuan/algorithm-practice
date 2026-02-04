In this problem, you need to find maximal and minimal elements of an array. What could be simpler?

You can imagine that the jury has an array, and initially you know only the number $n$ — the array's length.

Array elements are numbered from $1$ to $n$. You are allowed to compare two elements of the array by using their indices $i$ and $j$. There are three possible responses to this query: `<` (if $a_i$ is less than $a_j$), `=` (if $a_i$ is equal to $a_j$), and finally `>` (if $a_i$ is greater than $a_j$).

It's known that it's always possible to find both maximal and minimal elements of the array by using no more than $\lceil \frac{3n}{2} \rceil - 2$ comparisons, where $\lceil x \rceil$ is the result of rounding $x$ up.

Write the program that will find positions of the minimum and the maximum in the jury's array of length $n$, by using no more than $f(n)$ comparisons.

## Interaction

Each test for this problem will contain one or more arrays. You have to find positions of minimal and maximal elements for each of these arrays. The first line of the input contains integer $T$ ($1 \le T \le 1000$) — the number of arrays in the test.

Thus, at the beginning, your program should read number $T$, and then it should solve the problem for $T$ jury arrays one by one.

Then input for each array goes. First, your program has to read the number $n$ ($1 \le n \le 50$) — the length of the array. It will be provided in the next line of the input.

Further, your program can perform comparisons or report that the answer is found.

To perform a comparison, you have to output a string of the following pattern `? i j` ($i$ and $j$ must be integer numbers from $1$ to $n$) — the indices of the elements to compare in the current query.

To report the indices of minimal and maximal elements of the hidden array, your program has to output a line in the form `! i j` ($i$ and $j$ must be integer numbers from $1$ to $n$), where $i$ is an index of the minimal element of array, and $j$ is an index of the maximal element of the array. If there are several possible answers to the problem, you can output any of them.

There are several possible responses for a comparison:

- `<` — if $a_i$ is less than $a_j$
- `=` — if $a_i$ is equal to $a_j$
- `>` — if $a_i$ is greater than $a_j$

For an array of length $n$ your program can make at most $f(n)$ comparisons. Note that the operation of reporting an answer (`! i j`) is not included into the value of $f(n)$.

After the answer is reported, your program has to solve the problem for the next array, or it should terminate if all $T$ arrays are processed.

## Example

### Input
```
2
2
>
3
=
=
```

### Output
```
? 1 2
! 2 1
? 3 1
? 2 1
! 2 3
```
In this problem, you need to find maximal and minimal elements of an array. What could be simpler?

You can imagine that the jury has an array, and initially you know the only number n — array's length.

Array's elements are numbered from 1 to n. You are allowed to compare two elements of the array by using their indices i and j. There are three possible responses to this query: '<' (if ai is less than aj), '=' (if ai is equal to aj) and finally '>' (if ai is greater than aj).

It's known that it's always possible to find both maximal and minimal elements of the array by using no more than  comparisons, where ⌈ x⌉ is the result of rounding x up.

Write the program that will find positions of the minimum and the maximum in the jury's array of length n, by using no more than f(n) comparisons.

Interaction
Each test for this problem will contain one or more arrays. You have to find positions of minimal and maximal elements for each of these arrays. The first line of the input contains integer T (1 ≤ T ≤ 1000) — number of arrays in the test.

Thus, at the beginning, you program should read number T, and then it should solve the problem for T jury's arrays one by one.

Then input for each array goes. Firstly, your program has to read the number n (1 ≤ n ≤ 50) — the length of the array. It will be provided in the next line of the input.

Further, your program can perform comparisons or report that the answer is found.

To perform a comparison, you have to output string of the following pattern «? i j» (i and j must be integer numbers from 1 to n) — the indices of the elements to compare in the current query.
To report the indices of minimal and maximal elements of the hidden array, your program have to output a line in the form «! i j» (i and j must be integer numbers from 1 to n), where i is an index of the minimal element of array, and j is an index of the maximal element of the array. If there are several possible answers to the problem, you can output any of them.
There are several possible responses for a comparison:

'<' — if ai is less than aj,
'=' — if ai is equal to aj,
'>' — if ai is greater than aj.
For an array of length n your program can make at most  comparisons. Note that the operation of reporting an answer («! i j» ) is not included into the value of f(n).

After the answer is reported, your program has to solve the problem for the next array or it should terminate if all T arrays are processed.

Example
InputCopy
2
2
 
>
 
3
 
=
 
=
 
OutputCopy
 
 
? 1 2
 
! 2 1
 
? 3 1
 
? 2 1
 
! 2 3

### ideas
1. n-1次操作可以找到最大值（或者最小值）；但是再进一步就有点麻烦了
2. (1, 2) (3, 4), (...) n/2 (再剩下一个)
3. 这样子可以将它们分成两组, 大数组和小数组
4. 比如 n = 50, 那么25次比较后，得到（25， 25）
5. 然后 n = 12 + 12 => (6, 6)
6. 然后 n = (3 + 3) => (3, 3)
7. 然后 n = (2, 2) => (1, 1)
8. 25 + 24 + 6 + 4 = 59 * 2 < 3 * 50 - 2 好像是可以的？
9. 但是相等的数怎么处理呢？
10. 保留一个进入下一轮
11. 试试看