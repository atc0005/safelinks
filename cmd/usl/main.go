// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

package main

import (
	"fmt"
	"os"

	"github.com/atc0005/safelinks/internal/safelinks"
)

func main() {

	cfg := NewConfig()

	if cfg.Version {
		fmt.Println(Version())
		os.Exit(0)
	}

	var inputURLs []string

	switch {
	case cfg.Filename != "":
		f, err := os.Open(cfg.Filename)
		if err != nil {
			fmt.Printf("Failed to open %q: %v\n", cfg.Filename, err)
			os.Exit(1)
		}

		input, err := safelinks.ReadURLsFromFile(f)
		if err != nil {
			fmt.Printf("Failed to read URLs from %q: %v\n", cfg.Filename, err)
			os.Exit(1)
		}

		inputURLs = input

	default:
		input, err := safelinks.ProcessInputAsURL(cfg.URL)
		if err != nil {
			fmt.Printf("Failed to parse input as URL: %v\n", err)
			os.Exit(1)
		}

		inputURLs = input
	}

	hasErr := safelinks.ProcessInputURLs(inputURLs, os.Stdout, os.Stderr, cfg.Verbose)

	// Ensure unsuccessful error code if one encountered.
	if hasErr {
		os.Exit(1)
	}
}
