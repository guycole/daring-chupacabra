// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestFreshLocation(t *testing.T) {
	tests := []struct {
		yy, xx int
	}{
		{3, 3},
		{3, -3},
		{-3, 3},
		{-3, -3},
	}

	for _, ndx := range tests {
		result := newLocation(ndx.yy, ndx.xx)
		if result.XX != ndx.xx {
			t.Errorf("getFreshLocation(%d, %d) failure", ndx.yy, ndx.xx)
		}
		if result.YY != ndx.yy {
			t.Errorf("getFreshLocation(%d, %d) failure", ndx.yy, ndx.xx)
		}
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		y1, x1, y2, x2, target int
	}{
		{0, 0, 3, 3, 4},
		{0, 0, 3, -3, 4},
		{0, 0, -3, -3, 4},
		{0, 0, -3, 3, 4},
	}

	for _, ndx := range tests {
		loc1 := newLocation(ndx.y1, ndx.x1)
		loc2 := newLocation(ndx.y2, ndx.x2)
		result := loc1.getDistance(loc2)
		if result != ndx.target {
			t.Errorf("getDistance(%d, %d) failure expect %d got %d", ndx.y1, ndx.x1, ndx.target, result)
		}
	}
}

func TestAdjacency(t *testing.T) {
	origin := newLocation(4, 5)

	tests := []struct {
		y, x, target int
	}{
		{5, 4, 1},
		{5, 5, 2},
		{5, 6, 3},
		{4, 4, 4},
		{4, 5, 0},
		{4, 6, 5},
		{3, 4, 6},
		{3, 5, 7},
		{3, 6, 8},
		{7, 7, -1},
	}

	for _, ndx := range tests {
		newLoc := newLocation(ndx.y, ndx.x)
		result := origin.adjacencyTest(newLoc)
		if result != ndx.target {
			t.Errorf("adjancency(%d, %d) failure expect %d got %d", ndx.y, ndx.x, ndx.target, result)
		}
	}
}

func TestLegalLocation(t *testing.T) {
	origin := newLocation(4, 5)
	if origin.legalLocation(10, 10) == false {
		t.Errorf("legalLocation failure")
	}

	origin = newLocation(-4, -5)
	if origin.legalLocation(10, 10) == true {
		t.Errorf("legalLocation failure")
	}

	origin = newLocation(14, 15)
	if origin.legalLocation(10, 10) == true {
		t.Errorf("legalLocation failure")
	}
}
