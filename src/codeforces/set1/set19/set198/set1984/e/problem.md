### problem
Two hungry red pandas, Oscar and Lura, have a tree 𝑇
 with 𝑛
 nodes. They are willing to perform the following shuffle procedure on the whole tree 𝑇
 exactly once. With this shuffle procedure, they will create a new tree out of the nodes of the old tree.

Choose any node 𝑉
 from the original tree 𝑇
. Create a new tree 𝑇2
, with 𝑉
 as the root.
Remove 𝑉
 from 𝑇
, such that the original tree is split into one or more subtrees (or zero subtrees, if 𝑉
 is the only node in 𝑇
).
Shuffle each subtree with the same procedure (again choosing any node as the root), then connect all shuffled subtrees' roots back to 𝑉
 to finish constructing 𝑇2
.
After this, Oscar and Lura are left with a new tree 𝑇2
. They can only eat leaves and are very hungry, so please find the maximum number of leaves over all trees that can be created in exactly one shuffle.

Note that leaves are all nodes with degree 1
. Thus, the root may be considered as a leaf if it has only one child.

### solution
Excluding the root, the maximum number of leaves we have is the maximum independent set (MIS) of the rest of the tree. Why? First of all, after rooting tree 𝑇2
, no two adjacent nodes can both be leaves. This is because, no matter which one you choose to add to 𝑇2
 first, it must be the ancestor of the other one. Thus, the first chosen node will not be a leaf. Furthermore, any chosen MIS of nodes can all become leaves. This is because we can essentially choose to add all of the nodes not in the MIS first to the tree, leaving all of the remaining nodes to be childless nodes, making them all leaves.

To find the answer, we want to find the maximum MIS of all subtrees that are missing one leaf. The answer will be one greater than this maximum MIS due to the root of 𝑇2
 being also a leaf.

There are multiple dynamic programming approaches to compute this efficiently, including rerooting and performing dynamic programming on tree edges.