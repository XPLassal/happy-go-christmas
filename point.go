package main

type Point struct {
	x, y int
}

func (p *Point) isEquale(p1 Point) bool {
	if p.x == p1.x && p.y == p1.y {
		return true
	}
	return false
}
