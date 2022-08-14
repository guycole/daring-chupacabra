#!/bin/bash
#
# Title:base64_driver.sh
# Description: generate base64 encoded strings for k8s secrets
# Development Environment: OS X 10.13.6
# Author: G.S. Cole (guycole at gmail dot com)
#
echo -n "bigSekret" | base64
#
echo -n "redis-master.chupacabra.svc.cluster.local:6379" | base64
#
