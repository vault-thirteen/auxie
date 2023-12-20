package header

// Common message header names for various protocols.

// This package contains all IANA registered header names which have an RFC
// reference as of 2023-03-21.
// For more information visit this URL:
// https://www.iana.org/assignments/message-headers/message-headers.xml
// This list may contain some old deprecated and obsolete header names.
//
// HTTP field name registrations have been moved to:
// https://www.iana.org/assignments/http-fields/http-fields.xhtml

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
