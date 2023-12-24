// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type CellType struct {
	ItemID     string // same as catalog item ID
	OccupiedBy CatalogTokenEnum
}

func (ctt *CellType) clearToken() {
	ctt.ItemID = ""
	ctt.OccupiedBy = vacantToken
}

func (ctt *CellType) isVacant() bool {
	return ctt.OccupiedBy == vacantToken
}

func (ctt *CellType) updateToken(id string, occupiedBy CatalogTokenEnum) {
	ctt.ItemID = id
	ctt.OccupiedBy = occupiedBy
}
