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

// parseInputURL processes an input string via flag or stdin as a URL value.
// The URL is unescaped and quoting removed.
func parseInputURL(inputURL string) string {
	switch {

	// We received a URL via positional argument.
	case len(flag.Args()) > 0:

		if strings.TrimSpace(flag.Args()[0]) == "" {
			fmt.Println("Invalid URL provided.")
			os.Exit(1)
		}

		inputURL = cleanURL(flag.Args()[0])

	// We received a URL via flag.
	case inputURL != "":
		inputURL = cleanURL(flag.Args()[1])

	// Input URL not given via positional argument, not given via flag either.
	default:
		input, err := readURLFromUser()
		if err != nil {
			fmt.Println("Error reading URL:", err)
			os.Exit(1)
		}

		if strings.TrimSpace(input) == "" {
			fmt.Println("Invalid URL provided.")
			os.Exit(1)
		}

		inputURL = cleanURL(input)
	}

	return inputURL
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
