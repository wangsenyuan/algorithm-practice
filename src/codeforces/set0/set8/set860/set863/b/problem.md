# Kayaking Problem

## Problem Description

Vadim is really keen on travelling. Recently he heard about kayaking activity near his town and became very excited about it, so he joined a party of kayakers.

Now the party is ready to start its journey, but firstly they have to choose kayaks. There are **2·n** people in the group (including Vadim), and they have exactly:
- **n - 1** tandem kayaks (each of which can carry two people)
- **2** single kayaks

The i-th person's weight is **wᵢ**, and weight is an important matter in kayaking — if the difference between the weights of two people that sit in the same tandem kayak is too large, then it can crash. And, of course, people want to distribute their seats in kayaks in order to minimize the chances that kayaks will crash.

## Instability Definition

- The **instability of a single kayak** is always **0**
- The **instability of a tandem kayak** is the absolute difference between weights of the people that are in this kayak
- The **instability of the whole journey** is the total instability of all kayaks

## Task

Help the party to determine the **minimum possible total instability**!

## Input

- The first line contains one number **n** (2 ≤ n ≤ 50)
- The second line contains **2·n** integer numbers **w₁, w₂, ..., w₂ₙ**, where **wᵢ** is the weight of person i (1 ≤ wᵢ ≤ 1000)

## Output

Print the **minimum possible total instability**.

## ideas
1. 排序后，然后选择把i分配个单独的（那么它前面的必须是偶数，而且是依次配对）
2. 