# Flight Departure Schedule Problem

## Problem Description

Helen works in Metropolis airport. She is responsible for creating a departure schedule. There are n flights that must depart today, the i-th of them is planned to depart at the i-th minute of the day.

Metropolis airport is the main transport hub of Metropolia, so it is difficult to keep the schedule intact. This is exactly the case today: because of technical issues, no flights were able to depart during the first k minutes of the day, so now the new departure schedule must be created.

All n scheduled flights must now depart at different minutes between (k + 1)-th and (k + n)-th, inclusive. However, it's not mandatory for the flights to depart in the same order they were initially scheduled to do so — their order in the new schedule can be different. There is only one restriction: **no flight is allowed to depart earlier than it was supposed to depart in the initial schedule**.

Helen knows that each minute of delay of the i-th flight costs airport ci burles. Help her find the order for flights to depart in the new schedule that minimizes the total cost for the airport.

## Input

The first line contains two integers n and k (1 ≤ k ≤ n ≤ 300,000), where:
- n is the number of flights
- k is the number of minutes in the beginning of the day that the flights did not depart

The second line contains n integers c₁, c₂, ..., cₙ (1 ≤ cᵢ ≤ 10⁷), where cᵢ is the cost of delaying the i-th flight for one minute.

## Output

The first line must contain the minimum possible total cost of delaying the flights.

The second line must contain n different integers t₁, t₂, ..., tₙ (k + 1 ≤ tᵢ ≤ k + n), where tᵢ is the minute when the i-th flight must depart. If there are several optimal schedules, print any of them.

## ideas
1. 假设最后的安排是 d[i] (表示i离开的时间)
2. d[i] >= i (d[i] > k)
3. and result = sum(d[i] - i) * c[i]
4. d[i] * c[i] - i * c[i] 后半部分是固定的
5. d[i] * c[i]越小越好（满足 d[i] >= max(i, k))
6. 也就是说c[i] 越大，理论上越应该早安排
7. 用优先队列