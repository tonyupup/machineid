//go:build windows
// +build windows

package machineid

import (
	"os/exec"
	"strings"
)

// If there is an error running the commad an empty string is returned.
func machineID() (string, error) {
	c := exec.Command("wmic", "csproduct", "get", "uuid")
	output, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}
	u := trim(strings.Trim(strings.Split(string(output), "UUID")[1], " "))
	return u, nil
}
