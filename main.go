package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	urlMap := make(map[string]string)

	hexRouteArg := regexp.MustCompile(`([0-9a-f]{2})+`)
	intRouteArg := regexp.MustCompile(`\d{2,}`)

	for sc.Scan() {
		u, err := url.Parse(sc.Text())
		if err != nil {
			continue
		}

		normalizeUrl := strings.Split(u.String(), "#")[0]
		normalizeUrl = strings.Split(normalizeUrl, "?")[0]
		normalizeUrl = hexRouteArg.ReplaceAllString(normalizeUrl, "_")
		normalizeUrl = intRouteArg.ReplaceAllString(normalizeUrl, "_")

		existingUrl, ok := urlMap[normalizeUrl]
		if !ok {
			urlMap[normalizeUrl] = u.String()
			continue
		}

		u2, _ := url.Parse(existingUrl)
		for key, val := range u.Query() {
			if !u2.Query().Has(key) {
				q := u2.Query()
				q.Add(key, strings.Join(val, ","))
				u2.RawQuery = q.Encode()
			}
		}

		urlMap[normalizeUrl] = u2.String()
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	for _, v := range urlMap {
		fmt.Println(v)
	}
}
