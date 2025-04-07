C*++ language is quite similar to C++. The similarity manifests itself in the fact that the programs written in C*++ sometimes behave unpredictably and lead to absolutely unexpected effects. For example, let's imagine an arithmetic expression in C*++ that looks like this (expression is the main term):

expression ::= summand | expression + summand | expression - summand
summand ::= increment | coefficient*increment
increment ::= a++ | ++a
coefficient ::= 0|1|2|...|1000
For example, "5*a++-3*++a+a++" is a valid expression in C*++.

Thus, we have a sum consisting of several summands divided by signs "+" or "-". Every summand is an expression "a++" or "++a" multiplied by some integer coefficient. If the coefficient is omitted, it is suggested being equal to 1.

The calculation of such sum in C*++ goes the following way. First all the summands are calculated one after another, then they are summed by the usual arithmetic rules. If the summand contains "a++", then during the calculation first the value of the "a" variable is multiplied by the coefficient, then value of "a" is increased by 1. If the summand contains "++a", then the actions on it are performed in the reverse order: first "a" is increased by 1, then — multiplied by the coefficient.

The summands may be calculated in any order, that's why sometimes the result of the calculation is completely unpredictable! Your task is to find its largest possible value.

### ideas
1. 这个expression应该是没有歧义的吧？如果有歧义的话，那就不大对了
2. 只是执行顺序造成结果不确定吧？
3. 如果成立，那么可以先解析成树
4. 然后，如果是 c * a++ 的 或者 c * ++a 的
5. 每次a变化后，对于剩余的 sum要重新排序
6. 越大的，应该放在越后面