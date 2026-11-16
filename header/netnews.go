package header

// Netnews protocol header field names.

// Status: (Empty).
const (
	NetnewsHeaderFace     = "Face"      // [https://quimby.gnus.org/circus/face]
	NetnewsHeaderJabberID = "Jabber-ID" // [RFC 7259]
	NetnewsHeaderXFace    = "X-Face"    // [https://purl.org/x-face-spec]
	NetnewsHeaderXPGPSig  = "X-PGP-Sig" // [ftp://ftp.isc.org/pub/pgpcontrol/FORMAT]; [https://ftp.isc.org/pub/pgpcontrol/FORMAT]
)

// Status: Deprecated.
const (
	NetnewsHeaderLines       = "Lines"         // [RFC 5536]; [RFC 3977]
	NetnewsHeaderXArchivedAt = "X-Archived-At" // [RFC 5064]
)

// Status: Experimental.
const ()

// Status: Informational.
const ()

// Status: Obsoleted.
const (
	NetnewsHeaderAlsoControl     = "Also-Control"      // [RFC 1849]; [RFC 5536]
	NetnewsHeaderArticleNames    = "Article-Names"     // [RFC 1849]; [RFC 5536]
	NetnewsHeaderArticleUpdates  = "Article-Updates"   // [RFC 1849]; [RFC 5536]
	NetnewsHeaderDateReceived    = "Date-Received"     // [RFC 850]; [RFC 5536]
	NetnewsHeaderNNTPPostingDate = "NNTP-Posting-Date" // [RFC 5536]
	NetnewsHeaderNNTPPostingHost = "NNTP-Posting-Host" // [RFC 2980]; [RFC 5536]
	NetnewsHeaderPostingVersion  = "Posting-Version"   // [RFC 850]; [RFC 5536]
	NetnewsHeaderRelayVersion    = "Relay-Version"     // [RFC 850]; [RFC 5536]
	NetnewsHeaderSeeAlso         = "See-Also"          // [RFC 1849]; [RFC 5536]
)

// Status: Permanent.
const ()

// Status: Provisional.
const ()

// Status: Reserved.
const ()

// Status: Standard.
const (
	NetnewsHeaderApproved       = "Approved"        // [RFC 5536]
	NetnewsHeaderArchive        = "Archive"         // [RFC 5536]
	NetnewsHeaderArchivedAt     = "Archived-At"     // [RFC 5064]
	NetnewsHeaderCancelKey      = "Cancel-Key"      // [RFC 8315]
	NetnewsHeaderCancelLock     = "Cancel-Lock"     // [RFC 8315]
	NetnewsHeaderComments       = "Comments"        // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderControl        = "Control"         // [RFC 5536]
	NetnewsHeaderDate           = "Date"            // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderDistribution   = "Distribution"    // [RFC 5536]
	NetnewsHeaderExpires        = "Expires"         // [RFC 5536]
	NetnewsHeaderFollowupTo     = "Followup-To"     // [RFC 5536]
	NetnewsHeaderFrom           = "From"            // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderInjectionDate  = "Injection-Date"  // [RFC 5536]
	NetnewsHeaderInjectionInfo  = "Injection-Info"  // [RFC 5536]
	NetnewsHeaderKeywords       = "Keywords"        // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderMessageID      = "Message-ID"      // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderNewsgroups     = "Newsgroups"      // [RFC 5536]
	NetnewsHeaderOrganization   = "Organization"    // [RFC 5536]
	NetnewsHeaderOriginalSender = "Original-Sender" // [RFC 5537]
	NetnewsHeaderPath           = "Path"            // [RFC 5536]
	NetnewsHeaderReferences     = "References"      // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderReplyTo        = "Reply-To"        // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderSender         = "Sender"          // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderSubject        = "Subject"         // [RFC 5536]; [RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderSummary        = "Summary"         // [RFC 5536]
	NetnewsHeaderSupersedes     = "Supersedes"      // [RFC 5536]; [RFC 2156]
	NetnewsHeaderUserAgent      = "User-Agent"      // [RFC 5536]; [RFC 2616]
	NetnewsHeaderXref           = "Xref"            // [RFC 5536]
)
