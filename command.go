package main

import "os/exec"

var (
	resultCache = make(map[string]string)
)

func CachedExec(command string) (string, error) {
	if result, ok := resultCache[command]; ok {
		return result, nil
	}

	result, err := Exec(command)
	if err != nil {
		return "", err
	}
	resultCache[command] = result

	return result, nil
}

func Exec(command string) (string, error) {
	cmd := exec.Command("ghq", "list")

	result, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(result), nil
}
