package functions

import (
	"os/exec"
)

func runBashScript(scriptPath string) error {
	cmd := exec.Command("bash", scriptPath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
