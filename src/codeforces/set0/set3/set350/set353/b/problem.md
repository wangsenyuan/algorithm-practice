# Codeforces 353B - Two Heaps

## Problem Description

Valera has 2·n cubes, each cube contains an integer from 10 to 99. He arbitrarily chooses n cubes and puts them in the first heap. The remaining cubes form the second heap.

Valera decided to play with cubes. During the game he takes a cube from the first heap and writes down the number it has. Then he takes a cube from the second heap and write out its two digits near two digits he had written (to the right of them). In the end he obtained a single four-digit integer — the first two digits of it is written on the cube from the first heap, and the second two digits of it is written on the second cube from the second heap.

Valera knows arithmetic very well. So, he can easily count the number of distinct four-digit numbers he can get in the game. The other question is: how to split cubes into two heaps so that this number (the number of distinct four-digit integers Valera can get) will be as large as possible?

## Input

The first line contains integer **n** (1 ≤ n ≤ 100). 

The second line contains **2·n** space-separated integers **aᵢ** (10 ≤ aᵢ ≤ 99), denoting the numbers on the cubes.

## Output

In the first line print a single number — the maximum possible number of distinct four-digit numbers Valera can obtain. 

In the second line print **2·n** numbers **bᵢ** (1 ≤ bᵢ ≤ 2). The numbers mean: the i-th cube belongs to the bᵢ-th heap in your division.

**Note:** If there are multiple optimal ways to split the cubes into the heaps, print any of them.