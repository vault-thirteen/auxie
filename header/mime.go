package header

// MIME protocol header field names.

// Status: Obsoleted.
const (
	MimeHeaderBase        = "Base"         // [RFC1808][RFC2068 Section 14.11]
	MimeHeaderContentBase = "Content-Base" // [RFC2110][RFC2557]
)

// Status: Standard.
const (
	MimeHeaderContentTranslationType = "Content-Translation-Type" // [RFC8255]
)

// Status: (Empty).
const (
	MimeHeaderContentAlternative      = "Content-Alternative"       // [RFC4021]
	MimeHeaderContentDescription      = "Content-Description"       // [RFC4021]
	MimeHeaderContentDisposition      = "Content-Disposition"       // [RFC4021]
	MimeHeaderContentDuration         = "Content-Duration"          // [RFC4021]
	MimeHeaderContentfeatures         = "Content-features"          // [RFC4021]
	MimeHeaderContentID               = "Content-ID"                // [RFC4021]
	MimeHeaderContentLanguage         = "Content-Language"          // [RFC4021]
	MimeHeaderContentLocation         = "Content-Location"          // [RFC4021]
	MimeHeaderContentMD5              = "Content-MD5"               // [RFC4021]
	MimeHeaderContentTransferEncoding = "Content-Transfer-Encoding" // [RFC4021]
	MimeHeaderContentType             = "Content-Type"              // [RFC4021][RFC-ietf-lamps-header-protection-25]
	MimeHeaderMIMEVersion             = "MIME-Version"              // [RFC4021]
)
