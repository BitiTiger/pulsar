/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2021, Lewis Cook <lcook@FreeBSD.org>
 * Copyright (c) 2023, Cameron Himes <cameron.himes@hotmail.com>
 * All rights reserved.
 */
package main

import (
	"flag"
	"fmt"

	"github.com/BitiTiger/pulsar/internal/bot"
	"github.com/BitiTiger/pulsar/internal/version"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		cfgFile     string
		color       bool
		showVersion bool
		verbosity   int
	)

	flag.IntVar(&verbosity, "V", 1, "Log verbosity level (1-3)")
	flag.StringVar(&cfgFile, "c", "config.json", "JSON configuration file path")
	flag.BoolVar(&showVersion, "v", false, "Display pulsar version")
	flag.BoolVar(&color, "d", false, "Disable color output in logs")
	flag.Parse()

	log.SetFormatter(&nested.Formatter{
		ShowFullLevel:   true,
		TrimMessages:    true,
		TimestampFormat: "[02/Jan/2006:15:04:05]",
		NoFieldsColors:  true,
		NoColors:        color,
	})

	if showVersion {
		fmt.Println(version.Build)
		return
	}
	/*
	 * Clamp the verbosity with an lower bound of 1 and
	 * upper bound of 3 (1-3).
	 */
	if verbosity < 1 {
		verbosity = 1
	}

	if verbosity > 3 {
		verbosity = 3
	}

	switch verbosity {
	case 1:
		log.SetLevel(log.InfoLevel)
	case 2:
		log.SetLevel(log.DebugLevel)
	case 3:
		log.SetLevel(log.TraceLevel)
	}

	bot.Run(cfgFile)
}
