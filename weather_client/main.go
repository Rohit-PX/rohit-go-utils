package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	baseUrl = "https://api.weatherstack.com/current"
)

type Request struct {
	Language string `json:"language"`
	Query    string `json:"query"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
}

type PartialResponse struct {
	Request Request `json:"request"`
}

func main() {

	rawURL := getRawUrl(baseUrl)

	weatherResp, err := CheckWeather(rawURL)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Response", weatherResp)
}

func getRawUrl(baseUrl string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Panic(err)
	}

	params := url.Values{}
	apiKey := "009b2c8207aa24ac8a63c6522131b1b1"
	params.Add("access_key", apiKey)
	params.Add("query", "New York")

	u.RawQuery = params.Encode()
	return u.String()

}

func CheckWeather(url string) (*PartialResponse, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Server returned status %d: %s\n", resp.StatusCode, string(body))
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("body", string(body))
	// Parse JSON into a generic map
	//var result map[string]interface{}
	var pr PartialResponse

	if err := json.Unmarshal(body, &pr); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	// if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
	// 	return nil, fmt.Errorf("failed to decode response: %w", err)
	// }

	return &pr, nil
}
