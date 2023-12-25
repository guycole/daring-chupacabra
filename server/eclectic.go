// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

func (at *AppType) scheduleCreateAction(id string, tokenType CatalogTokenEnum, turn int) {
	candidate := newEventNode(createAction, id, tokenType)
	at.EventArray.insertNode(candidate, turn)
}

func (at *AppType) scheduleDeleteAction(id string, tokenType CatalogTokenEnum, turn int) {
	candidate := newEventNode(deleteAction, id, tokenType)
	at.EventArray.insertNode(candidate, turn)
}

func (at *AppType) scheduleMoveAction(id string, tokenType CatalogTokenEnum, turn, yy, xx int) {
	candidate := newEventNode(moveAction, id, tokenType)
	candidate.MoveArgs = &MoveArgType{XX: xx, YY: yy}
	at.EventArray.insertNode(candidate, turn)
}

func (at *AppType) scheduleNominalAction(id string, tokenType CatalogTokenEnum, turn int) {
	candidate := newEventNode(nominalAction, id, tokenType)
	at.EventArray.insertNode(candidate, turn)
}

func (at *AppType) eclecticManager() {
	at.SugarLog.Infof("eclectic manager turn:%d", at.TurnCounter)

	eventFlag := true

	for eventFlag {
		ent, err := at.EventArray.selectNextNode(at.TurnCounter)
		if err != nil {
			// all events have been consumed
			eventFlag = false
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
