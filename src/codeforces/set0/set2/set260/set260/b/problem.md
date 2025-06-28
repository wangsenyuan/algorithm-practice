# Ancient Prophesy - Apocalypse Date

## Problem Description

A recently found Ancient Prophesy is believed to contain the exact Apocalypse date. The prophesy is a string that only consists of digits and characters "-".

We'll say that some date is mentioned in the Prophesy if there is a substring in the Prophesy that is the date's record in the format `dd-mm-yyyy`. We'll say that the number of the date's occurrences is the number of such substrings in the Prophesy.

### Example
For example, the Prophesy `"0012-10-2012-10-2012"` mentions date `12-10-2012` twice:
- First time as `"0012-10-2012-10-2012"`
- Second time as `"0012-10-2012-10-2012"`

The date of the Apocalypse is such correct date that the number of times it is mentioned in the Prophesy is strictly larger than that of any other correct date.

## Date Format Rules

A date is correct if:
- The year lies in the range from **2013 to 2015**
- The month is from **1 to 12**
- The number of the day is strictly more than zero and doesn't exceed the number of days in the current month

**Important**: A date is written in the format `dd-mm-yyyy`, which means that leading zeroes may be added to the numbers of the months or days if needed.

### Examples
- Date `"1-1-2013"` **isn't** recorded in the format `dd-mm-yyyy`
- Date `"01-01-2013"` **is** recorded in the correct format

**Note**: Any year between 2013 and 2015 is not a leap year.

## Input

The first line contains the Prophesy: a non-empty string that only consists of digits and characters "-". The length of the Prophesy doesn't exceed $10^5$ characters.

## Output

In a single line print the date of the Apocalypse. It is guaranteed that such date exists and is unique.