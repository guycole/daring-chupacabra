// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"go.uber.org/zap"
)

type EclecticType struct {
	EventArray  *EventArrayType    // scheduled events
	ResponseMap *ResponseMapType   // responses waiting for client
	SugarLog    *zap.SugaredLogger // logging
	TurnCounter int                // current turn
}

func initializeEclectic(sugarLog *zap.SugaredLogger) *EclecticType {
	et := EclecticType{SugarLog: sugarLog, TurnCounter: 0}
	et.EventArray = initializeEventArray()
	et.ResponseMap = initializeResponseMap()
	return &et
}

func (et *EclecticType) insertNodeNextTurn(candidate *EventNodeType) {
	et.SugarLog.Debug("insert node next turn")
	et.EventArray.insertNode(candidate, et.TurnCounter+1)
}

func (et *EclecticType) getResponse(clientId string) []*ResponseNodeType {
	et.SugarLog.Debug("get responses")
	return et.ResponseMap.getResponse(clientId)
}

/*
//// fix below

func (at *AppType) scheduleCreateAction(id string, tokenType CatalogTokenEnum, turn int) {
	// at.SugarLog.Infof("scheduler create id:%s tokenType:%s turn:%d", id, tokenType, turn)
	//candidate := newEventNode(createAction, id, tokenType)
	//at.EventArray.insertNode(candidate, turn)
}

func (at *AppType) scheduleDeleteAction(id string, tokenType CatalogTokenEnum, turn int) {
	//candidate := newEventNode(deleteAction, id, tokenType)
	//at.EventArray.insertNode(candidate, turn)
}

func (at *AppType) scheduleMoveAction(id string, tokenType CatalogTokenEnum, turn, yy, xx int) {
	candidate := newEventNode(moveAction, id, tokenType)
	candidate.MoveArgs = &MoveArgType{XX: xx, YY: yy}
	//at.EventArray.insertNode(candidate, turn)
}

func (at *AppType) scheduleNominalAction(id string, tokenType CatalogTokenEnum, turn int) {
	//candidate := newEventNode(nominalAction, id, tokenType)
	//at.EventArray.insertNode(candidate, turn)
}
*/

func (et *EclecticType) eclecticManager() {
	et.SugarLog.Infof("eclectic manager turn:%d", et.TurnCounter)

	eventFlag := true

	for eventFlag {
		ent, err := et.EventArray.selectNextNode(et.TurnCounter)
		if err != nil {
			// all events have been consumed
			et.SugarLog.Debugf("no events")
			eventFlag = false
		} else {
			et.SugarLog.Debugf("action:%s id:%s tokenType:%s turn:%d", ent.Action, ent.ItemID, ent.TokenType, et.TurnCounter)
			switch ent.Action {
			case parseAction:
				et.SugarLog.Debugf("parseAction clientId:%s receiptId:%s", ent.ClientID, ent.ReceiptID)
				// parse

			}

			/*
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
			*/
		}
	}
}
