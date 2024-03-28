// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package safelinks

import (
	"bufio"
	"flag"
	"fmt"
	"html"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// SafeLinksURLRequiredPrefix is the required prefix for all Safe Links URLs.
const SafeLinksURLRequiredPrefix = "https://"

// FoundURLPattern is an unvalidated URL pattern match found in given input.
type FoundURLPattern struct {
	// Input      *string
	startPosition int
	endPosition   int
	URLPattern    string
}

// SafeLinkURL contains the encoded and decoded URLs for a matched Safe Link.
type SafeLinkURL struct {
	EncodedURL string
	DecodedURL string
}

// ValidURL attempts to validate whether a given input string is a valid URL.
func ValidURL(input string) bool {
	if _, err := url.Parse(input); err != nil {
		return false
	}

	return true
}

// ValidSafeLinkURL validates whether a given url.URL is a valid Safe Links
// URL.
func ValidSafeLinkURL(input *url.URL) bool {
	if err := assertValidURLParameter(input); err != nil {
		return false
	}

	return true
}

// ReadURLFromUser attempts to read a given URL pattern from the user via
// stdin prompt.
func ReadURLFromUser() (string, error) {
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

// ReadURLsFromFile attempts to read URL patterns from a given file
// (io.Reader).
//
// The collection of input URLs is returned or an error if one occurs.
func ReadURLsFromFile(r io.Reader) ([]string, error) {
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

// ProcessInputAsURL processes a given input string as a URL value. This
// input string represents a single URL given via CLI flag.
//
// If an input string is not provided, this function will attempt to read
// input URLs from stdin. Each input URL is unescaped and quoting removed.
//
// The collection of input URLs is returned or an error if one occurs.
func ProcessInputAsURL(inputURL string) ([]string, error) {
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
		return ReadURLsFromFile(os.Stdin)

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

		input, err := ReadURLFromUser()
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

	// Strip of potential enclosing angle brackets.
	s = strings.Trim(s, `<>`)

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
		safelink, err := url.Parse(inputURL)
		if err != nil {
			fmt.Printf("Failed to parse URL: %v\n", err)

			errEncountered = true
			continue
		}

		if err := assertValidURLParameter(safelink); err != nil {
			fmt.Fprintf(errOut, "Invalid Safelinks URL %q: %v\n", safelink, err)

			errEncountered = true
			continue
		}

		emitOutput(safelink, okOut, verbose)
	}

	return errEncountered
}

// GetURLPatternsUsingRegex parses the given input and returns a collection of
// FoundURLPattern values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URL patterns begin with that protocol scheme. nil
// is returned if no patterns using that scheme are found.
//
// NOTE: Validation is not performed to ensure that matched patterns are valid
// URLs.
//
// Internal logic uses a regular expression to match URL patterns optionally
// beginning with a left angle bracket, then 'https://' and ending with a
// whitespace character or a right angle bracket. Any angle brackets present
// are trimmed from returned matches.
// Internal logic uses a regular expression to match URL patterns optionally
// beginning with a left angle bracket, then 'https://' and ending with a
// whitespace character or a right angle bracket. Any angle brackets present
// are trimmed from returned matches.
func GetURLPatternsUsingRegex(input string) ([]FoundURLPattern, error) {
	urlPatterns := make([]FoundURLPattern, 0, 5)

	if !strings.Contains(input, SafeLinksURLRequiredPrefix) {
		return nil, ErrNoURLsFound
	}

	urlRegex := `<?` + SafeLinksURLRequiredPrefix + `\S+>?`

	r := regexp.MustCompile(urlRegex)

	matches := r.FindAllString(input, -1)
	log.Printf("Matches: %d\n", len(matches))
	for _, up := range matches {
		log.Println(up)
	}

	log.Println("Cleaning URLs of enclosing angle brackets")
	for i := range matches {
		matches[i] = strings.Trim(matches[i], "<>")
	}
	log.Printf("Matches (%d) trimmed:", len(matches))
	for _, up := range matches {
		log.Println(up)
	}

	for _, match := range matches {
		urlPatterns = append(
			urlPatterns,
			FoundURLPattern{
				URLPattern: match,
			},
		)
	}

	return urlPatterns, nil
}

// GetURLPatternsUsingIndex parses the given input and returns a collection of
// FoundURLPattern values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URL patterns begin with that protocol scheme. nil
// is returned if no patterns using that scheme are found.
//
// NOTE: Validation has not been performed to ensure that matched patterns are
// valid URLs.
//
// Internal logic uses slice indexing/iteration to match URL patterns
// beginning with 'https://' and ending with a whitespace character or a right
// angle bracket. Any angle brackets present are trimmed from returned
// matches.
func GetURLPatternsUsingIndex(input string) ([]FoundURLPattern, error) {
	// urls := make([]url.URL, 0, 5)
	urlPatterns := make([]FoundURLPattern, 0, 5)

	if !strings.Contains(input, SafeLinksURLRequiredPrefix) {
		return nil, ErrNoURLsFound
	}

	remaining := input

	for {
		urlStart := strings.Index(remaining, SafeLinksURLRequiredPrefix)
		log.Println("urlStart:", urlStart)

		if urlStart == -1 {
			break
		}

		next := urlStart + len(SafeLinksURLRequiredPrefix) + 1

		// Sanity check to keep from indexing past remaining string length.
		if next >= len(remaining) {
			break
		}

		urlEnd := getURLIndexEndPosition(remaining, next)

		urlPatterns = append(
			urlPatterns,
			FoundURLPattern{
				// recording for later potential debugging
				startPosition: urlStart,
				endPosition:   urlEnd,

				URLPattern: remaining[urlStart:urlEnd],
			},
		)

		// Abort further processing if we're at the end of our original input
		// string.
		if urlEnd+1 >= len(input) {
			break
		}

		// Otherwise,  record the next position as the starting point for
		// further URL match evaluation.
		remaining = remaining[urlEnd+1:]

	}

	log.Printf("Total URL pattern matches: %d", len(urlPatterns))
	for _, up := range urlPatterns {
		log.Println(up.URLPattern)
	}

	return urlPatterns, nil
}

// GetURLPatternsUsingPrefixMatchingOnFields parses the given input and
// returns a collection of FoundURLPattern values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URL patterns begin with that protocol scheme. nil
// is returned if no patterns using that scheme are found.
//
// NOTE: Validation has not been performed to ensure that matched patterns are
// valid URLs.
//
// Internal logic uses string splitting on whitespace and prefix matching to
// match URL patterns optionally beginning a left angle bracket, then
// 'https://' and ending with a whitespace character.
func GetURLPatternsUsingPrefixMatchingOnFields(input string) ([]FoundURLPattern, error) {
	urlPatterns := make([]FoundURLPattern, 0, 5)

	if !strings.Contains(input, SafeLinksURLRequiredPrefix) {
		return nil, ErrNoURLsFound
	}

	fields := strings.Fields(input)
	for _, field := range fields {
		switch {
		case strings.HasPrefix(field, SafeLinksURLRequiredPrefix):
			urlPatterns = append(
				urlPatterns,
				FoundURLPattern{
					URLPattern: field,
				},
			)

		case strings.HasPrefix(field, "<"+SafeLinksURLRequiredPrefix):
			urlPatterns = append(
				urlPatterns,
				FoundURLPattern{
					URLPattern: strings.Trim(field, "<>"),
				},
			)
		}
	}

	if len(urlPatterns) == 0 {
		return nil, ErrNoURLsFound
	}

	return urlPatterns, nil
}

// URLs parses the given input and returns a collection of *url.URL values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URLs begin with that protocol scheme. nil is
// returned if no valid URLs using that scheme are found.
func URLs(input string) ([]*url.URL, error) {
	urls := make([]*url.URL, 0, 5)

	// NOTE: Confirmed working with either of:
	//
	// - whitespace delimited URLs
	// - angle bracket delimited URLs (with or without leading/trailing
	//   whitespace)
	//
	// GetURLPatternsUsingPrefixMatchingOnFields does not support matching on
	// URL patterns without a leading space, but GetURLPatternsUsingRegex and
	// GetURLPatternsUsingIndex do.
	//
	// urlPatterns, err := GetURLPatternsUsingPrefixMatchingOnFields(input)
	// urlPatterns, err := GetURLPatternsUsingIndex(input)

	log.Println("Calling GetURLPatternsUsingRegex(input)")
	urlPatterns, err := GetURLPatternsUsingRegex(input)

	if err != nil {
		return nil, err
	}

	for _, pattern := range urlPatterns {
		u, err := url.Parse(pattern.URLPattern)

		if err != nil {
			continue
		}
		urls = append(urls, u)
	}

	return urls, nil
}

// SafeLinkURLsFromURLs evaluates a given collection of URLs and returns any
// that are found to be encoded as Safe Links. Deduplication is *not*
// performed. An error is returned if no valid matches are found.
func SafeLinkURLsFromURLs(urls []*url.URL) ([]SafeLinkURL, error) {
	safeLinkURLs := make([]SafeLinkURL, 0, len(urls))

	for _, u := range urls {
		if !ValidSafeLinkURL(u) {
			continue
		}

		originalURL := u.Query().Get("url")

		safeLinkURLs = append(
			safeLinkURLs,
			SafeLinkURL{
				EncodedURL: u.String(),
				// DecodedURL: originalURL,
				DecodedURL: cleanURL(originalURL),
			},
		)
	}

	if len(safeLinkURLs) == 0 {
		return nil, ErrNoSafeLinkURLsFound
	}

	return safeLinkURLs, nil
}

// SafeLinkURLs parses the given input and returns a collection of parsed and
// decoded URLs. Deduplication is *not* performed.
//
// An error is returned if no valid matches are found.
func SafeLinkURLs(input string) ([]SafeLinkURL, error) {
	log.Println("Calling URLs(input)")
	urls, err := URLs(input)
	if err != nil {
		return nil, err
	}

	log.Println("Calling SafeLinkURLsFromURLs(urls)")
	return SafeLinkURLsFromURLs(urls)
}

// FromURLs evaluates a given collection of URLs and returns a collection of
// SafeLinkURL values for any that are found to be encoded as Safe Links.
// Deduplication is *not* performed.
//
// An error is returned if no valid matches are found.
func FromURLs(urls []*url.URL) ([]SafeLinkURL, error) {
	return SafeLinkURLsFromURLs(urls)
}

// getURLIndexEndPosition accepts an input string and a starting position and
// iterates until it finds the first space character or the first right angle
// bracket. Either is assumed to be the separator used to indicate the end of
// a URL pattern.
func getURLIndexEndPosition(input string, startPos int) int {
	endPos := startPos

	for _, char := range input[startPos:] {
		if unicode.IsSpace(char) || char == '>' {
			log.Printf("char %q caused us to break", char)
			break // we found end of URL pattern
		}
		endPos++
	}

	return endPos
}
