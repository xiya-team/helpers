package helpers

import(
	"os"
)

func FileExits(path string) bool {

	ret, _ := FileExitsDetail(path)

	return ret
}

func FileExitsDetail(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
