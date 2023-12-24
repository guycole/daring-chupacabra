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

func (obj2Map *Obj2MapType) deleteItem(id string) {
	delete(*obj2Map, id)
}

func (obj2Map *Obj2MapType) insertItem(id string, location LocationType) {
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
