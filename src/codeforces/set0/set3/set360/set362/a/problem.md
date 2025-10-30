### Semiknight Meeting

Petya loves chess and invented a piece called a semiknight. A semiknight can move in any of these four directions on a standard chessboard:

- 2 squares forward and 2 to the right
- 2 squares forward and 2 to the left
- 2 squares backward and 2 to the right
- 2 squares backward and 2 to the left

Naturally, the semiknight cannot move beyond the limits of the chessboard.

Petya places two semiknights on a standard 8×8 chessboard and moves both simultaneously. The squares are large enough that after some move, the semiknights may meet (i.e., land on the same square). After meeting, they may continue moving and might meet again. Some squares are considered bad; meetings on bad squares do not count, though semiknights may pass through them. Petya wants to know whether there exists a sequence of moves so the semiknights meet on a good square.

Help Petya determine this for multiple boards.

### Input

- The first line contains an integer \(t\) (\(1 \le t \le 50\)) — the number of boards.
- Each board is given by an 8×8 matrix of characters consisting of:
  - `.` — an empty good square
  - `#` — a bad square
  - `K` — a semiknight's position
- Each matrix contains exactly two `K` characters.
- The semiknight squares themselves are considered good for meeting.
- Tests are separated by an empty line.

### Output

For each test case, print `YES` if the semiknights can meet on some good square, and `NO` otherwise.

### Example

Input

```text
2
........
........
......#.
K..##..#
.......#
...##..#
......#.
K.......

........
........
..#.....
..#..#..
..####..
...##...
........
....K#K#
```

Output

```text
YES
NO
```

### Note

Consider the first board from the sample. Number rows and columns from 1 to 8 (top to bottom, left to right). The semiknights can meet at square \((2, 7)\). The semiknight from \((4, 1)\) goes to \((2, 3)\), and the semiknight from \((8, 1)\) goes to \((6, 3)\). Then both go to \((4, 5)\), but this square is bad, so they continue together to \((2, 7)\).

On the second board, the semiknights never meet on a good square.

A boy Petya loves chess very much. He even came up with a chess piece of his own, a semiknight. The semiknight can move in any of these four directions: 2 squares forward and 2 squares to the right, 2 squares forward and 2 squares to the left, 2 squares backward and 2 to the right and 2 squares backward and 2 to the left. Naturally, the semiknight cannot move beyond the limits of the chessboard.

Petya put two semiknights on a standard chessboard. Petya simultaneously moves with both semiknights. The squares are rather large, so after some move the semiknights can meet, that is, they can end up in the same square. After the meeting the semiknights can move on, so it is possible that they meet again. Petya wonders if there is such sequence of moves when the semiknights meet. Petya considers some squares bad. That is, they do not suit for the meeting. The semiknights can move through these squares but their meetings in these squares don't count.

Petya prepared multiple chess boards. Help Petya find out whether the semiknights can meet on some good square for each board.

Please see the test case analysis.

Input
The first line contains number t (1 ≤ t ≤ 50) — the number of boards. Each board is described by a matrix of characters, consisting of 8 rows and 8 columns. The matrix consists of characters ".", "#", "K", representing an empty good square, a bad square and the semiknight's position, correspondingly. It is guaranteed that matrix contains exactly 2 semiknights. The semiknight's squares are considered good for the meeting. The tests are separated by empty line.

Output
For each test, print on a single line the answer to the problem: "YES", if the semiknights can meet and "NO" otherwise.

Examples
InputCopy
2
........
........
......#.
K..##..#
.......#
...##..#
......#.
K.......

........
........
..#.....
..#..#..
..####..
...##...
........
....K#K#
OutputCopy
YES
NO
Note
Consider the first board from the sample. We will assume the rows and columns of the matrix to be numbered 1 through 8 from top to bottom and from left to right, correspondingly. The knights can meet, for example, in square (2, 7). The semiknight from square (4, 1) goes to square (2, 3) and the semiknight goes from square (8, 1) to square (6, 3). Then both semiknights go to (4, 5) but this square is bad, so they move together to square (2, 7).

On the second board the semiknights will never meet.

