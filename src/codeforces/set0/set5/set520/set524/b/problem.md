Many years passed, and n friends met again at a party. Technology had advanced greatly since their last meeting; cameras with self-timers had become available, eliminating the need for one friend to stand by with the camera, thereby missing the picture.

The process of photographing can be simplified as follows. In the photograph, each friend occupies a rectangle of pixels: when standing, the i-th person occupies a rectangle of width w_i pixels and height h_i pixels. However, when photographing, each person can lie down, in which case they occupy a rectangle of width h_i pixels and height w_i pixels.

The shared photograph will have dimensions W × H, where W is the combined width of all the people-shaped rectangles, and H is the maximum height. The friends want to determine the minimum area of the shared photograph. Help them with this.

### Input

The first line contains an integer n (1 ≤ n ≤ 1000) — the number of friends.

The next n lines contain two integers w_i, h_i (1 ≤ w_i, h_i ≤ 1000), denoting the dimensions of the rectangle corresponding to the i-th friend.

### Output

Print a single integer equal to the minimum possible area of the photograph that can accommodate all the friends.

### Examples

#### Input
```
3 
10 1 
20 2 
30 3
```

#### Output
```
180
```

#### Input
```
3 
3 1 
2 2 
4 3
```

#### Output
```
21
```

#### Input
```
1 
5 10
```

#### Output
```
50
```


### ideas
1. 在给定宽度的情况下，计算高度？
2. dp[w] = h, 表示在宽度为w时的最小高度
3. dp[w + wi] = max(dp[w], hi) 如果使用新的宽度，那么产生新的高度
4. 如果 wi <= w, dp[w] += hi ? 
5. 还有一种情况时，i可以*嵌入*一个空格区间内。
6. 但是这里是人，所以不存在叠起来的情况吧？
7. 