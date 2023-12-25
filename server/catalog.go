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

type LifeCycleEnum int

const (
	scheduled LifeCycleEnum = iota
	created
	deleted
)

func (lce LifeCycleEnum) String() string {
	return [...]string{"created", "deleted"}[lce]
}

type CatalogItemType struct {
	ItemID    string
	LifeCycle LifeCycleEnum
	Location  *LocationType
	TokenType CatalogTokenEnum
}

type CatalogMapType map[string]*CatalogItemType

func initializeCatalogMap() *CatalogMapType {
	catalogMap := make(CatalogMapType)
	return &catalogMap
}

func (catalogMap *CatalogMapType) insertItem(id string, location *LocationType, token CatalogTokenEnum) {
	catalogItem := CatalogItemType{ItemID: id, LifeCycle: created, Location: location, TokenType: token}
	(*catalogMap)[id] = &catalogItem
}

func (catalogMap *CatalogMapType) selectItem(id string) (*CatalogItemType, error) {
	result, ok := (*catalogMap)[id]
	if !ok {
		return nil, errors.New("not found in catalog")
	}

	return result, nil
}

func (catalogMap *CatalogMapType) updateItemLifeCycle(id string, lifeCycle LifeCycleEnum) {
	catalogItem, err := catalogMap.selectItem(id)
	if err != nil {
		return
	}

	catalogItem.LifeCycle = lifeCycle
}

func (catalogMap *CatalogMapType) updateItemLocation(id string, location *LocationType) {
	catalogItem, err := catalogMap.selectItem(id)
	if err != nil {
		return
	}

	catalogItem.Location = location
}
