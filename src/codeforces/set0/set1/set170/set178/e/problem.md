# ABBYY Cup - Smart Beaver Problem

## Problem Description

Offering the ABBYY Cup participants a problem written by the Smart Beaver is becoming a tradition. He proposed the following problem.

You are given a monochrome image, that is, an image that is composed of two colors (black and white). The image is given in raster form, that is, as a matrix of pixels' colors, and the matrix's size coincides with the size of the image.

The white color on the given image corresponds to the background. Also, the image contains several black geometric shapes. It is known that the image can contain only two types of shapes: **squares** and **circles**. Your task is to count the number of circles and the number of squares which the given image contains.

## Important Notes

- The squares on the image can be rotated arbitrarily
- The image can possibly contain some noise arranged as follows: each pixel of the original image can change its color to the opposite with the probability of 20%

## Examples

1. **No noise, squares parallel to coordinate axes**: Two circles and three squares
2. **No noise, squares rotated arbitrarily**: Two circles and three squares  
3. **With noise, squares rotated arbitrarily**: One circle and three squares

## Input

- **First line**: A single integer $n$ ($1000 \leq n \leq 2000$), which is the length and width of the original image
- **Next $n$ lines**: Each line contains exactly $n$ integers $a_{ij}$ ($0 \leq a_{ij} \leq 1$), separated by spaces
  - $a_{ij} = 0$ corresponds to a white pixel
  - $a_{ij} = 1$ corresponds to a black pixel

## Constraints

- The lengths of the sides of the squares and the diameters of the circles in the image are at least **15 pixels**
- The distance between any two figures is at least **10 pixels**
- A human can easily calculate the number of circles and squares in the original image
- The total number of figures in the image doesn't exceed **50**

## Scoring

### 20 points
- No noise
- Sides of squares are parallel to the coordinate axes

### 50 points  
- No noise
- Squares are rotated arbitrarily

### 100 points
- Noise present
- Squares are rotated arbitrarily

## Output

Print exactly two integers, separated by a single space — the number of circles and the number of squares in the given image, correspondingly.


### ideas
1. 