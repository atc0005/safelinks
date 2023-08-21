// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"html"
	"io"
	"net/url"
	"os"
	"strings"
)

// readURLFromUser attempts to read a given URL pattern from the user via
// stdin prompt.
func readURLFromUser() (string, error) {
	fmt.Print("Enter URL: ")

	// NOTE: fmt.Scanln does not seem to handle the length of the input URL
	// properly. We use bufio.NewScanner to work around this.
	//
	// var input string
	// _, err := fmt.Scanln(&input)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text(), scanner.Err()
}

// readURLsFromFile attempts to read URL patterns from a given file
// (io.Reader).
//
// The collection of input URLs is returned or an error if one occurs.
func readURLsFromFile(r io.Reader) ([]string, error) {
	var inputURLs []string

	// Loop over input "reader" and attempt to collect each item.
	scanner := bufio.NewScanner((r))
	for scanner.Scan() {
		txt := scanner.Text()

		if strings.TrimSpace(txt) == "" {
			continue
		}

		inputURLs = append(inputURLs, txt)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading URLs: %w", err)
	}

	if len(inputURLs) == 0 {
		return nil, ErrInvalidURL
	}

	return inputURLs, nil
}

// processInputAsURL processes a given input string as a URL value. This
// input string represents a single URL given via CLI flag.
//
// If an input string is not provided, this function will attempt to read
// input URLs from stdin. Each input URL is unescaped and quoting removed.
//
// The collection of input URLs is returned or an error if one occurs.
func processInputAsURL(inputURL string) ([]string, error) {
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
		return readURLsFromFile(os.Stdin)

	// We received a URL via positional argument. We ignore all but the first
	// one.
	case len(flag.Args()) > 0:
		// fmt.Fprintln(os.Stderr, "Received URL via positional argument")

		if strings.TrimSpace(flag.Args()[0]) == "" {
			return nil, ErrInvalidURL
		}

		inputURLs = append(inputURLs, cleanURL(flag.Args()[0]))

	// We received a URL via flag.
	case inputURL != "":
		// fmt.Fprintln(os.Stderr, "Received URL via flag")

		inputURLs = append(inputURLs, cleanURL(inputURL))

	// Input URL not given via positional argument, not given via flag either.
	// We prompt the user for a single input value.
	default:
		// fmt.Fprintln(os.Stderr, "default switch case triggered")

		input, err := readURLFromUser()
		if err != nil {
			return nil, fmt.Errorf("error reading URL: %w", err)
		}

		if strings.TrimSpace(input) == "" {
			return nil, ErrInvalidURL
		}

		inputURLs = append(inputURLs, cleanURL(input))
	}

	return inputURLs, nil
}

// cleanURL strips away quoting or escaping of characters in a given URL.
func cleanURL(s string) string {
	// Strip off any quoting that may be present.
	s = strings.ReplaceAll(s, `'`, "")
	s = strings.ReplaceAll(s, `"`, "")

	// Replace escaped ampersands with literal ampersands.
	// inputURL = strings.ReplaceAll(flag.Args()[1], "&amp;", "&")

	// Use html package to handle ampersand escaping *and* any edge cases that
	// I may be unaware of.
	s = html.UnescapeString(s)

	return s
}

// assertValidURLParameter requires that the given url.URL contains a
// non-empty parameter named url.
func assertValidURLParameter(u *url.URL) error {
	urlValues := u.Query()
	if urlValues.Get("url") == "" {
		return ErrOriginalURLNotResolved
	}

	return nil
}
