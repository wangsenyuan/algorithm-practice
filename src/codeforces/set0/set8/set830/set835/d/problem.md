Palindromic characteristics of string s with length |s| is a sequence of |s| integers, where k-th number is the total number of non-empty substrings of s which are k-palindromes.

A string is 1-palindrome if and only if it reads the same backward as forward.

A string is k-palindrome (k > 1) if and only if:

Its left half equals to its right half.
Its left and right halfs are non-empty (k - 1)-palindromes.
The left half of string t is its prefix of length ⌊|t| / 2⌋, and right half — the suffix of the same length. ⌊|t| / 2⌋ denotes the length of string t divided by 2, rounded down.

Note that each substring is counted as many times as it appears in the string. For example, in the string "aaa" the substring "a" appears 3 times.