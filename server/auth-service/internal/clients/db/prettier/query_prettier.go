package prettier

import (
	"fmt"
	"strconv"
	"strings"
)

// Placeholder constants for query parameter formatting.
const (
	PlaceholderDollar   = "$"
	PlaceholderQuestion = "?"
)

// Pretty formats the query string by replacing placeholders with provided argument values.
// It also removes tabs and newlines for cleaner output.
func Pretty(query string, placeholder string, args ...any) string {
	for i, param := range args {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("%q", v)
		case []byte:
			value = fmt.Sprintf("%q", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}

		query = strings.Replace(query, fmt.Sprintf("%s%s", placeholder, strconv.Itoa(i+1)), value, -1)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.TrimSpace(query)
}
