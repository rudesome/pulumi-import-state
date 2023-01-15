package main

import (
  "bytes"
  "context"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "os/exec"
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

  resp, err := c.Get(fmt.Sprintf("%s/user/repos?per_page=100", c.BaseURL))
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data Repos
	json.Unmarshal(body, &data)
	return &data, nil
}

func PulumiImport(r *Repos) {
	for _, v := range *r {
		// dont import forks..
		if v.Fork == true {
			continue
		}

		// Pulumi define shell command
		cmdStruct := exec.Command("pulumi", "import", "github:index/repository:Repository", v.Name, v.Name, "-y", "--skip-preview", "--protect=false")
    
		// TODO: Set the -absolute path- argument to the function
		cmdStruct.Dir = "/home/rudesome/git/pulumi-github/"

		// Execute command
		_, err := cmdStruct.Output()

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {

	token, err := token("API_KEY")
	if err != nil {
		fmt.Println(err)
	}
	c := NewClient(token)

	res, err := c.GetRepos(nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	PulumiImport(res)
}
