// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"go.uber.org/zap"

	"gopkg.in/yaml.v3"
)

type ConfigurationType struct {
	ConfigurationFilename string
	ClockStart            int `yaml:"clockStart"`
	ClockStop             int `yaml:"clockStop"`
}

type Car struct {
	TopSpeed   int      `yaml:"topspeed"`
	Name       string   `yaml:"name"`
	Cool       bool     `yaml:"cool"`
	Passengers []string `yaml:"passengers"`
}

func (ct *ConfigurationType) initialize(sugarLog *zap.SugaredLogger) {
	sugarLog.Info("configure")

	f, err := os.ReadFile(ct.ConfigurationFilename)
	if err != nil {
		log.Fatal(err)
	}

	var c Car

	// Unmarshal our input YAML file into empty Car (var c)
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	log.Print(c)
}
