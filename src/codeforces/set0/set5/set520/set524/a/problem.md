# Social Network - Putative Friends

The foundation of any social network is a friendship relationship between two users, in one sense or another. In one well-known social network, friendship is symmetrical, meaning that if a is a friend of b, then b is also a friend of a.

This same network also has a function that displays the set of people who are highly likely to be known to a user. This function works as follows. Let's fix a user x. Let's assume that some other person y, who is not currently a friend of x, is a friend of at least k% of x's friends. Then y is a putative friend of x.

Each person on a social network has a unique identifier—an integer between 1 and 10^9. You are given a list of pairs of users who are friends. For each mentioned user, determine the set of their supposed friends.

## Input data

The first line contains two integers m and k (1 ≤ m ≤ 100, 0 ≤ k ≤ 100) — the number of pairs of friends and the required percentage of mutual friends to be considered a prospective friend.

The next m lines contain two numbers a_i, b_i (1 ≤ a_i, b_i ≤ 10^9, a_i ≠ b_i), denoting the identifiers of users who are friends.

It is guaranteed that each pair of people appears in the list no more than once.

## Output data

For all mentioned people, in ascending order of id, output information about their suggested friends. The information should be of the form "id: k id1 id2 ... idk", where id is the person's id, k is the number of their suggested friends, and id1, id2, ..., idk are the identifiers of their suggested friends in ascending order.

## Examples

### Example 1

**Input:**

```text
5 51
10 23
23 42
39 42
10 39
39 58
```

**Output:**

```text
10: 1 42
23: 1 39
39: 1 23
42: 1 10
58: 2 10 42
```

### Example 2

**Input:**

```text
5 100
1 2
1 3
1 4
2 3
2 4
```

**Output:**

```text
1: 0
2: 0
3: 1 4
4: 1 3
```
