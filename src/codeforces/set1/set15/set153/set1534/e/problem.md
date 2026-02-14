# Problem E

**Note:** The XOR-sum of an array a₁, a₂, …, aₙ (1 ≤ aᵢ ≤ 10⁹) is defined as a₁ ⊕ a₂ ⊕ … ⊕ aₙ, where ⊕ denotes the bitwise XOR operation.

Little Dormi received an array of n integers a₁, a₂, …, aₙ for Christmas. However, while playing with it over the winter break, he accidentally dropped it into his XOR machine, and the array got lost.

The XOR machine is currently configured with a query size of k (which you cannot change), and allows you to perform the following type of query: by giving the machine k distinct indices x₁, x₂, …, xₖ, it will output aₓ₁ ⊕ aₓ₂ ⊕ … ⊕ aₓₖ.

As Little Dormi's older brother, you would like to help him recover the XOR-sum of his array a₁, a₂, …, aₙ by querying the XOR machine.

Little Dormi isn't very patient, so to be as fast as possible, you must query the XOR machine the minimum number of times to find the XOR-sum of his array. Formally, let d be the minimum number of queries needed to find the XOR-sum of any array of length n with a query size of k. Your program will be accepted if you find the correct XOR-sum in at most d queries.

Lastly, you also noticed that with certain configurations of the machine k and values of n, it may not be possible to recover the XOR-sum of Little Dormi's lost array. If that is the case, you should report it as well.

The array a₁, a₂, …, aₙ is fixed before you start querying the XOR machine and does not change with the queries.

## Input

The only line of input contains the integers n and k (1 ≤ n ≤ 500, 1 ≤ k ≤ n), the length of the lost array and the configured query size of the XOR machine.

Elements of the original array satisfy 1 ≤ aᵢ ≤ 10⁹.

It can be proven that if it is possible to recover the XOR sum under the given constraints, it can be done in at most 500 queries. That is, d ≤ 500.

After taking n and k, begin interaction.

## Output

If it is impossible to recover the XOR-sum of the array, output -1 immediately after taking n and k. Do not begin interaction.

Otherwise, when your program finds the XOR-sum of the lost array a₁, a₂, …, aₙ, report the answer in the following format: "! x", where x is the XOR sum of the array a₁, a₂, …, aₙ, and terminate your program normally immediately after flushing the output stream.

Note that answering does not count as a query.

## Interaction

Each query is made in the format "? b", where b is an array of exactly k distinct integers from 1 to n denoting the indices of the elements in the lost array that you want to query the XOR sum of.

You will then receive an integer x, the XOR sum of the queried elements. It can be proven that 0 ≤ x ≤ 2·10⁹ will always be true.

After printing a query do not forget to output end of line and flush the output. Otherwise, you will get Idleness limit exceeded. To do this, use:

- `fflush(stdout)` or `cout.flush()` in C++;
- `System.out.flush()` in Java;
- `flush(output)` in Pascal;
- `stdout.flush()` in Python;
- see documentation for other languages.

If at any point you make an invalid query or try to make more than 500 queries (which is the hard limit), the interaction will terminate immediately and give you a Wrong Answer verdict. Note that if you exceed d queries, the interaction will continue normally unless you also exceed the 500 query hard limit, though you will still receive a Wrong Answer verdict either way.

## Hacks

To hack a solution, use the following format.

The first line contains the integers n and k (1 ≤ n ≤ 500, 1 ≤ k ≤ n).

The second line contains the array a₁, a₂, …, aₙ (1 ≤ aᵢ ≤ 10⁹).

## Examples

### Example 1

**Input:**
```
5 3

4

0

1
```

**Output:**
```
? 1 2 3

? 2 3 5

? 4 1 5

! 7
```

### Example 2

**Input:**
```
3 2
```

**Output:**
```
-1
```

## Note

In the first example interaction, the array a₁, a₂, …, aₙ is [2, 1, 7, 5, 6] and its XOR-sum is 7.

The first query made asks for indices 1, 2, 3, so the response is a₁ ⊕ a₂ ⊕ a₃ = 2 ⊕ 1 ⊕ 7 = 4.

The second query made asks for indices 2, 3, 5, so the response is a₂ ⊕ a₃ ⊕ a₅ = 1 ⊕ 7 ⊕ 6 = 0.

The third query made asks for indices 4, 1, 5, so the response is a₄ ⊕ a₁ ⊕ a₅ = 5 ⊕ 2 ⊕ 6 = 1. Note that the indices may be output in any order.

Additionally, even though three queries were made in the example interaction, it is just meant to demonstrate the interaction format and does not necessarily represent an optimal strategy.

In the second example interaction, there is no way to recover the XOR-sum of Little Dormi's array no matter what is queried, so the program immediately outputs -1 and exits.

### ideas
1. 如果 n % k = 0, very good
2. else let w = n % k ， 如果 w % 2 = 0, 
3. 在剩余w的时候，分成两次问, 公共的部分k - w/2个
4. 有区分的是 w/2个，这样子，两次结果 xor 下，就是剩余w个xor值
5. 如果w可以3等分，也是可以的，比如w = 3, 前两次xor的结果，和最后一次xor，就是最后一部分
6. w5等分（好像都行)
7. 只有w = 1不行
8. 但是w = 1的时候，如果+k后，还有额外的，也是可以的
9. 