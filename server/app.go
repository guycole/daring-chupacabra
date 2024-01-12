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
	FeatureFlags uint32             // control run time features
	GrpcPort     int                // gRPC port
	SugarLog     *zap.SugaredLogger // logging

	Quantum time.Duration
	RunFlag bool // true while scheduler runs

	CatalogMap *CatalogMapType // catalog of all items
	Eclectic   *EclecticType
}

func (at *AppType) timeKeeper() {
	var sleepTime time.Duration

	at.Quantum = 3 * time.Second

	for at.RunFlag {
		startTime := time.Now()

		at.Eclectic.TurnCounter++
		at.Eclectic.eclecticManager()

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

	// in the beginning...
	at.CatalogMap = initializeCatalogMap()
	at.Eclectic = initializeEclectic(at.SugarLog)
	//at.Obj1StateMap = initializeObj1Map()
	//at.Obj2StateMap = initializeObj2Map()

	at.RunFlag = true

	//at.genesis(movingToken)
}

// Run pacifier
func (at *AppType) run() {
	at.SugarLog.Info("run run run")

	//at.SugarLog.Fatal(http.ListenAndServe(":"+address, at.Router))

	if false {
		time.Sleep(10 * time.Second)
		at.RunFlag = false
	} else {
		go at.timeKeeper()

		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", at.GrpcPort))
		if err != nil {
			at.SugarLog.Fatalf("failed to listen: %v", err)
		}

		st := ServerType{Eclectic: at.Eclectic, SugarLog: at.SugarLog}

		grpcServer := grpc.NewServer()
		//pb.RegisterChupacabraServer(grpcServer, &ServerType{})
		pb.RegisterChupacabraServer(grpcServer, &st)
		at.SugarLog.Infof("server listening at %v", listener.Addr())

		if err := grpcServer.Serve(listener); err != nil {
			at.SugarLog.Fatalf("failed to serve: %v", err)
		}
	}
}
