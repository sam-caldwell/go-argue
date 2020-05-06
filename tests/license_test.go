/*
	AdrestiaAssert (license_text.go)
	(c) 2020 Sam Caldwell.  See LICENSE.txt

	Test the license file.
*/
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	Testing "testing"
)

func TestLicenseHash(t *Testing.T) {
	/*
		Hash the License file.  Ensure it matches the expected hash.
	*/
	const licenseFile = "../LICENSE.txt"
	const expectedHash = "72c552ad631e5d855b6e57ad38afce1dbec6807aec98232d599fb821d5285fcf"
	fileContent, err := ioutil.ReadFile(licenseFile)
	if err != nil {
		t.Errorf("File failed hash validation '%s'(%s): %v", licenseFile, expectedHash, err)
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(fileContent))
	if hash != expectedHash {
		t.Errorf("File hash mismatch: %s (%s)(%s)", licenseFile, expectedHash, hash)
	}
}
