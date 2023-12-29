package main

import (
	"fmt"

	"github.com/rudesome/pulumi-import-state/pkg/github"
)

func main() {

	token, err := github.Token("API_KEY")
  //os.Getenv(github.EnvironmentVariableToken)
	if err != nil {
		fmt.Println(err)
	}
	c := github.NewClient(token)

	repos, err := c.GetRepos(nil)

	if err != nil {
		fmt.Println(err.Error())
	}

  //github.PrettyJSON(repos)
  fmt.Println(repos)

	// TODO:
	// Check for pulumi prerequisites
	// Login, Evaluated folder, Stack

	// Path as user input
	//github.PulumiImport(repos, "/home/rudesome/github/pulumi-github")
}
