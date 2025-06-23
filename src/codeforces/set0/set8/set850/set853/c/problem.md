# Problem C: Beautiful Rectangles

## Problem Description

Ilya is sitting in a waiting area of Metropolis airport and is bored of looking at time table that shows again and again that his plane is delayed. So he took out a sheet of paper and decided to solve some problems.

First Ilya has drawn a grid of size $n \times n$ and marked $n$ squares on it, such that no two marked squares share the same row or the same column. He calls a rectangle on a grid with sides parallel to grid sides **beautiful** if exactly two of its corner squares are marked. There are exactly $\frac{n \cdot (n - 1)}{2}$ beautiful rectangles.

Ilya has chosen $q$ query rectangles on a grid with sides parallel to grid sides (not necessarily beautiful ones), and for each of those rectangles he wants to find its **beauty degree**. Beauty degree of a rectangle is the number of beautiful rectangles that share at least one square with the given one.

Now Ilya thinks that he might not have enough time to solve the problem till the departure of his flight. You are given the description of marked cells and the query rectangles, help Ilya find the beauty degree of each of the query rectangles.

## Input

The first line of input contains two integers $n$ and $q$ ($2 \leq n \leq 200\,000$, $1 \leq q \leq 200\,000$) — the size of the grid and the number of query rectangles.

The second line contains $n$ integers $p_1, p_2, \ldots, p_n$, separated by spaces ($1 \leq p_i \leq n$, all $p_i$ are different), they specify grid squares marked by Ilya: in column $i$ he has marked a square at row $p_i$, rows are numbered from $1$ to $n$, bottom to top, columns are numbered from $1$ to $n$, left to right.

The following $q$ lines describe query rectangles. Each rectangle is described by four integers: $l, d, r, u$ ($1 \leq l \leq r \leq n$, $1 \leq d \leq u \leq n$), here $l$ and $r$ are the leftmost and the rightmost columns of the rectangle, $d$ and $u$ the bottommost and the topmost rows of the rectangle.

## Output

For each query rectangle output its beauty degree on a separate line.