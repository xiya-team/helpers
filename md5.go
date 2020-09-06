package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

// Md5 - Calculate the md5 hash of a string
func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Md5File md5_file()
func Md5File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}