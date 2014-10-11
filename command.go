package main

import (
	"fmt"
	"os/exec"
)

func execCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)

	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf(string(result))
}
