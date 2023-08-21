// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"io"
	"net/url"
	"sort"
)

// emitOutput emits a given URL to the specified output sink. If specified,
// verbose output is used.
func emitOutput(u *url.URL, w io.Writer, verbose bool) {
	switch {
	case verbose:
		verboseOutput(u, w)

	default:
		simpleOutput(u, w)
	}
}

// simpleOutput handles generating reduced or "simple" output when verbose
// mode is not invoked.
func simpleOutput(u *url.URL, w io.Writer) {
	urlValues := u.Query()
	maskedURL := urlValues.Get("url")

	fmt.Fprintln(w, maskedURL)
}

// verboseOutput handles generating extended or "verbose" output when
// requested.
func verboseOutput(u *url.URL, w io.Writer) {
	urlValues := u.Query()
	urlValues.Add("host", u.Host)

	keys := make([]string, 0, len(urlValues))
	for k := range urlValues {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Fprintf(w, "\nExpanded values from the given link:\n\n")

	for _, key := range keys {
		if len(urlValues[key]) > 0 {
			fmt.Fprintf(w, "  %-10s: %s\n", key, urlValues[key][0])
		}
	}
}
