# Problem D

Petya and Vasya are inventing a new game that requires a rectangular board and one chess piece. At the beginning of the game the piece stands in the upper-left corner of the board. Two players move the piece in turns. Each turn the chess piece can be moved either one square to the right or one square down or jump k squares diagonally down and to the right. The player who can't move the piece loses.

The guys haven't yet thought what to call the game or the best size of the board for it. Your task is to write a program that can determine the outcome of the game depending on the board size.

## Input

The first input line contains two integers t and k (1 ≤ t ≤ 20, 1 ≤ k ≤ 10⁹). Each of the following t lines contains two numbers n, m — the board's length and width (1 ≤ n, m ≤ 10⁹).

## Output

Output t lines that can determine the outcomes of the game on every board. Write «+» if the first player is a winner, and «-» otherwise.

## Examples

**Input:**
```
10 2
1 1
1 2
2 1
2 2
1 3
2 3
3 1
3 2
3 3
4 3
```

**Output:**
```
-
+
+
-
-
+
-
+
+
+
```


### solution

In the problem we have a game on an acyclic graph. Positions in the game (vertices of the graph) are pairs (n, m), where n and m are distances from the current position of the piece to the bottom and to the rights borders of the board. For every position (n, m) there are at most three moves: (n - 1, m), (n, m - 1), (n - k, m - k).  The graph is acyclic, because at each move the sum n+m strictly decreases.

If the numbers n, m and k do not exeed, say, 1000, the problem is solved by easy dynamics on the acyclic graph, standard for such games. Let d(n, m) = 1, if the position (n, m) is winning, and d(n, m) = 0, if the position (n, m) is losing. The value of d(n, m) is calculated using the following considerations. If there exists a move from the current position to a losing one, then the current position is winning, otherwise it is losing.

But conditions of the problem do not allow us to solve it by the standard dynamics. The solution is to implement the dynamics for small values of n and m and to find a pattern! For example, build the matrix of values d(n, m) for k = 2: 

01010101010101
10101010101010
01111111111111
10110101010101
01101010101010
10110111111111
01101101010101
10110110101010
01101101............


Consider the corner stripes in this picture. Start with a corner stripe with width 2 in the upper-left corner (the 'corner stripe' in this case is the set of squares with n <= 2 || m <= 2). Starting from the upper-left corner, we have a stripe with width 2 (k for k > 2) such that 0 and 1 alternate.  That is no wonder, because in that stripe we cannot jump by k.  Next we have a stripe of only 1, then there is a stripe with width k like the first one, but it is inverted (1 are changed by 0, and 0 by 1), then we have a stripe of 1, and then there is exactly the same stripe as the one in the very beginning! Since every next corner stripe depends only on the previous one with width k, we have the pattern: the following stripe of 1, the inverted stripe, the stripe of 1 again, etc. In fact we get d(n, m) = d(n - (2 k + 2), m - (2 k + 2)) for n, m > 2 k + 2. It yields formulas to calculate d(n, m) for every n and m. 

To understand given explanations more carefully, implement the dynamics and find the pattern yourself.

There is also a tricky case k = 1. It yields to a bit different pattern:)

Remark.  I want to emphasize that we have not just found the pattern and made a shamanistic hypothesis that it will be repeated. We have proved this. The stripes will alternate further, because every next stripe is uniquely determined by the previous stripe with width k.  