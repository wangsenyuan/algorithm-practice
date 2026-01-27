This time, the Berland Team Olympiad in Informatics is held in a remote city that can only be reached by one small bus. The bus has n passenger seats, seat i can be occupied only by a participant from the city aᵢ.

Today, the bus has completed m trips, each time bringing n participants. The participants were then aligned in one line in the order they arrived, with people from the same bus standing in the order of their seats (i.e. if we write down the cities where the participants came from, we get the sequence a₁, a₂, …, aₙ repeated m times).

After that some teams were formed, each consisting of k participants from the same city standing next to each other in the line. Once formed, teams left the line. The teams were formed until there were no k neighboring participants from the same city.

Help the organizers determine how many participants are left in the line after that process ended. We can prove that answer doesn't depend on the order in which teams were selected.

## Input

The first line contains three integers n, k and m (1 ≤ n ≤ 10⁵, 2 ≤ k ≤ 10⁹, 1 ≤ m ≤ 10⁹).

The second line contains n integers a₁, a₂, …, aₙ (1 ≤ aᵢ ≤ 10⁵), where aᵢ is the index of city where the person occupying the i-th seat in the bus must be from.

## Output

Output the number of remaining participants in the line.

## Examples

### Input
```
4 2 5
1 2 3 1
```

### Output
```
12
```

### Input
```
1 9 10
1
```

### Output
```
1
```

### Input
```
3 2 10
1 2 1
```

### Output
```
0
```

## Note

In the second example, the line consists of ten participants from the same city. Nine of them will form a team. At the end, only one participant will stay in the line.


### ideas
1. 咋读不懂这个题目呢？

#### First Example Explanation (n=4, k=2, m=5, a=[1,2,3,1])

**Initial Setup:**
- Bus has 4 seats: [city1, city2, city3, city1]
- Bus makes 5 trips
- After all trips, the line is: `1, 2, 3, 1, 1, 2, 3, 1, 1, 2, 3, 1, 1, 2, 3, 1, 1, 2, 3, 1` (20 participants total)

**Key Observation:**
The pattern `[1, 2, 3, 1]` repeats 5 times. Notice that:
- Within each repetition: no consecutive same-city pairs exist
- **At boundaries between trips**: the last element (1) of one trip and first element (1) of next trip form consecutive pairs: `...1, 1...`
- There are **m-1 = 4 boundaries** between 5 trips

**Team Formation Process (k=2):**

1. **Boundary pairs:** At each of the 4 boundaries, we have `[1, 1]` forming a team
   - Remove 4 teams = 8 participants
   - Remaining: 20 - 8 = **12 participants**

2. **After boundary removals:** The line becomes fragmented, but no new consecutive pairs of same city are formed because:
   - The pattern `[1, 2, 3, 1]` has no internal consecutive same-city pairs
   - After removing boundary pairs, the remaining segments don't create new consecutive pairs

**Result:** 12 participants remain in the line.

**Why this works:**
- The sequence `[1, 2, 3, 1]` ends with 1 and starts with 1, creating boundary pairs when repeated
- Each boundary pair forms exactly one team (k=2)
- After removing all boundary pairs, no more k-consecutive same-city groups can be formed


### ideas
1. 分情况讨论
2. 如果全部是都是一样的， 那么答案 = n * m % k
3. 不是一样的情况; 且 m >= 2
4. 然后处理第一段 + 第二段；然后只需要考虑第二段剩余部分重复的情况