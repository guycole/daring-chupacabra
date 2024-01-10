// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc"

	pb "github.com/guycole/daring-chupacabra/proto"
)

type AppType struct {
	FeatureFlags  uint32             // control run time features
	Configuration *ConfigurationType // configuration parameters
	GrpcPort      int                // gRPC port
	SugarLog      *zap.SugaredLogger // logging

	Quantum     time.Duration
	RunFlag     bool // true while scheduler runs
	TurnCounter int  // current turn

	CellArray    *CellArrayType  // 2D game board
	EventArray   *EventArrayType // scheduled events
	CatalogMap   *CatalogMapType // catalog of all items
	Obj1StateMap *Obj1MapType    // state of all obj1 items
	Obj2StateMap *Obj2MapType    // state of all obj2 items
}

func (at *AppType) timeKeeper() {
	var sleepTime time.Duration

	at.Quantum = 3 * time.Second

	for at.RunFlag {
		startTime := time.Now()

		at.TurnCounter++
		at.eclecticManager()

		stopTime := time.Now()
		deltaTime := stopTime.Sub(startTime)
		at.SugarLog.Debugf("duration:%v", deltaTime)

		if deltaTime > at.Quantum {
			at.SugarLog.Warnf("overrun:%v", deltaTime)
			sleepTime = 0
		} else {
			sleepTime = at.Quantum - deltaTime
			at.SugarLog.Debugf("sleep:%v", sleepTime)
		}

		time.Sleep(sleepTime)
	}
}

func (at *AppType) initialize(configurationFilename string) {
	if isDevelopmentModeLogging(at.FeatureFlags) {
		at.SugarLog = zapSetup(true)
		at.SugarLog.Debug("debug level log entry")
	}

	at.Configuration = &ConfigurationType{ConfigurationFilename: configurationFilename}
	at.Configuration.initialize(at.SugarLog)

	// in the beginning...
	at.CatalogMap = initializeCatalogMap()
	at.CellArray = initializeCellArray()
	at.EventArray = initializeEventArray()
	at.Obj1StateMap = initializeObj1Map()
	at.Obj2StateMap = initializeObj2Map()

	at.RunFlag = true
	at.TurnCounter = 0

	at.genesis(movingToken)
}

// Run pacifier
func (at *AppType) run() {
	at.SugarLog.Info("run run run")

	go at.timeKeeper()

	//at.SugarLog.Fatal(http.ListenAndServe(":"+address, at.Router))

	if false {
		time.Sleep(10 * time.Second)
		at.RunFlag = false
	} else {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", at.GrpcPort))
		if err != nil {
			at.SugarLog.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterChupacabraServer(grpcServer, &ServerType{})
		at.SugarLog.Infof("server listening at %v", listener.Addr())

		if err := grpcServer.Serve(listener); err != nil {
			at.SugarLog.Fatalf("failed to serve: %v", err)
		}
	}
}
