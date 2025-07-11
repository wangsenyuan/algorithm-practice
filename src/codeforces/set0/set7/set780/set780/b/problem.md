# Problem Description

The main road in Bytecity is a straight line from south to north. Conveniently, there are coordinates measured in meters from the southernmost building in north direction.

At some points on the road there are **n** friends, and the **i-th** of them is standing at the point **xi** meters and can move with any speed no greater than **vi** meters per second in any of the two directions along the road: south or north.

You are to compute the **minimum time** needed to gather all the **n** friends at some point on the road. Note that the point they meet at doesn't need to have integer coordinate.

## Input

- The first line contains single integer **n** (2 ≤ n ≤ 60,000) — the number of friends.
- The second line contains **n** integers **x₁, x₂, ..., xₙ** (1 ≤ xᵢ ≤ 10⁹) — the current coordinates of the friends, in meters.
- The third line contains **n** integers **v₁, v₂, ..., vₙ** (1 ≤ vᵢ ≤ 10⁹) — the maximum speeds of the friends, in meters per second.

## Output

Print the **minimum time** (in seconds) needed for all the **n** friends to meet at some point on the road.

## Precision Requirements

Your answer will be considered correct if its absolute or relative error isn't greater than **10⁻⁶**. 

Formally, let your answer be **a**, while jury's answer be **b**. Your answer will be considered correct if:

|a - b| ≤ 10⁻⁶ or |a - b| / max(|a|, |b|) ≤ 10⁻⁶