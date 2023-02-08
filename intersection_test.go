package main

import "testing"

func TestIntersectionPoint(t *testing.T) {
	for _, tc := range []struct {
		name   string
		l1, l2 l
		want   v
	}{
		{"1", l{a: v{x: 4, y: 3}, b: v{x: -3, y: 3}}, l{a: v{x: -3, y: 2}, b: v{x: 2, y: 2}}, v{x: 1, y: 6}},
		{"2", l{a: v{x: 1, y: 1}, b: v{x: 1, y: 2}}, l{a: v{x: 1, y: 3}, b: v{x: 1, y: 1}}, v{x: 3, y: 5}},
		{"3", l{a: v{x: 3, y: 1}, b: v{x: -4, y: 4}}, l{a: v{x: -3, y: 1}, b: v{x: 6, y: 3}}, v{x: 3, y: 1}},  //!
		{"4", l{a: v{x: -5, y: 2}, b: v{x: 4, y: 7}}, l{a: v{x: 7, y: 2}, b: v{x: -7, y: 4}}, v{x: -5, y: 2}}, //!
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := intersectionPoint(tc.l1, tc.l2)
			if got.x != tc.want.x || got.y != tc.want.y {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
