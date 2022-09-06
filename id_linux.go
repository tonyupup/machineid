//go:build linux
// +build linux

package machineid

import (
	"os/exec"
)

const (
	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc          = "/etc/machine-id"
	DMI_UUID_PATH string = "/sys/class/dmi/id/product_uuid"
)

// machineID returns the uuid specified at `/var/lib/dbus/machine-id` or `/etc/machine-id`.
// If there is an error reading the files an empty string is returned.
// See https://unix.stackexchange.com/questions/144812/generate-consistent-machine-unique-id
func machineID() (string, error) {
	cmd := exec.Command("dmidecode", "-s", "system-uuid")
	outPut, err := cmd.CombinedOutput()
	if err == nil {
		goto done
	}

	outPut, err = readFile(DMI_UUID_PATH)
	if err == nil {
		goto done
	}

	outPut, err = readFile(dbusPath)
	if err == nil {
		goto done
	}

	outPut, err = readFile(dbusPathEtc)
	if err == nil {
		goto done
	}

	return "", err
done:
	return trim(string(outPut)), nil
}
