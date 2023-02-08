package main

import "testing"

func TestIntersectionPoint(t *testing.T) {
	for _, tc := range []struct {
		name   string
		l1, l2 line
		want   vector
	}{
		{"1", line{a: vector{x: 4, y: 3}, b: vector{x: -3, y: 3}}, line{a: vector{x: -3, y: 2}, b: vector{x: 2, y: 2}}, vector{x: 1, y: 6}},
		{"2", line{a: vector{x: 1, y: 1}, b: vector{x: 1, y: 2}}, line{a: vector{x: 1, y: 3}, b: vector{x: 1, y: 1}}, vector{x: 3, y: 5}},
		{"3", line{a: vector{x: 3, y: 1}, b: vector{x: -4, y: 4}}, line{a: vector{x: -3, y: 1}, b: vector{x: 6, y: 3}}, vector{x: 3, y: 1}},  //!
		{"4", line{a: vector{x: -5, y: 2}, b: vector{x: 4, y: 7}}, line{a: vector{x: 7, y: 2}, b: vector{x: -7, y: 4}}, vector{x: -5, y: 2}}, //!
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := intersectionPoint(tc.l1, tc.l2)
			if got.x != tc.want.x || got.y != tc.want.y {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
