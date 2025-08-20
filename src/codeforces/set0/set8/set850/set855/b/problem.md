# Professor Dumbledore and the Horcrux Curse

## Problem Description

Professor Dumbledore is helping Harry destroy the Horcruxes. He went to Gaunt Shack as he suspected a Horcrux to be present there. He saw Marvolo Gaunt's Ring and identified it as a Horcrux. Although he destroyed it, he is still affected by its curse. Professor Snape is helping Dumbledore remove the curse. For this, he wants to give Dumbledore exactly **x** drops of the potion he made.

The value of **x** is calculated as the maximum of **p·ai + q·aj + r·ak** for given **p**, **q**, **r** and array **a₁, a₂, ..., an** such that **1 ≤ i ≤ j ≤ k ≤ n**.

Help Snape find the value of **x**. Note that the value of **x** may be negative.

## Input

- First line contains 4 integers: **n**, **p**, **q**, **r** 
  - **-10⁹ ≤ p, q, r ≤ 10⁹**
  - **1 ≤ n ≤ 10⁵**
- Second line contains **n** space-separated integers: **a₁, a₂, ..., an**
  - **-10⁹ ≤ ai ≤ 10⁹**

## Output

Output a single integer representing the maximum value of **p·ai + q·aj + r·ak** that can be obtained, provided **1 ≤ i ≤ j ≤ k ≤ n**.

## Examples

### Example 1
**Input:**
```
5 1 2 3
1 2 3 4 5
```

**Output:**
```
30
```

**Explanation:** We can take **i = j = k = 5**, thus making the answer as **1·5 + 2·5 + 3·5 = 30**.

### Example 2
**Input:**
```
5 1 2 -3
-1 -2 -3 -4 -5
```

**Output:**
```
12
```

**Explanation:** Selecting **i = j = 1** and **k = 5** gives the answer **12**.