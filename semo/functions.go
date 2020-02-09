package semo

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"text/template"
	"time"
)

var defaultFuncMap = template.FuncMap{
	"toUpper":      strings.ToUpper,
	"toLower":      strings.ToLower,
	"noWhitespace": removeWhitespace,
	"runID":        runID,
}

// A random string that is constant per semo
// invocation
var runIdentifier *string

// Remove whitespace and replace it with
// an underscore
func removeWhitespace(str string) string {
	re := regexp.MustCompile(`\s`)
	return re.ReplaceAllString(str, "_")
}

// Returns the id of this run with a maximum
// length of maxLen characters.
// If no id has been set, generate a new one
func runID(maxLen int) string {
	if runIdentifier == nil {
		baseString := fmt.Sprintf("%d%d",
			time.Now().UnixNano(), rand.Uint32())
		sum := sha256.Sum256([]byte(baseString))
		idHash := fmt.Sprintf("%x", sum)
		runIdentifier = &idHash
	}

	if maxLen > len(*runIdentifier) {
		maxLen = len(*runIdentifier)
	}

	return (*runIdentifier)[:maxLen]
}
