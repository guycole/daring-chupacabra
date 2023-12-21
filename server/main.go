package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//var (
//	rpcPort = flag.Int("port", 50051, "server port")
//)

const banner = "chupacapra-server 0.0"

func main() {
	flag.Parse()

	rand.New(rand.NewSource(time.Now().UnixNano()))

	log.Println(banner)

	sugarLog := zapSetup(false)
	sugarLog.Info(banner)

	var configurationFilename string

	app := AppType{SugarLog: sugarLog}

	envVars := [...]string{"CONFIGURATION_FILENAME", "FEATURE_FLAGS", "GRPC_PORT"}

	for index, element := range envVars {
		temp, err := os.LookupEnv(element)
		if err {
			sugarLog.Infof("%d:%s:%s", index, element, temp)
		} else {
			sugarLog.Fatal("missing:", element)
		}

		switch element {
		case "CONFIGURATION_FILENAME":
			configurationFilename = temp
		case "FEATURE_FLAGS":
			temp, err := strconv.Atoi(temp)
			if err == nil {
				app.FeatureFlags = uint32(temp)
			} else {
				sugarLog.Fatal("bad featureFlags")
			}
		case "GRPC_PORT":
			temp, err := strconv.Atoi(temp)
			if err == nil {
				app.GrpcPort = temp
			} else {
				sugarLog.Fatal("bad grpcPort")
			}
		default:
			sugarLog.Fatal("unknown environment var:", element)
		}
	}

	app.initialize(configurationFilename)
	app.run()
}
