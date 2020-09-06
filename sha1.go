package helpers

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

// Sha1 - Calculate the sha1 hash of a string
func Sha1(s string) string {
	digest := sha1.Sum([]byte(s))
	return hex.EncodeToString(digest[:])
}

// Sha1File sha1_file()
func Sha1File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}