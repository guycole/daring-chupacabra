// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestCatalogItemType(t *testing.T) {
	tests := []struct {
		candidate string
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83"},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc"},
	}

	for _, ndx := range tests {
		result := CatalogItemType{ItemID: ndx.candidate}
		if result.ItemID != ndx.candidate {
			t.Errorf("TestCatalogType failure")
		}
	}
}

func TestCatalogMapOperations(t *testing.T) {
	catalogMap := initializeCatalogMap()

	tests := []struct {
		candidate string
		tokenType CatalogTokenEnum
		yy, xx    int
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83", obj1Token, 3, 5},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc", obj2Token, 7, 9},
	}

	for _, ndx := range tests {
		location := LocationType{YY: ndx.yy, XX: ndx.xx}
		catalogMap.insertItem(ndx.candidate, &location, ndx.tokenType)
	}

	if len(*catalogMap) != 2 {
		t.Errorf("TestCatalogMapOperations failure")
	}

	// select test
	for _, ndx := range tests {
		catalogItem, err := catalogMap.selectItem(ndx.candidate)
		if err != nil {
			t.Errorf("TestCatalogMapOperations failure")
		}
		if catalogItem.ItemID != ndx.candidate {
			t.Errorf("TestCatalogMapOperations failure")
		}
		if catalogItem.TokenType != ndx.tokenType {
			t.Errorf("TestCatalogMapOperations failure")
		}
	}

	// update test
	location := LocationType{YY: 13, XX: 13}

	for _, ndx := range tests {
		catalogMap.updateItemLifeCycle(ndx.candidate, deleted)
		catalogMap.updateItemLocation(ndx.candidate, &location)

		catalogItem, err := catalogMap.selectItem(ndx.candidate)
		if err != nil {
			t.Errorf("TestCatalogMapOperations failure")
		}
		if catalogItem.LifeCycle != deleted {
			t.Errorf("TestCatalogMapOperations failure")
		}
		if sameLocation(catalogItem.Location, &location) != true {
			t.Errorf("TestCatalogMapOperations failure")
		}
	}
}
