package main

import "fmt"

type Point149 struct {
	x int
	y int
}

// y = kx + b, k = (y2 - y1)/(x2 - x1)
// y = (y2 - y1)/(x2 - x1)x + b => b = (x1x2 - y2x1)/(x2 - x1)
// y = (y2 - y1)/(x2 - x1)x + (x1x2 - y2x1)/(x2 - x1)
// at the same time:
// ax + by + c = 0
// we could get (y2 - y1)x + (x1 - x2)y + (y1x2 - y2x1)
// where a = y2 - y1, b = x1 - x2, c = (y1x2 - y2x1)
// assume gcd(a, b, c) = t
// (a/t, b/t, c/t) can determine a line as the signature of the line
func maxPoints(points [][]int) int {
	// corner cases
	if points == nil || len(points) == 0 {
		return 0
	}
	if len(points) == 1 {
		return 1
	}

	// create point array
	ps := []*Point149{}
	for _, point := range points {
		ps = append(ps, &Point149{point[0], point[1]})
	}
	max := 0

	// linesMap
	// key = string(slope + "," + intersect)
	// value = set of Point
	linesMap := make(map[string]map[*Point149]bool)
	for _, p1 := range ps {
		for _, p2 := range ps {
			if p1 == p2 {
				continue
			}
			a, b, c := p2.y-p1.y, p1.x-p2.x, p1.y*p2.x-p2.y*p1.x
			factor := gcd(gcd(a, b), c)
			a /= factor
			b /= factor
			c /= factor
			if a < 0 { // make sure a > 0
				b = -b
				c = -c
			} else if a == 0 && b < 0 { // if a = 0, make sure b > 0
				b = -b
				c = -c
			}
			lineKey := fmt.Sprintf("%d,%d,%d", a, b, c)
			// add those two points to corresponding line set
			if set, ok := linesMap[lineKey]; ok {
				if _, ok := set[p1]; !ok {
					set[p1] = true
				}
				if _, ok := set[p2]; !ok {
					set[p2] = true
				}
			} else { // or create the line, and add those two lines into set line set
				linesMap[lineKey] = map[*Point149]bool{
					p1: true,
					p2: true,
				}
			}
			if max < len(linesMap[lineKey]) {
				max = len(linesMap[lineKey])
			}
		}
	}
	return max
}

// euclidean algorithm for gcd
// O(n)
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func main() {
	fmt.Println(maxPoints([][]int{{2, 2}, {0, 0}, {-1, -1}}))
}
