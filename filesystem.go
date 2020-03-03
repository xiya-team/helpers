package helpers

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// Chgrp - Changes file group
func Chgrp(name string, uid, gid int) error {
	return Chown(name, uid, gid)
}

// Chmod chmod()
func Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// Chown chown()
func Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}


// Copy copy()
func Copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

// Delete delete()
func Delete(filename string) error {
	return os.Remove(filename)
}

// Dirname - Returns a parent directory's path
func Dirname(dirPth string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirPth)
}

// Fclose fclose()
func Fclose(handle *os.File) error {
	return handle.Close()
}

// FileExists file_exists()
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// Filemtime filemtime()
func Filemtime(filename string) (int64, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	fileinfo, err := fd.Stat()
	if err != nil {
		return 0, err
	}
	return fileinfo.ModTime().Unix(), nil
}

// Glob - Find pathnames matching a pattern
func Glob(pattern string) (matches []string, err error) {
	return filepath.Glob(pattern)
}

// IsDir is_dir()
func IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// FileSize filesize()
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// IsReadable is_readable()
func IsReadable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_RDONLY, 0)
	if err != nil {
		return false
	}
	return true
}

// IsWritable - Tells whether the filename is writable
func IsWritable(name string) bool {
	_, err := syscall.Open(name, syscall.O_WRONLY, 0)
	return err == nil
}

// IsWriteable is_writeable()
func IsWriteable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_WRONLY, 0)
	if err != nil {
		return false
	}
	return true
}

// Mkdir - Makes directory
// Mkdir mkdir()
func Mkdir(filename string, mode os.FileMode) error {
	return os.Mkdir(filename, mode)
}

func MkdirAll(filename string, mode os.FileMode)  error {
	return os.MkdirAll(filename, mode)
}

// Realpath - Returns canonicalized absolute pathname
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// Rename - Renames a file or directory
func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// Rmdir — Removes directory
func Rmdir(path string) error {
	return os.RemoveAll(path)
}

// Stat - Gives information about a file
func Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Unlink - Deletes a file
func Unlink(name string) error {
	return os.Remove(name)
}
