// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// HTAreFilesEqual compares two files by their SHA-256 digest.
func HTAreFilesEqual(fFile string, sFile string) (bool, error) {
	f, err := os.Open(fFile)
	if err != nil {
		return false, err
	}
	defer f.Close()

	hf := sha256.New()
	if _, err := io.Copy(hf, f); err != nil {
		return false, err
	}

	s, err := os.Open(sFile)
	if err != nil {
		return false, err
	}
	defer s.Close()

	hs := sha256.New()
	if _, err := io.Copy(hs, s); err != nil {
		return false, err
	}

	fstr := hex.EncodeToString(hf.Sum(nil))
	sstr := hex.EncodeToString(hs.Sum(nil))

	return fstr == sstr, nil
}
