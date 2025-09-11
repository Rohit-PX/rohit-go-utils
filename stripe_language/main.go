package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Supported languages", ParseAcceptLanguage("en-US", []string{"en-US", "fr-CA"}))

}

func ParseAcceptLanguage(languages string, supportedLanguages []string) []string {
	langMap := make(map[string]bool)

	for _, lang := range strings.Split(languages, ",") {
		if _, ok := langMap[lang]; !ok {
			langMap[lang] = true
		}
	}
	var supported []string

	for _, l := range supportedLanguages {
		if _, ok := langMap[l]; ok {
			supported = append(supported, l)
		}
	}
	return supported
}
