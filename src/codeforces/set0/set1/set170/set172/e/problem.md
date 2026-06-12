# E. BHTML+BCSS

[Problem link](https://codeforces.com/problemset/problem/172/E)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

This problem is about imaginary languages BHTML and BCSS, which slightly resemble
HTML and CSS. Read the statement carefully: the resemblance is slight and the
problem uses very simplified analogs.

You are given a BHTML document that resembles HTML but is much simpler. It is
recorded as a sequence of opening and closing tags. A tag that looks like
`<tagname>` is called an opening tag, and a tag that looks like `</tagname>` is
called a closing tag. There are also self-closing tags written as `<tagname/>`;
in this problem they are fully equivalent to `<tagname></tagname>`. All tagnames
are strings of lowercase Latin letters with length from 1 to 10 characters.
Tagnames of different tags may coincide.

The document tags form a correct bracket sequence. That is, we can obtain an
empty sequence from the given one using these operations:

- remove any self-closing tag `<tagname/>`;
- remove a pair of consecutive opening and closing tags with the same name, i.e.
  remove substring `<tagname></tagname>`.

For example, `<a><b></b></a>` is valid, but `<a></b>`, `<a><b>`, `<a><b></a>`,
and `<a><b></b></c>` are not.

For every opening tag there is exactly one matching closing tag; each such pair
is called an **element**. A self-closing tag is also an element. One element is
**nested** inside another if the tags of the first element lie between the tags
of the second one. An element is not nested in itself.

Each BCSS rule is written as a subsequence of words `x1 x2 ... xn`. The rule
matches an element `t` if both hold:

1. there exists a chain of nested elements with tagnames `x1, x2, ..., xn`
   (each next element is nested in the previous one);
2. this chain ends with element `t`, so the tagname of `t` equals `xn`.

For example, element `b` matches rule `a b` if some element `a` exists in which
`b` is nested. Element `c` matches rule `a b b c` if elements `a`, `b`, `b`, `c`
exist and form a nesting chain `a` → `b` → `b` → `c`.

Given a BHTML document and a set of BCSS rules, determine how many document
elements match each rule.

## Input

The first line contains a BHTML document. Its length is from 4 to `10^6`
characters. The document has correct structure, contains no spaces or other
unnecessary characters, and uses tagnames of length from 1 to 10 consisting of
lowercase Latin letters.

The second line contains an integer `m` (`1 <= m <= 200`) — the number of
queries. Each of the next `m` lines contains one query: a sequence
`x1 x2 ... xn`, where `n` (`1 <= n <= 200`) is the number of words in the query
and each `x_i` is a non-empty lowercase Latin string of length at most 10. Words
are separated by single spaces. A query does not start or end with a space.

## Output

Print `m` lines. The `j`-th line must contain the number of document elements
that match the `j`-th BCSS rule. Print `0` if no element matches.

## Example

### Input

```text
<a><b><b></b></b></a><a><b></b><b><v/></b></a><b></b>
4
a
a b
b a
b b a
```

### Output

```text
2
1
4
0
```

### Input

```text
<b><aa/></b><aa><b/><b/></aa>
5
aa
b
b a
a b
a a a
```

### Output

```text
2
3
2
1
0
```

### ideas
1. 组成一棵树, 
2. 考虑根据规则进行匹配, a -> b -> c 
3. 到达a的时候,应该进入最近的b(但是可能有多条子树,都有b节点)
4. 然后到达b的时候,应该进入最近的c, 然后+c中所有和他同名的tag的数量
5. 这样的复杂性是 m * n * k (m * n是问题的规模, k是节点的数量)
6. 有可能到1e9. 感觉不大行~
7. 根据T来判断,迭代每个T,判断它是否满足规则
8. 如果这样的判断能够在k次内完成,那么就是 m * k ~ 2e8, 似乎可行
9. 

## Solution Review

### Matching Idea

Parse the document into a forest. Each element is one node, and its children are
the elements directly nested inside it.

For a rule such as:

```text
a b c
```

traverse a root-to-node path while maintaining how many words of the rule have
already been matched. When the current node name equals the next required word,
advance the matched prefix by one. Otherwise, keep the same prefix and continue
into the descendants.

Greedily using the earliest possible matching ancestor is valid. Replacing a
chosen rule element with an earlier matching ancestor only gives the remaining
rule words a larger subtree in which to find their matches, so it cannot destroy
an existing nesting chain.

### Subtree Aggregation

For every node `u`, define:

```text
same[u] = number of nodes in u's subtree whose tagname equals tagname(u)
```

If the last word of a rule becomes matched at `u`, then every same-named node in
`u`'s subtree also matches that rule. The chain used to reach `u` can use that
descendant as its final element instead. Therefore, add `same[u]` to the answer
and do not continue processing this completed rule inside that subtree. This
counts every matching element once.

The current implementation computes this value as `dp[u]` while parsing the
tree.

### Flaw in the Current Traversal

The logical counting is correct, but carrying all rules together has a serious
memory risk. At every recursive node, the implementation allocates a new
`[]pair` containing every unfinished rule and passes it to the children.

With up to `200` rules and a document containing roughly `10^5` deeply nested
elements, all these slices remain live along the recursion path. Their backing
arrays alone can require hundreds of megabytes, exceeding the `256 MB` memory
limit. A deeply nested document also creates recursion-depth risk.

### Recommended Traversal

Process each rule separately. For one rule, the traversal state is only one
integer: the length of the matched prefix.

```text
dfs(node, matchedPrefix)
```

At each node:

1. Advance `matchedPrefix` if the node matches the next rule word.
2. If the whole rule is matched, add `same[node]` and skip this subtree.
3. Otherwise, visit the children with the updated prefix.

This preserves the same `O(number_of_nodes * number_of_rules)` time complexity,
but removes the `O(depth * number_of_rules)` live slice allocations. Using an
iterative DFS instead of recursion also avoids stack-depth problems.
