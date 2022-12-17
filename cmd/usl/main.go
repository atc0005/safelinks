// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {

	cfg := NewConfig()

	if cfg.Version {
		fmt.Println(Version())
		os.Exit(0)
	}

	inputURL, err := processInputAsURL(cfg.URL)
	if err != nil {
		fmt.Printf("Failed to parse input as URL: %v\n", err)
		os.Exit(1)
	}

	safelink, err := url.Parse(inputURL)
	if err != nil {
		fmt.Printf("Failed to parse URL: %v\n", err)
		os.Exit(1)
	}

	if err := assertValidURLParameter(safelink); err != nil {
		fmt.Printf("Invalid Safelinks URL: %v\n", err)
		os.Exit(1)
	}

	switch {
	case cfg.Verbose:
		verboseOutput(safelink, os.Stdout)

	default:
		simpleOutput(safelink, os.Stdout)
	}

}
