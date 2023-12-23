// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import "errors"

type CatalogTokenEnum int

const (
	vacantToken CatalogTokenEnum = iota
	obj1Token
	obj2Token
)

func (cte CatalogTokenEnum) String() string {
	return [...]string{"vacant", "obj1", "obj2"}[cte]
}

type CatalogItemType struct {
	ItemID   string
	Location *LocationType
	Token    CatalogTokenEnum
}

type CatalogMapType map[string]*CatalogItemType

func initializeCatalogMap() *CatalogMapType {
	catalogMap := make(CatalogMapType)
	return &catalogMap
}

/*
type StateMapInterface interface {
	getItemID() string
	getLocation() LocationType
	getToken() CatalogTokenEnum
}
*/

func (catalogMap *CatalogMapType) deleteItem(itemID string) {
	delete(*catalogMap, itemID)
}

func (catalogMap *CatalogMapType) insertItem(itemID string, location *LocationType, token CatalogTokenEnum) {
	catalogItem := CatalogItemType{ItemID: itemID, Location: location, Token: token}
	(*catalogMap)[itemID] = &catalogItem
}

func (catalogMap *CatalogMapType) selectItem(itemID string) (*CatalogItemType, error) {
	result, ok := (*catalogMap)[itemID]
	if !ok {
		return nil, errors.New("not found in catalog")
	}

	return result, nil
}
