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
	"sort"
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
	case !cfg.Verbose:
		urlValues := safelink.Query()
		maskedURL := urlValues.Get("url")

		if maskedURL != "" {
			fmt.Printf("\nOriginal URL:\n\n%v\n", maskedURL)
			return
		}

		fmt.Println("Unable to resolve original URL")
		os.Exit(1)

	default:
		urlValues := safelink.Query()
		urlValues.Add("host", safelink.Host)
		keys := make([]string, 0, len(urlValues))
		for k := range urlValues {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		fmt.Printf("\nExpanded values from the given link:\n\n")

		for _, key := range keys {
			if len(urlValues[key]) > 0 {
				fmt.Printf("  %-10s: %s\n", key, urlValues[key][0])
			}
		}
	}

}
