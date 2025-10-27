Winter holidays are coming up. They are going to last for 𝑛 days.

During the holidays, Monocarp wants to try all of these activities exactly once with his friends:

go skiing;
watch a movie in a cinema;
play board games.

Monocarp knows that, on the 𝑖-th day, exactly 𝑎𝑖 friends will join him for skiing, 𝑏𝑖 friends will join him for a movie and 𝑐𝑖 friends will join him for board games.

Monocarp also knows that he can't try more than one activity in a single day.

Thus, he asks you to help him choose three distinct days 𝑥, 𝑦, 𝑧 in such a way that the total number of friends to join him for the activities (𝑎𝑥 + 𝑏𝑦 + 𝑐𝑧) is maximized.

## Input

The first line contains a single integer 𝑡 (1 ≤ 𝑡 ≤ 10⁴) — the number of testcases.

The first line of each testcase contains a single integer 𝑛 (3 ≤ 𝑛 ≤ 10⁵) — the duration of the winter holidays in days.

The second line contains 𝑛 integers 𝑎₁, 𝑎₂, …, 𝑎𝑛 (1 ≤ 𝑎ᵢ ≤ 10⁸) — the number of friends that will join Monocarp for skiing on the 𝑖-th day.

The third line contains 𝑛 integers 𝑏₁, 𝑏₂, …, 𝑏𝑛 (1 ≤ 𝑏ᵢ ≤ 10⁸) — the number of friends that will join Monocarp for a movie on the 𝑖-th day.

The fourth line contains 𝑛 integers 𝑐₁, 𝑐₂, …, 𝑐𝑛 (1 ≤ 𝑐ᵢ ≤ 10⁸) — the number of friends that will join Monocarp for board games on the 𝑖-th day.

The sum of 𝑛 over all testcases doesn't exceed 10⁵.

## Output

For each testcase, print a single integer — the maximum total number of friends that can join Monocarp for the activities on three distinct days.

## Example

### Input
```
4
3
1 10 1
10 1 1
1 1 10
4
30 20 10 1
30 5 15 20
30 25 10 10
10
5 19 12 3 18 18 6 17 10 13
15 17 19 11 16 3 11 17 17 17
1 17 18 10 15 8 17 3 13 12
10
17 5 4 18 12 4 11 2 16 16
8 4 14 19 3 12 6 7 5 16
3 4 8 11 10 8 10 2 20 3
```

### Output
```
30
75
55
56
```

## Note

In the first testcase, Monocarp can choose day 2 for skiing, day 1 for a movie and day 3 for board games. This way, 𝑎₂ = 10 friends will join him for skiing, 𝑏₁ = 10 friends will join him for a movie and 𝑐₃ = 10 friends will join him for board games. The total number of friends is 30.

In the second testcase, Monocarp can choose day 1 for skiing, day 4 for a movie and day 2 for board games. 30 + 20 + 25 = 75 friends in total. Note that Monocarp can't choose day 1 for all activities, because he can't try more than one activity in a single day.

In the third testcase, Monocarp can choose day 2 for skiing, day 3 for a movie and day 7 for board games. 19 + 19 + 17 = 55 friends in total.

In the fourth testcase, Monocarp can choose day 1 for skiing, day 4 for a movie and day 9 for board games. 17 + 19 + 20 = 56 friends in total.