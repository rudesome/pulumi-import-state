package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type Repos []Repo

type Repo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Fork     bool   `json:"fork"`
	Archived bool   `json:"archived"`
}

func (c *Client) GetRepos(ctx context.Context) (*Repos, error) {
	fmt.Println("Getting your repos")

	resp, err := c.Get(fmt.Sprintf("%s/user/repos?per_page=100", c.BaseURL))
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

	return &data, nil
}
