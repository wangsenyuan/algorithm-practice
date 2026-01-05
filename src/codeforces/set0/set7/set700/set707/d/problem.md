# Problem Description

Recently in school Alina has learned what are the persistent data structures: they are data structures that always preserves the previous version of itself and access to it when it is modified.

After reaching home Alina decided to invent her own persistent data structure. Inventing didn't take long: there is a bookcase right behind her bed. Alina thinks that the bookcase is a good choice for a persistent data structure. Initially the bookcase is empty, thus there is no book at any position at any shelf.

The bookcase consists of n shelves, and each shelf has exactly m positions for books at it. Alina enumerates shelves by integers from 1 to n and positions at shelves — from 1 to m. Initially the bookcase is empty, thus there is no book at any position at any shelf in it.

Alina wrote down q operations, which will be consecutively applied to the bookcase. Each of the operations has one of four types:

1. `1 i j` — Place a book at position j at shelf i if there is no book at it.
2. `2 i j` — Remove the book from position j at shelf i if there is a book at it.
3. `3 i` — Invert book placing at shelf i. This means that from every position at shelf i which has a book at it, the book should be removed, and at every position at shelf i which has not book at it, a book should be placed.
4. `4 k` — Return the books in the bookcase in a state they were after applying k-th operation. In particular, k = 0 means that the bookcase should be in initial state, thus every book in the bookcase should be removed from its position.

After applying each of operation Alina is interested in the number of books in the bookcase. Alina got 'A' in the school and had no problem finding this values. Will you do so?

## Input

The first line of the input contains three integers n, m and q (1 ≤ n, m ≤ 10³, 1 ≤ q ≤ 10⁵) — the bookcase dimensions and the number of operations respectively.

The next q lines describes operations in chronological order — i-th of them describes i-th operation in one of the four formats described in the statement.

It is guaranteed that shelf indices and position indices are correct, and in each of fourth-type operation the number k corresponds to some operation before it or equals to 0.

## Output

For each operation, print the number of books in the bookcase after applying it in a separate line. The answers should be printed in chronological order.

## Examples

### Example 1

**Input:**

```text
2 3 3
1 1 1
3 2
4 0
```

**Output:**

```text
1
4
0
```

### Example 2

**Input:**

```text
4 2 6
3 2
2 2 2
3 3
3 2
2 2 2
3 2
```

**Output:**

```text
2
1
3
3
2
4
```

### Example 3

**Input:**

```text
2 2 2
3 2
2 2 1
```

**Output:**

```text
2
1
```


### ideas
1. 对于query 4， 只需要记录下每个状态后的数量，就可以了
2. 关键是操作3， invert的时候，需要记录快速的进行当时的操作；使用lazy update貌似是可以的