// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package safelinks

import (
	"log"
	"net/url"
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
// Internal logic uses a regular expression to match URL patterns beginning
// with 'https://' and ending with a whitespace character.
func GetURLPatternsUsingRegex(input string) ([]FoundURLPattern, error) {
	// urls := make([]url.URL, 0, 5)

	urlPatterns := make([]FoundURLPattern, 0, 5)

	if !strings.Contains(input, SafeLinksURLRequiredPrefix) {
		return nil, ErrNoURLsFound
	}

	// This works but would match regular http:// prefixes:
	//
	// https://www.honeybadger.io/blog/a-definitive-guide-to-regular-expressions-in-go/
	// urlRegex := `https?://\S+|www\.\S+`

	urlRegex := SafeLinksURLRequiredPrefix + `\S+|www\.\S+`

	r := regexp.MustCompile(urlRegex)

	matches := r.FindAllString(input, -1)
	log.Println("Matches:", matches)

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
// beginning with 'https://' and ending with a whitespace character.
func GetURLPatternsUsingIndex(input string) ([]FoundURLPattern, error) {
	// urls := make([]url.URL, 0, 5)
	urlPatterns := make([]FoundURLPattern, 0, 5)

	if !strings.Contains(input, SafeLinksURLRequiredPrefix) {
		return nil, ErrNoURLsFound
	}

	remaining := input

	for {
		urlStart := strings.Index(remaining, SafeLinksURLRequiredPrefix)

		if urlStart == -1 {
			break
		}

		next := urlStart + len(SafeLinksURLRequiredPrefix) + 1

		// Sanity check to keep from indexing past remaining string length.
		if next >= len(remaining) {
			break
		}

		// Assume we found ending point until proven otherwise.
		// urlEnd := next

		// for _, char := range remaining[next:] {
		// 	if unicode.IsSpace(char) {
		// 		break // we found end of URL pattern
		// 	}
		// 	urlEnd++
		// }

		urlEnd := getURLIndexEndPosition(remaining[next:], next)

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
// match URL patterns beginning with 'https://' and ending with a whitespace
// character.
func GetURLPatternsUsingPrefixMatchingOnFields(input string) ([]FoundURLPattern, error) {
	urlPatterns := make([]FoundURLPattern, 0, 5)

	if !strings.Contains(input, SafeLinksURLRequiredPrefix) {
		return nil, ErrNoURLsFound
	}

	fields := strings.Fields(input)
	for _, field := range fields {
		if strings.HasPrefix(field, SafeLinksURLRequiredPrefix) {
			urlPatterns = append(
				urlPatterns,
				FoundURLPattern{
					URLPattern: field,
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

	// NOTE: Confirmed working:
	//
	// urlPatterns, err := GetURLPatternsUsingIndex(input)
	// urlPatterns, err := GetURLPatternsUsingPrefixMatchingOnFields(input)
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
	urls, err := URLs(input)
	if err != nil {
		return nil, err
	}

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
// iterates until it finds the first space character. This is assumed to be
// the separator used to indicate the end of a URL pattern.
func getURLIndexEndPosition(input string, startPos int) int {
	endPos := startPos

	for _, char := range input[startPos:] {
		if unicode.IsSpace(char) {
			break // we found end of URL pattern
		}
		endPos++
	}

	return endPos
}
