package header

// Common message header names for various protocols.

// Delimiters.
const (
	DelimiterCommaSpace = ", "
)

// MakeListOfHeaders composes a list of headers delimited by a comma and
// space.
func MakeListOfHeaders(headers []string) (headersList string) {
	if len(headers) == 0 {
		return ""
	}

	for _, header := range headers {
		headersList = headersList + header + DelimiterCommaSpace
	}

	return headersList[:len(headersList)-2]
}
