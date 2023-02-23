package main

import "testing"

func TestIncludesPoint(t *testing.T) {
	for _, tc := range []struct {
		a, v, b, u vector
		want       vector
		intersect  bool
	}{} {
		got, intersect := IntersectionPoint(tc.a, tc.v, tc.b, tc.u)
		if tc.intersect != intersect {
			switch intersect {
			case true:
				t.Errorf("Lines intersect when they must not")
			case false:
				t.Errorf("Lines are parallel when they must intersect")
			}
		}
		if got.x != tc.want.x || got.y != tc.want.y {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}

	}
}
