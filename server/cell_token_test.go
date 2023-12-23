// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestCellToken(t *testing.T) {
	tests := []struct {
		candidate  string
		occupiedBy CatalogTokenEnum
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83", obj1Token},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc", obj2Token},
	}

	for _, ndx := range tests {
		cellTokenType := &CellTokenType{ItemID: ndx.candidate, OccupiedBy: ndx.occupiedBy}
		if cellTokenType.ItemID != ndx.candidate {
			t.Errorf("TestCellToken failure")
		}
		if cellTokenType.OccupiedBy != ndx.occupiedBy {
			t.Errorf("TestCellToken failure")
		}

		cellTokenType.updateToken("bogus", obj1Token)
		if cellTokenType.ItemID != "bogus" {
			t.Errorf("TestCellToken failure")
		}
		if cellTokenType.OccupiedBy != obj1Token {
			t.Errorf("TestCellToken failure")
		}

		cellTokenType.clearToken()
		if cellTokenType.ItemID != "" {
			t.Errorf("TestCellToken failure")
		}
		if cellTokenType.OccupiedBy != vacantToken {
			t.Errorf("TestCellToken failure")
		}
	}
}
