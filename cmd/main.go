package main

import (
	"fmt"
  "github/rudesome/pulumi-import-state/pkg/github"
)

const (
	baseURL = "https://api.github.com"
)

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
