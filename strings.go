package helpers

import (
	"bytes"
	"html"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Constants for StrPad
const (
	StrPadRight = "STR_PAD_RIGHT"
	StrPadLeft  = "STR_PAD_LEFT"
)

// Bin2hex bin2hex()
func Bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

// Bindec bindec()
func Bindec(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 10), nil
}

// Hex2bin hex2bin()
func Hex2bin(data string) (string, error) {
	i, err := strconv.ParseInt(data, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

// Chr - Return a specific character
func Chr(ascii int) string {
	for ascii < 0 {
		ascii += 256
	}
	return string(ascii % 256)
}

// Ord ord()
func Ord(char string) int {
	r, _ := utf8.DecodeRune([]byte(char))
	return int(r)
}

// Explode explode()
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

// GetHtmlTranslationTable - Returns the translation table used by htmlspecialchars() and htmlentities()
func GetHtmlTranslationTable() map[string]string {
	return map[string]string{
		`"`: "&quot;",
		`&`: "&amp;",
		`<`: "&lt;",
		`>`: "&gt;",
	}
}

// Htmlspecialchars - Convert special characters to HTML entities
func Htmlspecialchars(s string) string {
	return html.EscapeString(s)
}

// HtmlspecialcharsDecode - Convert special HTML entities back to characters
func HtmlspecialcharsDecode(s string) string {
	return html.UnescapeString(s)
}

// Implode implode()
func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// Join - Alias of implode()
func Join(a []string, sep string) string {
	return Implode(a, sep)
}

//StripTags - Strip HTML and PHP tags from a string
func StripTags(s string) string {
	reg, _ := regexp.Compile(`<[\S\s]+?>`)
	s = reg.ReplaceAllStringFunc(s, strings.ToLower)
	//remove style
	reg, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	s = reg.ReplaceAllString(s, "")
	//remove script
	reg, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	s = reg.ReplaceAllString(s, "")

	reg, _ = regexp.Compile(`<[\S\s]+?>`)
	s = reg.ReplaceAllString(s, "\n")

	reg, _ = regexp.Compile(`\s{2,}`)
	s = reg.ReplaceAllString(s, "\n")

	return strings.TrimSpace(s)
}

// Trim trim()
func Trim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, characterMask[0])
}

// Ltrim ltrim()
func Ltrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	return strings.TrimLeft(str, characterMask[0])
}

// Rtrim rtrim()
func Rtrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

// Nl2br - Inserts HTML line breaks before all newlines in a string
func Nl2br(str string, isXhtml bool) string {
	r, n, runes := '\r', '\n', []rune(str)
	var br []byte
	if isXhtml {
		br = []byte("<br />")
	} else {
		br = []byte("<br>")
	}
	skip := false
	length := len(runes)
	var buf bytes.Buffer
	for i, v := range runes {
		if skip {
			skip = false
			continue
		}
		switch v {
		case n, r:
			if (i+1 < length) && (v == r && runes[i+1] == n) || (v == n && runes[i+1] == r) {
				buf.Write(br)
				skip = true
				continue
			}
			buf.Write(br)
		default:
			buf.WriteRune(v)
		}
	}
	return buf.String()
}

// StrPad - Pad a string to a certain length with another string
func StrPad(s string, length int, args ...string) string {
	runes := []rune(s)
	l := len(runes)
	if l > length {
		return s
	}
	padString := " "
	padType := StrPadRight
	if len(args) > 1 {
		padString = args[0]
		padType = args[1]
	} else if len(args) > 0 {
		padString = args[0]
	}

	padStringLen := len([]rune(padString))
	count := (length-l)/padStringLen + 1
	out := ""
	padString = strings.Repeat(padString, count)
	if padType == StrPadLeft {
		out = string([]rune(padString)[:length-l]) + s
	} else {
		out = s + string([]rune(padString)[:length-l])
	}
	return out
}

// StrRepeat - Repeat a string
func StrRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// StrReplace str_replace()
func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

// Strtolower strtolower()
func Strtolower(str string) string {
	return strings.ToLower(str)
}

// Strtoupper strtoupper()
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}



// Strpos strpos()
func Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}


// Stripos stripos()
func Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	haystack = haystack[offset:]
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}


// Strrpos strrpos()
func Strrpos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(haystack, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// Strripos strripos()
func Strripos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(strings.ToLower(haystack), strings.ToLower(needle))
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// Strrchr - Find the last occurrence of a character in a string
func Strrchr(s, substr string) string {
	i := strings.LastIndex(s, substr)
	if i < 0 {
		return ""
	}
	return s[i:]
}

// Strlen - Get string length
func Strlen(s string) int {
	return len(s)
}

// MbStrlen - Get string length
func MbStrlen(s string) int {
	return utf8.RuneCountInString(s)
}

// Strrev strrev()
func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Substr substr()
func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}


// MbSubstr - Get part of string
func MbSubstr(s string, start int, length ...int) string {
	runes := []rune(s)
	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(runes) + l
			if end < 0 {
				end = 0
			}
			return string(runes[start:end])
		}
		end := start + l
		if end > len(runes) {
			end = len(runes)
		}
		return string(runes[start:end])
	}
	return string(runes[start:])
}

// SubstrCount - Count the number of substring occurrences
func SubstrCount(s, substr string) int {
	return strings.Count(s, substr)
}

// Ucfirst ucfirst()
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

// Ucwords ucwords()
func Ucwords(str string) string {
	return strings.Title(str)
}

