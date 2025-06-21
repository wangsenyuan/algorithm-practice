# Problem Description

A sum of **p** rubles is charged from Arkady's mobile phone account every day in the morning. Among the following **m** days, there are **n** days when Arkady will top up the account: on day **di** he will deposit **ti** rubles on his mobile phone account. Arkady will always top up the account before the daily payment is done. There will be no other payments nor top-ups in the following **m** days.

Determine the number of days starting from the 1st to the **m**-th such that the account will have a negative amount on it after the daily payment (i.e., in the evening). Initially, the account's balance is zero rubles.

## Input

The first line contains three integers **n**, **p**, and **m**:
- **n** (1 ≤ n ≤ 100,000) — the number of days Arkady will top up the account
- **p** (1 ≤ p ≤ 10^9) — the amount of the daily payment
- **m** (1 ≤ m ≤ 10^9) — the number of days you should check
- **n ≤ m**

The **i**-th of the following **n** lines contains two integers **di** and **ti**:
- **di** (1 ≤ di ≤ m) — the index of the day when Arkady will make the **i**-th top-up
- **ti** (1 ≤ ti ≤ 10^9) — the amount he will deposit on this day

**Note:** It is guaranteed that the indices of the days are distinct and are given in increasing order, i.e., di > di-1 for all i from 2 to n.

## Output

Print the number of days from the 1st to the **m**-th such that the account will have a negative amount on it after the daily payment.