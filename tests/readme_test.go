/*
	AdrestiaAssert (license_text.go)
	(c) 2020 Sam Caldwell.  See LICENSE.txt

	Test the readme file
*/
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	Testing "testing"
)

func TestReadMeHash(t *Testing.T) {
	/*
		Hash the License file.  Ensure it matches the expected hash.
	*/
	const licenseFile = "../README.md"
	const expectedHash = "64214bf43ac9e2bddc2f942b37e64db52dada361d169e747a8f9d06bd85e5708"
	fileContent, err := ioutil.ReadFile(licenseFile)
	if err != nil {
		t.Errorf("File failed hash validation '%s'(%s): %v", licenseFile, expectedHash, err)
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(fileContent))
	if hash != expectedHash {
		t.Errorf("File hash mismatch: %s (%s)(%s)", licenseFile, expectedHash, hash)
	}
}

// ToDo: add more tests for individual documentation files.