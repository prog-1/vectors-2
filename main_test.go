package main

import "testing"

func TestIncludesPoint(t *testing.T) {
	for _, tc := range []struct {
		l1, l2    line
		want      vector
		intersect bool
	}{
		{line{vector{-5, -2}, vector{4, 7}}, line{vector{7, 2}, vector{-7, 4}}, vector{0, 3}, true},
	} {
		got, intersect := IntersectionPoint(&tc.l1, &tc.l2)
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
