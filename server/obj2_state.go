// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import "errors"

type Obj2ItemType struct {
	ItemID   string
	Location LocationType
}

type Obj2MapType map[string]*Obj2ItemType

func initializeObj2Map() *Obj2MapType {
	result := make(Obj2MapType)
	return &result
}

func (obj2Map *Obj2MapType) deleteItem(itemID string) {
	delete(*obj2Map, itemID)
}

func (obj2Map *Obj2MapType) insertItem(itemID string, location LocationType) {
	catalogItem := Obj2ItemType{ItemID: itemID, Location: location}
	(*obj2Map)[itemID] = &catalogItem
}

func (obj2Map *Obj2MapType) selectItem(itemID string) (*Obj2ItemType, error) {
	result, ok := (*obj2Map)[itemID]
	if !ok {
		return nil, errors.New("not found in catalog")
	}

	return result, nil
}
