package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

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

type Repos []Repo

type Repo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Fork     bool   `json:"fork"`
	Archived bool   `json:"archived"`
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

func PrettyJSON(jsonData []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(jsonData), "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out.String())
}

func (c *Client) GetRepos(ctx context.Context) (*Repos, error) {
	fmt.Println("Getting your repos")

	resp, err := c.Get(fmt.Sprintf("%s/user/repos", c.BaseURL)) // ?per_page=100", c.BaseURL))
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	//fmt.Println(strings.Split(resp.Header.Get("Link"), ";")[0])

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	var data Repos
	json.Unmarshal(body, &data)
	fmt.Println(&data)

	return &data, nil
}

func PulumiImport(r *Repos, path string) {
	for _, v := range *r {
		// dont import forks..
		if v.Fork == true {
			continue
		}

		// Pulumi define shell command:
		// `pulumi import [type] [name] [id] [flags]`
		cmdStruct := exec.Command(
			"pulumi",
			"import",
			"github:index/repository:Repository",
			v.Name,
			v.Name,
			"-y",
			"--skip-preview",
			"--protect=false",
		)

    if !(filepath.IsAbs(path)) {
      fmt.Println("Please enter the absolute path to the pulumi github directory")
      return
    }

		cmdStruct.Dir = path

    fmt.Println("cmdStruct is: ", cmdStruct.Dir)

		//Execute command
    //_, err := cmdStruct.Output()

    //if err != nil {
      //fmt.Println(err.Error())
    //}
	}
}

func main() {

	token, err := token("API_KEY")
	if err != nil {
		fmt.Println(err)
	}
	c := NewClient(token)

	repos, err := c.GetRepos(nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(repos)

	// TODO:
	// Check for pulumi prerequisites
	// Login, Evaluated folder, Stack

	// Path as user input
	PulumiImport(repos, "/home/rudesome/github/pulumi-github")
}
