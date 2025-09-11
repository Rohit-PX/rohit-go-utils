package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewYorWeather(t *testing.T) {
	url := getRawUrl(baseUrl)
	weather, err := CheckWeather(url)
	assert.NoError(t, err)
	assert.NotNil(t, weather)
	assert.Equal(t, "en", weather.Request.Language)
}
