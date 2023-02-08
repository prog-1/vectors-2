package main

type Vector struct {
	x, y int
}

type Line struct {
	a, v Vector
}

func Intersects(l1, l2 *Line) bool {
	v := l1.v
	u := l2.v
	return u.x*v.y != u.y*v.x
}

func IntersectionPoint(l1, l2 *Line) *Vector {
	a, v, b, u := l1.a, l1.v, l2.a, l2.v
	t1 := (u.x*(a.y-b.y) + u.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	// t2 := (a.x + v.x*t1 - b.x) / u.x
	// if 0 <= t1 && t1 <= 1 && 0 <= t2 && t2 <= 1 {
	// 	return &Vector{a.x + v.x*t1, a.y + v.y*t1}, true
	// }
	// return &Vector{}, false
	return &Vector{a.x + v.x*t1, a.y + v.y*t1}
}

func main() {

}
