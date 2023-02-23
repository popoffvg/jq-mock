package jq

import (
	"errors"
	"fmt"
	"os/exec"
)

func ExecForPath(path, expression string) (string, error) {
	var cmd = fmt.Sprintf("cat %s | jq %s", path, expression)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		var execErr *exec.ExitError
		if errors.As(err, &execErr) {
			return "", fmt.Errorf("%w: %s %s", err, cmd, string(execErr.Stderr))
		}
		return "", fmt.Errorf("%w: %s %s", err, cmd, string(out))
	}

	return string(out), nil
}
