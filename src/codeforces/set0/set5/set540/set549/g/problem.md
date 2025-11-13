# Problem

Do you like summer? Residents of Berland do. They especially love eating ice cream in the hot summer. So this summer day a large queue of n Berland residents lined up in front of the ice cream stall. We know that each of them has a certain amount of berland dollars with them. The residents of Berland are nice people, so each person agrees to swap places with the person right behind him for just 1 dollar. More formally, if person a stands just behind person b, then person a can pay person b 1 dollar, then a and b get swapped. Of course, if person a has zero dollars, he can not swap places with person b.

Residents of Berland are strange people. In particular, they get upset when there is someone with a strictly smaller sum of money in the line in front of them.

Can you help the residents of Berland form such order in the line so that they were all happy? A happy resident is the one who stands first in the line or the one in front of who another resident stands with not less number of dollars. Note that the people of Berland are people of honor and they agree to swap places only in the manner described above.

## Input

The first line contains integer n (1 ≤ n ≤ 200 000) — the number of residents who stand in the line.

The second line contains n space-separated integers ai (0 ≤ ai ≤ 109), where ai is the number of Berland dollars of a man standing on the i-th position in the line. The positions are numbered starting from the end of the line.

## Output

If it is impossible to make all the residents happy, print ":(" without the quotes. Otherwise, print in the single line n space-separated integers, the i-th of them must be equal to the number of money of the person on position i in the new line. If there are multiple answers, print any of them.

## Examples

### Example 1

**Input**
```
2
11 8
```

**Output**
```
9 10
```

### Example 2

**Input**
```
5
10 9 7 10 6
```

**Output**
```
:(
```

### Example 3

**Input**
```
3
12 3 3
```

**Output**
```
4 4 10
```

## Note

In the first sample two residents should swap places, after that the first resident has 10 dollars and he is at the head of the line and the second resident will have 9 coins and he will be at the end of the line.

In the second sample it is impossible to achieve the desired result.

In the third sample the first person can swap with the second one, then they will have the following numbers of dollars: 4 11 3, then the second person (in the new line) swaps with the third one, and the resulting numbers of dollars will equal to: 4 4 10. In this line everybody will be happy.


### ideas
1. 最后的结果肯定是， a[1] <= a[2] <= ... <= a[n]
2. 两个相邻位置，进行交换，只会交换1, 如果j移动到了位置i， 那么j获得 j - i 的钱，如果i移动到了j, 必须给出 j - i 的钱
3. 且b[i] = a[j] + j - i <= b[i+1] <= b[i+2] <= ..
4. 如果 i < j, 且 a[i] > a[j], 那么i必须往前移动，j必须往后移动，两个必须相向移动，i消耗，j增加，知道 a[i] <= a[j]
5. 如果有一个i, a[i] - (n - i) > 它后面所有 a[j] - (n - j) 那么i肯定会移动到位置n
6. 也就是a[i] + i - n 最大的i，必然移动到n,
7. i后面的所有的+1,（但是位置-1了）所以，a[i] + i 整体没有变的
8. 然后a[i] + i - (n - 1)最大的i，移动到n-1 (所以，就是a[i] + i， 最大的那个数) 
9. 但是这样子，岂不是一定有答案了？
10. [10 9 7 10 6], 那么必然是 10 + 4 移动到最后, 
11. [10, 9, 7, 7, 9]
12. [10 + 1, 9 + 2, 7 + 3, 7 + 4, 9 + 5]
13. [10, 8, 8, 7, 9] => false