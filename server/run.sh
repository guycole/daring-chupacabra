#!/bin/bash
#
# Title:run.sh
# Description:
#
PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
export CONFIGURATION_FILENAME="config1.yaml"
export FEATURE_FLAGS="25"
export GRPC_PORT="50051"
#
./server
#