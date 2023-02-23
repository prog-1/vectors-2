package main

import "testing"

type Want struct {
	p         vec
	intersect bool
}

func TestIncludesPoint(t *testing.T) {
	for _, tc := range []struct {
		a, a1, b, b1 vec
		want         Want
	}{
		{vec{-5, -2}, vec{4, 7}, vec{7, 2}, vec{-7, 4}, Want{vec{0, 3}, true}},
		{vec{2, 2}, vec{4, 4}, vec{2, 4}, vec{4, 6}, Want{vec{}, false}},
		{vec{-50 + halfScreenWidth, -20 + halfScreenHeight}, vec{40 + halfScreenWidth, 70 + halfScreenHeight}, vec{70 + halfScreenWidth, 20 + halfScreenHeight}, vec{-70 + halfScreenHeight, 40 + halfScreenWidth}, Want{vec{335, 285}, true}},
	} {
		got, intersect := IntersectionPoint(tc.a, sub(tc.a1, tc.a), tc.b, sub(tc.b1, tc.b))
		if tc.want.intersect != intersect {
			switch intersect {
			case true:
				t.Errorf("Lines intersect when they must not")
			case false:
				t.Errorf("Lines are parallel when they must intersect")
			}
		}
		if got.x != tc.want.p.x || got.y != tc.want.p.y {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}

	}
}
