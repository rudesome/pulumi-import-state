package github

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyJSON(jsonData []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(jsonData), "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out.String())
}
