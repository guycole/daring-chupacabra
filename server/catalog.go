// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import "errors"

type CatalogItemType struct {
	ItemID string
}

type CatalogMapType map[string]*CatalogItemType

func initializeCatalogMap() *CatalogMapType {
	catalogMap := make(CatalogMapType)
	return &catalogMap
}

func (catalogMap *CatalogMapType) deleteItem(itemID string) {
	delete(*catalogMap, itemID)
}

func (catalogMap *CatalogMapType) insertItem(itemID string) {
	catalogItem := CatalogItemType{ItemID: itemID}
	(*catalogMap)[itemID] = &catalogItem
}

func (catalogMap *CatalogMapType) selectItem(itemID string) (*CatalogItemType, error) {
	result, ok := (*catalogMap)[itemID]
	if !ok {
		return nil, errors.New("not found in catalog")
	}

	return result, nil
}
