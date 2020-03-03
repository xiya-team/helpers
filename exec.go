package helpers

import (
	"os/exec"
	"strings"
	"unicode"
)

// Exec - Execute an external program
func Exec(command string, output *[]string, returnVar *int) string {
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})
	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}
	cmd := exec.Command(parts[0], parts[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		*returnVar = 1
		return ""
	}
	*returnVar = 0
	*output = strings.Split(strings.TrimRight(string(out), "\n"), "\n")
	if l := len(*output); l > 0 {
		return (*output)[l-1]
	}
	return ""
}