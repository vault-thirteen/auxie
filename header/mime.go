package header

// MIME protocol header field names.

// Status: (Empty).
const (
	MimeHeaderContentAlternative      = "Content-Alternative"       // [RFC 4021]
	MimeHeaderContentDescription      = "Content-Description"       // [RFC 4021]
	MimeHeaderContentDisposition      = "Content-Disposition"       // [RFC 4021]
	MimeHeaderContentDuration         = "Content-Duration"          // [RFC 4021]
	MimeHeaderContentID               = "Content-ID"                // [RFC 4021]
	MimeHeaderContentLanguage         = "Content-Language"          // [RFC 4021]
	MimeHeaderContentLocation         = "Content-Location"          // [RFC 4021]
	MimeHeaderContentMD5              = "Content-MD5"               // [RFC 4021]
	MimeHeaderContentTransferEncoding = "Content-Transfer-Encoding" // [RFC 4021]
	MimeHeaderContentType             = "Content-Type"              // [RFC 4021]; [RFC 9788]
	MimeHeaderContentfeatures         = "Content-features"          // [RFC 4021]
	MimeHeaderMIMEVersion             = "MIME-Version"              // [RFC 4021]
)

// Status: Deprecated.
const ()

// Status: Experimental.
const ()

// Status: Informational.
const ()

// Status: Obsoleted.
const (
	MimeHeaderBase        = "Base"         // [RFC 1808]; [RFC 2068, Section 14.11]
	MimeHeaderContentBase = "Content-Base" // [RFC 2110]; [RFC 2557]
)

// Status: Permanent.
const ()

// Status: Provisional.
const ()

// Status: Reserved.
const ()

// Status: Standard.
const (
	MimeHeaderContentTranslationType = "Content-Translation-Type" // [RFC 8255]
)
