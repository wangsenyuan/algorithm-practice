# Problem A: Doubly Linked Lists

## Problem Statement

A doubly linked list is one of the fundamental data structures. A doubly linked list is a sequence of elements, each containing information about the previous and the next elements of the list. In this problem, all lists have linear structure. I.e., each element except the first has exactly one previous element, and each element except the last has exactly one next element. The list is not closed in a cycle.

In this problem, you are given `n` memory cells forming one or more doubly linked lists. Each cell contains information about an element from some list. Memory cells are numbered from 1 to `n`.

For each cell `i`, you are given two values:
- `li` — cell containing the previous element for the element in cell `i`
- `ri` — cell containing the next element for the element in cell `i`

If cell `i` contains information about an element which has no previous element, then `li = 0`. Similarly, if cell `i` contains information about an element which has no next element, then `ri = 0`.

### Example
Three lists are shown in the picture. For example, for the picture above, the values of `l` and `r` are the following:
- `l1 = 4, r1 = 7`
- `l2 = 5, r2 = 0`
- `l3 = 0, r3 = 0`
- `l4 = 6, r4 = 1`
- `l5 = 0, r5 = 2`
- `l6 = 0, r6 = 4`
- `l7 = 1, r7 = 0`

## Task
Your task is to unite all given lists into a single list, joining them to each other in any order. In particular, if the input data already contains a single list, then there is no need to perform any actions. Print the resulting list in the form of values `li, ri`.

**Note:** Any action other than joining the beginning of one list to the end of another cannot be performed.

## Input
The first line contains a single integer `n` (`1 ≤ n ≤ 100`) — the number of memory cells where the doubly linked lists are located.

Each of the following `n` lines contains two integers `li, ri` (`0 ≤ li, ri ≤ n`) — the cells of the previous and the next element of the list for cell `i`. 
- Value `li = 0` if the element in cell `i` has no previous element in its list
- Value `ri = 0` if the element in cell `i` has no next element in its list

It is guaranteed that the input contains the correct description of a single or more doubly linked lists. All lists have linear structure: each element of a list except the first has exactly one previous element; each element of a list except the last has exactly one next element. Each memory cell contains information about one element from some list, and each element of each list is written in one of the `n` given cells.

## Output
Print `n` lines, where the `i`-th line must contain two integers `li` and `ri` — the cells of the previous and the next element of the list for cell `i` after all lists from the input are united into a single list. If there are many solutions, print any of them.

## Example

**Input:**
```
7
4 7
5 0
0 0
6 1
0 2
0 4
1 0
```

**Output:**
```
4 7
5 6
0 5
6 1
3 2
2 4
1 0
```