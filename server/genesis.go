// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"github.com/google/uuid"
)

func (at *AppType) createObj1() {
	id := uuid.NewString()

	at.SugarLog.Infof("create obj1:%s", id)

	//location := randomLocation(maxCellArraySideY, maxCellArraySideX)
	//at.Obj1StateMap.insertItem(id, location)

	at.scheduleEvent(createAction, id, obj1Token, at.TurnCounter+1)
}

func (at *AppType) createObj2() {
	id := uuid.NewString()

	at.SugarLog.Infof("create obj2:%s", id)

	at.scheduleEvent(createAction, id, obj2Token, at.TurnCounter+1)
}

/*
at.CatalogMap.insertItem(id, location, obj1Token)

	at.CellArray.updateCell(target, &location, obj1Token)
*/

func (at *AppType) genesis() {
	for ndx := 0; ndx < 3; ndx++ {
		at.createObj1()
		at.createObj2()
	}
}
