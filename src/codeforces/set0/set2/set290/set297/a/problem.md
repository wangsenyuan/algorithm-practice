# Problem

You are fishing with polar bears Alice and Bob. While waiting for the fish to bite, the polar bears get bored. They come up with a game. First Alice and Bob each writes a 01-string (strings that only contain character "0" and "1") a and b. Then you try to turn a into b using two types of operations:

1. Write parity(a) to the end of a. For example, .
2. Remove the first character of a. For example, . You cannot perform this operation if a is empty.

You can use as many operations as you want. The problem is, is it possible to turn a into b?

The parity of a 01-string is 1 if there is an odd number of "1"s in the string, and 0 otherwise.

## Input

The first line contains the string a and the second line contains the string b ($1 \leq |a|, |b| \leq 1000$). Both strings contain only the characters "0" and "1". Here $|x|$ denotes the length of the string x.

## Output

Print "YES" (without quotes) if it is possible to turn a into b, and "NO" (without quotes) otherwise.

## Examples

### Example 1

**Input:**

```text
01011
0110
```

**Output:**

```text
YES
```

### Example 2

**Input:**

```text
0011
1110
```

**Output:**

```text
NO
```

## Note

In the first sample, the steps are as follows: $01011 \to 1011 \to 011 \to 0110$


### ideas
1. 操作1以后，a的parity是偶数
2. 操作2以后，a的parity有可能变化
3. 如果a中只有0，那么除非b = 0, 否则 false
4. 考虑a = 1, 那么如果再加一个1，是可以的，但是不能增加两个1
5. 因为增加一个1以后，它就变成了偶数个，如果要增加一个，必须删掉一个，然后增加一个
6. 所以1的个数，只会多一个出来；且如果目前是偶数，那1不会更多
7. 如果是奇数，可以多一个出来, 
8. 考虑偶数个的情况，头部删掉一个1，可以在尾部增加一个1（所以可以认为是循环了一下）
9. 且可以在中间添加0（偶数的情况下，始终可以添加0）
10. 而且也可以把多余的0给删掉
11. 所以如果b的1个数是偶数，且 b <= a, 那么肯定可以
12. 如果b的个数是奇数，且b <= a也是可以的，先得到和b相同形状（通过偶数转化）然后删掉头部就可以了