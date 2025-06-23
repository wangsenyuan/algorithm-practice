# Olympiad of Metropolises

## Problem Description

Country of Metropolia is holding Olympiad of Metropolises soon. It means that all jury members of the olympiad should meet together in Metropolis (the capital of the country) for the problem preparation process.

There are $n + 1$ cities consecutively numbered from $0$ to $n$. City $0$ is Metropolis that is the meeting point for all jury members. For each city from $1$ to $n$ there is exactly one jury member living there. Olympiad preparation is a long and demanding process that requires $k$ days of work. For all of these $k$ days each of the $n$ jury members should be present in Metropolis to be able to work on problems.

You know the flight schedule in the country (jury members consider themselves important enough to only use flights for transportation). All flights in Metropolia are either going to Metropolis or out of Metropolis. There are no night flights in Metropolia, or in other words, plane always takes off at the same day it arrives. On his arrival day and departure day jury member is not able to discuss the olympiad. All flights in Megapolia depart and arrive at the same day.

Gathering everybody for $k$ days in the capital is a hard objective, doing that while spending the minimum possible money is even harder. Nevertheless, your task is to arrange the cheapest way to bring all of the jury members to Metropolis, so that they can work together for $k$ days and then send them back to their home cities. Cost of the arrangement is defined as a total cost of tickets for all used flights. It is allowed for jury member to stay in Metropolis for more than $k$ days.

## Input

The first line of input contains three integers $n$, $m$ and $k$ ($1 \leq n \leq 10^5$, $0 \leq m \leq 10^5$, $1 \leq k \leq 10^6$).

The $i$-th of the following $m$ lines contains the description of the $i$-th flight defined by four integers $d_i$, $f_i$, $t_i$ and $c_i$ ($1 \leq d_i \leq 10^6$, $0 \leq f_i \leq n$, $0 \leq t_i \leq n$, $1 \leq c_i \leq 10^6$, exactly one of $f_i$ and $t_i$ equals zero), the day of departure (and arrival), the departure city, the arrival city and the ticket cost.

## Output

Output the only integer that is the minimum cost of gathering all jury members in city $0$ for $k$ days and then sending them back to their home cities.

If it is impossible to gather everybody in Metropolis for $k$ days and then send them back to their home cities, output "-1" (without the quotes).

