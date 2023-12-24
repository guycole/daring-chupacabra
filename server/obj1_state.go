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

/*
func (obj1Map *Obj1MapType) updateItem(action EventActionEnum, itemID string) error {
	switch action {
	case createAction:
		// fresh item
		at.SugarLog.Infof("create obj1:%s", itemID)
	case deleteAction:
		// delete item
		at.SugarLog.Infof("delete obj1:%s", itemID)
	case houseKeepingAction:
		// house keeping
		at.SugarLog.Infof("housekeeping obj1:%s", itemID)
	case moveAction:
		// move item
		at.SugarLog.Infof("move obj1:%s", itemID)
	default:
		return errors.New("unknown action")
	}
}
*/
