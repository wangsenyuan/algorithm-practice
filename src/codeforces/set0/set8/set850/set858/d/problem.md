# Problem D: Phone Numbers

## Problem Description

There are n phone numbers in Polycarp's contacts on his phone. Each number is a 9-digit integer, starting with a digit different from 0. All the numbers are distinct.

There is the latest version of Berdroid OS installed on Polycarp's phone. If some number is entered, it shows up all the numbers in the contacts for which there is a substring equal to the entered sequence of digits.

### Example

If there are three phone numbers in Polycarp's contacts: `123456789`, `100000000` and `100123456`, then:

- If he enters `00` → two numbers will show up: `100000000` and `100123456`
- If he enters `123` → two numbers will show up: `123456789` and `100123456`
- If he enters `01` → there will be only one number: `100123456`

## Task

For each of the phone numbers in Polycarp's contacts, find the minimum length sequence of digits such that if Polycarp enters this sequence, Berdroid shows only this phone number.

## Input

- The first line contains a single integer **n** (1 ≤ n ≤ 70000) — the total number of phone contacts in Polycarp's contacts.
- The phone numbers follow, one in each line. Each number is a positive 9-digit integer starting with a digit from 1 to 9. All the numbers are distinct.

## Output

Print exactly **n** lines: the i-th of them should contain the shortest non-empty sequence of digits, such that if Polycarp enters it, the Berdroid OS shows up only the i-th number from the contacts. If there are several such sequences, print any of them.

## ideas
1. 9位长的，共有9 * 9 / 2 = 40个子串