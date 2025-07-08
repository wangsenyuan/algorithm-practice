# Hobby-Droning Problem

In 2077, a sport called hobby-droning is gaining popularity among robots.

You already have a drone, and you want to win. For this, your drone needs to fly through a course with $n$ obstacles.

## Problem Description

The $i$-th obstacle is defined by two numbers $l_i, r_i$. Let the height of your drone at the $i$-th obstacle be $h_i$. Then the drone passes through this obstacle if $l_i \leq h_i \leq r_i$. Initially, the drone is on the ground, meaning $h_0 = 0$.

The flight program for the drone is represented by an array $d_1, d_2, \ldots, d_n$, where $h_i - h_{i-1} = d_i$, and $0 \leq d_i \leq 1$. This means that your drone either does not change height between obstacles or rises by 1. You already have a flight program, but some $d_i$ in it are unknown and marked as $-1$. Replace the unknown $d_i$ with numbers 0 and 1 to create a flight program that passes through the entire obstacle course, or report that it is impossible.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

In the first line of each test case, an integer $n$ ($1 \leq n \leq 2 \cdot 10^5$) is given — the size of the array $d$.

In the second line of each test case, there are $n$ integers $d_1, d_2, \ldots, d_n$ ($-1 \leq d_i \leq 1$) — the elements of the array $d$. $d_i = -1$ means that this $d_i$ is unknown to you.

Next, there are $n$ lines containing 2 integers $l_i, r_i$ ($0 \leq l_i \leq r_i \leq n$) — descriptions of the obstacles.

It is guaranteed that the sum of $n$ across all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output $n$ integers $d_1, d_2, \ldots, d_n$, if it is possible to correctly restore the array $d$, or $-1$ if it is not possible.

### ideas
1. c * x 