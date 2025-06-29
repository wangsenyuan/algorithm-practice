# Problem Description

Inna and Dima bought a table of size $n \times m$ in the shop. Each cell of the table contains a single letter: "D", "I", "M", "A".

Inna loves Dima, so she wants to go through his name as many times as possible as she moves through the table. For that, Inna acts as follows:

1. Initially, Inna chooses some cell of the table where letter "D" is written;
2. Then Inna can move to some side-adjacent table cell that contains letter "I";
3. Then from this cell she can go to one of the side-adjacent table cells that contains the written letter "M";
4. Then she can go to a side-adjacent cell that contains letter "A". Then Inna assumes that she has gone through her sweetheart's name;
5. Inna's next move can be going to one of the side-adjacent table cells that contains letter "D" and then walk on through name DIMA in the similar manner.

**Important:** Inna never skips a letter. So, from the letter "D" she always goes to the letter "I", from the letter "I" she always goes to the letter "M", from the letter "M" she always goes to the letter "A", and from the letter "A" she always goes to the letter "D".

Depending on the choice of the initial table cell, Inna can go through name DIMA either:
- An infinite number of times, or
- Some positive finite number of times, or
- She can't go through his name once

Help Inna find out what maximum number of times she can go through name DIMA.

## Input

The first line of the input contains two integers $n$ and $m$ ($1 \leq n, m \leq 10^3$).

Then follow $n$ lines that describe Inna and Dima's table. Each line contains $m$ characters. Each character is one of the following four characters: "D", "I", "M", "A".

**Note:** It is not guaranteed that the table contains at least one letter "D".

## Output

- If Inna cannot go through name DIMA once, print on a single line `"Poor Dima!"` without the quotes.
- If there is the infinite number of names DIMA Inna can go through, print `"Poor Inna!"` without the quotes.
- Otherwise print a single integer — the maximum number of times Inna can go through name DIMA.