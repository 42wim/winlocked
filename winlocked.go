package winlocked

import (
	"strings"

	"github.com/mitchellh/go-ps"
)

// Locked checks if the loginui process is running, this is the simplest way
// to detect if a users desktop is locked. Doesn't work with multiple users on the
// same desktop.
// See https://stackoverflow.com/a/61681203
func Locked() (bool, error) {
	processes, err := ps.Processes()
	if err != nil {
		return false, err
	}

	for _, process := range processes {
		if strings.ToLower(process.Executable()) == "logonui.exe" {
			return true, nil
		}
	}

	return false, nil
}
