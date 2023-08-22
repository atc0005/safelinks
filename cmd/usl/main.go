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
	"net/url"
	"os"
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

		input, err := readURLsFromFile(f)
		if err != nil {
			fmt.Printf("Failed to read URLs from %q: %v\n", cfg.Filename, err)
			os.Exit(1)
		}

		inputURLs = input

	default:
		input, err := processInputAsURL(cfg.URL)
		if err != nil {
			fmt.Printf("Failed to parse input as URL: %v\n", err)
			os.Exit(1)
		}

		inputURLs = input
	}

	var errEncountered bool
	for _, inputURL := range inputURLs {
		safelink, err := url.Parse(inputURL)
		if err != nil {
			fmt.Printf("Failed to parse URL: %v\n", err)

			errEncountered = true
			continue
		}

		if err := assertValidURLParameter(safelink); err != nil {
			fmt.Fprintf(os.Stderr, "Invalid Safelinks URL %q: %v\n", safelink, err)

			errEncountered = true
			continue
		}

		switch {
		case cfg.Verbose:
			verboseOutput(safelink, os.Stdout)

		default:
			simpleOutput(safelink, os.Stdout)
		}
	}

	// Ensure unsuccessful error code if one encountered.
	if errEncountered {
		os.Exit(1)
	}
}
