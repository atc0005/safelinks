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

// processInputAsURL processes a given input string as a URL value. If not
// provided, this function will attempt to read the input URL from the first
// positional argument. The URL is unescaped and quoting removed.
func processInputAsURL(inputURL string) (string, error) {
	switch {

	// We received a URL via positional argument.
	case len(flag.Args()) > 0:

		if strings.TrimSpace(flag.Args()[0]) == "" {
			return "", ErrInvalidURL
		}

		inputURL = cleanURL(flag.Args()[0])

	// We received a URL via flag.
	case inputURL != "":
		inputURL = cleanURL(inputURL)

	// Input URL not given via positional argument, not given via flag either.
	default:
		input, err := readURLFromUser()
		if err != nil {
			return "", fmt.Errorf("error reading URL: %w", err)
		}

		if strings.TrimSpace(input) == "" {
			return "", ErrInvalidURL
		}

		inputURL = cleanURL(input)
	}

	return inputURL, nil
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
