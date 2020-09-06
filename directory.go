package helpers

import (
	"os"
	"syscall"
)

// Chdir - Change directory
func Chdir(dir string) error {

	return os.Chdir(dir)
}

// Getcwd getcwd()
func Getcwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

// Closedir - Close directory's handle
func Closedir(fd int) (err error) {

	return syscall.Close(fd)
}
