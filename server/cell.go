// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type CellType struct {
	ItemID    string // same as catalog item ID
	TokenType CatalogTokenEnum
}

func (ctt *CellType) clearToken() {
	ctt.ItemID = ""
	ctt.TokenType = vacantToken
}

func (ctt *CellType) isVacant() bool {
	return ctt.TokenType == vacantToken
}

func (ctt *CellType) updateToken(id string, tokenType CatalogTokenEnum) {
	ctt.ItemID = id
	ctt.TokenType = tokenType
}
