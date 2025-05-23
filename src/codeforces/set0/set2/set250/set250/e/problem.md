# Joe's House Rampage

Joe has been hurt on the Internet. Now he is storming around the house, destroying everything in his path.

Joe's house has **n** floors, each floor is a segment of **m** cells. Each cell either contains:
- nothing (it is an *empty* cell),
- a *brick*, or
- a *concrete wall* (always one of these three).

It is believed that each floor is surrounded by a concrete wall on the left and on the right.

Joe starts on the **n-th floor** in the **first cell** (counting from left to right). At each moment, Joe has a direction of gaze: **right** or **left** (initially, Joe looks to the right).

Every second, Joe performs one of the following actions:

1. **If the cell directly under Joe is empty:**  
   Joe falls down to this cell (gaze direction is preserved).
2. **Otherwise, consider the next cell in the current gaze direction:**
   - If the cell is empty, Joe moves into it (gaze direction is preserved).
   - If the cell has bricks, Joe breaks them with his forehead (the cell becomes empty), and changes the direction of his gaze to the opposite.
   - If the cell has a concrete wall, Joe just changes the direction of his gaze to the opposite (concrete can withstand any number of forehead hits).

Joe calms down as soon as he reaches any cell of the **first floor**.

*The figure below shows an example of Joe's movements around the house.*