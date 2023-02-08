package main

import "fmt"

type v struct { //vector
	x, y int
}

type l struct { //line
	a, b v
}

func main() {
	l1 := l{a: v{x: 4, y: 3}, b: v{x: -3, y: 3}}
	l2 := l{a: v{x: -3, y: 2}, b: v{x: 2, y: 2}}

	// l1 := l{a: v{x: 3, y: 1}, v: v{x: -4, y: 4}}
	// l2 := l{a: v{x: -3, y: 1}, v: v{x: 6, y: 3}}

	// l1 := l{a: v{x: -5, y: 2}, v: v{x: 4, y: 7}}
	// l2 := l{a: v{x: 7, y: 2}, v: v{x: -7, y: 4}}

	// l1 := l{a: v{x: 1, y: 1}, v: v{x: 1, y: 2}}
	// l2 := l{a: v{x: 1, y: 3}, v: v{x: 1, y: 1}}

	fmt.Println(intersectionPoint(l1, l2))
}

func intersectionPoint(l1, l2 l) (p v) {
	a, v, b, u := l1.a, l1.b, l2.a, l2.b //for convinience

	//t1 := (u.x*a.y - u.x*b.y) - (u.y*a.x+u.y*b.x)/(u.y*v.x-u.x*v.y)//not right - wroong
	t1 := (u.x*(a.y-b.y) + u.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	p.x = a.x + v.x*t1
	p.y = a.y + v.y*t1
	return p
}
