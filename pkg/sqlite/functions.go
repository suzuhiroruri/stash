package sqlite

import (
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func durationToTinyIntFn(str string) (int64, error) {
	splits := strings.Split(str, ":")

	if len(splits) > 3 {
		return 0, nil
	}

	seconds := 0
	factor := 1
	for len(splits) > 0 {
		// pop the last split
		var thisSplit string
		thisSplit, splits = splits[len(splits)-1], splits[:len(splits)-1]

		thisInt, err := strconv.Atoi(thisSplit)
		if err != nil {
			return 0, nil
		}

		seconds += factor * thisInt
		factor *= 60
	}

	return int64(seconds), nil
}

func basenameFn(str string) (string, error) {
	return filepath.Base(str), nil
}

// NormalizeNFC normalizes a string to NFC (Normalization Form Canonical Composition).
// This is needed to handle macOS NFD (Normalization Form Decomposed) file names,
// where combining characters like Japanese voiced marks are stored separately.
// See https://github.com/stashapp/stash/issues/4425
func NormalizeNFC(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	return norm.NFC.String(s)
}
