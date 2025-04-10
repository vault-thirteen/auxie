package header

// Mail protocol header field names.

// Status: Deprecated.
const (
	MailHeaderXArchivedAt = "X-Archived-At" // [RFC5064]
)

// Status: Experimental.
const (
	MailHeaderARCAuthenticationResults = "ARC-Authentication-Results" // [RFC8617]
	MailHeaderARCMessageSignature      = "ARC-Message-Signature"      // [RFC8617]
	MailHeaderARCSeal                  = "ARC-Seal"                   // [RFC8617]
)

// Status: Informational.
const (
	MailHeaderOrganization = "Organization" // [RFC7681]
)

// Status: Obsoleted.
const (
	MailHeaderDowngradedBcc                       = "Downgraded-Bcc"                         // [RFC5504][RFC6857]
	MailHeaderDowngradedCc                        = "Downgraded-Cc"                          // [RFC5504][RFC6857]
	MailHeaderDowngradedDispositionNotificationTo = "Downgraded-Disposition-Notification-To" // [RFC5504][RFC6857]
	MailHeaderDowngradedFrom                      = "Downgraded-From"                        // [RFC5504][RFC6857 Section 3.1.10]
	MailHeaderDowngradedMailFrom                  = "Downgraded-Mail-From"                   // [RFC5504][RFC6857 Section 3.1.10]
	MailHeaderDowngradedRcptTo                    = "Downgraded-Rcpt-To"                     // [RFC5504][RFC6857]
	MailHeaderDowngradedReplyTo                   = "Downgraded-Reply-To"                    // [RFC5504][RFC6857]
	MailHeaderDowngradedResentBcc                 = "Downgraded-Resent-Bcc"                  // [RFC5504][RFC6857]
	MailHeaderDowngradedResentCc                  = "Downgraded-Resent-Cc"                   // [RFC5504][RFC6857]
	MailHeaderDowngradedResentFrom                = "Downgraded-Resent-From"                 // [RFC5504][RFC6857]
	MailHeaderDowngradedResentReplyTo             = "Downgraded-Resent-Reply-To"             // [RFC5504][RFC6857]
	MailHeaderDowngradedResentSender              = "Downgraded-Resent-Sender"               // [RFC5504][RFC6857]
	MailHeaderDowngradedResentTo                  = "Downgraded-Resent-To"                   // [RFC5504][RFC6857]
	MailHeaderDowngradedReturnPath                = "Downgraded-Return-Path"                 // [RFC5504][RFC6857]
	MailHeaderDowngradedSender                    = "Downgraded-Sender"                      // [RFC5504][RFC6857]
	MailHeaderDowngradedTo                        = "Downgraded-To"                          // [RFC5504][RFC6857]
	MailHeaderResentReplyTo                       = "Resent-Reply-To"                        // [RFC-ietf-emailcore-rfc5322bis-12, 4.5.6]
)

// Status: Standard.
const (
	MailHeaderArchivedAt                  = "Archived-At"                   // [RFC5064]
	MailHeaderAuthenticationResults       = "Authentication-Results"        // [RFC8601]
	MailHeaderAutoSubmitted               = "Auto-Submitted"                // [RFC3834 section 5]
	MailHeaderBcc                         = "Bcc"                           // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.3]
	MailHeaderCc                          = "Cc"                            // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.3]
	MailHeaderComments                    = "Comments"                      // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.5]
	MailHeaderDate                        = "Date"                          // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.1]
	MailHeaderDKIMSignature               = "DKIM-Signature"                // [RFC6376]
	MailHeaderDowngradedFinalRecipient    = "Downgraded-Final-Recipient"    // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedInReplyTo         = "Downgraded-In-Reply-To"        // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedMessageId         = "Downgraded-Message-Id"         // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedOriginalRecipient = "Downgraded-Original-Recipient" // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedReferences        = "Downgraded-References"         // [RFC6857 Section 3.1.10]
	MailHeaderFrom                        = "From"                          // [RFC6854][RFC-ietf-emailcore-rfc5322bis-12, 3.6.2]
	MailHeaderHPOuter                     = "HP-Outer"                      // [RFC-ietf-lamps-header-protection-25 Section 2.2.1]
	MailHeaderInReplyTo                   = "In-Reply-To"                   // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.4]
	MailHeaderKeywords                    = "Keywords"                      // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.5]
	MailHeaderListUnsubscribePost         = "List-Unsubscribe-Post"         // [RFC8058]
	MailHeaderMessageID                   = "Message-ID"                    // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.4]
	MailHeaderMTPriority                  = "MT-Priority"                   // [RFC6758]
	MailHeaderOriginalFrom                = "Original-From"                 // [RFC5703]
	MailHeaderOriginalRecipient           = "Original-Recipient"            // [RFC3798][RFC5337]
	MailHeaderOriginalSubject             = "Original-Subject"              // [RFC5703]
	MailHeaderReceived                    = "Received"                      // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.7][draft-ietf-emailcore-rfc5321bis-42]
	MailHeaderReceivedSPF                 = "Received-SPF"                  // [RFC7208]
	MailHeaderReferences                  = "References"                    // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.4]
	MailHeaderReplyTo                     = "Reply-To"                      // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.2]
	MailHeaderRequireRecipientValidSince  = "Require-Recipient-Valid-Since" // [RFC7293]
	MailHeaderResentBcc                   = "Resent-Bcc"                    // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.6]
	MailHeaderResentCc                    = "Resent-Cc"                     // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.6]
	MailHeaderResentDate                  = "Resent-Date"                   // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.6]
	MailHeaderResentFrom                  = "Resent-From"                   // [RFC6854][RFC-ietf-emailcore-rfc5322bis-12, 3.6.6]
	MailHeaderResentMessageID             = "Resent-Message-ID"             // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.6]
	MailHeaderResentSender                = "Resent-Sender"                 // [RFC6854][RFC-ietf-emailcore-rfc5322bis-12, 3.6.6]
	MailHeaderResentTo                    = "Resent-To"                     // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.6]
	MailHeaderReturnPath                  = "Return-Path"                   // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.7][draft-ietf-emailcore-rfc5321bis-42]
	MailHeaderSender                      = "Sender"                        // [RFC6854][RFC-ietf-emailcore-rfc5322bis-12, 3.6.2]
	MailHeaderSubject                     = "Subject"                       // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.5]
	MailHeaderTLSReportDomain             = "TLS-Report-Domain"             // [RFC8460]
	MailHeaderTLSReportSubmitter          = "TLS-Report-Submitter"          // [RFC8460]
	MailHeaderTLSRequired                 = "TLS-Required"                  // [RFC8689]
	MailHeaderTo                          = "To"                            // [RFC-ietf-emailcore-rfc5322bis-12, 3.6.3]
	MailHeaderVBRInfo                     = "VBR-Info"                      // [RFC5518]
)

// Status: (Empty).
const (
	MailHeaderAcceptLanguage                  = "Accept-Language"                    // [RFC4021]
	MailHeaderAlternateRecipient              = "Alternate-Recipient"                // [RFC4021]
	MailHeaderApparentlyTo                    = "Apparently-To"                      // [RFC2076]
	MailHeaderAuthor                          = "Author"                             // [RFC9057]
	MailHeaderAutoforwarded                   = "Autoforwarded"                      // [RFC4021]
	MailHeaderAutosubmitted                   = "Autosubmitted"                      // [RFC4021]
	MailHeaderCFBLAddress                     = "CFBL-Address"                       // [RFC9477]
	MailHeaderCFBLFeedbackID                  = "CFBL-Feedback-ID"                   // [RFC9477]
	MailHeaderContentIdentifier               = "Content-Identifier"                 // [RFC4021]
	MailHeaderContentReturn                   = "Content-Return"                     // [RFC4021]
	MailHeaderConversion                      = "Conversion"                         // [RFC4021]
	MailHeaderConversionWithLoss              = "Conversion-With-Loss"               // [RFC4021]
	MailHeaderDeferredDelivery                = "Deferred-Delivery"                  // [RFC4021]
	MailHeaderDeliveredTo                     = "Delivered-To"                       // [RFC9228]
	MailHeaderDeliveryDate                    = "Delivery-Date"                      // [RFC4021]
	MailHeaderDiscardedX400IPMSExtensions     = "Discarded-X400-IPMS-Extensions"     // [RFC4021]
	MailHeaderDiscardedX400MTSExtensions      = "Discarded-X400-MTS-Extensions"      // [RFC4021]
	MailHeaderDiscloseRecipients              = "Disclose-Recipients"                // [RFC4021]
	MailHeaderDispositionNotificationOptions  = "Disposition-Notification-Options"   // [RFC4021]
	MailHeaderDispositionNotificationTo       = "Disposition-Notification-To"        // [RFC4021]
	MailHeaderDLExpansionHistory              = "DL-Expansion-History"               // [RFC4021]
	MailHeaderEDIINTFeatures                  = "EDIINT-Features"                    // [RFC6017]
	MailHeaderEesstVersion                    = "Eesst-Version"                      // [RFC7681]
	MailHeaderEncoding                        = "Encoding"                           // [RFC4021]
	MailHeaderEncrypted                       = "Encrypted"                          // [RFC4021]
	MailHeaderErrorsTo                        = "Errors-To"                          // [RFC2076]
	MailHeaderExpires                         = "Expires"                            // [RFC4021]
	MailHeaderExpiryDate                      = "Expiry-Date"                        // [RFC4021]
	MailHeaderFace                            = "Face"                               // [https://quimby.gnus.org/circus/face]
	MailHeaderFormSub                         = "Form-Sub"                           // [draft-levine-mailbomb-header-00]
	MailHeaderGenerateDeliveryReport          = "Generate-Delivery-Report"           // [RFC4021]
	MailHeaderImportance                      = "Importance"                         // [RFC4021]
	MailHeaderIncompleteCopy                  = "Incomplete-Copy"                    // [RFC4021]
	MailHeaderJabberID                        = "Jabber-ID"                          // [RFC7259]
	MailHeaderLanguage                        = "Language"                           // [RFC4021]
	MailHeaderLatestDeliveryTime              = "Latest-Delivery-Time"               // [RFC4021]
	MailHeaderListArchive                     = "List-Archive"                       // [RFC4021]
	MailHeaderListHelp                        = "List-Help"                          // [RFC4021]
	MailHeaderListID                          = "List-ID"                            // [RFC4021]
	MailHeaderListOwner                       = "List-Owner"                         // [RFC4021]
	MailHeaderListPost                        = "List-Post"                          // [RFC4021]
	MailHeaderListSubscribe                   = "List-Subscribe"                     // [RFC4021]
	MailHeaderListUnsubscribe                 = "List-Unsubscribe"                   // [RFC4021]
	MailHeaderMessageContext                  = "Message-Context"                    // [RFC4021]
	MailHeaderMessageType                     = "Message-Type"                       // [RFC4021]
	MailHeaderMMHSAcp127MessageIdentifier     = "MMHS-Acp127-Message-Identifier"     // [RFC6477][ACP123 Appendix A1.14 and Appendix B.116]
	MailHeaderMMHSAuthorizingUsers            = "MMHS-Authorizing-Users"             // [RFC7912]
	MailHeaderMMHSCodressMessageIndicator     = "MMHS-Codress-Message-Indicator"     // [RFC6477][ACP123 Appendix A1.6 and Appendix B.110]
	MailHeaderMMHSCopyPrecedence              = "MMHS-Copy-Precedence"               // [RFC6477][ACP123 Appendix A1.9 and Appendix B.102]
	MailHeaderMMHSExemptedAddress             = "MMHS-Exempted-Address"              // [RFC6477][ACP123 Appendix A1.1 and Appendix B.105]
	MailHeaderMMHSExtendedAuthorisationInfo   = "MMHS-Extended-Authorisation-Info"   // [RFC6477][ACP123 Appendix A1.2 and Appendix B.106]
	MailHeaderMMHSHandlingInstructions        = "MMHS-Handling-Instructions"         // [RFC6477][ACP123 Appendix A1.4 and Appendix B.108]
	MailHeaderMMHSMessageInstructions         = "MMHS-Message-Instructions"          // [RFC6477][ACP123 Appendix A1.5 and Appendix B.109]
	MailHeaderMMHSMessageType                 = "MMHS-Message-Type"                  // [RFC6477][ACP123 Appendix A1.10 and Appendix B.103]
	MailHeaderMMHSOriginatorPLAD              = "MMHS-Originator-PLAD"               // [RFC6477][ACP123 Appendix A1.15 and Appendix B.117]
	MailHeaderMMHSOriginatorReference         = "MMHS-Originator-Reference"          // [RFC6477][ACP123 Appendix A1.7 and Appendix B.111]
	MailHeaderMMHSOtherRecipientsIndicatorCC  = "MMHS-Other-Recipients-Indicator-CC" // [RFC6477][ACP123 Appendix A1.12 and Appendix B.113]
	MailHeaderMMHSOtherRecipientsIndicatorTo  = "MMHS-Other-Recipients-Indicator-To" // [RFC6477][ACP123 Appendix A1.12 and Appendix B.113]
	MailHeaderMMHSPrimaryPrecedence           = "MMHS-Primary-Precedence"            // [RFC6477][ACP123 Appendix A1.8 and Appendix B.101]
	MailHeaderMMHSSubjectIndicatorCodes       = "MMHS-Subject-Indicator-Codes"       // [RFC6477][ACP123 Appendix A1.3 and Appendix B.107]
	MailHeaderObsoletes                       = "Obsoletes"                          // [RFC4021]
	MailHeaderOriginalEncodedInformationTypes = "Original-Encoded-Information-Types" // [RFC4021]
	MailHeaderOriginalMessageID               = "Original-Message-ID"                // [RFC4021]
	MailHeaderOriginatorReturnAddress         = "Originator-Return-Address"          // [RFC4021]
	MailHeaderPICSLabel                       = "PICS-Label"                         // [RFC4021]
	MailHeaderPreventNonDeliveryReport        = "Prevent-NonDelivery-Report"         // [RFC4021]
	MailHeaderPriority                        = "Priority"                           // [RFC4021]
	MailHeaderPrivicon                        = "Privicon"                           // [draft-koenig-privicons-01]
	MailHeaderReplyBy                         = "Reply-By"                           // [RFC4021]
	MailHeaderSensitivity                     = "Sensitivity"                        // [RFC4021]
	MailHeaderSIOLabel                        = "SIO-Label"                          // [RFC7444]
	MailHeaderSIOLabelHistory                 = "SIO-Label-History"                  // [RFC7444]
	MailHeaderSolicitation                    = "Solicitation"                       // [RFC3865]
	MailHeaderSupersedes                      = "Supersedes"                         // [RFC4021]
	MailHeaderWrongRecipient                  = "Wrong-Recipient"                    // [draft-ietf-mailmaint-wrong-recipient-00]
	MailHeaderXFace                           = "X-Face"                             // [https://purl.org/x-face-spec]
	MailHeaderXMittente                       = "X-Mittente"                         // [RFC6109]
	MailHeaderXRicevuta                       = "X-Ricevuta"                         // [RFC6109]
	MailHeaderXRiferimentoMessageID           = "X-Riferimento-Message-ID"           // [RFC6109]
	MailHeaderXTipoRicevuta                   = "X-TipoRicevuta"                     // [RFC6109]
	MailHeaderXTrasporto                      = "X-Trasporto"                        // [RFC6109]
	MailHeaderXVerificaSicurezza              = "X-VerificaSicurezza"                // [RFC6109]
	MailHeaderX400ContentIdentifier           = "X400-Content-Identifier"            // [RFC4021]
	MailHeaderX400ContentReturn               = "X400-Content-Return"                // [RFC4021]
	MailHeaderX400ContentType                 = "X400-Content-Type"                  // [RFC4021]
	MailHeaderX400MTSIdentifier               = "X400-MTS-Identifier"                // [RFC4021]
	MailHeaderX400Originator                  = "X400-Originator"                    // [RFC4021]
	MailHeaderX400Received                    = "X400-Received"                      // [RFC4021]
	MailHeaderX400Recipients                  = "X400-Recipients"                    // [RFC4021]
	MailHeaderX400Trace                       = "X400-Trace"                         // [RFC4021]
)
