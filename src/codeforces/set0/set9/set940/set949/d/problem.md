Instructors of Some Informatics School make students go to bed.

The house contains $n$ rooms, in each room exactly $b$ students were supposed to sleep. However, at the time of curfew it happened that many students are not located in their assigned rooms. The rooms are arranged in a row and numbered from $1$ to $n$. Initially, in the $i$-th room there are $a_i$ students. All students are currently somewhere in the house, therefore
\[
a_1 + a_2 + \dots + a_n = n b.
\]
Also 2 instructors live in this house.

The process of curfew enforcement is the following. One instructor starts near room $1$ and moves toward room $n$, while the second instructor starts near room $n$ and moves toward room $1$. After processing the current room, each instructor moves on to the next one. Both instructors enter rooms and move simultaneously; if $n$ is odd, then only the first instructor processes the middle room. When all rooms are processed, the process ends.

When an instructor processes a room, she counts the number of students in the room, then turns off the light, and locks the room. Also, if the number of students inside the processed room is not equal to $b$, the instructor writes down the number of this room into her notebook. Instructors are in a hurry (to prepare the study plan for the next day), so they don't care about who is in the room, but only about the number of students.

While instructors are inside the rooms, students can run between rooms that are not locked and not being processed. A student can run by at most $d$ rooms, that is she can move to a room with number that differs by at most $d$. Also, after (or instead of) running each student can hide under a bed in a room she is in. In this case the instructor will not count her during the processing. In each room any number of students can hide simultaneously.

Formally, here is what's happening:

1. A curfew is announced, at this point in room $i$ there are $a_i$ students.
2. Each student can run to another room but not further than $d$ rooms away from her initial room, or stay in place. After that each student can optionally hide under a bed.
3. Instructors enter room $1$ and room $n$, they count students there and lock the room (after it no one can enter or leave this room).
4. Each student from rooms with numbers from $2$ to $n-1$ can run to another room but not further than $d$ rooms away from her current room, or stay in place. Each student can optionally hide under a bed.
5. Instructors move from room $1$ to room $2$ and from room $n$ to room $n-1$.
6. This process continues until all rooms are processed.

Let $x_1$ denote the number of rooms in which the first instructor counted the number of non-hidden students different from $b$, and $x_2$ be the same number for the second instructor. Students know that the principal will only listen to one complaint, therefore they want to minimize $\max(x_1, x_2)$. Help them find this value if they use the optimal strategy.

## Input

The first line contains three integers $n$, $d$ and $b$ ($2 \le n \le 100\,000$, $1 \le d \le n-1$, $1 \le b \le 10\,000$), number of rooms in the house, running distance of a student, official number of students in a room.

The second line contains $n$ integers $a_1, a_2, \dots, a_n$ ($0 \le a_i \le 10^9$), the $i$-th of which stands for the number of students in the $i$-th room before curfew announcement.

It is guaranteed that $a_1 + a_2 + \dots + a_n = n b$.

## Output

Output one integer, the minimal possible value of $\max(x_1, x_2)$.

## Examples

### Input
```
5 1 1
1 0 0 0 4
```

### Output
```
1
```

### Input
```
6 1 2
3 8 0 1 0 0
```

### Output
```
2
```

## Note

In the first sample the first three rooms are processed by the first instructor, and the last two are processed by the second instructor. One of the optimal strategies is the following: firstly three students run from room $5$ to room $4$, on the next stage two of them run to room $3$, and one of those two hides under a bed. This way, the first instructor writes down room $2$, and the second writes down nothing.

In the second sample one of the optimal strategies is the following: firstly all students in room $1$ hide, all students from room $2$ run to room $3$. On the next stage one student runs from room $3$ to room $4$, and 5 students hide. This way, the first instructor writes down rooms $1$ and $2$, the second instructor writes down rooms $5$ and $6$.


### ideas
1. x是那些a[i] != b的room的数量
2. 如果一个房间里面，超过了b个人，可以算做b个人（他们可以藏起来）
3. 如果一个房间里面，不足b个人（无论如何都不行）， 那么就往中间移动
4. 用二分; 而且，不足的部分，应该尽量的在前部完成（因为后面后足够的时间安排）
5. 理解错了。 是可以多次移动的