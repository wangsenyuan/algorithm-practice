# Problem E

Polycarp has recently got himself a new job. He now earns so much that his old wallet can't even store all the money he has.

Berland bills somehow come in lots of different sizes. However, all of them are shaped as rectangles (possibly squares). All wallets are also produced in form of rectangles (possibly squares).

A bill $x \times y$ fits into some wallet $h \times w$ if either $x \leq h$ and $y \leq w$ or $y \leq h$ and $x \leq w$. Bills can overlap with each other in a wallet and an infinite amount of bills can fit into a wallet. That implies that all the bills Polycarp currently have fit into a wallet if every single one of them fits into it independently of the others.

Now you are asked to perform the queries of two types:

- `+ x y` — Polycarp earns a bill of size $x \times y$;
- `? h w` — Polycarp wants to check if all the bills he has earned to this moment fit into a wallet of size $h \times w$.

It is guaranteed that there is at least one query of type 1 before the first query of type 2 and that there is at least one query of type 2 in the input data.

For each query of type 2 print "YES" if all the bills he has earned to this moment fit into a wallet of given size. Print "NO" otherwise.

## Input

The first line contains a single integer $n$ $(2 \leq n \leq 5 \cdot 10^5)$ — the number of queries.

Each of the next $n$ lines contains a query of one of these two types:

- `+ x y` $(1 \leq x, y \leq 10^9)$ — Polycarp earns a bill of size $x \times y$;
- `? h w` $(1 \leq h, w \leq 10^9)$ — Polycarp wants to check if all the bills he has earned to this moment fit into a wallet of size $h \times w$.

It is guaranteed that there is at least one query of type 1 before the first query of type 2 and that there is at least one query of type 2 in the input data.

## Output

For each query of type 2 print "YES" if all the bills he has earned to this moment fit into a wallet of given size. Print "NO" otherwise.

## Example

**Input:**
```
9
+ 3 2
+ 2 3
? 1 20
? 3 3
? 2 3
+ 1 5
? 10 10
? 1 5
+ 1 1
```

**Output:**
```
NO
YES
YES
YES
NO
```

## Note

The queries of type 2 of the example:

- Neither bill fits;
- Both bills fit (just checking that you got that bills can overlap);
- Both bills fit (both bills are actually the same);
- All bills fit (too much of free space in a wallet is not a problem);
- Only bill $1 \times 5$ fit (all the others don't, thus it's "NO").


### ideas
1. let p = {min(x, y), max(x, y)}
2. 组成一个arr， 在这个arr中，按照p.first升序排列，且按照p.second降序排列