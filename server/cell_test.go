// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestCell(t *testing.T) {
	tests := []struct {
		candidate string
		tokenType CatalogTokenEnum
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83", obj1Token},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc", obj2Token},
	}

	for _, ndx := range tests {
		// create cell and test
		cell := &CellType{ItemID: ndx.candidate, TokenType: ndx.tokenType}
		if cell.ItemID != ndx.candidate || cell.TokenType != ndx.tokenType {
			t.Errorf("TestCell failure")
		}

		// bad token id
		cell.updateToken("bogus", obj1Token)
		if cell.ItemID != "bogus" {
			t.Errorf("TestCell failure")
		}

		// reset test
		cell.clearToken()
		if cell.ItemID != "" {
			t.Errorf("TestCell failure")
		}
		if cell.TokenType != vacantToken {
			t.Errorf("TestCell failure")
		}
	}
}
