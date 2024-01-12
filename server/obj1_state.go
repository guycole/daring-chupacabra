// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import "errors"

type Obj1ItemType struct {
	ItemID   string
	Location *LocationType
}

type Obj1MapType map[string]*Obj1ItemType

func initializeObj1Map() *Obj1MapType {
	result := make(Obj1MapType)
	return &result
}

func (obj1Map *Obj1MapType) deleteItem(itemID string) {
	delete(*obj1Map, itemID)
}

func (obj1Map *Obj1MapType) insertItem(id string, location *LocationType) {
	catalogItem := Obj1ItemType{ItemID: id, Location: location}
	(*obj1Map)[id] = &catalogItem
}

func (obj1Map *Obj1MapType) selectItem(id string) (*Obj1ItemType, error) {
	result, ok := (*obj1Map)[id]
	if !ok {
		return nil, errors.New("not found in catalog")
	}

	return result, nil
}

func (at *AppType) serviceObj1(ent *EventNodeType) {
	switch ent.Action {
	case nothingAction:
		at.SugarLog.Debug("obj1Token/nothingAction")
		/*
			case createAction:
				var location *LocationType
				for flag := true; flag; flag = !at.CellArray.isVacant(location) {
					location = randomLocation(maxCellArraySideY, maxCellArraySideX)
				}

				at.CellArray.updateCell(ent.ItemID, location, obj1Token)
				at.CatalogMap.insertItem(ent.ItemID, location, obj1Token)
				at.Obj1StateMap.insertItem(ent.ItemID, location)

				//at.scheduleNominalAction(ent.ItemID, obj1Token, at.TurnCounter+1)
			case deleteAction:
				at.SugarLog.Debug("obj1Token/deleteAction")

				target, err := at.Obj1StateMap.selectItem(ent.ItemID)
				if err == nil {
					at.CellArray.clearCell(target.Location)
					at.Obj1StateMap.deleteItem(ent.ItemID)
					at.CatalogMap.updateItemLifeCycle(ent.ItemID, deleted)
				}
			case moveAction:
				at.SugarLog.Debug("obj1Token/moveAction")
			case nominalAction:
				//at.scheduleNominalAction(ent.ItemID, obj1Token, at.TurnCounter+1)
		*/
	default:
		at.SugarLog.Fatal("uknown action")
	}
}
