# Problem C

General Payne has a battalion of n soldiers. The soldiers' beauty contest is coming up, it will last for k days. Payne decided that his battalion will participate in the pageant. Now he has to choose the participants.

All soldiers in the battalion have different beauty that is represented by a positive integer. The value aᵢ represents the beauty of the i-th soldier.

On each of k days General has to send a detachment of soldiers to the pageant. The beauty of the detachment is the sum of the beauties of the soldiers who are part of this detachment. Payne wants to surprise the jury of the beauty pageant, so each of k days the beauty of the sent detachment should be unique. In other words, all k beauties of the sent detachments must be distinct numbers.

Help Payne choose k detachments of different beauties for the pageant. Please note that Payne cannot just forget to send soldiers on one day, that is, the detachment of soldiers he sends to the pageant should never be empty.

## Input

The first line contains two integers n, k (1 ≤ n ≤ 50; 1 ≤ k ≤ 2ⁿ − 1) — the number of soldiers and the number of days in the pageant, correspondingly. The second line contains space-separated integers a₁, a₂, …, aₙ (1 ≤ aᵢ ≤ 10⁷) — the beauties of the battalion soldiers.

It is guaranteed that Payne's battalion doesn't have two soldiers with the same beauty.

## Output

Print k lines: in the i-th line print the description of the detachment that will participate in the pageant on the i-th day. The description consists of integer cᵢ (1 ≤ cᵢ ≤ n) — the number of soldiers in the detachment on the i-th day of the pageant and cᵢ distinct integers p₁,ᵢ, p₂,ᵢ, …, p_{cᵢ},ᵢ — the beauties of the soldiers in the detachment on the i-th day of the pageant. The beauties of the soldiers are allowed to print in any order.

Separate numbers on the lines by spaces. It is guaranteed that there is a solution that meets the problem conditions. If there are multiple solutions, print any of them.

## Examples

**Input:**
```
3 3
1 2 3
```

**Output:**
```
1 1
1 2
2 3 2
```

**Input:**
```
2 1
7 12
```

**Output:**
```
1 12
```

### ideas
1. k > n 的时候？
2. pow(2, n) - 1 个组合，所以找到k个不同的，应该是很容易的
3. 