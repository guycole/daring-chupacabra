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

func (obj1Map *Obj1MapType) insertItem(itemID string, location *LocationType) {
	catalogItem := Obj1ItemType{ItemID: itemID, Location: location}
	(*obj1Map)[itemID] = &catalogItem
}

func (obj1Map *Obj1MapType) selectItem(itemID string) (*Obj1ItemType, error) {
	result, ok := (*obj1Map)[itemID]
	if !ok {
		return nil, errors.New("not found in catalog")
	}

	return result, nil
}
