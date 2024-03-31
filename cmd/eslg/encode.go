// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"log"
	"strings"

	"github.com/atc0005/safelinks/internal/safelinks"
)

// encodeInput processes given input replacing any normal URL with an encoded
// value similar to a real Safe Links value. Other input is returned
// unmodified.
func encodeInput(txt string) (string, error) {
	log.Println("Calling URLs(input)")
	urls, err := safelinks.URLs(txt)
	if err != nil {
		return "", err
	}

	nonSafeLinkURLs := safelinks.FilterURLs(urls, true)

	log.Printf("%d URLs identified as nonSafeLinkURLs", len(nonSafeLinkURLs))

	if len(nonSafeLinkURLs) == 0 {
		return "", safelinks.ErrNoNonSafeLinkURLsFound
	}

	log.Printf("nonSafeLinkURLs URLs (%d):", len(nonSafeLinkURLs))
	for i, u := range nonSafeLinkURLs {
		log.Printf("(%2.2d) %s", i+1, u.String())
	}

	modifiedInput := txt
	log.Println("Replacing original unencoded URLs")
	for _, link := range nonSafeLinkURLs {
		fauxSafeLinksURL := safelinks.EncodeURLAsFauxSafeLinksURL(link)
		modifiedInput = strings.Replace(modifiedInput, link.String(), fauxSafeLinksURL, 1)
	}

	return modifiedInput, nil
}
