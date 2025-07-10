# Problem Description

Something happened in Uzhlyandia again... There are riots on the streets... Famous Uzhlyandian superheroes Shean the Sheep and Stas the Giraffe were called in order to save the situation. Upon the arriving, they found that citizens are worried about maximum values of the Main Uzhlyandian Function f, which is defined as follows:

## Function Definition

The Main Uzhlyandian Function f is defined as:

$f(l, r) = |a_l - a_r| + |a_{l+1} - a_{r-1}| + |a_{l+2} - a_{r-2}| + ... + |a_{r-1} - a_{l+1}| + |a_r - a_l|$

In the above formula, $1 \leq l < r \leq n$ must hold, where $n$ is the size of the Main Uzhlyandian Array $a$, and $|x|$ means absolute value of $x$.

But the heroes skipped their math lessons in school, so they asked you for help. Help them calculate the maximum value of $f$ among all possible values of $l$ and $r$ for the given array $a$.

## Input

The first line contains single integer $n$ ($2 \leq n \leq 10^5$) — the size of the array $a$.

The second line contains $n$ integers $a_1, a_2, ..., a_n$ ($-10^9 \leq a_i \leq 10^9$) — the array elements.

## Output

Print the only integer — the maximum value of $f$.


### ideas
1. 先处理奇数位的，这样子后续 d1, d2, d3....的符号可以被复用
2. +d0, -d1, +d2, -d3, ... 那么就是最大和

