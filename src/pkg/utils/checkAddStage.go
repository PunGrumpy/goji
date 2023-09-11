package utils

import (
	"os/exec"
	"strings"
)

var execCommand = exec.Command

func CheckAddStage() (bool, error) {
	cmd := execCommand("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "A ") || strings.HasPrefix(line, "M ") || strings.HasPrefix(line, "R ") {
			return true, nil
		}
	}

	return false, nil
}
