package helpers

import "os"

/**
获取当前环境
*/
func GetENV() (env string) {
	env = os.Getenv("CURRENTENV")
	if Empty(env) {
		env = "develop"
	}
	return
}

