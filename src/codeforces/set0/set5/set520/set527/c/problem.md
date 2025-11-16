### Glass Carving

Leonid wants to become a glass carver (the person who creates beautiful artworks by cutting the glass). He already has a rectangular sheet of glass sized \(w \times h\) millimeters, a diamond glass cutter, and lots of enthusiasm.

To practice, he makes vertical and horizontal cuts through the entire sheet. This results in smaller rectangular fragments of glass. Leonid does not move the newly made fragments; each cut divides every fragment it passes through into smaller fragments.

After each cut, determine the area of the largest currently available glass fragment.

### Input
- The first line contains three integers \(w, h, n\)  
  with constraints: \(2 \le w, h \le 200{,}000\), \(1 \le n \le 200{,}000\).
- The next \(n\) lines each describe a cut in one of the forms:
  - `H y`: a horizontal cut at distance \(y\) millimeters (\(1 \le y \le h-1\)) from the lower edge.
  - `V x`: a vertical cut at distance \(x\) millimeters (\(1 \le x \le w-1\)) from the left edge.
- No two cuts are identical.

### Output
After each cut, print on a single line the area of the maximum available glass fragment in mm².

### Examples

Input
```
4 3 4
H 2
V 2
V 3
V 1
```
Output
```
8
4
4
2
```

Input
```
7 6 5
H 4
V 3
V 5
H 2
V 1
```
Output
```
28
16
12
6
4
```