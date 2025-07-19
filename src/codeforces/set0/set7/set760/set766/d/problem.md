# Mahmoud and Ehab and the Dictionary

## Problem Description

Mahmoud wants to write a new dictionary that contains n words and relations between them. There are two types of relations:
- **Synonymy**: Two words mean the same
- **Antonymy**: Two words mean the opposite

From time to time he discovers a new relation between two words.

### Relation Propagation Rules

He knows that if two words have a relation between them, then each of them has relations with the words that have relations with the other. For example:
- If "like" means "love" and "love" is the opposite of "hate", then "like" is also the opposite of "hate"
- If "love" is the opposite of "hate" and "hate" is the opposite of "like", then "love" means "like"

### Wrong Relations

Sometimes Mahmoud discovers a wrong relation. A wrong relation is a relation that makes two words equal and opposite at the same time. For example:
- If he knows that "love" means "like" and "like" is the opposite of "hate", and then he figures out that "hate" means "like", the last relation is absolutely wrong because it makes "hate" and "like" opposite and have the same meaning at the same time.

After Mahmoud figured out many relations, he was worried that some of them were wrong so that they will make other relations also wrong, so he decided to tell every relation he figured out to his coder friend Ehab and for every relation he wanted to know is it correct or wrong, basing on the previously discovered relations. If it is wrong he ignores it, and doesn't check with following relations.

After adding all relations, Mahmoud asked Ehab about relations between some words based on the information he had given to him. Ehab is busy making a Codeforces round so he asked you for help.

## Input

The first line of input contains three integers n, m and q (2 ≤ n ≤ 10^5, 1 ≤ m, q ≤ 10^5) where:
- n is the number of words in the dictionary
- m is the number of relations Mahmoud figured out
- q is the number of questions Mahmoud asked after telling all relations

The second line contains n distinct words a1, a2, ..., an consisting of small English letters with length not exceeding 20, which are the words in the dictionary.

Then m lines follow, each of them contains:
- An integer t (1 ≤ t ≤ 2) followed by two different words xi and yi which have appeared in the dictionary words
- If t = 1, that means xi has a synonymy relation with yi
- If t = 2, that means xi has an antonymy relation with yi

Then q lines follow, each of them contains two different words which have appeared in the dictionary. These are the pairs of words Mahmoud wants to know the relation between basing on the relations he had discovered.

**Note**: All words in input contain only lowercase English letters and their lengths don't exceed 20 characters. In all relations and in all questions the two words are different.

## Output

First, print m lines, one per each relation:
- If some relation is wrong (makes two words opposite and have the same meaning at the same time) you should print "NO" (without quotes) and ignore it
- Otherwise print "YES" (without quotes)

After that print q lines, one per each question:
- If the two words have the same meaning, output 1
- If they are opposites, output 2
- If there is no relation between them, output 3

## Examples

### Example 1

**Input:**
```
3 3 4
hate love like
1 love like
2 love hate
1 hate like
love like
love hate
like hate
hate like
```

**Output:**
```
YES
YES
NO
1
2
2
2
```

### Example 2

**Input:**
```
8 6 5
hi welcome hello ihateyou goaway dog cat rat
1 hi welcome
1 ihateyou goaway
2 hello ihateyou
2 hi goaway
2 hi hello
1 hi hello
dog cat
dog hi
hi hello
ihateyou goaway
welcome ihateyou
```

**Output:**
```
YES
YES
YES
YES
NO
YES
3
3
1
1
2
```

