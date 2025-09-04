# Problem Description

Nearly each project of the F company has a whole team of developers working on it. They often are in different rooms of the office in different cities and even countries. To keep in touch and track the results of the project, the F company conducts shared online meetings in a Spyke chat.

One day the director of the F company got hold of the records of a part of an online meeting of one successful team. The director watched the record and wanted to talk to the team leader. But how can he tell who the leader is? The director logically supposed that the leader is the person who is present at any conversation during a chat meeting. In other words, if at some moment of time at least one person is present on the meeting, then the leader is present on the meeting.

You are the assistant director. Given the 'user logged on'/'user logged off' messages of the meeting in the chronological order, help the director determine who can be the leader. Note that the director has the record of only a continuous part of the meeting (probably, it's not the whole meeting).

## Input

The first line contains integers $n$ and $m$ ($1 \leq n, m \leq 10^5$) — the number of team participants and the number of messages. Each of the next $m$ lines contains a message in the format:

- `+ id`: the record means that the person with number $id$ ($1 \leq id \leq n$) has logged on to the meeting.
- `- id`: the record means that the person with number $id$ ($1 \leq id \leq n$) has logged off from the meeting.

Assume that all the people of the team are numbered from 1 to $n$ and the messages are given in the chronological order. It is guaranteed that the given sequence is the correct record of a continuous part of the meeting. It is guaranteed that no two log on/log off events occurred simultaneously.

## Output

In the first line print integer $k$ ($0 \leq k \leq n$) — how many people can be leaders. In the next line, print $k$ integers in the increasing order — the numbers of the people who can be leaders.

If the data is such that no member of the team can be a leader, print a single number 0.


### ideas
1. 如果目前有一个人y在场，在他的on/off中间，那么leader必须也在场
2. 这些记录是这个meetting的部分记录，并不是全部
3. 那些没有登录过的，有可能是leader
4. 如果一个人先出现了log off, 说明他在之前都是在场的，那么在他之前log on的都不是leader
5. 居然可以重复进入/离开的～
6. 