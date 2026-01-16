package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	t := readNum(reader)

	for i := 0; i < t; i++ {
		seg1 := readSegment(reader)
		seg2 := readSegment(reader)
		seg3 := readSegment(reader)

		res := solve(seg1, seg2, seg3)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type Point struct {
	x, y int
}

type Segment struct {
	p1, p2 Point
}

func readSegment(reader *bufio.Reader) Segment {
	coords := readNNums(reader, 4)
	return Segment{
		p1: Point{x: coords[0], y: coords[1]},
		p2: Point{x: coords[2], y: coords[3]},
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(seg1, seg2, seg3 Segment) bool {
	// Try all 6 permutations of which segments are first, second, and third
	segs := []Segment{seg1, seg2, seg3}
	permutations := [][]int{
		{0, 1, 2}, {0, 2, 1},
		{1, 0, 2}, {1, 2, 0},
		{2, 0, 1}, {2, 1, 0},
	}

	for _, perm := range permutations {
		first := segs[perm[0]]
		second := segs[perm[1]]
		third := segs[perm[2]]

		if checkLetterA(first, second, third) {
			return true
		}
	}

	return false
}

func checkLetterA(first, second, third Segment) bool {
	// Condition 1: Two segments have common endpoint
	commonPoint, hasCommon := findCommonEndpoint(first, second)
	if !hasCommon {
		return false
	}

	// Condition 2: The third segment connects two points on the different segments
	// Find the endpoints of third segment
	p1 := third.p1
	p2 := third.p2

	// Check if p1 and p2 lie on first and second segments (or vice versa)
	// and they are on different segments
	var p1OnFirst, p1OnSecond, p2OnFirst, p2OnSecond bool

	// Check which segments p1 and p2 belong to
	if pointOnSegment(first, p1) && !pointsEqual(p1, commonPoint) {
		p1OnFirst = true
	}
	if pointOnSegment(second, p1) && !pointsEqual(p1, commonPoint) {
		p1OnSecond = true
	}
	if pointOnSegment(first, p2) && !pointsEqual(p2, commonPoint) {
		p2OnFirst = true
	}
	if pointOnSegment(second, p2) && !pointsEqual(p2, commonPoint) {
		p2OnSecond = true
	}

	// Third segment should connect points on different segments
	if !((p1OnFirst && p2OnSecond) || (p1OnSecond && p2OnFirst)) {
		return false
	}

	// Condition 3: The angle between first and second segments is > 0 and <= 90 degrees
	angle := calculateAngle(first, second, commonPoint)
	if angle <= 0 || angle > 90 {
		return false
	}

	// Condition 4: The third segment divides each of the first two segments in proportion >= 1/4
	// Determine which point is on which segment
	var pointOnFirst, pointOnSecond Point
	if p1OnFirst {
		pointOnFirst = p1
		pointOnSecond = p2
	} else {
		pointOnFirst = p2
		pointOnSecond = p1
	}

	// Check proportion for first segment
	if !checkProportion(first, pointOnFirst, commonPoint) {
		return false
	}

	// Check proportion for second segment
	if !checkProportion(second, pointOnSecond, commonPoint) {
		return false
	}

	return true
}

func findCommonEndpoint(seg1, seg2 Segment) (Point, bool) {
	if pointsEqual(seg1.p1, seg2.p1) || pointsEqual(seg1.p1, seg2.p2) {
		return seg1.p1, true
	}
	if pointsEqual(seg1.p2, seg2.p1) || pointsEqual(seg1.p2, seg2.p2) {
		return seg1.p2, true
	}
	return Point{}, false
}

func pointsEqual(p1, p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func pointOnSegment(seg Segment, p Point) bool {
	// Check if point p lies on segment seg
	// First check if point equals one of the endpoints
	if pointsEqual(seg.p1, p) || pointsEqual(seg.p2, p) {
		return true
	}

	// Using cross product to check collinearity
	dx1 := seg.p2.x - seg.p1.x
	dy1 := seg.p2.y - seg.p1.y
	dx2 := p.x - seg.p1.x
	dy2 := p.y - seg.p1.y

	// Cross product should be 0 for collinearity
	cross := dx1*dy2 - dy1*dx2
	if cross != 0 {
		return false
	}

	// Check if point is within segment bounds using dot product
	dot := dx1*dx2 + dy1*dy2
	lenSq := dx1*dx1 + dy1*dy1

	if lenSq == 0 {
		return false
	}

	// Parameter t should be between 0 and 1 (inclusive) for point to be on segment
	t := float64(dot) / float64(lenSq)
	return t > 0 && t < 1
}

func calculateAngle(seg1, seg2 Segment, commonPoint Point) float64 {
	// Calculate vectors from common point
	var v1, v2 Point

	if pointsEqual(seg1.p1, commonPoint) {
		v1 = Point{x: seg1.p2.x - seg1.p1.x, y: seg1.p2.y - seg1.p1.y}
	} else {
		v1 = Point{x: seg1.p1.x - seg1.p2.x, y: seg1.p1.y - seg1.p2.y}
	}

	if pointsEqual(seg2.p1, commonPoint) {
		v2 = Point{x: seg2.p2.x - seg2.p1.x, y: seg2.p2.y - seg2.p1.y}
	} else {
		v2 = Point{x: seg2.p1.x - seg2.p2.x, y: seg2.p1.y - seg2.p2.y}
	}

	// Calculate angle using dot product
	dot := v1.x*v2.x + v1.y*v2.y
	len1 := math.Sqrt(float64(v1.x*v1.x + v1.y*v1.y))
	len2 := math.Sqrt(float64(v2.x*v2.x + v2.y*v2.y))

	if len1 == 0 || len2 == 0 {
		return 0
	}

	cosAngle := float64(dot) / (len1 * len2)
	// Clamp to [-1, 1] to avoid numerical errors
	if cosAngle > 1 {
		cosAngle = 1
	}
	if cosAngle < -1 {
		cosAngle = -1
	}

	angle := math.Acos(cosAngle) * 180.0 / math.Pi
	return angle
}

func checkProportion(seg Segment, pointOnSeg Point, endpoint Point) bool {
	// Check if the ratio of the shortest part to the longest part is >= 1/4
	// The segment is divided by pointOnSeg, with endpoint being one end

	// Find the other endpoint
	var otherEnd Point
	if pointsEqual(seg.p1, endpoint) {
		otherEnd = seg.p2
	} else {
		otherEnd = seg.p1
	}

	// Calculate actual distances (not squared)
	dist1 := distanceFloat(endpoint, pointOnSeg)
	dist2 := distanceFloat(pointOnSeg, otherEnd)

	if dist1 == 0 || dist2 == 0 {
		return false
	}

	// Find shorter and longer parts
	shorter := dist1
	longer := dist2
	if dist1 > dist2 {
		shorter = dist2
		longer = dist1
	}

	// Check if ratio >= 1/4
	// shorter / longer >= 1/4
	// 4 * shorter >= longer
	return 4*shorter >= longer
}

func distanceFloat(p1, p2 Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	return math.Sqrt(dx*dx + dy*dy)
}
