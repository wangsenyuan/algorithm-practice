### Problem

Valera is at the airport with `n` robots and wants to move as many as possible to the luggage compartment, which is at distance `d` meters away. He has `S` liters of fuel.

Each robot `i` has three parameters:

- `c_i` — how many other robots it can carry (capacity of compartments),
- `f_i` — liters of fuel required if this robot moves by itself,
- `l_i` — maximum distance this robot can travel on its own.

Robots behave as follows:

- First, Valera chooses some robots that will travel to the luggage compartment **on their own**.  
  The total fuel needed for these robots must not exceed `S`.
- Then he arranges robots inside other robots’ compartments, forming a hierarchy.  
  A robot that does **not** move on its own can be placed inside another robot (even inside another non-moving robot), as long as the entire nested chain is ultimately carried by a moving robot.
- All moving robots, plus all robots (directly or indirectly) inside them, travel straight to the luggage compartment; the rest are lost.
- Any robot that participates in carrying others (directly or indirectly) must have `l_i ≥ d` (it can cover the distance `d`).
- During movement, Valera cannot stop or rearrange the robots.

Your task is to:

- **maximize** the total number of robots that reach the luggage compartment;
- among all such ways, **minimize** the total fuel spent (sum of `f_i` of robots that move on their own).

If Valera cannot deliver any robots, the answer is `0 0`.

### Input

- The first line contains three integers `n`, `d`, `S`  
  - `1 ≤ n ≤ 10^5` — number of robots  
  - `1 ≤ d, S ≤ 10^9` — distance to luggage compartment and available fuel
- Each of the next `n` lines contains three integers `c_i`, `f_i`, `l_i`  
  - `0 ≤ c_i, f_i, l_i ≤ 10^9`

### Output

Print two space-separated integers:

- the **maximum number of robots** that can be transported to the luggage compartment;
- the **minimum total fuel** required to achieve that maximum.

If it is impossible to transport any robots, print `0 0`.

### Examples

**Input**
```
3 10 10
0 12 10
1 6 10
0 1 1
```

**Output**
```
2 6
```

**Input**
```
2 7 10
3 12 10
5 16 8
```

**Output**
```
0 0
```

**Input**
```
4 8 10
0 12 3
1 1 0
0 3 11
1 6 9
```

**Output**
```
4 9
```

### editorial

In this problem, there are two cases to consider.

The first is to let all robots drive independently. Then we'll select a set of robots with l i  ≥  d , sort them by increasing amount of fuel required, and then take them in that order and send them to the luggage compartment until the fuel runs out. This case is similar to the solution to Problem C in the same round.

The second case considers robots that nest within each other. Let's select a set of robots with c i  > 0 , let there be k robots, and let . It's easy to see that if we pack all these robots so that only one of them drives independently, while the rest are nested within other robots, we'll still have some free slots for other robots.

Thus, we select a set of robots for which c i  > 0 , l i  ≥  d - that is, these are the robots that can carry other robots and at the same time can independently reach the luggage compartment.

Let's check how many of these robots will actually travel under their own power (remembering the fuel restrictions), and use the formula above to find the number of free slots. Note that we've already dispatched all the robots with c i  > 0 , meaning only those with c i  = 0 remain; let's say there are num of them .

Each robot can have 3 states: it will occupy one of the slots, it will move under its own power if it has l i  ≥  d and if it has enough fuel, or it will not move at all.

We know how many slots there are, which means we can determine how many robots will either go on their own or stay behind, or more precisely, how many . Thus, among all robots with c i  = 0 and l i  ≥  d, we need to find the maximum number of robots (but no more than g ) that have enough fuel remaining. The remaining robots will either go on free slots or stay behind (their number is known).

This subproblem is solved by pre-computing the value f ( i ) , which represents the minimum amount of fuel required to dispatch i robots among those for which c i  = 0 and l i  ≥  d . A binary search can obviously be performed on this function.

Thus, we will collect the elements of the solution: first, we will enumerate the number of robots that will go under their own power among those for which c i  > 0 , and then, using a binary search for the pre-calculated value, we will find how many of the remaining robots will fit into the free slots, how many will also go under their own power, and how many will remain.

Let's not forget the case when all robots drive independently.

We obtain an algorithm that can be implemented in O ( n log ( n ) ) time .

### step-by-step explanation

The key is to separate the problem into two layers:

1. Which robots must spend fuel and drive by themselves.
2. Which robots can be taken for free by placing them inside other robots.

The answer is the best balance between these two.

#### Step 1: Observe what matters structurally

If a robot has `c_i > 0`, it can contribute carrying capacity.

If we take several such robots and make exactly `t` of them drive by themselves, then:

- all selected driving robots must satisfy `l_i >= d`;
- the non-driving `c_i > 0` robots may still be used, because they can be nested inside a driving robot;
- every `c_i > 0` robot contributes `c_i` slots, but if that robot itself is transported inside the structure, it occupies one slot.

So across all robots with `c_i > 0`, the net number of free places is:

`sum(c_i - 1) + t`

Why?

- each positive-capacity robot contributes `c_i`;
- each such robot also needs one place unless it is itself a root that drives;
- if `t` of them drive, then `t` robots no longer need to occupy another slot.

This is exactly why the code stores:

- `carriers = number of robots with c_i > 0`
- `baseSlots = sum(c_i - 1)` over all robots with `c_i > 0`

and later uses:

- `freeSlots = baseSlots + t`

#### Step 2: Split robots into useful groups

We classify robots into three fuel-sorted groups:

1. `allDrive`
These are all robots with `l_i >= d`.
If we ignore nesting completely and let everybody move independently, then the best strategy is to take the cheapest ones first.

2. `goodCarry`
These are robots with `c_i > 0` and `l_i >= d`.
Only these robots can be chosen as driving roots for a nested structure.

3. `simpleDrive`
These are robots with `c_i = 0` and `l_i >= d`.
They cannot carry others, so they are only useful as:
- self-driving robots, or
- robots occupying already available free slots.

Also:

- `others = number of robots with c_i = 0`

because after deciding to take all `c_i > 0` robots into the structure, the only remaining robots to possibly add are the `c_i = 0` ones.

#### Step 3: Handle the easy case first

Ignore nesting and let every transported robot drive by itself.

Among all robots with `l_i >= d`, sort fuel costs increasingly and take the longest affordable prefix under budget `S`.

This gives one valid candidate answer:

- transported robots = number of affordable drivable robots
- fuel = sum of their costs

That is what `allDrive` is for.

#### Step 4: Now force a nested structure

Suppose exactly `t` robots from `goodCarry` drive by themselves.

To minimize fuel for this fixed `t`, we should always choose the `t` cheapest robots from `goodCarry`.

Let:

- `carryFuel = fuel of the cheapest t driving carriers`
- `remain = S - carryFuel`

If `carryFuel > S`, this choice is impossible.

At this point:

- all `carriers` robots with `c_i > 0` are already included in the transported set;
- they create `freeSlots = baseSlots + t` available places;
- we can now try to add robots from the `c_i = 0` group.

#### Step 5: How many `c_i = 0` robots can we additionally take?

Each `c_i = 0` robot has only two possible useful roles:

1. occupy one free slot, costing `0` extra fuel;
2. drive by itself, which requires `l_i >= d` and costs fuel.

So with `freeSlots` free positions and `remain` fuel, the best extra robots come from the `others` group:

- first fill as many free slots as possible;
- if we still want more robots, some `simpleDrive` robots must drive.

Among `simpleDrive`, if we want `k` of them to drive, the minimum fuel is the sum of the `k` smallest fuel values.

That is why we sort `simpleDrive` and build prefix sums.

#### Step 6: Why prefix sums + binary search?

For a sorted fuel array:

- prefix sum `pref[k]` = minimum fuel needed to make exactly `k` cheapest robots drive.

So for a remaining budget `remain`, the maximum number of `simpleDrive` robots that can drive is the largest `k` such that:

`pref[k] <= remain`

Because prefix sums are increasing, this can be found by binary search in `O(log n)`.

That is exactly what `maximizeWithBudget()` does.

#### Step 7: Compute the result for one fixed `t`

For a chosen `t`:

- `afford` = maximum number of `simpleDrive` robots that can self-drive within `remain`
- `freeSlots = baseSlots + t`

Then the total number of extra `c_i = 0` robots we can take is:

`additional = min(others, freeSlots + afford)`

Explanation:

- at most `others` such robots exist;
- among them, up to `freeSlots` can be carried for free;
- up to `afford` more can drive by themselves.

But not all `additional` robots need fuel. Only those beyond the free slots need to drive:

`needDrive = max(0, additional - freeSlots)`

So the total fuel for this configuration is:

`totalFuel = carryFuel + fuel of cheapest needDrive robots in simpleDrive`

and the total transported robots are:

`totalCnt = carriers + additional`

We try every possible `t` and keep:

- the largest `totalCnt`
- and among equal counts, the smallest `totalFuel`

#### Step 8: Why this is optimal

For any fixed number `t` of driving carriers:

- choosing any carrier other than the cheapest `t` only increases fuel;
- choosing any self-driving `c_i = 0` robots other than the cheapest possible set also only increases fuel;
- all `c_i > 0` robots should be included, because once we have at least one valid driving carrier root, these robots never hurt and only increase carrying capacity.

So the search over all `t` covers all optimal nested configurations.

Together with the independent-driving case, this gives the global optimum.

#### Step 9: Complexity

We sort a few arrays and do a binary search for each candidate `t`.

Total complexity:

- `O(n log n)` time
- `O(n)` memory

#### Important subtle point

A robot with `c_i > 0` but `l_i < d` can still be transported inside another robot.

It cannot be a driving root, but it still contributes capacity once nested.

This is the main detail that sample 3 depends on, and it is why the implementation counts:

- every robot with `c_i > 0` in `carriers` and `baseSlots`,
- but only robots with both `c_i > 0` and `l_i >= d` in `goodCarry`.