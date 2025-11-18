# Problem Description

Polycarp is flying in the airplane. Finally, it is his favorite time — the lunchtime. The BerAvia company stewardess is giving food consecutively to all the passengers from the 1-th one to the last one. Polycarp is sitting on seat m, that means, he will be the m-th person to get food.

The flight menu has k dishes in total and when Polycarp boarded the flight, he had time to count the number of portions of each dish on board. Thus, he knows values a1, a2, ..., ak, where ai is the number of portions of the i-th dish.

The stewardess has already given food to m - 1 passengers, gave Polycarp a polite smile and asked him what he would prefer. That's when Polycarp realized that they might have run out of some dishes by that moment. For some of the m - 1 passengers ahead of him, he noticed what dishes they were given. Besides, he's heard some strange mumbling from some of the m - 1 passengers ahead of him, similar to phrase 'I'm disappointed'. That happened when a passenger asked for some dish but the stewardess gave him a polite smile and said that they had run out of that dish. In that case the passenger needed to choose some other dish that was available. If Polycarp heard no more sounds from a passenger, that meant that the passenger chose his dish at the first try.

Help Polycarp to find out for each dish: whether they could have run out of the dish by the moment Polyarp was served or that dish was definitely available.

## Input

Each test in this problem consists of one or more input sets. First goes a string that contains a single integer t (1 ≤ t ≤ 100 000) — the number of input data sets in the test. Then the sets follow, each set is preceded by an empty line.

The first line of each set of the input contains integers m, k (2 ≤ m ≤ 100 000, 1 ≤ k ≤ 100 000) — the number of Polycarp's seat and the number of dishes, respectively.

The second line contains a sequence of k integers a1, a2, ..., ak (1 ≤ ai ≤ 100 000), where ai is the initial number of portions of the i-th dish.

Then m - 1 lines follow, each line contains the description of Polycarp's observations about giving food to a passenger sitting in front of him: the j-th line contains a pair of integers tj, rj (0 ≤ tj ≤ k, 0 ≤ rj ≤ 1), where tj is the number of the dish that was given to the j-th passenger (or 0, if Polycarp didn't notice what dish was given to the passenger), and rj — a 1 or a 0, depending on whether the j-th passenger was or wasn't disappointed, respectively.

We know that sum ai equals at least m, that is, Polycarp will definitely get some dish, even if it is the last thing he wanted. It is guaranteed that the data is consistent.

Sum m for all input sets doesn't exceed 100 000. Sum k for all input sets doesn't exceed 100 000.

## Output

For each input set print the answer as a single line. Print a string of k letters "Y" or "N". Letter "Y" in position i should be printed if they could have run out of the i-th dish by the time the stewardess started serving Polycarp.

## Examples

### Input Example

```text
2

3 4
2 3 2 1
1 0
0 0

5 5
1 2 1 3 1
3 0
0 0
2 1
4 0
```

### Output Example

```text
YNNY
YYYNY
```

## Note

In the first input set depending on the choice of the second passenger the situation could develop in different ways:

- If he chose the first dish, then by the moment the stewardess reaches Polycarp, they will have run out of the first dish;
- If he chose the fourth dish, then by the moment the stewardess reaches Polycarp, they will have run out of the fourth dish;
- Otherwise, Polycarp will be able to choose from any of the four dishes.

Thus, the answer is "YNNY".

In the second input set there is, for example, the following possible scenario. First, the first passenger takes the only third dish, then the second passenger takes the second dish. Then, the third passenger asks for the third dish, but it is not available, so he makes disappointed muttering and ends up with the second dish. Then the fourth passenger takes the fourth dish, and Polycarp ends up with the choice between the first, fourth and fifth dish.

Likewise, another possible scenario is when by the time the stewardess comes to Polycarp, they will have run out of either the first or the fifth dish (this can happen if one of these dishes is taken by the second passenger). It is easy to see that there is more than enough of the fourth dish, so Polycarp can always count on it. Thus, the answer is "YYYNY".


### ideas
1. 不懂～
2. r的作用是什么呢？如果 t[i] > 0, 那么不管r[i]是什么， t[i]都被用掉了
3. 当t[i] = 0 的时候， r[i] = 0 代表什么呢？
4. 如果r[i] = 1, 那么假设j = t[i]， 那么表示很， 第i个人心仪的那份，已经run out了
5. 如果r[i] = 0, 那么 j = t[i]这个时候，还是有的
6. 还有一个点，就是如果考虑i是否run out, 就应该假设所有的t[?] = 0, 都是i
7. 对于i来说，假设位置r[pos] = 0, 且t[pos] = i, pos是最后一次这样的位置
8. 且i一共出现了x次， t[?] = i 的次数 = x
9. let y = a[i] - x 剩余没有出现i的次数 如果 t[?] = 0 的数量 >= y, 貌似肯定可以把i给消耗掉
10. 好像不一定。 如果 0的数量 < y, 肯定是N
11. 如果 r[?] = 1, 但是他得到的是t[?] = i， 那么在他之前，必须有一个被用完了，才可以
12. 假设目前考虑的是i，如果t[?] != i, 那么可以考虑在这个位置前是否能满足用完i的情况，如果不可以，必须使用另外一个j != t[?], j != i 的去满足这个位置
13. 假设使用的是j，那么就会消耗掉一些0的位置
14. 如果t[?] = i, 也需要找到这个一个j，先消耗掉
15. 好复杂～
16. 假设r[?] = 1, 且t[?] = i, 
17. 所以，此时要找到一个后面没有出现过的j，(且y <= 前面的0的那些)， 然后，要用最少的那个j（留下足够的0）
18. 如果r[?] = 0, 且t[?] = i, 且i是最后一次出现，那么就需要判断，前面是否有足够的0，消耗掉i
19. 如果r[?] = 1, 且t[?] = 0; 这个貌似和第一个case一样，也需要找到符合条件的j，且需要消耗掉以后; 然后0的计数+1
20. 如果r[?] = 0, 且t[?] = 0, 那么此时0计数+1
21. 找到最小的j，没有问题（用个heap）就可以做到
22. 只有当一个j，后面不再出现时，它才能进入那个去消耗0的队列