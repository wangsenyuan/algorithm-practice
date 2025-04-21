Serval has just found a Kaitenzushi buffet restaurant. Kaitenzushi means that there is a conveyor belt in the restaurant, delivering plates of sushi in front of the customer, Serval.

In this restaurant, each plate contains exactly 𝑘
 pieces of sushi and the 𝑖
-th plate has a deliciousness 𝑑𝑖
. Serval will have a meal in this restaurant for 𝑛
 minutes, and within the 𝑛
 minutes, he must eat up all the pieces of sushi he took from the belt.

Denote the counter for uneaten taken pieces of sushi as 𝑟
. Initially, 𝑟=0
. In the 𝑖
-th minute (1≤𝑖≤𝑛
), only the 𝑖
-th plate of sushi will be delivered in front of Serval, and he can do one of the following:

Take the 𝑖
-th plate of sushi (whose deliciousness is 𝑑𝑖
) from the belt, and 𝑟
 will be increased by 𝑘
;
Eat one uneaten piece of sushi that he took from the belt before, and 𝑟
 will be decreased by 1
. Note that you can do this only if 𝑟>0
;
Or, do nothing, and 𝑟
 will remain unchanged.
Note that after the 𝑛
 minutes, the value of 𝑟
 must be 0
.

Serval wants to maximize the sum of the deliciousnesses of all the plates he took. Help him find it out!

### ideas
1. 如果take了一个plate，那么必须把它吃完；
2. 感觉要从后往前处理。假设目前已经积累了r个寿司，遇到了一个更美味的i，如果r + k > n - i (时间不够吃完)
3. 那么就交换最不美味的部分
4. 似乎可以work