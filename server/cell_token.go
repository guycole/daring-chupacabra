// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type CellTokenType struct {
	ItemID     string // same as catalog item ID
	OccupiedBy CatalogTokenEnum
}

func (ctt *CellTokenType) clearToken() {
	ctt.ItemID = ""
	ctt.OccupiedBy = vacantToken
}

func (ctt *CellTokenType) isVacant() bool {
	if ctt.OccupiedBy == vacantToken {
		return true
	}

	return false
}

func (ctt *CellTokenType) updateToken(id string, occupiedBy CatalogTokenEnum) {
	ctt.ItemID = id
	ctt.OccupiedBy = occupiedBy
}
