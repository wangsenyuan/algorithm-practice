# Levko and Array Beauty

Levko has an array consisting of integers: $a_1, a_2, \ldots, a_n$. But he doesn't like this array at all.

Levko thinks that the beauty of the array $a$ directly depends on the value $c(a)$, which can be calculated by the formula:

$$
c(a) = \max_{1 \leq i < j \leq n} |a_i - a_j|
$$

The less the value of $c(a)$ is, the more beautiful the array is.

It's time to change the world, and Levko is going to change his array for the better. To be exact, Levko wants to change the values of at most $k$ array elements (it is allowed to replace the values by any integers). Of course, the changes should make the array as beautiful as possible.

Help Levko and calculate what minimum number $c(a)$ he can reach.

---

## Input

- The first line contains two integers $n$ and $k$ ($1 \leq k \leq n \leq 2000$).
- The second line contains space-separated integers $a_1, a_2, \ldots, a_n$ ($-10^9 \leq a_i \leq 10^9$).

## Output

- A single number — the minimum value of $c(a)$ Levko can get.


## ideas
1. 二分+dp?
2. 在给定x的时候，如何进行修改呢？
3. 这里的难点是，对a[i]的修改，貌似会影响到a[i+1]的选择；a[i]的修改，也受到a[i-1]修改后的结果影响
4. 假设选定x的时候，肯定有一个位置是不会变的（变化不会超过n-1)
5. 那是不是可以固定一个位置，贪心的进行选择呢？就是能不变的时候，不改变，如果要改变，就根据后一个数，看要改变成什么？
6.  