package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAcceptLanguage(t *testing.T) {
	lang := "en-US"
	supp := []string{"en-US", "fr-CA"}

	op := ParseAcceptLanguage(lang, supp)
	assert.Equal(t, len(op), 1)

	lang = "fr-CA,fr-FR,en-US"
	supp = []string{"fr-CA", "en-US"}

	op = ParseAcceptLanguage(lang, supp)
	assert.Equal(t, len(op), 2)
}

func TestParseAcceptLanguageNegative(t *testing.T) {
	lang := ""
	supp := []string{"fr-CA", "en-US"}

	op := ParseAcceptLanguage(lang, supp)
	assert.Equal(t, 0, len(op))

	lang = "en-US"
	supp = []string{}

	op = ParseAcceptLanguage(lang, supp)
	assert.Equal(t, 0, len(op))

	lang = "en-US"
	supp = []string{}

	op = ParseAcceptLanguage(lang, supp)
	assert.Equal(t, 0, len(op))
}
