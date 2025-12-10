The king of Berland organizes a ball! 𝑛
 pair are invited to the ball, they are numbered from 1
 to 𝑛
. Each pair consists of one man and one woman. Each dancer (either man or woman) has a monochrome costume. The color of each costume is represented by an integer from 1
 to 𝑘
, inclusive.

Let 𝑏𝑖
 be the color of the man's costume and 𝑔𝑖
 be the color of the woman's costume in the 𝑖
-th pair. You have to choose a color for each dancer's costume (i.e. values 𝑏1,𝑏2,…,𝑏𝑛
 and 𝑔1,𝑔2,…𝑔𝑛
) in such a way that:

for every 𝑖
: 𝑏𝑖
 and 𝑔𝑖
 are integers between 1
 and 𝑘
, inclusive;
there are no two completely identical pairs, i.e. no two indices 𝑖,𝑗
 (𝑖≠𝑗
) such that 𝑏𝑖=𝑏𝑗
 and 𝑔𝑖=𝑔𝑗
 at the same time;
there is no pair such that the color of the man's costume is the same as the color of the woman's costume in this pair, i.e. 𝑏𝑖≠𝑔𝑖
 for every 𝑖
;
for each two consecutive (adjacent) pairs both man's costume colors and woman's costume colors differ, i.e. for every 𝑖
 from 1
 to 𝑛−1
 the conditions 𝑏𝑖≠𝑏𝑖+1
 and 𝑔𝑖≠𝑔𝑖+1
 hold.
Let's take a look at the examples of bad and good color choosing (for 𝑛=4
 and 𝑘=3
, man is the first in a pair and woman is the second):

Bad color choosing:

(1,2)
, (2,3)
, (3,2)
, (1,2)
 — contradiction with the second rule (there are equal pairs);
(2,3)
, (1,1)
, (3,2)
, (1,3)
 — contradiction with the third rule (there is a pair with costumes of the same color);
(1,2)
, (2,3)
, (1,3)
, (2,1)
 — contradiction with the fourth rule (there are two consecutive pairs such that colors of costumes of men/women are the same).
Good color choosing:

(1,2)
, (2,1)
, (1,3)
, (3,1)
;
(1,2)
, (3,1)
, (2,3)
, (3,2)
;
(3,1)
, (1,2)
, (2,3)
, (3,2)
.
You have to find any suitable color choosing or say that no suitable choosing exists.

Input
The only line of the input contains two integers 𝑛
 and 𝑘
 (2≤𝑛,𝑘≤2⋅105
) — the number of pairs and the number of colors.

Output
If it is impossible to find any suitable colors choosing, print "NO".

Otherwise print "YES" and then the colors of the costumes of pairs in the next 𝑛
 lines. The 𝑖
-th line should contain two integers 𝑏𝑖
 and 𝑔𝑖
 — colors of costumes of man and woman in the 𝑖
-th pair, respectively.

You can print each letter in any case (upper or lower). For example, "YeS", "no" and "yES" are all acceptable.

Examples
InputCopy
4 3
OutputCopy
YES
3 1
1 3
3 2
2 3
InputCopy
10 4
OutputCopy
YES
2 1
1 3
4 2
3 4
4 3
3 2
2 4
4 1
1 4
3 1
InputCopy
13 4
OutputCopy
NO