# Problem Description

In 2077, the robots that took over the world realized that human music wasn't that great, so they started composing their own.

To write music, the robots have a special musical instrument capable of producing $n$ different sounds. Each sound is characterized by its volume and pitch. A sequence of sounds is called music. Music is considered beautiful if any two consecutive sounds differ either only in volume or only in pitch. Music is considered boring if the volume or pitch of any three consecutive sounds is the same.

You want to compose beautiful, non-boring music that contains each sound produced by your musical instrument exactly once.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

In the first line of each test case, there is a number $n$ ($1 \leq n \leq 2 \cdot 10^5$) — the number of sounds that the musical instrument can produce.

Next, there are $n$ lines, where the $i$-th line contains a pair of numbers $v_i, p_i$ ($1 \leq v_i, p_i \leq 10^9$) — the volume and pitch of the $i$-th sound, respectively. It is guaranteed that among all $n$ sounds, there are no duplicates, meaning for any $i \neq j$, at least one of the conditions $v_i \neq v_j$ or $p_i \neq p_j$ holds.

The sum of $n$ across all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case:
- If it is possible to compose such music, output "YES", and on the next line, output $n$ numbers — the indices of the sounds in the order that forms beautiful non-boring music.
- Otherwise, output "NO".

You may output each letter in any case (lowercase or uppercase). For example, the strings "yEs", "yes", "Yes", and "YES" will be accepted as a positive answer.

## ideas
1. 如果存在，那么就是假设s[1], s[2]是参数0相同，那么s[2], s[3]是参数2相同
2. 把volume和pitch看作节点，把song看作是边，那么就是访问所有边一次的，欧拉图
3. 还有两个限制，两条相邻的边，必须一条是访问volume，一条是访问pitch
4. 第二个限制是，同一个volume（或者pitch）的边，不能连续出现
5. 起点很容易找，就是deg为基数的那个点（如果都是偶数的话，起点可以随便确定吗）
6. 第一个限制，似乎也比较好处理，根据访问边的数量的parity就可以知道应该访问什么样的节点
7. 第二个限制不好处理呐，挑出一个不连续出现的，比较容易。但是如何保证，肯定有答案，却比较头疼
8. 