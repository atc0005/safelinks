// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/atc0005/safelinks/internal/safelinks"
)

// ReadURLsFromInput processes a given input string as a URL value. This
// input string represents a single URL given via CLI flag.
//
// If an input string is not provided, this function will attempt to read
// input URLs from stdin. Each input URL is unescaped and quoting removed.
//
// The collection of input URLs is returned or an error if one occurs.
func ReadURLsFromInput(inputURL string) ([]string, error) {
	var inputURLs []string

	// https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang
	// https://stackoverflow.com/a/26567513/903870
	// stat, _ := os.Stdin.Stat()
	// if (stat.Mode() & os.ModeCharDevice) == 0 {
	// 	fmt.Println("data is being piped to stdin")
	// } else {
	// 	fmt.Println("stdin is from a terminal")
	// }

	stat, _ := os.Stdin.Stat()

	switch {

	// We received one or more URLs via standard input.
	case (stat.Mode() & os.ModeCharDevice) == 0:
		// fmt.Fprintln(os.Stderr, "Received URL via standard input")
		return safelinks.ReadFromFile(os.Stdin)

	// We received a URL via positional argument. We ignore all but the first
	// one.
	case len(flag.Args()) > 0:
		// fmt.Fprintln(os.Stderr, "Received URL via positional argument")

		if strings.TrimSpace(flag.Args()[0]) == "" {
			return nil, safelinks.ErrInvalidURL
		}

		inputURLs = append(inputURLs, flag.Args()[0])

	// We received a URL via flag.
	case inputURL != "":
		// fmt.Fprintln(os.Stderr, "Received URL via flag")

		inputURLs = append(inputURLs, inputURL)

	// Input URL not given via positional argument, not given via flag either.
	// We prompt the user for a single input value.
	default:
		// fmt.Fprintln(os.Stderr, "default switch case triggered")

		input, err := safelinks.ReadURLFromUser()
		if err != nil {
			return nil, fmt.Errorf("error reading URL: %w", err)
		}

		if strings.TrimSpace(input) == "" {
			return nil, safelinks.ErrInvalidURL
		}

		inputURLs = append(inputURLs, input)
	}

	return inputURLs, nil
}

// ProcessInputURLs processes a given collection of input URL strings and
// emits successful decoding results to the specified results output sink.
// Errors are emitted to the specified error output sink if encountered but
// bulk processing continues until all input URLs have been evaluated.
//
// If requested, decoded URLs are emitted in verbose format.
//
// A boolean value is returned indicating whether any errors occurred.
func ProcessInputURLs(inputURLs []string, okOut io.Writer, errOut io.Writer, verbose bool) bool {
	var errEncountered bool

	for _, inputURL := range inputURLs {
		cleanedURL := safelinks.CleanURL(inputURL)
		safelink, err := url.Parse(cleanedURL)
		if err != nil {
			fmt.Fprintf(errOut, "Failed to parse URL: %v\n", err)

			errEncountered = true
			continue
		}

		if !safelinks.ValidSafeLinkURL(safelink) {
			fmt.Fprintf(errOut, "Invalid Safelinks URL %q\n", safelink)

			errEncountered = true
			continue
		}

		emitOutput(safelink, okOut, verbose)
	}

	return errEncountered
}
