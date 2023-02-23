package jq

import (
	"fmt"
	"os/exec"
)

func ExecForPath(path, expression string) (string, error) {
	const cmd = "cat %s | jq %s"
	out, err := exec.Command("bash", "-c", fmt.Sprintf(cmd, path, expression)).Output()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(out))
	}

	return string(out), nil
}
