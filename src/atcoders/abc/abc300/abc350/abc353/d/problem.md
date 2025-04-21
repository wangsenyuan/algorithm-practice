For positive integers x and y, define f(x,y) as follows:

Interpret the decimal representations of x and y as strings and concatenate them in this order to obtain a string z. The value of f(x,y) is the value of z when interpreted as a decimal integer.

For example, f(3,14)=314 and f(100,1)=1001.

You are given a sequence of positive integers A=(A₁,...,Aₙ) of length N. Find the value of the following expression modulo 998244353:

∑(i=1 to N-1) ∑(j=i+1 to N) f(Aᵢ,Aⱼ).