Ivan wants to write a letter to his friend. The letter is a string s consisting of lowercase Latin letters.

Unfortunately, when Ivan started writing the letter, he realised that it is very long and writing the whole letter may take extremely long time. So he wants to write the compressed version of string s instead of the string itself.

The compressed version of string s is a sequence of strings c1, s1, c2, s2, ..., ck, sk, where ci is the decimal representation of number ai (without any leading zeroes) and si is some string consisting of lowercase Latin letters. If Ivan writes string s1 exactly a1 times, then string s2 exactly a2 times, and so on, the result will be string s.

The length of a compressed version is |c1| + |s1| + |c2| + |s2|... |ck| + |sk|. Among all compressed versions Ivan wants to choose a version such that its length is minimum possible. Help Ivan to determine minimum possible length.

### ideas
1. abcabcabc