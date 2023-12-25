package github

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	baseURL = "https://api.github.com"
)

type Client struct {
	BaseURL    string
	apiToken   string
	HTTPClient *http.Client
}

func token(envName string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	apiToken := fmt.Sprintf("token %s", os.Getenv(envName))

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

	return c.HTTPClient.Do(req)

}
