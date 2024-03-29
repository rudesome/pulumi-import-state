package main

import (
	"fmt"

	"github.com/rudesome/pulumi-import-state/pkg/github"
	"github.com/rudesome/pulumi-import-state/pkg/cli"
)

func main() {

	token, err := github.Token()
	if err != nil {
		fmt.Println(err)
	}
	c := github.NewClient(token)

	repos, raw, err := c.GetRepos(nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	github.PrettyJSON(raw)

  cli.Start()

	// TODO:
	// Check for pulumi prerequisites
	// Login, Evaluated folder, Stack

	// Path as user input
	github.PulumiImport(repos, "/home/rudesome/github/pulumi-github")
}
