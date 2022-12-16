// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {

	cfg := NewConfig()

	if cfg.Version {
		fmt.Println(Version())
		os.Exit(0)
	}

	inputURL := parseInputURL(cfg.URL)

	safelink, err := url.Parse(inputURL)
	if err != nil {
		log.Fatal("Failed to parse URL")
	}

	switch {
	case cfg.Verbose:
		verboseOutput(safelink, os.Stdout)

	default:
		simpleOutput(safelink, os.Stdout)
	}

}
