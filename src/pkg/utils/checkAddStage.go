package utils

import (
	"os/exec"
)

var execCommand = exec.Command

func CheckAddStage() (bool, error) {
	cmd := execCommand("git", "status")
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	if string(out) == "Changes to be committed" {
		return true, nil
	}
	return false, nil
}
