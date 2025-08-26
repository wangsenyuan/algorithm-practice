# Problem C: Movie Company Rating

## Problem Description

A movie company has released 2 movies. These 2 movies were watched by $n$ people. For each person, we know their attitude towards the first movie (liked it, neutral, or disliked it) and towards the second movie.

If a person is asked to leave a review for the movie, then:

- **If that person liked the movie**: they will leave a positive review, and the movie's rating will increase by 1
- **If that person disliked the movie**: they will leave a negative review, and the movie's rating will decrease by 1  
- **Otherwise (neutral)**: they will leave a neutral review, and the movie's rating will not change

Every person will review exactly one movie — and for every person, you can choose which movie they will review.

The company's rating is the **minimum** of the ratings of the two movies. Your task is to calculate the **maximum possible rating** of the company.

## Input

- The first line contains a single integer $t$ $(1 \leq t \leq 10^4)$ — the number of test cases.
- The first line of each test case contains a single integer $n$ $(1 \leq n \leq 2 \cdot 10^5)$.
- The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ $(-1 \leq a_i \leq 1)$, where $a_i$ is equal to:
  - $-1$ if the first movie was disliked by the $i$-th viewer
  - $1$ if the first movie was liked
  - $0$ if the attitude is neutral
- The third line contains $n$ integers $b_1, b_2, \ldots, b_n$ $(-1 \leq b_i \leq 1)$, where $b_i$ is equal to:
  - $-1$ if the second movie was disliked by the $i$-th viewer
  - $1$ if the second movie was liked
  - $0$ if the attitude is neutral

**Additional constraint**: The sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, print a single integer — the maximum possible rating of the company, if for each person, you choose which movie to leave a review on.

## Example

### Input
```
4
2
-1 1
-1 -1
1
-1
-1
5
0 -1 1 0 1
-1 1 0 0 1
4
-1 -1 -1 1
-1 1 1 1
```

### Output
```
0
-1
1
1
```

## Explanation

- **Test case 1**: With 2 people, optimal strategy gives company rating 0
- **Test case 2**: With 1 person, company rating is -1
- **Test case 3**: With 5 people, optimal strategy gives company rating 1
- **Test case 4**: With 4 people, optimal strategy gives company rating 1