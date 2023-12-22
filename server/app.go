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
	FeatureFlags  uint32
	Configuration *ConfigurationType
	GrpcPort      int
	SugarLog      *zap.SugaredLogger

	TurnCounter int
	EventArray  *EventArrayType
}

func (at *AppType) runaturn() time.Duration {
	at.SugarLog.Infof("turn:%d", at.TurnCounter)

	startTime := time.Now()
	at.SugarLog.Debugf("start:%v", startTime)

	//	discoverCandidates()

	stopTime := time.Now()
	at.SugarLog.Debugf("stop:%v", stopTime)

	deltaTime := stopTime.Sub(startTime)

	return deltaTime
}

func (at *AppType) timeKeeper() {
	go func() {
		for {
			duration := at.runaturn()
			at.SugarLog.Debugf("duration:%v", duration)
			time.Sleep(1 * time.Second)
			at.TurnCounter++

			if at.TurnCounter > 10 {
				break
			}
		}
	}()

	at.SugarLog.Infof("timeKeeper exit")
}

func (at *AppType) initialize(configurationFilename string) {
	if isDevelopmentModeLogging(at.FeatureFlags) {
		at.SugarLog = zapSetup(true)
		at.SugarLog.Debug("debug level log entry")
	}

	at.Configuration = &ConfigurationType{ConfigurationFilename: configurationFilename}
	at.Configuration.initialize(at.SugarLog)

	at.EventArray = initializeEventArray()
	/*
		at.EventArray.insert("item1", 101)
		at.EventArray.insert("item2", 12)
		at.EventArray.insert("item3", 12)
		at.EventArray.dumper()
	*/
}

// Run pacifier
func (at *AppType) run() {
	at.SugarLog.Info("run run run")

	//at.timeKeeper()

	//at.SugarLog.Fatal(http.ListenAndServe(":"+address, at.Router))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", at.GrpcPort))
	if err != nil {
		at.SugarLog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChupacabraServer(grpcServer, &serverType{})
	at.SugarLog.Infof("server listening at %v", listener.Addr())

	//	if err := grpcServer.Serve(listener); err != nil {
	//		at.SugarLog.Fatalf("failed to serve: %v", err)
	//	}

	//time.Sleep(22 * time.Second)
}