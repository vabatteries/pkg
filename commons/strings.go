package commons

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var caserTitle cases.Caser = cases.Title(language.English, cases.NoLower)

var whitespace = regexp.MustCompile(`\s+`)

func StringTitle(s string) string {
	return caserTitle.String(s)
}

// StringNormalize will replace arbitrary lengths of consecutive white space with a single one.
func StringNormalize(s string) string {
	s = whitespace.ReplaceAllString(s, " ")

  return strings.TrimSpace(s)
}

//StringIsBlank checks if a string is blank. Will trim the string before the check takes place.
func StringIsBlank(s string) bool {
  return len(strings.TrimSpace(s)) == 0
}

// StringIsNotBlank check if a string is NOT blank.
func StringIsNotBlank(s string) bool {
  return !StringIsBlank(s)
}


func StringMissing(a, b string) string {
	dmp := diffmatchpatch.New()
  diffs := dmp.DiffMain(a, b, false)

  var missing string

  for _, d := range diffs {
    if d.Type == diffmatchpatch.DiffDelete {
      missing += d.Text
    }
  }

  return missing
}
