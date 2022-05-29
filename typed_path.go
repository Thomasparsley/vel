package vel

import (
	"regexp"
	"strings"
)

var TypedParamRE = regexp.MustCompile(`\<[a-z]*\:[a-z]*\>`)

// TypedPath is a path string with a typed params. Example: `/user/<string:uuid>/<int:id>`
type TypedPath string

func (p TypedPath) UntypedPath() string {
	result := string(p)

	for _, path := range TypedParamRE.FindAllString(result, -1) {
		// Remove the <> and evrything before :
		toReplace := strings.Split(path, ":")[0]
		toReplace = strings.Replace(toReplace, ">", "", 1)

		// Replace the typed param with a untyped param
		result = strings.Replace(result, path, toReplace, 1)
	}

	return result
}
