// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Package safelinks_test provides test coverage for exported package
// functionality.
//
//nolint:dupl,gocognit // ignore "lines are duplicate of" and function complexity
package safelinks_test

import (
	_ "embed"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	"github.com/atc0005/safelinks/internal/safelinks"
	"github.com/google/go-cmp/cmp"
)

/*
// Package main assists with generating test var names
package main

import (
	"fmt"
	"strings"
)

func main() {
	list := []string{
		"email-with-angle-brackets-with-crlf-eol.txt",
		"email-with-angle-brackets-with-lf-eol.txt",
		"email-without-angle-brackets-with-crlf-eol.txt",
		"email-without-angle-brackets-with-lf-eol.txt",
		"list-of-broken-urls-with-crlf-eol.txt",
		"list-of-broken-urls-with-lf-eol.txt",
		"single-broken-url-with-crlf-eol.txt",
		"single-broken-url-with-lf-eol.txt",
	}

	for _, item := range list {
		item = strings.Replace(item, ".txt", "", -1)
		item = strings.Replace(item, "urls", "URLs", -1)
		item = strings.Replace(item, "url", "URL", -1)
		item = strings.Replace(item, "crlf", "CRLF", -1)
		item = strings.Replace(item, "lf", "LF", -1)
		item = strings.Title(item)
		item = strings.Replace(item, "-", "", -1)
		fmt.Println(item)
	}

}
*/

// The format used by the test files is VERY specific; trailing space plus
// newline patterns are intentional. Because "format on save" editor
// functionality easily breaks this explicitly formatted text, test input and output is stored in
// separate files to reduce test breakage due to editors "helping".
var (
	//go:embed testdata/output/embedded/decoded-email-with-angle-brackets-with-crlf-eol.txt
	outputDecodedEmailWithAngleBracketsWithCRLFEol string

	//go:embed testdata/output/embedded/decoded-email-with-angle-brackets-with-lf-eol.txt
	outputDecodedEmailWithAngleBracketsWithLFEol string

	//go:embed testdata/output/embedded/decoded-email-without-angle-brackets-with-crlf-eol.txt
	outputDecodedEmailWithoutAngleBracketsWithCRLFEol string

	//go:embed testdata/output/embedded/decoded-email-without-angle-brackets-with-lf-eol.txt
	outputDecodedEmailWithoutAngleBracketsWithLFEol string

	//go:embed testdata/output/standalone/decoded-list-of-urls-with-crlf-eol.txt
	outputDecodedListOfURLsWithCRLFEol string

	//go:embed testdata/output/standalone/decoded-list-of-urls-with-lf-eol.txt
	outputDecodedListOfURLsWithLFEol string

	//go:embed testdata/output/standalone/decoded-single-url-with-crlf-eol.txt
	outputDecodedSingleURLWithCRLFEol string

	//go:embed testdata/output/standalone/decoded-single-url-with-lf-eol.txt
	outputDecodedSingleURLWithLFEol string
)

// The format used by the test files is VERY specific; trailing space plus
// newline patterns are intentional. Because "format on save" editor
// functionality easily breaks this explicitly formatted text, test input and
// output is stored in separate files to reduce test breakage due to editors
// "helping".
var (
	//go:embed testdata/input/encoded-all/email-with-angle-brackets-with-crlf-eol.txt
	inputEncodedAllEmailWithAngleBracketsWithCRLFEol string

	//go:embed testdata/input/encoded-all/email-with-angle-brackets-with-lf-eol.txt
	inputEncodedAllEmailWithAngleBracketsWithLFEol string

	//go:embed testdata/input/encoded-all/email-without-angle-brackets-with-crlf-eol.txt
	inputEncodedAllEmailWithoutAngleBracketsWithCRLFEol string

	//go:embed testdata/input/encoded-all/email-without-angle-brackets-with-lf-eol.txt
	inputEncodedAllEmailWithoutAngleBracketsWithLFEol string

	//go:embed testdata/input/encoded-all/single-safelinks-url-with-crlf-eol.txt
	inputEncodedSingleSafelinksURLWithCRLFEol string

	//go:embed testdata/input/encoded-all/list-of-urls-with-crlf-eol.txt
	inputEncodedAllListOfURLsWithCRLFEol string

	//go:embed testdata/input/encoded-all/list-of-urls-with-lf-eol.txt
	inputEncodedAllListOfURLsWithLFEol string

	//go:embed testdata/input/encoded-all/single-safelinks-url-with-lf-eol.txt
	inputEncodedSingleSafelinksURLWithLFEol string
)

// The format used by the test files is VERY specific; trailing space plus
// newline patterns are intentional. Because "format on save" editor
// functionality easily breaks this explicitly formatted text, test input and
// output is stored in separate files to reduce test breakage due to editors
// "helping".
var (
	//go:embed testdata/input/encoded-mixed/email-with-angle-brackets-with-crlf-eol.txt
	inputEncodedMixedEmailWithAngleBracketsWithCRLFEol string

	//go:embed testdata/input/encoded-mixed/email-with-angle-brackets-with-lf-eol.txt
	inputEncodedMixedEmailWithAngleBracketsWithLFEol string

	//go:embed testdata/input/encoded-mixed/email-without-angle-brackets-with-crlf-eol.txt
	inputEncodedMixedEmailWithoutAngleBracketsWithCRLFEol string

	//go:embed testdata/input/encoded-mixed/email-without-angle-brackets-with-lf-eol.txt
	inputEncodedMixedEmailWithoutAngleBracketsWithLFEol string

	//go:embed testdata/input/encoded-mixed/list-of-urls-with-crlf-eol.txt
	inputEncodedMixedListOfURLsWithCRLFEol string

	//go:embed testdata/input/encoded-mixed/list-of-urls-with-lf-eol.txt
	inputEncodedMixedListOfURLsWithLFEol string
)

// All entries in these test files have patterns similar to URLs but are all
// considered "invalid" URLs.
//
// The format used by the test files is VERY specific; trailing space plus
// newline patterns are intentional. Because "format on save" editor
// functionality easily breaks this explicitly formatted text, test input and
// output is stored in separate files to reduce test breakage due to editors
// "helping".
var (
	//go:embed testdata/input/invalid-all/email-with-angle-brackets-with-crlf-eol.txt
	inputInvalidAllEmailWithAngleBracketsWithCRLFEol string

	//go:embed testdata/input/invalid-all/email-with-angle-brackets-with-lf-eol.txt
	inputInvalidAllEmailWithAngleBracketsWithLFEol string

	//go:embed testdata/input/invalid-all/email-without-angle-brackets-with-crlf-eol.txt
	inputInvalidAllEmailWithoutAngleBracketsWithCRLFEol string

	//go:embed testdata/input/invalid-all/email-without-angle-brackets-with-lf-eol.txt
	inputInvalidAllEmailWithoutAngleBracketsWithLFEol string

	//go:embed testdata/input/invalid-all/list-of-urls-with-crlf-eol.txt
	inputInvalidAllListOfURLsWithCRLFEol string

	//go:embed testdata/input/invalid-all/list-of-urls-with-lf-eol.txt
	inputInvalidAllListOfURLsWithLFEol string

	//go:embed testdata/input/invalid-all/single-url-with-crlf-eol.txt
	inputInvalidAllSingleURLWithCRLFEol string

	//go:embed testdata/input/invalid-all/single-url-with-lf-eol.txt
	inputInvalidAllSingleURLWithLFEol string
)

// Some entries in these test files have patterns similar to URLs but are
// considered "invalid". Other URLs within the files are valid.
//
// The format used by the test files is VERY specific; trailing space plus
// newline patterns are intentional. Because "format on save" editor
// functionality easily breaks this explicitly formatted text, test input and
// output is stored in separate files to reduce test breakage due to editors
// "helping".
var (
	//go:embed testdata/input/invalid-mixed/email-with-angle-brackets-with-crlf-eol.txt
	inputInvalidMixedEmailWithAngleBracketsWithCRLFEol string

	//go:embed testdata/input/invalid-mixed/email-with-angle-brackets-with-lf-eol.txt
	inputInvalidMixedEmailWithAngleBracketsWithLFEol string

	//go:embed testdata/input/invalid-mixed/email-without-angle-brackets-with-crlf-eol.txt
	inputInvalidMixedEmailWithoutAngleBracketsWithCRLFEol string

	//go:embed testdata/input/invalid-mixed/email-without-angle-brackets-with-lf-eol.txt
	inputInvalidMixedEmailWithoutAngleBracketsWithLFEol string

	//go:embed testdata/input/invalid-mixed/list-of-urls-with-crlf-eol.txt
	inputInvalidMixedListOfURLsWithCRLFEol string

	//go:embed testdata/input/invalid-mixed/list-of-urls-with-lf-eol.txt
	inputInvalidMixedListOfURLsWithLFEol string
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// The safelinks package disables "debug" logging by default. Here we have
	// the option of explicitly enabling that logging for troubleshooting
	// purposes.

	// log.SetOutput(os.Stderr)
	log.SetOutput(io.Discard)

	os.Exit(m.Run())
}

// TestURLsFindsAllValidURLs asserts that all valid URLs (encoded or
// unencoded) are found within given input.
func TestURLsFindsAllValidURLs(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input          string
		foundURLsCount int
	}{
		"Encoded list of URLs with LF EOL": {
			input:          inputEncodedAllListOfURLsWithLFEol,
			foundURLsCount: 15,
		},
		"Encoded list of URLs with CRLF EOL": {
			input:          inputEncodedAllListOfURLsWithCRLFEol,
			foundURLsCount: 15,
		},
		"Encoded single URL with LF EOL": {
			input:          inputEncodedSingleSafelinksURLWithLFEol,
			foundURLsCount: 1,
		},
		"Encoded single URL with CRLF EOL": {
			input:          inputEncodedSingleSafelinksURLWithCRLFEol,
			foundURLsCount: 1,
		},
		"Encoded email with angle brackets with CRLF EOL": {
			input:          inputEncodedAllEmailWithAngleBracketsWithCRLFEol,
			foundURLsCount: 5,
		},
		"Encoded email with angle brackets with LF EOL": {
			input:          inputEncodedAllEmailWithAngleBracketsWithLFEol,
			foundURLsCount: 5,
		},
		"Encoded email without angle brackets with CRLF EOL": {
			input:          inputEncodedAllEmailWithoutAngleBracketsWithCRLFEol,
			foundURLsCount: 5,
		},
		"Encoded email without angle brackets with LF EOL": {
			input:          inputEncodedAllEmailWithoutAngleBracketsWithLFEol,
			foundURLsCount: 5,
		},

		"Mixed encoded email with angle brackets with CRLF EOL": {
			input:          inputEncodedMixedEmailWithAngleBracketsWithCRLFEol,
			foundURLsCount: 5,
		},
		"Mixed encoded email with angle brackets with LF EOL": {
			input:          inputEncodedMixedEmailWithAngleBracketsWithLFEol,
			foundURLsCount: 5,
		},
		"Mixed encoded email without angle brackets with CRLF EOL": {
			input:          inputEncodedMixedEmailWithoutAngleBracketsWithCRLFEol,
			foundURLsCount: 5,
		},
		"Mixed encoded email without angle brackets with LF EOL": {
			input:          inputEncodedMixedEmailWithoutAngleBracketsWithLFEol,
			foundURLsCount: 5,
		},
		"Mixed encoded list of URLs with CRLF EOL": {
			input:          inputEncodedMixedListOfURLsWithCRLFEol,
			foundURLsCount: 15,
		},
		"Mixed encoded list of URLs with LF EOL": {
			input:          inputEncodedMixedListOfURLsWithLFEol,
			foundURLsCount: 15,
		},

		"Mixed invalid list of URLs with LF EOL": {
			input:          inputInvalidMixedListOfURLsWithLFEol,
			foundURLsCount: 7,
		},
		"Mixed invalid list of URLs with CRLF EOL": {
			input:          inputInvalidMixedListOfURLsWithCRLFEol,
			foundURLsCount: 7,
		},
		"Mixed invalid URLs email with angle brackets with CRLF EOL": {
			input:          inputInvalidMixedEmailWithAngleBracketsWithCRLFEol,
			foundURLsCount: 2,
		},
		"Mixed invalid URLs email with with angle brackets with LF EOL": {
			input:          inputInvalidMixedEmailWithAngleBracketsWithLFEol,
			foundURLsCount: 2,
		},
		"Mixed invalid URLs email without angle brackets with CRLF EOL": {
			input:          inputInvalidMixedEmailWithoutAngleBracketsWithCRLFEol,
			foundURLsCount: 2,
		},
		"Mixed invalid URLs email without angle brackets with LF EOL": {
			input:          inputInvalidMixedEmailWithoutAngleBracketsWithLFEol,
			foundURLsCount: 2,
		},
	}

	for name, tt := range tests {
		// Guard against referencing the loop iterator variable directly.
		//
		// https://stackoverflow.com/questions/68559574/using-the-variable-on-range-scope-x-in-function-literal-scopelint
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tt := tt

		t.Run(name, func(t *testing.T) {
			want := tt.foundURLsCount
			urls, err := safelinks.URLs(tt.input, true)
			// t.Log("tt.input:", tt.input)
			if err != nil {
				t.Fatalf("failed to find URLs in input: %v", err)
			}

			got := len(urls)

			if want != got {
				t.Errorf("\nwant %d\ngot %d", want, got)
			} else {
				t.Logf("OK: Found expected number of URLs within given input.")
			}
		})
	}

}

// TestURLsFailsForInvalidURLs asserts that no URLs (encoded or unencoded) are
// found within given invalid input.
func TestURLsFailsForInvalidURLs(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
	}{
		"All invalid list of URLs with LF EOL": {
			input: inputInvalidAllListOfURLsWithLFEol,
		},
		"All invalid list of URLs with CRLF EOL": {
			input: inputInvalidAllListOfURLsWithCRLFEol,
		},
		"All invalid email with angle brackets with LF EOL": {
			input: inputInvalidAllEmailWithAngleBracketsWithLFEol,
		},
		"All invalid email with angle brackets with CRLF EOL": {
			input: inputInvalidAllEmailWithAngleBracketsWithCRLFEol,
		},
		"All invalid email without angle brackets with LF EOL": {
			input: inputInvalidAllEmailWithoutAngleBracketsWithLFEol,
		},
		"All invalid email without angle brackets with CRLF EOL": {
			input: inputInvalidAllEmailWithoutAngleBracketsWithCRLFEol,
		},
		"Single invalid URL with LF EOL": {
			input: inputInvalidAllSingleURLWithLFEol,
		},
		"Single invalid URL with CRLF EOL": {
			input: inputInvalidAllSingleURLWithCRLFEol,
		},
	}

	for name, tt := range tests {
		// Guard against referencing the loop iterator variable directly.
		//
		// https://stackoverflow.com/questions/68559574/using-the-variable-on-range-scope-x-in-function-literal-scopelint
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := safelinks.URLs(tt.input, true)
			if err == nil {
				t.Logf("result (%d matches): %+v:", len(result), result)
				t.Fatalf(
					"\nwant error when searching for valid URLs within input\n"+
						"got %d successful URL matches for invalid input",
					len(result),
				)
			} else {
				t.Logf("result: %+v:", result)
				t.Logf(
					"OK: \nwant error when searching for valid URLs within input\ngot error as expected: %v", err,
				)
			}
		})
	}

}

func TestFilterURLsCorrectlyFiltersByType(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input                      string
		foundEncodedLinksURLsCount int
		foundUnencodedURLsCount    int
	}{
		"Mixed encoded email with angle brackets with CRLF EOL": {
			input:                      inputEncodedMixedEmailWithAngleBracketsWithCRLFEol,
			foundEncodedLinksURLsCount: 2,
			foundUnencodedURLsCount:    3,
		},
		"Mixed encoded email with angle brackets with LF EOL": {
			input:                      inputEncodedMixedEmailWithAngleBracketsWithLFEol,
			foundEncodedLinksURLsCount: 2,
			foundUnencodedURLsCount:    3,
		},
		"Mixed encoded email without angle brackets with CRLF EOL": {
			input:                      inputEncodedMixedEmailWithoutAngleBracketsWithCRLFEol,
			foundEncodedLinksURLsCount: 2,
			foundUnencodedURLsCount:    3,
		},
		"Mixed encoded email without angle brackets with LF EOL": {
			input:                      inputEncodedMixedEmailWithoutAngleBracketsWithLFEol,
			foundEncodedLinksURLsCount: 2,
			foundUnencodedURLsCount:    3,
		},
		"Mixed encoded list of URLs with CRLF EOL": {
			input:                      inputEncodedMixedListOfURLsWithCRLFEol,
			foundEncodedLinksURLsCount: 8,
			foundUnencodedURLsCount:    7,
		},
		"Mixed encoded list of URLs with LF EOL": {
			input:                      inputEncodedMixedListOfURLsWithLFEol,
			foundEncodedLinksURLsCount: 8,
			foundUnencodedURLsCount:    7,
		},

		"Encoded email with angle brackets with CRLF EOL": {
			input:                      inputEncodedAllEmailWithAngleBracketsWithCRLFEol,
			foundEncodedLinksURLsCount: 5,
			foundUnencodedURLsCount:    0,
		},
		"Encoded email with angle brackets with LF EOL": {
			input:                      inputEncodedAllEmailWithAngleBracketsWithLFEol,
			foundEncodedLinksURLsCount: 5,
			foundUnencodedURLsCount:    0,
		},
		"Encoded email without angle brackets with CRLF EOL": {
			input:                      inputEncodedAllEmailWithoutAngleBracketsWithCRLFEol,
			foundEncodedLinksURLsCount: 5,
			foundUnencodedURLsCount:    0,
		},
		"Encoded email without angle brackets with LF EOL": {
			input:                      inputEncodedAllEmailWithoutAngleBracketsWithLFEol,
			foundEncodedLinksURLsCount: 5,
			foundUnencodedURLsCount:    0,
		},
		"Encoded list of URLs with CRLF EOL": {
			input:                      inputEncodedAllListOfURLsWithCRLFEol,
			foundEncodedLinksURLsCount: 15,
			foundUnencodedURLsCount:    0,
		},
		"Encoded list of URLs with LF EOL": {
			input:                      inputEncodedAllListOfURLsWithLFEol,
			foundEncodedLinksURLsCount: 15,
			foundUnencodedURLsCount:    0,
		},
		"Encoded single URL with CRLF EOL": {
			input:                      inputEncodedSingleSafelinksURLWithCRLFEol,
			foundEncodedLinksURLsCount: 1,
			foundUnencodedURLsCount:    0,
		},
		"Encoded single URL with LF EOL": {
			input:                      inputEncodedSingleSafelinksURLWithLFEol,
			foundEncodedLinksURLsCount: 1,
			foundUnencodedURLsCount:    0,
		},

		"Unencoded list of URLs with LF EOL": {
			input:                      outputDecodedListOfURLsWithLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    15,
		},
		"Unencoded list of URLs with CRLF EOL": {
			input:                      outputDecodedListOfURLsWithCRLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    15,
		},
		"Unencoded single URL with LF EOL": {
			input:                      outputDecodedSingleURLWithLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    1,
		},
		"Unencoded single URL with CRLF EOL": {
			input:                      outputDecodedSingleURLWithCRLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    1,
		},
		"Unencoded email with angle brackets with LF EOL": {
			input:                      outputDecodedEmailWithAngleBracketsWithLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    5,
		},
		"Unencoded email with angle brackets with CRLF EOL": {
			input:                      outputDecodedEmailWithAngleBracketsWithCRLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    5,
		},
		"Unencoded email without angle brackets with LF EOL": {
			input:                      outputDecodedEmailWithoutAngleBracketsWithLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    5,
		},
		"Unencoded email without angle brackets with CRLF EOL": {
			input:                      outputDecodedEmailWithoutAngleBracketsWithCRLFEol,
			foundEncodedLinksURLsCount: 0,
			foundUnencodedURLsCount:    5,
		},
	}

	for name, tt := range tests {
		// Guard against referencing the loop iterator variable directly.
		//
		// https://stackoverflow.com/questions/68559574/using-the-variable-on-range-scope-x-in-function-literal-scopelint
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tt := tt

		t.Run(name, func(t *testing.T) {
			allURLs, err := safelinks.URLs(tt.input, true)
			if err != nil {
				t.Fatalf("failed to find URLs in input: %v", err)
			}

			encodedURLs := safelinks.FilterParsedURLs(allURLs, false)
			wantEncodedURLsCount := tt.foundEncodedLinksURLsCount
			gotEncodedURLsCount := len(encodedURLs)

			if wantEncodedURLsCount != gotEncodedURLsCount {
				t.Errorf("\nwant %d Safe Links URLs\ngot %d Safe Links URLs", wantEncodedURLsCount, gotEncodedURLsCount)
			} else {
				t.Logf("OK: Found expected number (%d) of Safe Links URLs within given input.", wantEncodedURLsCount)
			}

			unencodedURLs := safelinks.FilterParsedURLs(allURLs, true)
			wantUnencodedURLsCount := tt.foundUnencodedURLsCount
			gotUnencodedURLsCount := len(unencodedURLs)

			if wantUnencodedURLsCount != gotUnencodedURLsCount {
				t.Errorf("\nwant %d unencoded URLs\ngot %d unencoded URLs", wantUnencodedURLsCount, gotUnencodedURLsCount)
			} else {
				t.Logf("OK: Found expected number (%d) of unencoded URLs within given input.", wantUnencodedURLsCount)
			}

		})
	}

}

// TestSafeLinkURLsFindsAllValidSafeLinks evaluates mixed encoded email input
// for this test and matches on encoded URLs, not counting unencoded URLs.
func TestSafeLinkURLsFindsAllValidSafeLinks(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input                 string
		foundEncodedURLsCount int
	}{
		"Mixed invalid list of URLs with LF EOL": {
			input:                 inputInvalidMixedListOfURLsWithLFEol,
			foundEncodedURLsCount: 0,
		},
		"Mixed invalid list of URLs with CRLF EOL": {
			input:                 inputInvalidMixedListOfURLsWithCRLFEol,
			foundEncodedURLsCount: 0,
		},
		"Mixed invalid URLs email with angle brackets with CRLF EOL": {
			input:                 inputInvalidMixedEmailWithAngleBracketsWithCRLFEol,
			foundEncodedURLsCount: 0,
		},
		"Mixed invalid URLs email with with angle brackets with LF EOL": {
			input:                 inputInvalidMixedEmailWithAngleBracketsWithLFEol,
			foundEncodedURLsCount: 0,
		},
		"Mixed invalid URLs email without angle brackets with CRLF EOL": {
			input:                 inputInvalidMixedEmailWithoutAngleBracketsWithCRLFEol,
			foundEncodedURLsCount: 0,
		},
		"Mixed invalid URLs email without angle brackets with LF EOL": {
			input:                 inputInvalidMixedEmailWithoutAngleBracketsWithLFEol,
			foundEncodedURLsCount: 0,
		},

		"Encoded email with angle brackets with CRLF EOL": {
			input:                 inputEncodedAllEmailWithAngleBracketsWithCRLFEol,
			foundEncodedURLsCount: 5,
		},
		"Encoded email with angle brackets with LF EOL": {
			input:                 inputEncodedAllEmailWithAngleBracketsWithLFEol,
			foundEncodedURLsCount: 5,
		},
		"Encoded email without angle brackets with CRLF EOL": {
			input:                 inputEncodedAllEmailWithoutAngleBracketsWithCRLFEol,
			foundEncodedURLsCount: 5,
		},
		"Encoded email without angle brackets with LF EOL": {
			input:                 inputEncodedAllEmailWithoutAngleBracketsWithLFEol,
			foundEncodedURLsCount: 5,
		},
		"Encoded list of URLs with CRLF EOL": {
			input:                 inputEncodedAllListOfURLsWithCRLFEol,
			foundEncodedURLsCount: 15,
		},
		"Encoded list of URLs with LF EOL": {
			input:                 inputEncodedAllListOfURLsWithLFEol,
			foundEncodedURLsCount: 15,
		},
		"Encoded single URL with CRLF EOL": {
			input:                 inputEncodedSingleSafelinksURLWithCRLFEol,
			foundEncodedURLsCount: 1,
		},
		"Encoded single URL with LF EOL": {
			input:                 inputEncodedSingleSafelinksURLWithLFEol,
			foundEncodedURLsCount: 1,
		},

		"Mixed encoded email with angle brackets with CRLF EOL": {
			input:                 inputEncodedMixedEmailWithAngleBracketsWithCRLFEol,
			foundEncodedURLsCount: 2,
		},
		"Mixed encoded email with angle brackets with LF EOL": {
			input:                 inputEncodedMixedEmailWithAngleBracketsWithLFEol,
			foundEncodedURLsCount: 2,
		},
		"Mixed encoded email without angle brackets with CRLF EOL": {
			input:                 inputEncodedMixedEmailWithoutAngleBracketsWithCRLFEol,
			foundEncodedURLsCount: 2,
		},
		"Mixed encoded email without angle brackets with LF EOL": {
			input:                 inputEncodedMixedEmailWithoutAngleBracketsWithLFEol,
			foundEncodedURLsCount: 2,
		},
		"Mixed encoded list of URLs with CRLF EOL": {
			input:                 inputEncodedMixedListOfURLsWithCRLFEol,
			foundEncodedURLsCount: 8,
		},
		"Mixed encoded list of URLs with LF EOL": {
			input:                 inputEncodedMixedListOfURLsWithLFEol,
			foundEncodedURLsCount: 8,
		},
	}

	for name, tt := range tests {
		// Guard against referencing the loop iterator variable directly.
		//
		// https://stackoverflow.com/questions/68559574/using-the-variable-on-range-scope-x-in-function-literal-scopelint
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tt := tt

		t.Run(name, func(t *testing.T) {
			want := tt.foundEncodedURLsCount
			urls, err := safelinks.SafeLinkURLs(tt.input)

			switch {
			case err == nil:

			case errors.Is(err, safelinks.ErrNoSafeLinkURLsFound):
				t.Log("Ignore ErrNoSafeLinkURLsFound error so that " +
					"we can evaluate the results directly later")

			default:
				t.Fatalf("failed to find Safe Links URLs in input: %v", err)
			}

			got := len(urls)

			if want != got {
				t.Errorf("\nwant %d\ngot %d", want, got)
			} else {
				t.Logf("OK: Found expected number of URLs within given input.")
			}
		})
	}

}

// TestDecodeInputSucceedsForValidInput asserts that a given collection of
// encoded URLs is decoded as a specific collection of URLs with the same line
// ending.
func TestDecodeInputSucceedsForValidInput(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input  string
		result string
	}{
		"Encoded list of URLs with LF EOL": {
			input:  inputEncodedAllListOfURLsWithLFEol,
			result: outputDecodedListOfURLsWithLFEol,
		},
		"Encoded list of URLs with CRLF EOL": {
			input:  inputEncodedAllListOfURLsWithCRLFEol,
			result: outputDecodedListOfURLsWithCRLFEol,
		},
		"Encoded single URL with LF EOL": {
			input:  inputEncodedSingleSafelinksURLWithLFEol,
			result: outputDecodedSingleURLWithLFEol,
		},
		"Encoded single URL with CRLF EOL": {
			input:  inputEncodedSingleSafelinksURLWithCRLFEol,
			result: outputDecodedSingleURLWithCRLFEol,
		},
		"Encoded email with angle brackets with CRLF EOL": {
			input:  inputEncodedAllEmailWithAngleBracketsWithCRLFEol,
			result: outputDecodedEmailWithAngleBracketsWithCRLFEol,
		},
		"Encoded email with angle brackets with LF EOL": {
			input:  inputEncodedAllEmailWithAngleBracketsWithLFEol,
			result: outputDecodedEmailWithAngleBracketsWithLFEol,
		},
		"Encoded email without angle brackets with CRLF EOL": {
			input:  inputEncodedAllEmailWithoutAngleBracketsWithCRLFEol,
			result: outputDecodedEmailWithoutAngleBracketsWithCRLFEol,
		},
		"Encoded email without angle brackets with LF EOL": {
			input:  inputEncodedAllEmailWithoutAngleBracketsWithLFEol,
			result: outputDecodedEmailWithoutAngleBracketsWithLFEol,
		},
		"Mixed encoded list of URLs with LF EOL": {
			input:  inputEncodedMixedListOfURLsWithLFEol,
			result: outputDecodedListOfURLsWithLFEol,
		},
		"Mixed encoded list of URLs with CRLF EOL": {
			input:  inputEncodedMixedListOfURLsWithCRLFEol,
			result: outputDecodedListOfURLsWithCRLFEol,
		},
		"Mixed encoded email with angle brackets with CRLF EOL": {
			input:  inputEncodedMixedEmailWithAngleBracketsWithCRLFEol,
			result: outputDecodedEmailWithAngleBracketsWithCRLFEol,
		},
		"Mixed encoded email with angle brackets with LF EOL": {
			input:  inputEncodedMixedEmailWithAngleBracketsWithLFEol,
			result: outputDecodedEmailWithAngleBracketsWithLFEol,
		},
		"Mixed encoded email without angle brackets with CRLF EOL": {
			input:  inputEncodedMixedEmailWithoutAngleBracketsWithCRLFEol,
			result: outputDecodedEmailWithoutAngleBracketsWithCRLFEol,
		},
		"Mixed encoded email without angle brackets with LF EOL": {
			input:  inputEncodedMixedEmailWithoutAngleBracketsWithLFEol,
			result: outputDecodedEmailWithoutAngleBracketsWithLFEol,
		},
	}

	for name, tt := range tests {
		// Guard against referencing the loop iterator variable directly.
		//
		// https://stackoverflow.com/questions/68559574/using-the-variable-on-range-scope-x-in-function-literal-scopelint
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tt := tt

		t.Run(name, func(t *testing.T) {
			want := tt.result
			got, err := safelinks.DecodeInput(tt.input)
			if err != nil {
				t.Fatalf("failed to decode input: %v", err)
			}

			if d := cmp.Diff(want, got); d != "" {
				t.Errorf("(-want, +got)\n:%s", d)
			}
		})
	}

}

// TestDecodeInputFailsForInvalidInput asserts that a given collection of
// unencoded, malformed or other problematic URLs fail to decode.
func TestDecodeInputFailsForInvalidInput(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
	}{
		"Unencoded list of URLs with LF EOL": {
			input: outputDecodedListOfURLsWithLFEol,
		},
		"Unencoded list of URLs with CRLF EOL": {
			input: outputDecodedListOfURLsWithCRLFEol,
		},
		"Unencoded single URL with LF EOL": {
			input: outputDecodedSingleURLWithLFEol,
		},
		"Unencoded single URL with CRLF EOL": {
			input: outputDecodedSingleURLWithCRLFEol,
		},
		"Unencoded email with angle brackets with LF EOL": {
			input: outputDecodedEmailWithAngleBracketsWithLFEol,
		},
		"Unencoded email with angle brackets with CRLF EOL": {
			input: outputDecodedEmailWithAngleBracketsWithCRLFEol,
		},
		"Unencoded email without angle brackets with LF EOL": {
			input: outputDecodedEmailWithoutAngleBracketsWithLFEol,
		},
		"Unencoded email without angle brackets with CRLF EOL": {
			input: outputDecodedEmailWithoutAngleBracketsWithCRLFEol,
		},

		"Mixed invalid list of URLs with LF EOL": {
			input: inputInvalidMixedListOfURLsWithLFEol,
		},
		"Mixed invalid list of URLs with CRLF EOL": {
			input: inputInvalidMixedListOfURLsWithCRLFEol,
		},
		"Mixed invalid email with angle brackets with LF EOL": {
			input: inputInvalidMixedEmailWithAngleBracketsWithLFEol,
		},
		"Mixed invalid email with angle brackets with CRLF EOL": {
			input: inputInvalidMixedEmailWithAngleBracketsWithCRLFEol,
		},
		"Mixed invalid email without angle brackets with LF EOL": {
			input: inputInvalidMixedEmailWithoutAngleBracketsWithLFEol,
		},
		"Mixed invalid email without angle brackets with CRLF EOL": {
			input: inputInvalidMixedEmailWithoutAngleBracketsWithCRLFEol,
		},

		"All invalid list of URLs with LF EOL": {
			input: inputInvalidAllListOfURLsWithLFEol,
		},
		"All invalid list of URLs with CRLF EOL": {
			input: inputInvalidAllListOfURLsWithCRLFEol,
		},
		"All invalid email with angle brackets with LF EOL": {
			input: inputInvalidAllEmailWithAngleBracketsWithLFEol,
		},
		"All invalid email with angle brackets with CRLF EOL": {
			input: inputInvalidAllEmailWithAngleBracketsWithCRLFEol,
		},
		"All invalid email without angle brackets with LF EOL": {
			input: inputInvalidAllEmailWithoutAngleBracketsWithLFEol,
		},
		"All invalid email without angle brackets with CRLF EOL": {
			input: inputInvalidAllEmailWithoutAngleBracketsWithCRLFEol,
		},
		"Single invalid URL with LF EOL": {
			input: inputInvalidAllSingleURLWithLFEol,
		},
		"Single invalid URL with CRLF EOL": {
			input: inputInvalidAllSingleURLWithCRLFEol,
		},
	}

	for name, tt := range tests {
		// Guard against referencing the loop iterator variable directly.
		//
		// https://stackoverflow.com/questions/68559574/using-the-variable-on-range-scope-x-in-function-literal-scopelint
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := safelinks.DecodeInput(tt.input)
			if err == nil {
				t.Logf("result: %+v:", result)
				t.Fatalf(
					"\nwant error when decoding invalid input\ngot successful decoding result",
				)
			} else {
				t.Logf("result: %+v:", result)
				t.Logf(
					"OK: \nwant error when decoding invalid input\ngot error as expected: %v", err,
				)
			}

		})
	}
}
