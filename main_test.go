package main

import "testing"

func TestIncludesPoint(t *testing.T) {
	for _, tc := range []struct {
		l1, l2 Line
		want   Vector
	}{
		// {Line{Vector{2, 4}, Vector{3, 1}}, Line{Vector{3, 1}, Vector{2, 2}}, Vector{8, 6}},
		// {Line{Vector{1, 1}, Vector{1, 2}}, Line{Vector{1, 3}, Vector{1, 1}}, Vector{3, 5}},
		{Line{Vector{5, 3}, Vector{6, 8}}, Line{Vector{3, 5}, Vector{10, 4}}, Vector{8, 7}},
	} {
		if got := IntersectionPoint(&tc.l1, &tc.l2); got.x != tc.want.x || got.y != tc.want.y {
			t.Errorf("got = %v, want = %v", got, tc.want)
		} else {
			t.Errorf("Not right, wrong)")
		}

	}
}
