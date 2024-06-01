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
	"io"
	"log"
	"os"

	"github.com/atc0005/safelinks/internal/safelinks"
)

func main() {
	userFeedbackOut := os.Stderr

	// use io.Discard for normal operation
	// switch to os.Stderr for debugging
	debugLoggingOut := io.Discard

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(debugLoggingOut)

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
			_, _ = fmt.Fprintf(userFeedbackOut, "Failed to open %q: %v\n", cfg.Filename, err)
			os.Exit(1)
		}

		input, err := safelinks.ReadFromFile(f)
		if err != nil {
			_, _ = fmt.Fprintf(userFeedbackOut, "Failed to read URLs from %q: %v\n", cfg.Filename, err)
			os.Exit(1)
		}

		inputURLs = input

	default:
		input, err := ReadURLsFromInput(cfg.URL)
		if err != nil {
			_, _ = fmt.Fprintf(userFeedbackOut, "Failed to parse input as URL: %v\n", err)
			os.Exit(1)
		}

		inputURLs = input
	}

	hasErr := ProcessInputURLs(inputURLs, os.Stdout, userFeedbackOut, cfg.Verbose)

	// Ensure unsuccessful error code if one encountered.
	if hasErr {
		os.Exit(1)
	}
}
