// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type CellTokenEnum int

const (
	vacantToken CellTokenEnum = iota
	obj1Token
	obj2Token
)

func (cte CellTokenEnum) String() string {
	return [...]string{"vacant", "obj1", "obj2"}[cte]
}

type CellTokenType struct {
	ItemID     string // same as catalog item ID
	OccupiedBy CellTokenEnum
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

func (ctt *CellTokenType) updateToken(itemID string, occupiedBy CellTokenEnum) {
	ctt.ItemID = itemID
	ctt.OccupiedBy = occupiedBy
}
