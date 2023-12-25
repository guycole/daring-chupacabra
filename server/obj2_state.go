// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"errors"
)

type Obj2ItemType struct {
	ItemID   string
	Location *LocationType
}

type Obj2MapType map[string]*Obj2ItemType

func initializeObj2Map() *Obj2MapType {
	result := make(Obj2MapType)
	return &result
}

func (obj2Map *Obj2MapType) deleteItem(id string) {
	delete(*obj2Map, id)
}

func (obj2Map *Obj2MapType) insertItem(id string, location *LocationType) {
	catalogItem := Obj2ItemType{ItemID: id, Location: location}
	(*obj2Map)[id] = &catalogItem
}

func (obj2Map *Obj2MapType) selectItem(id string) (*Obj2ItemType, error) {
	result, ok := (*obj2Map)[id]
	if !ok {
		return nil, errors.New("not found in catalog")
	}

	return result, nil
}

func (at *AppType) serviceObj2(ent *EventNodeType) {
	switch ent.Action {
	case nothingAction:
		at.SugarLog.Debug("obj2Token/nothingAction")
	case createAction:
		var location *LocationType
		for flag := true; flag; flag = !at.CellArray.isVacant(location) {
			location = randomLocation(maxCellArraySideY, maxCellArraySideX)
		}

		at.CellArray.updateCell(ent.ItemID, location, obj2Token)
		at.CatalogMap.insertItem(ent.ItemID, location, obj2Token)
		at.Obj2StateMap.insertItem(ent.ItemID, location)
	case deleteAction:
		at.SugarLog.Debug("obj2Token/deleteAction")

		target, err := at.Obj2StateMap.selectItem(ent.ItemID)
		if err == nil {
			at.CellArray.clearCell(target.Location)
			at.Obj2StateMap.deleteItem(ent.ItemID)
			at.CatalogMap.updateItemLifeCycle(ent.ItemID, deleted)
		}
	case moveAction:
		target, err := at.Obj2StateMap.selectItem(ent.ItemID)
		if err != nil {
			at.SugarLog.Fatal("missing item in obj2StateMap")
		}

		at.SugarLog.Debugf("id:%s locationy:%d locationx:%d turn:%d", ent.ItemID, target.Location.YY, target.Location.XX, at.TurnCounter)

		newLocation := &LocationType{YY: target.Location.YY + ent.MoveArgs.YY, XX: target.Location.XX + ent.MoveArgs.XX}
		if !newLocation.legalLocation(maxCellArraySideY, maxCellArraySideX) {
			at.SugarLog.Fatal("bad location")
		}
		if !at.CellArray.isVacant(newLocation) {
			at.SugarLog.Fatal("not vacant")
		}

		at.CellArray.moveCell(target.Location, newLocation)

		target.Location = newLocation
		at.SugarLog.Debugf("id:%s locationy:%d locationx:%d turn:%d", ent.ItemID, target.Location.YY, target.Location.XX, at.TurnCounter)

		at.CatalogMap.updateItemLocation(ent.ItemID, newLocation)

		//at.scheduleMoveAction(ent.ItemID, obj2Token, at.TurnCounter+1, 3, 3)
	case nominalAction:
		at.SugarLog.Debug("obj2Token/nominalAction")
		at.scheduleNominalAction(ent.ItemID, obj2Token, at.TurnCounter+1)
	default:
		at.SugarLog.Fatal("uknown action")
	}
}
