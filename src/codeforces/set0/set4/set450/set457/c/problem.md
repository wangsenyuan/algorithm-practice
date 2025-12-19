### Problem

You are running for a governor in a small city in Russia. You ran some polls and did some research, and for every person in the city you know whom he will vote for, and how much it will cost to bribe that person to vote for you instead of whomever he wants to vote for right now. You are curious, what is the smallest amount of money you need to spend on bribing to win the elections. To win elections you need to have strictly more votes than any other candidate.

### Input

First line contains one integer $n$ $(1 \le n \le 10^5)$ — number of voters in the city. Each of the next $n$ lines describes one voter and contains two integers $a_i$ and $b_i$ $(0 \le a_i \le 10^5; 0 \le b_i \le 10^4)$ — number of the candidate that voter is going to vote for and amount of money you need to pay him to change his mind. You are the candidate $0$ (so if a voter wants to vote for you, $a_i$ is equal to zero, in which case $b_i$ will also be equal to zero).

### Output

Print one integer — smallest amount of money you need to spend to win the elections.

### Examples

**Input**

```text
5
1 2
1 2
1 2
2 1
0 0
```

**Output**

```text
3
```

**Input**

```text
4
1 2
1 2
2 1
0 0
```

**Output**

```text
2
```

**Input**

```text
1
100000 0
```

**Output**

```text
0
```
