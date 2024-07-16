// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package safelinks

import (
	"bufio"
	crand "crypto/rand"
	"fmt"
	"html"
	"io"
	"log"
	"math/big"
	mrand "math/rand"
	"net/url"
	"os"
	"regexp"
	"strings"
	"unicode"
)

const (
	// SafeLinksURLRequiredPrefix is the required prefix for all Safe Links
	// URLs.
	SafeLinksURLRequiredPrefix string = "https://"

	// SafeLinksBaseDomain is the common component of a fully qualified domain
	// name used in Safe Links encoded URLs.
	SafeLinksBaseDomain string = "safelinks.protection.outlook.com"

	// SafeLinksURLTemplate is a template for observed SafeLinks URLs.
	// https://SUBDOMAIN.safelinks.protection.outlook.com/?url=ENCODED_URL&data=data_placeholder&sdata=sdata_placeholder&reserved=0
	SafeLinksURLTemplate string = "%s%s/?url=%s&data=%s&sdata=%s&reserved=%s"
)

const (
	// HTTPPlainURLPrefix is the plaintext prefix used for unencrypted
	// connections to a HTTP-enabled site/service.
	HTTPPlainURLPrefix string = "http://"
)

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

// ParsedURL contains the original matched URL pattern and a parsed URL.
type ParsedURL struct {
	// Original is the untrimmed, unmodified original match pattern.
	Original string

	// Parsed is the parsed form of the original match pattern *after* it has
	// been trimmed of unwanted/invalid characters (e.g., angle brackets,
	// period).
	Parsed *url.URL
}

// Trimmed is a copy of the original URL match pattern with unwanted/invalid
// leading/trailing characters removed.
func (pURL ParsedURL) Trimmed() string {
	return trimEnclosingURLCharacters(pURL.Original)
}

// Parsed is a trimmed copy of the original URL match pattern parsed as a
// url.URL value.
// func (pURL ParsedURL) Parsed() string {
// 	return trimEnclosingURLCharacters(pURL.Original)
// }

// ValidURLPattern attempts to validate whether a given input string is a
// valid URL.
func ValidURLPattern(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		// url.Parse is *very* liberal; any parse failure is an immediate
		// validation failure.
		return false
	}

	return ValidURL(u)
}

// ValidURL attempts to validate whether a given url.URL value is a valid /
// usable URL. On its down url.Parse is *very* forgiving so we apply
// additional checks to ensure the url.URL value meets our minimum
// requirements.
func ValidURL(u *url.URL) bool {
	switch {
	case u.Host == "":
		return false
	case u.Scheme == "":
		return false
	default:
		return true
	}
}

// ValidSafeLinkURL validates whether a given url.URL is a valid Safe Links
// URL.
func ValidSafeLinkURL(u *url.URL) bool {
	if !strings.Contains(u.Host, SafeLinksBaseDomain) {
		log.Printf("URL %q fails base domain check", u.String())
		return false
	}

	if err := assertValidURLParameter(u); err != nil {
		log.Printf("URL %q fails %q parameter check", u.String(), "url")
		return false
	}

	return true
}

// ReadURLFromUser attempts to read input from the user via stdin prompt. The
// user is prompted for a URL but validation of that input is left to the
// caller to perform.
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

// ReadFromFile attempts to read newline separated entries from a given file
// (io.Reader).
//
// The collection of entries is returned or an error if one occurs.
func ReadFromFile(r io.Reader) ([]string, error) {
	var entries []string

	// Loop over input "reader" and attempt to collect each item.
	scanner := bufio.NewScanner((r))
	for scanner.Scan() {
		txt := scanner.Text()

		if strings.TrimSpace(txt) == "" {
			continue
		}

		entries = append(entries, txt)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	if len(entries) == 0 {
		return nil, ErrMissingValue
	}

	return entries, nil
}

// CleanURL strips away quoting, escaping of characters or other problematic
// leading or trailing characters in a given URL.
func CleanURL(s string) string {
	// Remove potential quoting.
	s = strings.ReplaceAll(s, `'`, "")
	s = strings.ReplaceAll(s, `"`, "")

	// Remove potential invalid leading or trailing characters from URL.
	s = trimEnclosingURLCharacters(s)

	// Replace escaped ampersands with literal ampersands.
	// inputURL = strings.ReplaceAll(flag.Args()[1], "&amp;", "&")

	// Use html package to handle ampersand escaping *and* any edge cases that
	// I may be unaware of.
	s = html.UnescapeString(s)

	return s
}

func randomBool() bool {
	//nolint:gosec,nolintlint // G404: Use of weak random number generator
	return mrand.Intn(2) == 0
}

// trimEnclosingURLCharacters trims invalid leading or trailing characters
// from given URL.
func trimEnclosingURLCharacters(url string) string {
	// Remove potential leading/trailing period.
	url = strings.Trim(url, `.`)

	// Remove potential enclosing angle brackets.
	url = strings.Trim(url, `<>`)

	// Remove potential enclosing parenthesis used with Markdown formatted
	// URLs.
	url = strings.Trim(url, `()`)

	return url
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

// GetURLPatternsUsingRegex parses the given input and returns a collection of
// FoundURLPattern values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URL patterns begin with that protocol scheme. If
// specified, non-HTTPS URLs are evaluated also. nil is returned if no
// matching patterns are found.
//
// NOTE: Validation is not performed to ensure that matched patterns are valid
// URLs.
//
// Internal logic uses a regular expression to match URL patterns optionally
// beginning with a left angle bracket, then 'https://' (or 'http://' if
// specified) and ending with a whitespace character or a right angle bracket.
// The caller is responsible for trimming angle brackets and other unwanted
// characters.
func GetURLPatternsUsingRegex(input string, evalPlainHTTP bool) ([]FoundURLPattern, error) {
	urlPatterns := make([]FoundURLPattern, 0, 5)

	log.Println("Evaluating plain HTTP URLs:", evalPlainHTTP)

	if !hasAcceptableURLPrefix(input, evalPlainHTTP) {
		return nil, ErrNoURLsFound
	}

	var urlRegex string
	switch {
	case evalPlainHTTP:
		// urlRegex = `<?` + SafeLinksURLRequiredPrefix + `|` + HTTPPlainURLPrefix + `\S+>?`
		urlRegex = fmt.Sprintf(
			`<?(?:%s|%s)\S+>?`,
			SafeLinksURLRequiredPrefix,
			HTTPPlainURLPrefix,
		)
		log.Printf("urlRegex set to also allow plain HTTP prefixes: %q", urlRegex)

	default:
		urlRegex = `<?` + SafeLinksURLRequiredPrefix + `\S+>?`
		log.Printf("urlRegex set to disallow plain HTTP prefixes: %q", urlRegex)
	}

	r := regexp.MustCompile(urlRegex)

	matches := r.FindAllString(input, -1)
	log.Printf("Matches (%d) untrimmed:\n", len(matches))
	for _, m := range matches {
		log.Println(m)
	}

	// log.Println("Cleaning URLs of invalid leading/trailing characters")
	// for i := range matches {
	// 	matches[i] = trimEnclosingURLCharacters(matches[i])
	// }
	// log.Printf("Matches (%d) trimmed:", len(matches))
	// for _, m := range matches {
	// 	log.Println(m)
	// }

	for _, m := range matches {
		urlPatterns = append(
			urlPatterns,
			FoundURLPattern{
				URLPattern: m, // the caller will handle trimming
			},
		)
	}

	return urlPatterns, nil
}

// GetURLPatternsUsingIndex parses the given input and returns a collection of
// FoundURLPattern values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URL patterns begin with that protocol scheme. If
// specified, non-HTTPS URLs are evaluated also. nil is returned if no
// matching patterns are found.
//
// NOTE: Validation has not been performed to ensure that matched patterns are
// valid URLs.
//
// Internal logic uses slice indexing/iteration to match URL patterns
// beginning with 'https://' (or optionally 'http://') and ending with a
// whitespace character or a right angle bracket. The caller is responsible
// for trimming angle brackets and other unwanted characters.
func GetURLPatternsUsingIndex(input string, evalAllHTTPURLs bool) ([]FoundURLPattern, error) {
	if !hasAcceptableURLPrefix(input, evalAllHTTPURLs) {
		return nil, ErrNoURLsFound
	}

	matches, err := getURLPatternsUsingIndex(input, SafeLinksURLRequiredPrefix)
	if err != nil {
		return nil, err
	}

	if evalAllHTTPURLs {
		log.Println("Evaluating plain HTTP URLs also")

		additionalMatches, err := getURLPatternsUsingIndex(input, HTTPPlainURLPrefix)
		if err != nil {
			return nil, err
		}

		matches = append(matches, additionalMatches...)
	}

	log.Printf("Total URL pattern matches: %d", len(matches))
	for _, up := range matches {
		log.Println(up.URLPattern)
	}

	return matches, nil
}

// GetURLPatternsUsingPrefixMatchingOnFields parses the given input and
// returns a collection of FoundURLPattern values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URL patterns begin with that protocol scheme. If
// specified, non-HTTPS URLs are evaluated also. nil is returned if no
// matching patterns are found.
//
// NOTE: Validation has not been performed to ensure that matched patterns are
// valid URLs.
//
// Internal logic uses string splitting on whitespace and prefix matching to
// match URL patterns optionally beginning with a left angle bracket, then
// 'https://' (or 'http://' if specified) and ending with a whitespace
// character. The caller is responsible for trimming angle brackets and other
// unwanted characters.
func GetURLPatternsUsingPrefixMatchingOnFields(input string, evalPlainHTTP bool) ([]FoundURLPattern, error) {
	urlPatterns := make([]FoundURLPattern, 0, 5)

	log.Println("Evaluating plain HTTP URLs:", evalPlainHTTP)

	if !hasAcceptableURLPrefix(input, evalPlainHTTP) {
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
		case evalPlainHTTP && strings.HasPrefix(field, HTTPPlainURLPrefix):
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
					// URLPattern: strings.Trim(field, "<>"),
					URLPattern: field, // the caller will handle trimming
				},
			)

		case evalPlainHTTP && strings.HasPrefix(field, "<"+HTTPPlainURLPrefix):
			urlPatterns = append(
				urlPatterns,
				FoundURLPattern{
					// URLPattern: strings.Trim(field, "<>"),
					URLPattern: field, // the caller will handle trimming
				},
			)
		}
	}

	if len(urlPatterns) == 0 {
		return nil, ErrNoURLsFound
	}

	// log.Println("Cleaning URLs of invalid leading/trailing characters")
	// for i := range urlPatterns {
	// 	urlPatterns[i].URLPattern = trimEnclosingURLCharacters(urlPatterns[i].URLPattern)
	// }

	return urlPatterns, nil
}

// URLs parses the given input and returns a collection of ParsedURL values.
//
// Since all Safe Links URLs observed in the wild begin with a HTTPS scheme we
// require that all matched URLs begin with that protocol scheme. If
// specified, non-HTTPS URLs are also evaluated. nil is returned if no
// matching patterns are found.
//
// The result is a collection of ParsedURL values containing the original URL
// match pattern and a parsed
func URLs(input string, evalPlainHTTP bool) ([]ParsedURL, error) {
	parsedURLs := make([]ParsedURL, 0, 5)

	log.Println("Evaluating plain HTTP URLs:", evalPlainHTTP)

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
	// urlPatterns, err := GetURLPatternsUsingPrefixMatchingOnFields(input, evalPlainHTTP)
	// urlPatterns, err := GetURLPatternsUsingIndex(input, evalPlainHTTP)

	log.Println("Calling GetURLPatternsUsingRegex")
	urlPatterns, err := GetURLPatternsUsingRegex(input, evalPlainHTTP)
	if err != nil {
		return nil, err
	}

	log.Printf("Processing %d matched patterns", len(urlPatterns))

	for _, pattern := range urlPatterns {
		trimmedPattern := trimEnclosingURLCharacters(pattern.URLPattern)
		u, err := url.Parse(trimmedPattern)
		if err != nil {
			// url.Parse is *very* lenient. Any failure at this point is a
			// reliable "skip" indication.
			continue
		}

		if !ValidURL(u) {
			continue
		}

		pURL := ParsedURL{
			Original: pattern.URLPattern,
			Parsed:   u,
		}

		log.Printf("Original URL match: %q", pURL.Original)
		log.Printf("Trimmed URL match: %q", pURL.Trimmed())
		log.Printf("Parsed URL: %+v", pURL.Parsed.String())

		parsedURLs = append(parsedURLs, pURL)
	}

	return parsedURLs, nil
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
				DecodedURL: CleanURL(originalURL),
			},
		)
	}

	if len(safeLinkURLs) == 0 {
		return nil, ErrNoSafeLinkURLsFound
	}

	return safeLinkURLs, nil
}

// SafeLinkURLsFromParsedURLs evaluates a given collection of parsed URLs and
// returns any that are found to be encoded as Safe Links. Deduplication is
// *not* performed. An error is returned if no valid matches are found.
func SafeLinkURLsFromParsedURLs(parsedURLs []ParsedURL) ([]SafeLinkURL, error) {
	safeLinkURLs := make([]SafeLinkURL, 0, len(parsedURLs))

	for _, u := range parsedURLs {
		if !ValidSafeLinkURL(u.Parsed) {
			continue
		}

		originalURL := u.Parsed.Query().Get("url")

		safeLinkURLs = append(
			safeLinkURLs,
			SafeLinkURL{
				EncodedURL: u.Parsed.String(),
				// DecodedURL: originalURL,
				DecodedURL: CleanURL(originalURL),
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
	log.Println("Calling URLs")
	urls, err := URLs(input, false)
	if err != nil {
		return nil, err
	}

	log.Println("Calling SafeLinkURLsFromURLs")
	return SafeLinkURLsFromParsedURLs(urls)
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

// getURLPatternsUsingIndex performs the bulk of the work for the exported
// GetURLPatternsUsingIndex function. See that function's doc comments for
// further details.
func getURLPatternsUsingIndex(input string, urlPrefix string) ([]FoundURLPattern, error) {
	urlPatterns := make([]FoundURLPattern, 0, 5)

	remaining := input

	for {
		urlStart := strings.Index(remaining, urlPrefix)
		log.Println("urlStart:", urlStart)

		if urlStart == -1 {
			break
		}

		next := urlStart + len(urlPrefix) + 1

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

				URLPattern: remaining[urlStart:urlEnd], // the caller will handle trimming
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

	// log.Println("Cleaning URLs of invalid leading/trailing characters")
	// for i := range urlPatterns {
	// 	urlPatterns[i].URLPattern = trimEnclosingURLCharacters(urlPatterns[i].URLPattern)
	// }

	return urlPatterns, nil
}

// hasAcceptableURLPrefix accepts an input string and an indication of whether
// a plain HTTP prefix should be considered OK alongside the existing
// Safe Links URL required prefix.
func hasAcceptableURLPrefix(input string, evalPlainHTTP bool) bool {
	hasSafeLinksURL := strings.Contains(input, SafeLinksURLRequiredPrefix)
	hasPlainHTTPURL := strings.Contains(input, HTTPPlainURLPrefix)

	switch {
	case hasSafeLinksURL:
		return true
	case hasPlainHTTPURL && evalPlainHTTP:
		return true
	default:
		return false
	}
}

// GetRandomSafeLinksFQDN returns a pseudorandom FQDN from a list observed to
// be associated with Safe Links URLs. Entries in the list have a naming
// pattern of *.safelinks.protection.outlook.com.
func GetRandomSafeLinksFQDN() string {
	subdomains := []string{
		"emea01",
		"eur04",
		"na01",
		"nam01",
		"nam02",
		"nam04",
		"nam10",
		"nam11",
		"nam12",
	}

	var subdomain string
	crandMax := big.NewInt(int64(len(subdomains)))
	mrandMax := len(subdomains)

	// Attempt crypto/rand based PRNG sourced value first before falling back
	// to math/rand sourced value. This should probably just be ignored
	// (without bothering to apply fallback behavior) as the use case does not
	// require a cryptographically strong pseudo-random number generator.
	n, err := crand.Int(crand.Reader, crandMax)
	switch {
	case err != nil:
		//nolint:gosec,nolintlint // G404: Use of weak random number generator
		subdomain = subdomains[mrand.Intn(mrandMax)]
	default:
		subdomain = subdomains[n.Int64()]
	}

	return strings.Join([]string{subdomain, SafeLinksBaseDomain}, ".")
}

// EncodeParsedURLAsFauxSafeLinksURL encodes the provided ParsedURL in a
// format that mimics real Safe Links encoded URLs observed in the wild. This
// output is intended for use with testing encoding/decoding behavior.
func EncodeParsedURLAsFauxSafeLinksURL(pURL ParsedURL) string {
	return fmt.Sprintf(
		SafeLinksURLTemplate,
		SafeLinksURLRequiredPrefix,
		GetRandomSafeLinksFQDN(),
		url.QueryEscape(pURL.Trimmed()),
		"data_placeholder",
		"sdata_placeholder",
		"0", // 0 is the only value observed in the wild thus far.
	)
}

// FilterURLs filters the given collection of URLs, returning the remaining
// URLs or an error if none remain after filtering.
//
// If specified, Safe Link URLs are excluded from the collection returning
// only URLs that have not been encoded as Safe Links URLs. Otherwise, only
// URLs that have been encoded as Safe Links URLs are returned.
//
// An empty collection is returned if no URLs remain after filtering.
func FilterURLs(urls []*url.URL, excludeSafeLinkURLs bool) []*url.URL {
	remaining := make([]*url.URL, 0, len(urls))

	keepSafeLinksURLs := !excludeSafeLinkURLs
	keepNonSafeLinksURLs := excludeSafeLinkURLs

	for _, u := range urls {
		if ValidSafeLinkURL(u) {
			log.Printf("URL identified as Safe Links encoded: %q", u.String())

			switch {
			case keepSafeLinksURLs:
				log.Printf("Retaining Safe Links encoded URL %q as requested", u.String())
				remaining = append(remaining, u)
			default:
				log.Printf("Skipping Safe Links encoded URL %q as requested", u.String())
			}

			continue
		}

		log.Printf("URL not identified as Safe Links encoded: %q", u.String())

		switch {
		case keepNonSafeLinksURLs:
			log.Printf("Retaining unencoded URL %q as requested", u.String())
			remaining = append(remaining, u)

			continue
		default:
			log.Printf("Skipping unencoded URL %q as requested", u.String())
		}
	}

	return remaining
}

// FilterParsedURLs filters the given collection of parsed URLs, returning the
// remaining parsed URLs or an error if none remain after filtering.
//
// If specified, Safe Link URLs are excluded from the collection returning
// only URLs that have not been encoded as Safe Links URLs. Otherwise, only
// URLs that have been encoded as Safe Links URLs are returned.
//
// An empty collection is returned if no URLs remain after filtering.
func FilterParsedURLs(parsedURLs []ParsedURL, excludeSafeLinkURLs bool) []ParsedURL {
	remaining := make([]ParsedURL, 0, len(parsedURLs))

	keepSafeLinksURLs := !excludeSafeLinkURLs
	keepNonSafeLinksURLs := excludeSafeLinkURLs

	for _, pURL := range parsedURLs {
		if ValidSafeLinkURL(pURL.Parsed) {
			log.Printf("URL identified as Safe Links encoded (orig): %q", pURL.Original)

			switch {
			case keepSafeLinksURLs:
				log.Printf("Retaining Safe Links encoded URL %q as requested", pURL.Original)
				remaining = append(remaining, pURL)
			default:
				log.Printf("Skipping Safe Links encoded URL %q as requested", pURL.Original)
			}

			continue
		}

		log.Printf("URL not identified as Safe Links encoded: %q", pURL.Original)

		switch {
		case keepNonSafeLinksURLs:
			log.Printf("Retaining unencoded URL %q as requested", pURL.Original)
			remaining = append(remaining, pURL)

			continue
		default:
			log.Printf("Skipping unencoded URL %q as requested", pURL.Original)
		}
	}

	return remaining
}

// DecodeInput processes given input replacing any Safe Links encoded URL
// with the original decoded value. Other input is returned unmodified.
func DecodeInput(txt string) (string, error) {
	log.Println("Calling SafeLinkURLs")

	safeLinks, err := SafeLinkURLs(txt)
	if err != nil {
		return "", err
	}

	modifiedInput := txt

	// URLs are "cleaned" of problematic leading and trailing characters as
	// part of retrieving them and asserting that they're in the expected
	// format of Safe Links URLs. In order to safely match and replace the
	// original encoded URL we also have to perform that same URL cleaning
	// step. This helps handle edge cases where an original URL match applies
	// to more characters than intended.
	for _, sl := range safeLinks {
		cleanedOriginalURL := trimEnclosingURLCharacters(sl.EncodedURL)
		modifiedInput = strings.Replace(modifiedInput, cleanedOriginalURL, sl.DecodedURL, 1)
	}

	return modifiedInput, nil
}

// EncodeInput processes given input replacing any normal URL with an encoded
// value similar to a real Safe Links value. Other input is returned
// unmodified.
func EncodeInput(txt string, randomlyEncode bool) (string, error) {
	log.Println("Calling URLs")
	urls, err := URLs(txt, true)
	if err != nil {
		return "", err
	}

	nonSafeLinkURLs := FilterParsedURLs(urls, true)

	log.Printf("%d URLs identified as nonSafeLinkURLs", len(nonSafeLinkURLs))

	if len(nonSafeLinkURLs) == 0 {
		return "", ErrNoNonSafeLinkURLsFound
	}

	log.Printf("nonSafeLinkURLs URLs (%d):", len(nonSafeLinkURLs))
	for i, u := range nonSafeLinkURLs {
		log.Printf("(%2.2d) %s", i+1, u.Original)
	}

	modifiedInput := txt
	log.Printf("Replacing original unencoded URLs (randomly: %t)", randomlyEncode)
	shouldEncode := true
	for _, pURL := range nonSafeLinkURLs {
		if randomlyEncode {
			shouldEncode = randomBool()
		}

		if shouldEncode {
			cleanedOriginalURL := pURL.Trimmed()
			fauxSafeLinksURL := EncodeParsedURLAsFauxSafeLinksURL(pURL)
			modifiedInput = strings.Replace(modifiedInput, cleanedOriginalURL, fauxSafeLinksURL, 1)
		}
	}

	if modifiedInput == txt {
		return "", fmt.Errorf("encoded output matches input: %w", ErrEncodingUnsuccessful)
	}

	return modifiedInput, nil
}

// QueryEscapeInput processes given input replacing any normal URL with an
// escaped string so it can be safely placed inside a URL query. Other input
// is returned unmodified.
func QueryEscapeInput(txt string, randomlyEscape bool) (string, error) {
	log.Println("Calling URLs")
	urls, err := URLs(txt, true)
	if err != nil {
		return "", err
	}

	nonSafeLinkURLs := FilterParsedURLs(urls, true)

	log.Printf("%d URLs identified as nonSafeLinkURLs", len(nonSafeLinkURLs))

	if len(nonSafeLinkURLs) == 0 {
		return "", ErrNoNonSafeLinkURLsFound
	}

	log.Printf("nonSafeLinkURLs URLs (%d):", len(nonSafeLinkURLs))
	for i, u := range nonSafeLinkURLs {
		log.Printf("(%2.2d) %s", i+1, u.Original)
	}

	modifiedInput := txt
	log.Printf("Replacing original unencoded URLs (randomly: %t)", randomlyEscape)
	shouldQueryEscape := true
	for _, pURL := range nonSafeLinkURLs {
		if randomlyEscape {
			shouldQueryEscape = randomBool()
		}

		if shouldQueryEscape {
			cleanedOriginalURL := pURL.Trimmed()
			queryEscapedURL := url.QueryEscape(cleanedOriginalURL)
			modifiedInput = strings.Replace(modifiedInput, cleanedOriginalURL, queryEscapedURL, 1)
		}
	}

	if modifiedInput == txt {
		return "", fmt.Errorf("encoded output matches input: %w", ErrQueryEscapingUnsuccessful)
	}

	return modifiedInput, nil
}
