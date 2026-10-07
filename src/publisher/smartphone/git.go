// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// gitModified returns the files reported by "git ls-files -m" so that their
// last_update timestamp can be refreshed during publishing. The command runs
// inside SrcPath, so the returned keys are absolute-ish paths prefixed with
// SrcPath. Failure to read the VCS state is non-fatal.
func gitModified(cfg Config) map[string]bool {
	cmd := exec.Command("git", "ls-files", "-m")
	cmd.Dir = cfg.SrcPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		if cfg.Verbose {
			fmt.Fprintln(os.Stderr, "WARNING (non-fatal): git VCS info unavailable, skipping modified-file tracking: ", err)
		}
		return nil
	}

	modified := make(map[string]bool)
	for _, file := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		if file == "" {
			continue
		}

		pathFile := fmt.Sprintf("%s%s", cfg.SrcPath, file)
		modified[pathFile] = true

		if cfg.Verbose {
			fmt.Println("GIT: The file ", pathFile, " was modified.")
		}
	}
	return modified
}
