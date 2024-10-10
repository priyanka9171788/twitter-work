// main.go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dghubble/oauth1"
)

const (
	// Replace these with your API keys and tokens
	apiKey            = "He6XLgcjugj0UlMWiYcf0krnG"
	apiSecretKey      = "lBIk21SpxmUqJ1E9WrpTbIkEkqpe2C2lclci79ua8i17e3kQwS"
	accessToken       = "1844361275645554693-4JsZMJaAMNI43AZDAZrBj8v6nWOPa2"
	accessTokenSecret = "k0XyLKdQmJSrddDVJ5jz9G9ywVqnMqDLD5nnv8f5QYlY3"
	delaySecond       = 10
)

func main() {
	// Example of posting a tweet
	tweetID, err := postTweet("Hello from Twitter API fourth tweet!")
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}
	fmt.Printf("Posted tweet with ID: %s\n", tweetID)
	fmt.Printf("Waiting for %d seconds before deleting a tweet...\n", delaySecond)
	time.Sleep(time.Duration(delaySecond) * time.Second)

	//Deleting the tweet after delay
	err = deleteTweet(tweetID)
	if err != nil {
		log.Fatalf("Error deleting tweet: %v", err)

	}
	fmt.Printf("Deleted tweet with ID: %s\n", tweetID)

}

// postTweet sends a tweet to Twitter
func postTweet(content string) (string, error) {
	// OAuth1 authentication
	config := oauth1.NewConfig(apiKey, apiSecretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	// Create HTTP client with OAuth1
	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	// Create tweet content
	tweetData := map[string]interface{}{
		"text": content,
	}
	tweetJSON, err := json.Marshal(tweetData)
	if err != nil {
		return "", fmt.Errorf("error marshaling tweet content: %w", err)
	}

	// Make POST request to v2 tweets endpoint
	response, err := httpClient.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBuffer(tweetJSON))
	if err != nil {
		return "", fmt.Errorf("failed to post tweet: %w", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode == http.StatusCreated {
		var result map[string]interface{}
		body, _ := ioutil.ReadAll(response.Body)
		if err := json.Unmarshal(body, &result); err != nil {
			return "", fmt.Errorf("error unmarshaling response: %w", err)
		}
		// Extract tweet ID from the response
		return result["data"].(map[string]interface{})["id"].(string), nil
	}

	// Print detailed error information
	body, _ := ioutil.ReadAll(response.Body)
	return "", fmt.Errorf("failed to post tweet: %s, response: %s", response.Status, string(body))
}

func deleteTweet(tweetID string) error {
	// OAuth1 authentication
	config := oauth1.NewConfig(apiKey, apiSecretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	// Create HTTP client with OAuth1
	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID), nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	// Execute the DELETE request
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute delete request: %w", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode != http.StatusOK {
		// Read and print detailed error information from the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to delete tweet: unable to read response body: %v", err)
		}

		return fmt.Errorf("failed to delete tweet: %s, response: %s", response.Status, string(body))
	}

	// Success case
	fmt.Println("Tweet successfully deleted.")
	return nil

}