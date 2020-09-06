package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// FilePutContents file_put_contents()
func FilePutContents(filename string, data string, mode os.FileMode) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, []byte(data), mode)
}
