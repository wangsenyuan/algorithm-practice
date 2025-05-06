Vadim loves filling square tables with integers. But today he came up with a way to do it for fun! Let's take, for example, a table of size 2×2
, with rows numbered from top to bottom and columns numbered from left to right. We place 1
 in the top left cell, 2
 in the bottom right, 3
 in the bottom left, and 4
 in the top right. That's all he needs for fun!

Fortunately for Vadim, he has a table of size 2𝑛×2𝑛
. He plans to fill it with integers from 1
 to 22𝑛
 in ascending order. To fill such a large table, Vadim will divide it into 4
 equal square tables, filling the top left one first, then the bottom right one, followed by the bottom left one, and finally the top right one. Each smaller table will be divided into even smaller ones as he fills them until he reaches tables of size 2×2
, which he will fill in the order described above.

Now Vadim is eager to start filling the table, but he has 𝑞
 questions of two types:

what number will be in the cell at the 𝑥
-th row and 𝑦
-th column;
in which cell coordinates will the number 𝑑
 be located.