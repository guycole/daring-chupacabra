// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"math"
	"math/rand"
	"strconv"
)

// row column origin = 0,0 lower left corner of map
type LocationType struct {
	XX int // column
	YY int // row
}

func sameLocation(loc1, loc2 *LocationType) bool {
	if loc1.XX == loc2.XX && loc1.YY == loc2.YY {
		return true
	}

	return false
}

func newLocation(y, x int) *LocationType {
	result := LocationType{YY: y, XX: x}
	return &result
}

func randomLocation(limitY, limitX int) *LocationType {
	xx := rand.Intn(limitX)
	yy := rand.Intn(limitY)

	return newLocation(yy, xx)
}

func stringLocation(y, x string) *LocationType {
	yy, err1 := strconv.Atoi(y)
	xx, err2 := strconv.Atoi(x)

	if err1 == nil && err2 == nil {
		return newLocation(yy, xx)
	}

	return nil
}

func (origin *LocationType) getDistance(destination *LocationType) int {
	dx := float64(destination.XX - origin.XX)
	dy := float64(destination.YY - origin.YY)
	result := math.Hypot(dx, dy)
	return int(result)
}

func (origin *LocationType) legalLocation(limitY, limitX int) bool {
	if origin.XX < 0 || origin.XX >= limitX {
		return false
	}

	if origin.YY < 0 || origin.YY >= limitY {
		return false
	}

	return true
}

/*
   1 2 3 (indices and relative locations)
   4 0 5 (0 is origin)
   6 7 8
*/

// return index of target location relative to origin location
func (origin *LocationType) adjacencyTest(target *LocationType) int {
	var x, y int

	for ndx := 0; ndx < 9; ndx++ {
		switch ndx {
		case 0:
			x = origin.XX
			y = origin.YY
		case 1:
			x = origin.XX - 1
			y = origin.YY + 1
		case 2:
			x = origin.XX
			y = origin.YY + 1
		case 3:
			x = origin.XX + 1
			y = origin.YY + 1
		case 4:
			x = origin.XX - 1
			y = origin.YY
		case 5:
			x = origin.XX + 1
			y = origin.YY
		case 6:
			x = origin.XX - 1
			y = origin.YY - 1
		case 7:
			x = origin.XX
			y = origin.YY - 1
		case 8:
			x = origin.XX + 1
			y = origin.YY - 1
		}

		if x == target.XX && y == target.YY {
			return ndx
		}
	}

	return -1
}
