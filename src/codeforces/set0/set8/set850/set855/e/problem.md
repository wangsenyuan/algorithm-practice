# Harry Potter and the Magic Numbers

## Problem Description

Harry came to know from Dumbledore that Salazar Slytherin's locket is a horcrux. This locket was present earlier at 12 Grimmauld Place, the home of Sirius Black's mother. It was stolen from there and is now present in the Ministry of Magic in the office of Dolorous Umbridge, Harry's former Defense Against the Dark Arts teacher.

Harry, Ron and Hermione are infiltrating the Ministry. Upon reaching Umbridge's office, they observed a code lock with a puzzle asking them to calculate count of magic numbers between two integers $l$ and $r$ (both inclusive).

Harry remembered from his detention time with Umbridge that she defined a **magic number** as a number which when converted to a given base $b$, all the digits from $0$ to $b-1$ appear even number of times in its representation without any leading zeros.

You have to answer $q$ queries to unlock the office. Each query has three integers $b_i$, $l_i$ and $r_i$, the base and the range for which you have to find the count of magic numbers.

## Input

- First line of input contains $q$ $(1 \leq q \leq 10^5)$ — number of queries.
- Each of the next $q$ lines contain three space separated integers $b_i$, $l_i$, $r_i$ $(2 \leq b_i \leq 10, 1 \leq l_i \leq r_i \leq 10^{18})$.

## Output

You have to output $q$ lines, each containing a single integer, the answer to the corresponding query.

## Examples

### Example 1

**Input:**
```
2
2 4 9
3 1 10
```

**Output:**
```
1
2
```

### Example 2

**Input:**
```
2
2 1 100
5 1 100
```

**Output:**
```
21
4
```

## Note

In sample test case 1, for first query, when we convert numbers 4 to 9 into base 2, we get:

- $4 = 100_2$
- $5 = 101_2$
- $6 = 110_2$
- $7 = 111_2$
- $8 = 1000_2$
- $9 = 1001_2$

Out of these, only base 2 representation of 9 has even number of 1's and 0's. Thus, the answer is 1.


### ideas
1. 要在b进制数中，找出 x >= l, x <= r (l,r是10进制下的数)
2. x中的出现的digit，出现次数为偶数
3. 将r按照b进制表示出来, dp[i][mask][0/1]表示到第i位为止，由mask表示的数（出现了奇数次），且是否小于(0)或者等于(1) r的计数 