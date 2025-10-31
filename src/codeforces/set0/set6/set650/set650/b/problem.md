## Vasya's Telephone Photos

### Description
Vasya's telephone contains `n` photos. Photo number `1` is currently open. You can move left or right to the adjacent photo by swiping. If you swipe left from the first photo, you reach photo `n`. Similarly, swiping right from the last photo brings you to photo `1`. It takes `a` seconds to move between adjacent photos.

Each photo has an intended orientation — horizontal or vertical. The phone is in vertical orientation and cannot be rotated. It takes `b` seconds to rotate a photo if its orientation does not match the intended one.

Vasya has `T` seconds to watch photos and wants to watch as many as possible. When he opens a photo for the first time, he spends `1` second to view it. If the photo is in the wrong orientation, he spends an additional `b` seconds to rotate it before viewing. If a photo has already been opened before, he skips it without spending time. It is not allowed to skip unseen photos.

Help Vasya find the maximum number of photos he can watch within `T` seconds.

### Input
- **First line**: four integers `n`, `a`, `b`, `T` — the number of photos, time to move to an adjacent photo, time to rotate a photo, and the total available time.
  - Constraints: `1 ≤ n ≤ 5·10^5`, `1 ≤ a, b ≤ 1000`, `1 ≤ T ≤ 10^9`.
- **Second line**: a string of length `n` over characters `'w'` and `'h'`.
  - If the `i`-th character is `'w'`, then photo `i` should be seen in horizontal orientation.
  - If the `i`-th character is `'h'`, then photo `i` should be seen in vertical orientation.

### Output
Output a single integer — the maximum number of photos Vasya can watch during `T` seconds.

### Examples
Example 1

Input
```text
4 2 3 10
wwhw
```

Output
```text
2
```

Example 2

Input
```text
5 2 4 13
hhwhh
```

Output
```text
4
```

Example 3

Input
```text
5 2 4 1000
hhwhh
```

Output
```text
5
```

Example 4

Input
```text
3 1 100 10
whw
```

Output
```text
0
```

### Notes
- In the first sample, rotate the first photo (`3` seconds), watch it (`1` second), move left (`2` seconds), rotate the fourth photo (`3` seconds), watch it (`1` second). Total time is exactly `10` seconds.
- In the last sample, the time is not enough even to watch the first photo, and unseen photos cannot be skipped.


### ideas
1. 大体是一个区间，等待的时间可以不用考虑， 假设一共x张，那么就等待打开花费了x秒
2. 而且肯定是往一个方向，假设是l, 然后一直往后
3. 这样子，一共花费的时间 = (r - l) * a + (n - l) * a + 需要旋转的照片的数量 * b
4. 从n开始，这样子就一定是个区间