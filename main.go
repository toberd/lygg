package main

import (
	"fmt"
	"lygg/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
