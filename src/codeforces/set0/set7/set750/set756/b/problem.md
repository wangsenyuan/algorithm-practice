# Problem B: Innovative Ticketing System

## Problem Description

A new innovative ticketing system for public transport is introduced in Bytesburg. Now there is a single travel card for all transport. To make a trip, a passenger scans his card and then he is charged according to the fare.

## Fare Structure

The fare is constructed in the following manner. There are three types of tickets:

- **Single trip ticket**: costs 20 byteland rubles
- **90-minute ticket**: costs 50 byteland rubles  
- **One-day ticket (1440 minutes)**: costs 120 byteland rubles

**Note**: A ticket for x minutes activated at time t can be used for trips started in time range from t to t + x - 1, inclusive. Assume that all trips take exactly one minute.

## System Behavior

To simplify the choice for the passenger, the system automatically chooses the optimal tickets. After each trip starts, the system:

1. Analyses all the previous trips and the current trip
2. Chooses a set of tickets for these trips with a minimum total cost
3. Charges the passenger the difference between the new minimum cost and what was charged before

Let the minimum total cost of tickets to cover all trips from the first to the current be `a`, and the total sum charged before be `b`. Then the system charges the passenger the sum `a - b`.

## Task

You have to write a program that, for given trips made by a passenger, calculates the sum the passenger is charged after each trip.

## Input

- The first line contains integer `n` (1 ≤ n ≤ 10^5) — the number of trips made by passenger
- Each of the following `n` lines contains the time of trip `ti` (0 ≤ ti ≤ 10^9), measured in minutes from the time of starting the system
- All `ti` are different, given in ascending order, i.e. `ti+1 > ti` holds for all 1 ≤ i < n

## Output

Output `n` integers. For each trip, print the sum the passenger is charged after it.

## Examples

### Example 1

**Input:**
```
3
10
20
30
```

**Output:**
```
20
20
10
```

### Example 2

**Input:**
```
10
13
45
46
60
103
115
126
150
256
516
```

**Output:**
```
20
20
10
0
20
0
0
20
20
10
```

## Explanation

In the first example, the system works as follows:

- For the first and second trips, it is cheaper to pay for two one-trip tickets, so each time 20 rubles is charged
- After the third trip, the system understands that it would be cheaper to buy a ticket for 90 minutes
- This ticket costs 50 rubles, and the passenger had already paid 40 rubles, so it is necessary to charge 10 rubles only


