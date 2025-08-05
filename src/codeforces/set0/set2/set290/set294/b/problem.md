# Shaass and Bookshelf

## Problem Description

Shaass has n books. He wants to make a bookshelf for all his books. He wants the bookshelf's dimensions to be as small as possible. The thickness of the i-th book is ti and its pages' width is equal to wi. The thickness of each book is either 1 or 2. All books have the same page heights.

Shaass puts the books on the bookshelf in the following way:
1. First he selects some of the books and puts them vertically
2. Then he puts the rest of the books horizontally above the vertical books
3. The sum of the widths of the horizontal books must be no more than the total thickness of the vertical books

A sample arrangement of the books is depicted in the figure.

**Goal:** Help Shaass to find the minimum total thickness of the vertical books that we can achieve.

## Input

- The first line of the input contains an integer n, (1 ≤ n ≤ 100)
- Each of the next n lines contains two integers ti and wi denoting the thickness and width of the i-th book correspondingly, (1 ≤ ti ≤ 2, 1 ≤ wi ≤ 100)

## Output

On the only line of the output print the minimum total thickness of the vertical books that we can achieve.


### ideas. 
1. 每本书的厚度(t), 要么是1，要么是2
2. 每本书的宽度(w) <= 100
3. 假设选择了一些书水平放置，那么总的宽度是 sum(w), 那么 sum(w) <= sum(t 剩余的部分)
4. 有一个简单的观察是，w越小的，放在上面越好，越大的，放在下面越好
5. 如果把书放在上面，那么下面减少t[i], 上面增加w[i]
6. 所以，应该按照w[i] - t[i]升序排列
7.  
