# Tree Query Problem - Java Implementation

This is a Java implementation of a complex tree query problem that uses multiple advanced algorithms:

## Algorithm Overview

The solution combines several sophisticated algorithms:

1. **Euler Tour**: Flattens the tree into a linear array for range queries
2. **LCA (Lowest Common Ancestor)**: Uses binary lifting for efficient LCA computation
3. **Mo's Algorithm**: Offline range query algorithm for efficient processing
4. **Coordinate Compression**: Reduces the range of frequency values
5. **Tree Representation**: Adjacency list with edge-based graph structure

## Key Components

### Data Structures
- **Graph**: Adjacency list representation using arrays for efficiency
- **Query**: Encapsulates query information with range boundaries
- **FastReader**: Fast I/O for competitive programming

### Main Algorithm Steps

1. **Coordinate Compression**: Compress frequency values to reduce memory usage
2. **Tree Construction**: Build adjacency list from edge list
3. **Euler Tour + LCA Setup**: 
   - Perform DFS to create Euler tour
   - Build binary lifting table for LCA queries
4. **Query Processing**:
   - Convert tree path queries to range queries on Euler tour
   - Sort queries using Mo's algorithm ordering
   - Process queries with sliding window technique
5. **Result Calculation**: Handle special cases for LCA nodes

## Time Complexity
- **Preprocessing**: O(n log n) for Euler tour and LCA setup
- **Query Processing**: O((n + q) * sqrt(n)) where q is number of queries
- **Overall**: O(n log n + q * sqrt(n))

## Space Complexity
- O(n log n) for binary lifting table
- O(n) for Euler tour and other arrays
- O(q) for query storage

## Usage

```bash
# Compile
javac Solution.java

# Run (reads from stdin, outputs to stdout)
java Solution < input.txt > output.txt
```

## Input Format
```
n
a1 a2 ... an
f1 f2 ... fn
u1 v1
u2 v2
...
un-1 vn-1
m
a1 b1
a2 b2
...
am bm
```

Where:
- `n`: number of nodes
- `a[i]`: gender of node i (0 for boy, 1 for girl)
- `f[i]`: frequency value of node i
- `u[i] v[i]`: tree edges
- `m`: number of queries
- `a[i] b[i]`: query endpoints

## Output Format
One integer per line representing the answer for each query. 