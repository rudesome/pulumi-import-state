package github

import (
	"fmt"
	"net/http"
	"os"
)

const (
	baseURL = "https://api.github.com"
	//
	EnvironmentalVariableToken string = "API_KEY"
)

type Client struct {
	BaseURL    string
	apiToken   string
	HTTPClient *http.Client
}

func Token(envName string) (string, error) {
	token := os.Getenv(EnvironmentalVariableToken)
	if token == "" {
		fmt.Printf("Unable to read %s environment variable. Empty\n", EnvironmentalVariableToken)
	}
	apiToken := fmt.Sprintf("token %s", token)

	return apiToken, nil
}

func NewClient(apiToken string) *Client {
	return &Client{
		BaseURL:    baseURL,
		apiToken:   apiToken,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	return c.Do(req)

}

func (c *Client) Do(req *http.Request) (*http.Response, error) {

	req.Header.Set("Authorization", c.apiToken)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Content-Type", "application/json")

	// nested invocation of Get() function
	return c.HTTPClient.Do(req)

}
