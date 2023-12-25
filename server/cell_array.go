// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"errors"
)

const maxCellArraySideX = 50
const maxCellArraySideY = 50

type CellArrayType [maxCellArraySideX][maxCellArraySideY]*CellType

func initializeCellArray() *CellArrayType {
	cellArray := new(CellArrayType)

	for yy := 0; yy < maxCellArraySideY; yy++ {
		for xx := 0; xx < maxCellArraySideX; xx++ {
			cellArray[yy][xx] = &CellType{ItemID: "", TokenType: vacantToken}
		}
	}

	return cellArray
}

func (cat *CellArrayType) clearCell(location *LocationType) error {
	if !location.legalLocation(maxCellArraySideY, maxCellArraySideX) {
		return errors.New("bad cell location")
	}

	cat[location.YY][location.XX].clearToken()

	return nil
}

func (cat *CellArrayType) isVacant(location *LocationType) bool {
	if !location.legalLocation(maxCellArraySideY, maxCellArraySideX) {
		return false
	}

	return cat[location.YY][location.XX].isVacant()
}

func (cat *CellArrayType) moveCell(source, destination *LocationType) error {
	if !destination.legalLocation(maxCellArraySideY, maxCellArraySideX) {
		return errors.New("destination cell bad location")
	}

	destinationCell := cat[destination.YY][destination.XX]
	if !destinationCell.isVacant() {
		return errors.New("destination cell is occupied")
	}

	sourceCell := cat[source.YY][source.XX]
	if sourceCell.isVacant() {
		return errors.New("source cell is vacant")
	}

	destinationCell.updateToken(sourceCell.ItemID, sourceCell.TokenType)
	sourceCell.clearToken()

	return nil
}

func (cat *CellArrayType) updateCell(id string, location *LocationType, tokenType CatalogTokenEnum) error {
	if !location.legalLocation(maxCellArraySideY, maxCellArraySideX) {
		return errors.New("bad cell location")
	}

	target := cat[location.YY][location.XX]

	target.ItemID = id
	target.TokenType = tokenType

	return nil
}
