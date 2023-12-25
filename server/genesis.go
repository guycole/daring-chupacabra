// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"github.com/google/uuid"
)

type GenesisEnum int

const (
	noSelection GenesisEnum = iota
	conwayLife
	movingToken
)

func (ge GenesisEnum) String() string {
	return [...]string{"noSelection", "conwayLife", "movingToken"}[ge]
}

func (at *AppType) createObj1() {
	id := uuid.NewString()

	//at.SugarLog.Infof("create obj1:%s", id)

	at.scheduleCreateAction(id, obj1Token, at.TurnCounter+1)
}

func (at *AppType) createObj2() {
	id := uuid.NewString()

	//at.SugarLog.Infof("create obj2:%s", id)

	at.scheduleCreateAction(id, obj2Token, at.TurnCounter+1)
	at.scheduleMoveAction(id, obj2Token, at.TurnCounter+2, 3, 3)
	at.scheduleMoveAction(id, obj2Token, at.TurnCounter+3, 3, 3)
}

func (at *AppType) movingToken() {
	for ndx := 0; ndx < 2; ndx++ {
		at.createObj1()
		at.createObj2()
	}
}

func (at *AppType) genesis(selected GenesisEnum) {
	switch selected {
	case conwayLife:
		at.SugarLog.Debug("genesis: conwayLife")
		//		at.genesisConwayLife()
	case movingToken:
		at.SugarLog.Debug("genesis: movingToken")
		at.movingToken()
	default:
		at.SugarLog.Info("genesis: no selection")
	}
}
