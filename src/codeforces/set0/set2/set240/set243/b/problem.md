One day Petya got a birthday present from his mom: a book called "The Legends and Myths of Graph Theory". From this book Petya learned about a hydra graph.

A non-oriented graph is a hydra, if it has a structure, shown on the figure below. Namely, there are two nodes u and v connected by an edge, they are the hydra's chest and stomach, correspondingly. The chest is connected with h nodes, which are the hydra's heads. The stomach is connected with t nodes, which are the hydra's tails. Note that the hydra is a tree, consisting of h + t + 2 nodes.


Also, Petya's got a non-directed graph G, consisting of n nodes and m edges. Petya got this graph as a last year birthday present from his mom. Graph G contains no self-loops or multiple edges.

Now Petya wants to find a hydra in graph G. Or else, to make sure that the graph doesn't have a hydra.

### ideas
1. 如果存在，那么deg[u] >= h + 1
2. deg[v] >= t + 1
3. 所以，就应该在这样的pair中去查找
4. 假设x和u相连，y和v相连，那么 x != y
5. 