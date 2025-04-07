package header

// Netnews protocol header field names.

// Status: Deprecated.
const (
	NetnewsHeaderLines       = "Lines"         // [RFC5536][RFC3977]
	NetnewsHeaderXArchivedAt = "X-Archived-At" // [RFC5064]
)

// Status: Obsoleted.
const (
	NetnewsHeaderAlsoControl     = "Also-Control"      // [RFC1849][RFC5536]
	NetnewsHeaderArticleNames    = "Article-Names"     // [RFC1849][RFC5536]
	NetnewsHeaderArticleUpdates  = "Article-Updates"   // [RFC1849][RFC5536]
	NetnewsHeaderDateReceived    = "Date-Received"     // [RFC850][RFC5536]
	NetnewsHeaderNNTPPostingDate = "NNTP-Posting-Date" // [RFC5536]
	NetnewsHeaderNNTPPostingHost = "NNTP-Posting-Host" // [RFC2980][RFC5536]
	NetnewsHeaderPostingVersion  = "Posting-Version"   // [RFC850][RFC5536]
	NetnewsHeaderRelayVersion    = "Relay-Version"     // [RFC850][RFC5536]
	NetnewsHeaderSeeAlso         = "See-Also"          // [RFC1849][RFC5536]
)

// Status: Standard.
const (
	NetnewsHeaderApproved       = "Approved"        // [RFC5536]
	NetnewsHeaderArchive        = "Archive"         // [RFC5536]
	NetnewsHeaderArchivedAt     = "Archived-At"     // [RFC5064]
	NetnewsHeaderCancelKey      = "Cancel-Key"      // [RFC8315]
	NetnewsHeaderCancelLock     = "Cancel-Lock"     // [RFC8315]
	NetnewsHeaderComments       = "Comments"        // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderControl        = "Control"         // [RFC5536]
	NetnewsHeaderDate           = "Date"            // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderDistribution   = "Distribution"    // [RFC5536]
	NetnewsHeaderExpires        = "Expires"         // [RFC5536]
	NetnewsHeaderFollowupTo     = "Followup-To"     // [RFC5536]
	NetnewsHeaderFrom           = "From"            // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderInjectionDate  = "Injection-Date"  // [RFC5536]
	NetnewsHeaderInjectionInfo  = "Injection-Info"  // [RFC5536]
	NetnewsHeaderKeywords       = "Keywords"        // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderMessageID      = "Message-ID"      // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderNewsgroups     = "Newsgroups"      // [RFC5536]
	NetnewsHeaderOrganization   = "Organization"    // [RFC5536]
	NetnewsHeaderOriginalSender = "Original-Sender" // [RFC5537]
	NetnewsHeaderPath           = "Path"            // [RFC5536]
	NetnewsHeaderReferences     = "References"      // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderReplyTo        = "Reply-To"        // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderSender         = "Sender"          // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderSubject        = "Subject"         // [RFC5536][RFC-ietf-emailcore-rfc5322bis-12]
	NetnewsHeaderSummary        = "Summary"         // [RFC5536]
	NetnewsHeaderSupersedes     = "Supersedes"      // [RFC5536][RFC2156]
	NetnewsHeaderUserAgent      = "User-Agent"      // [RFC5536][RFC2616]
	NetnewsHeaderXref           = "Xref"            // [RFC5536]
)

// Status: (Empty).
const (
	NetnewsHeaderFace     = "Face"      // [https://quimby.gnus.org/circus/face]
	NetnewsHeaderJabberID = "Jabber-ID" // [RFC7259]
	NetnewsHeaderXFace    = "X-Face"    // [https://purl.org/x-face-spec]
	NetnewsHeaderXPGPSig  = "X-PGP-Sig" // [ftp://ftp.isc.org/pub/pgpcontrol/FORMAT][https://ftp.isc.org/pub/pgpcontrol/FORMAT]
)
