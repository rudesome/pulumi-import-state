package github

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

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
			fmt.Println(
				"Please enter the absolute path to the pulumi github directory",
			)
			return
		}

		cmdStruct.Dir = path

		//Execute command
		//_, err := cmdStruct.Output()

		//if err != nil {
		//fmt.Println(err.Error())
		//}
	}
}
