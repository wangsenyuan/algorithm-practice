# Problem C: Alice and Bob's Coloring Game

## Problem Description

Alice and Bob are playing a game using an integer array $a$ of size $n$.

Initially, all elements of the array are colorless. First, Alice chooses 3 elements and colors them red. Then Bob chooses any element and colors it blue (if it was red — recolor it). Alice wins if the sum of the red elements is strictly greater than the value of the blue element.

Your task is to calculate the number of ways that Alice can choose 3 elements in order to win regardless of Bob's actions.

## Input

- The first line contains a single integer $t$ ($1 \leq t \leq 1000$) — the number of test cases.
- The first line of each test case contains a single integer $n$ ($3 \leq n \leq 5000$).
- The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_1 \leq a_2 \leq \cdots \leq a_n \leq 10^5$).

**Additional constraint:** The sum of $n$ over all test cases doesn't exceed 5000.

## Output

For each test case, print a single integer — the number of ways that Alice can choose 3 elements in order to win regardless of Bob's actions.

## Example

### Input
```
6
3
1 2 3
4
1 1 2 4
5
7 7 7 7 7
5
1 1 2 2 4
6
2 3 3 4 5 5
5
1 1 1 1 3
```

### Output
```
0
0
10
2
16
0
```

## Explanation

- **Test case 1-2:** No matter which three elements Alice chooses, Bob will be able to paint one element blue so that Alice does not win.

- **Test case 3:** Alice can choose any three elements. If Bob colors one of the red elements, the sum of red elements will be 14, and the sum of blue elements will be 7. If Bob chooses an uncolored element, the sum of red elements will be 21, and the sum of blue elements will be 7.

- **Test case 4:** Alice can choose either the 1st, 3rd and 4th element, or the 2nd, 3rd and 4th element.

- **Test case 5:** Alice has 16 different ways to choose 3 elements that guarantee a win.

- **Test case 6:** No matter which three elements Alice chooses, Bob can always choose an element that prevents Alice from winning.