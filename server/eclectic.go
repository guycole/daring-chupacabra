// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

func (at *AppType) houseKeepingObj1(ent *EventNodeType) {
	at.SugarLog.Infof("houseKeepingObj1:%s", ent.ItemID)
}

func (at *AppType) serviceObj1(ent *EventNodeType) {
	/*
		switch ent.Action {
		case nothingAction:
			at.SugarLog.Debug("obj1Token/nothingAction")
		case createAction:
			at.SugarLog.Debug("obj1Token/createAction")

			location := randomLocation(maxCellArraySideY, maxCellArraySideX)
			at.Obj1StateMap.insertItem(ent.ItemID, location)

			at.CellArray.updateCell(ent.ItemID, location, obj1Token)
			at.scheduleEvent(houseKeepingAction, ent.ItemID, obj1Token, at.TurnCounter+1)
		case deleteAction:
			at.SugarLog.Debug("obj1Token/deleteAction")
		case houseKeepingAction:
			at.SugarLog.Debug("obj1Token/housekeepingAction")
			at.houseKeepingObj1(ent)
			at.scheduleEvent(houseKeepingAction, ent.ItemID, obj1Token, at.TurnCounter+1)
		case moveAction:
			at.SugarLog.Debug("obj1Token/moveAction")
		default:
			at.SugarLog.Fatal("uknown action")
		}
	*/
}

func (at *AppType) serviceObj2(ent *EventNodeType) {
	/*
		switch ent.Action {
		case nothingAction:
			at.SugarLog.Debug("obj2Token/nothingAction")
		case createAction:
			at.SugarLog.Debug("obj2Token/createAction")

			//location := randomLocation(maxCellArraySideY, maxCellArraySideX)
			//at.Obj2StateMap.insertItem(ent.ItemID, location)

			//at.CellArray.updateCell(ent.ItemID, location, obj1Token)
			at.scheduleEvent(houseKeepingAction, ent.ItemID, obj2Token, at.TurnCounter+1)
		case deleteAction:
			at.SugarLog.Debug("obj2Token/deleteAction")
		case houseKeepingAction:
			at.SugarLog.Debug("obj2Token/housekeepingAction")
			at.houseKeepingObj1(ent)
			at.scheduleEvent(houseKeepingAction, ent.ItemID, obj2Token, at.TurnCounter+1)
		case moveAction:
			at.SugarLog.Debug("obj2Token/moveAction")
		default:
			at.SugarLog.Fatal("uknown action")
		}
	*/
}

func (at *AppType) scheduleEvent(action EventActionEnum, id string, token CatalogTokenEnum, turn int) {
	at.EventArray.insertNode(action, id, token, turn)
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

			switch ent.TokenType {
			case vacantToken:
				at.SugarLog.Infof("vacant token")
			case obj1Token:
				at.serviceObj1(ent)
			case obj2Token:
				at.serviceObj2(ent)
			default:
				at.SugarLog.Fatal("uknown token")
			}
		}
	}
}
