Vasiliy has a car and he wants to get from home to the post office. The distance which he needs to pass equals to d kilometers.

Vasiliy's car is not new — it breaks after driven every k kilometers and Vasiliy needs t seconds to repair it. After repairing his car Vasiliy can drive again (but after k kilometers it will break again, and so on). In the beginning of the trip the car is just from repair station.

To drive one kilometer on car Vasiliy spends a seconds, to walk one kilometer on foot he needs b seconds (a < b).

Your task is to find minimal time after which Vasiliy will be able to reach the post office. Consider that in every moment of time Vasiliy can left his car and start to go on foot.

## Input

The first line contains 5 positive integers d, k, a, b, t (1 ≤ d ≤ 10^12; 1 ≤ k, a, b, t ≤ 10^6; a < b), where:

- d — the distance from home to the post office;
- k — the distance, which car is able to drive before breaking;
- a — the time, which Vasiliy spends to drive 1 kilometer on his car;
- b — the time, which Vasiliy spends to walk 1 kilometer on foot;
- t — the time, which Vasiliy spends to repair his car.

## Output

Print the minimal time after which Vasiliy will be able to reach the post office.

## Examples

### Example 1

**Input:**
```
5 2 1 4 10
```

**Output:**
```
14
```

### Example 2

**Input:**
```
5 2 1 4 5
```

**Output:**
```
13
```

## Note

In the first example Vasiliy needs to drive the first 2 kilometers on the car (in 2 seconds) and then to walk on foot 3 kilometers (in 12 seconds). So the answer equals to 14 seconds.

In the second example Vasiliy needs to drive the first 2 kilometers on the car (in 2 seconds), then repair his car (in 5 seconds) and drive 2 kilometers more on the car (in 2 seconds). After that he needs to walk on foot 1 kilometer (in 4 seconds). So the answer equals to 13 seconds.


### ideas
1. 分几种情况，完全使用车，完全步行，或者兼有
2. 假设车开了x公里，步行y = d - x
3. 那么 x % k == 0（车在正常行驶时，没有理由步行)
4. let w = x / k
5. 那么总花费时间 = x * a + (w - 1) * t + y * b
6. x * a + (d - x) * b + (x / k - 1) * t
7. (a - b + t / k) * x + d * b - t, a < b
8. 如果 a + t / k - b == 0, 貌似是固定的额b * d - t
9. 如果 a + t / k - b > 0, 那么 x越大，花费越长
10. 如果 a + t / k - b < 0, 那么x越小，花费越长
11. 所以，要么只有第一段开车（对应系数为正的情况）
12. 要么一直到最后一段都开车，剩下的部分走路（省掉一次修理时间）