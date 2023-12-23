// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestCellArray(t *testing.T) {
	cellArray := initializeCellArray()

	tests := []struct {
		candidate  string
		occupiedBy CatalogTokenEnum
		yy, xx     int
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83", obj1Token, 3, 3},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc", obj2Token, 5, 5},
	}

	for _, ndx := range tests {
		location1 := &LocationType{YY: ndx.yy, XX: ndx.xx}
		cellArray.updateCell(ndx.candidate, location1, ndx.occupiedBy)

		cellTokenType1 := cellArray[ndx.yy][ndx.xx]
		if cellTokenType1.ItemID != ndx.candidate {
			t.Errorf("TestCellToken failure")
		}
		if cellTokenType1.OccupiedBy != ndx.occupiedBy {
			t.Errorf("TestCellToken failure")
		}

		location2 := &LocationType{YY: ndx.yy + 2, XX: ndx.xx + 3}
		cellArray.moveCell(location1, location2)

		if !cellTokenType1.isVacant() {
			t.Errorf("TestCellToken failure")
		}

		cellTokenType2 := cellArray[ndx.yy+2][ndx.xx+3]
		if cellTokenType2.ItemID != ndx.candidate {
			t.Errorf("TestCellToken failure")
		}
		if cellTokenType2.OccupiedBy != ndx.occupiedBy {
			t.Errorf("TestCellToken failure")
		}
	}
}
