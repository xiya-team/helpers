package helpers

import "os"

// Exit exit()
func Exit(status int) {
	os.Exit(status)
}

// Die die()
func Die(status int) {
	os.Exit(status)
}