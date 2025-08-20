# Problem E: Comments Feed

## Problem Description

A rare article on the Internet is posted without the possibility to comment on it. On Polycarp's website, each article has a comments feed.

Each comment on Polycarp's website is a non-empty string consisting of uppercase and lowercase letters of the English alphabet. Comments have a tree-like structure, meaning each comment (except root comments at the highest level) has exactly one parent comment.

## Input Format

When Polycarp saves comments to his hard drive, he uses the following format:

1. **Comment text**: The text of the comment is written first
2. **Reply count**: The number of comments for which this comment is a parent comment (i.e., the number of replies to this comment)
3. **Replies**: The comments for which this comment is a parent comment are written (using the same algorithm recursively)

All elements in this format are separated by single commas. Similarly, comments at the first level are separated by commas.

### Example Structure

Given the following comment structure:
```
hello (2 replies)
├── ok (0 replies)
└── bye (0 replies)
test (0 replies)
one (1 reply)
└── two (2 replies)
    ├── a (0 replies)
    └── b (0 replies)
```

The first comment is written as: `"hello,2,ok,0,bye,0"`
The second comment is written as: `"test,0"`
The third comment is written as: `"one,1,two,2,a,0,b,0"`

The entire comments feed becomes: `"hello,2,ok,0,bye,0,test,0,one,1,two,2,a,0,b,0"`

## Output Format

For a given comments feed in the specified format, print the comments in a different format:

1. **Maximum depth**: Print an integer $d$ — the maximum depth of nesting comments
2. **Level-by-level output**: Print $d$ lines, where the $i$-th line corresponds to nesting level $i$
3. **Comment order**: For the $i$-th row, print comments of nesting level $i$ in the order of their appearance in Polycarp's comments feed, separated by spaces

## Input

- The first line contains a non-empty comments feed in the described format
- The feed consists of uppercase and lowercase letters of the English alphabet, digits, and commas
- Each comment is a non-empty string consisting of uppercase and lowercase English characters
- Each reply count is an integer (consisting of at least one digit) that equals 0 or does not contain leading zeros
- The length of the whole string does not exceed $10^6$
- The given structure of comments is guaranteed to be valid

## Output

Print comments in the format specified above. For each level of nesting, comments should be printed in the order they appear in the input.

## Examples

### Example 1

**Input:**
```
hello,2,ok,0,bye,0,test,0,one,1,two,2,a,0,b,0
```

**Output:**
```
3
hello test one 
ok bye two 
a b 
```

### Example 2

**Input:**
```
a,5,A,0,a,0,A,0,a,0,A,0
```

**Output:**
```
2
a 
A a A a A 
```

### Example 3

**Input:**
```
A,3,B,2,C,0,D,1,E,0,F,1,G,0,H,1,I,1,J,0,K,1,L,0,M,2,N,0,O,1,P,0
```

**Output:**
```
4
A K M 
B F H L N O 
C D G I P 
E J 
```

## Note

The first example is explained in the problem statement above.