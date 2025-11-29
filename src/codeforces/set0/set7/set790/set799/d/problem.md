In one of the games Arkady is fond of the game process happens on a rectangular field. In the game process Arkady can buy extensions for his field, each extension enlarges one of the field sizes in a particular number of times. Formally, there are n extensions, the i-th of them multiplies the width or the length (by Arkady's choice) by ai. Each extension can't be used more than once, the extensions can be used in any order.

Now Arkady's field has size h × w. He wants to enlarge it so that it is possible to place a rectangle of size a × b on it (along the width or along the length, with sides parallel to the field sides). Find the minimum number of extensions needed to reach Arkady's goal.

## Input

The first line contains five integers a, b, h, w and n (1 ≤ a, b, h, w, n ≤ 100 000) — the sizes of the rectangle needed to be placed, the initial sizes of the field and the number of available extensions.

The second line contains n integers a1, a2, ..., an (2 ≤ ai ≤ 100 000), where ai equals the integer a side multiplies by when the i-th extension is applied.

## Output

Print the minimum number of extensions needed to reach Arkady's goal. If it is not possible to place the rectangle on the field with all extensions, print -1. If the rectangle can be placed on the initial field, print 0.

## Examples

### Example 1

**Input:**
```
3 3 2 4 4
2 5 4 10
```

**Output:**
```
1
```

### Example 2

**Input:**
```
3 3 3 3 5
2 3 5 4 2
```

**Output:**
```
0
```

### Example 3

**Input:**
```
5 5 1 2 3
2 2 3
```

**Output:**
```
-1
```

### Example 4

**Input:**
```
3 4 1 1 3
2 3 2
```

**Output:**
```
3
```

## Note

In the first example it is enough to use any of the extensions available. For example, we can enlarge h in 5 times using the second extension. Then h becomes equal 10 and it is now possible to place the rectangle on the field.


### ideas
1. 一开始有个 h * w 的 field， 需要放置 a * b 的rect进去
2. 所以要增加这个field， 
3. 把a排序，使用前i个，使的高度为H时，所能得到的最大的W是多少
4. a的长度如果超过60， 那么肯定就可以了 2 ** 30 > a and b了
5. 所以，只需要考虑len(a) < 40
6. dp[i][s] = h (表示前i个， 其中s个用来作为高度) 