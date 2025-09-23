# Problem B: Calendar

## Description

The girl Taylor has a beautiful calendar for the year y. In the calendar all days are given with their days of week: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.

The calendar is so beautiful that she wants to know what is the next year after y when the calendar will be exactly the same. Help Taylor to find that year.

## Notes

- Leap years have 366 days
- A year is a leap year if:
  - It is divisible by 400, OR
  - It is divisible by 4, but not by 100
- Reference: [Leap year rules](https://en.wikipedia.org/wiki/Leap_year)

## Input

The only line contains integer y (1000 ≤ y < 100'000) — the year of the calendar.

## Output

Print the only integer y' — the next year after y when the calendar will be the same. Note that you should find the first year after y with the same calendar.

## Examples

### Example 1
**Input:**
```
2016
```

**Output:**
```
2044
```

### Example 2
**Input:**
```
2000
```

**Output:**
```
2028
```

### Example 3
**Input:**
```
50501
```

**Output:**
```
50507
```

## Additional Note

Today is Monday, the 13th of June, 2016.