package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rohit-go-utils/stripe/pkg/errorstripe"
)

func CallTransactionEndpoint(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return &errorstripe.ErrHttpServer{
			Type: "Server call error",
			Err:  err,
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Server returned status %d: %s\n", resp.StatusCode, string(body))
		return nil
	}
	// Parse JSON into a generic map
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// Pretty print the result
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))

	return nil
}
