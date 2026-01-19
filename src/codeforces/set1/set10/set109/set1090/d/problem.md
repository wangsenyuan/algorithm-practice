# Problem D

Vasya had an array of 𝑛 integers, each element of the array was from 1 to 𝑛. He chose 𝑚 pairs of different positions and wrote them down to a sheet of paper. Then Vasya compared the elements at these positions, and wrote down the results of the comparisons to another sheet of paper. For each pair he wrote either "greater", "less", or "equal".

After several years, he has found the first sheet of paper, but he couldn't find the second one. Also he doesn't remember the array he had. In particular, he doesn't remember if the array had equal elements. He has told this sad story to his informatics teacher Dr Helen.

She told him that it could be the case that even if Vasya finds his second sheet, he would still not be able to find out whether the array had two equal elements.

Now Vasya wants to find two arrays of integers, each of length 𝑛. All elements of the first array must be distinct, and there must be two equal elements in the second array. For each pair of positions Vasya wrote at the first sheet of paper, the result of the comparison must be the same for the corresponding elements of the first array, and the corresponding elements of the second array.

Help Vasya find two such arrays of length 𝑛, or find out that there are no such arrays for his sets of pairs.

## Input

The first line of input contains two integers 𝑛, 𝑚 — the number of elements in the array and number of comparisons made by Vasya (1≤𝑛≤10⁵, 0≤𝑚≤10⁵).

Each of the following 𝑚 lines contains two integers 𝑎ᵢ, 𝑏ᵢ — the positions of the 𝑖-th comparison (1≤𝑎ᵢ,𝑏ᵢ≤𝑛; 𝑎ᵢ≠𝑏ᵢ). It's guaranteed that any unordered pair is given in the input at most once.

## Output

The first line of output must contain "YES" if there exist two arrays, such that the results of comparisons would be the same, and all numbers in the first one are distinct, and the second one contains two equal numbers. Otherwise it must contain "NO".

If the arrays exist, the second line must contain the array of distinct integers, the third line must contain the array, that contains at least one pair of equal elements. Elements of the arrays must be integers from 1 to 𝑛.

## Examples

**Example 1**

Input:
```
1 0
```

Output:
```
NO
```

**Example 2**

Input:
```
3 1
1 2
```

Output:
```
YES
1 3 2 
1 3 1 
```

**Example 3**

Input:
```
4 3
1 2
1 3
2 4
```

Output:
```
YES
1 3 4 2 
1 3 4 1 
```
