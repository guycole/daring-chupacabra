// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc"

	pb "github.com/guycole/daring-chupacabra/proto"
)

type AppType struct {
	FeatureFlags  uint32
	Configuration *ConfigurationType
	GrpcPort      int
	SugarLog      *zap.SugaredLogger

	Quantum      time.Time
	RunFlag      bool
	TurnCounter  int
	CellArray    *CellArrayType
	EventArray   *EventArrayType
	CatalogMap   *CatalogMapType
	Obj1StateMap *Obj1MapType
}

func (at *AppType) timeKeeper() {
	for at.RunFlag {
		startTime := time.Now()
		fmt.Println(reflect.TypeOf(startTime))

		at.TurnCounter++
		at.eclecticManager()

		stopTime := time.Now()
		deltaTime := stopTime.Sub(startTime)
		at.SugarLog.Debugf("duration:%v", deltaTime)

		time.Sleep(1 * time.Second)
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

	at.RunFlag = true
	at.TurnCounter = 0

	at.genesis()
}

// Run pacifier
func (at *AppType) run() {
	at.SugarLog.Info("run run run")

	go at.timeKeeper()
	time.Sleep(10 * time.Second)
	at.RunFlag = false

	//at.SugarLog.Fatal(http.ListenAndServe(":"+address, at.Router))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", at.GrpcPort))
	if err != nil {
		at.SugarLog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChupacabraServer(grpcServer, &ServerType{})
	at.SugarLog.Infof("server listening at %v", listener.Addr())

	//	if err := grpcServer.Serve(listener); err != nil {
	//		at.SugarLog.Fatalf("failed to serve: %v", err)
	//	}

	//time.Sleep(22 * time.Second)
}
