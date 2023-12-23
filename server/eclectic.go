// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

func (at *AppType) scheduleEvent(action EventActionEnum, itemID string, turn int) {
	at.EventArray.insertNode(action, itemID, turn)
}

func (at *AppType) eclecticManager() {
	at.SugarLog.Infof("eclectic manager turn:%d", at.TurnCounter)

	runFlag := true
	for runFlag {
		ent, err := at.EventArray.selectNextNode(at.TurnCounter)
		if err != nil {
			runFlag = false
		} else {
			at.SugarLog.Infof("event:%v", ent)
		}
	}
}

func (at *AppType) seed() {
	target := "a6b1baf8-4eae-4d7a-ad8f-e31b8c7bf12e"

	location := LocationType{YY: 4, XX: 4}

	at.CatalogMap.insertItem(target, &location, obj1Token)
	at.Obj1StateMap.insertItem(target, &location)
	at.CellArray.updateCell(target, &location, obj1Token)

	at.scheduleEvent(createAction, target, at.TurnCounter+1)
}
