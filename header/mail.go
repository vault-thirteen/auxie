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
	MailHeaderResentReplyTo                       = "Resent-Reply-To"                        // [RFC5322]
)

// Status: Standard.
const (
	MailHeaderArchivedAt                  = "Archived-At"                   // [RFC5064]
	MailHeaderAuthenticationResults       = "Authentication-Results"        // [RFC8601]
	MailHeaderAutoSubmitted               = "Auto-Submitted"                // [RFC3834 section 5]
	MailHeaderBcc                         = "Bcc"                           // [RFC5322]
	MailHeaderCc                          = "Cc"                            // [RFC5322]
	MailHeaderComments                    = "Comments"                      // [RFC5322]
	MailHeaderDate                        = "Date"                          // [RFC5322]
	MailHeaderDKIMSignature               = "DKIM-Signature"                // [RFC6376]
	MailHeaderDowngradedFinalRecipient    = "Downgraded-Final-Recipient"    // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedInReplyTo         = "Downgraded-In-Reply-To"        // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedMessageId         = "Downgraded-Message-Id"         // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedOriginalRecipient = "Downgraded-Original-Recipient" // [RFC6857 Section 3.1.10]
	MailHeaderDowngradedReferences        = "Downgraded-References"         // [RFC6857 Section 3.1.10]
	MailHeaderFrom                        = "From"                          // [RFC5322][RFC6854]
	MailHeaderInReplyTo                   = "In-Reply-To"                   // [RFC5322]
	MailHeaderKeywords                    = "Keywords"                      // [RFC5322]
	MailHeaderListUnsubscribePost         = "List-Unsubscribe-Post"         // [RFC8058]
	MailHeaderMessageID                   = "Message-ID"                    // [RFC5322]
	MailHeaderMTPriority                  = "MT-Priority"                   // [RFC6758]
	MailHeaderOriginalFrom                = "Original-From"                 // [RFC5703]
	MailHeaderOriginalRecipient           = "Original-Recipient"            // [RFC3798][RFC5337]
	MailHeaderOriginalSubject             = "Original-Subject"              // [RFC5703]
	MailHeaderReceived                    = "Received"                      // [RFC5322][RFC5321]
	MailHeaderReceivedSPF                 = "Received-SPF"                  // [RFC7208]
	MailHeaderReferences                  = "References"                    // [RFC5322]
	MailHeaderReplyTo                     = "Reply-To"                      // [RFC5322]
	MailHeaderRequireRecipientValidSince  = "Require-Recipient-Valid-Since" // [RFC7293]
	MailHeaderResentBcc                   = "Resent-Bcc"                    // [RFC5322]
	MailHeaderResentCc                    = "Resent-Cc"                     // [RFC5322]
	MailHeaderResentDate                  = "Resent-Date"                   // [RFC5322]
	MailHeaderResentFrom                  = "Resent-From"                   // [RFC5322][RFC6854]
	MailHeaderResentMessageID             = "Resent-Message-ID"             // [RFC5322]
	MailHeaderResentSender                = "Resent-Sender"                 // [RFC5322][RFC6854]
	MailHeaderResentTo                    = "Resent-To"                     // [RFC5322]
	MailHeaderReturnPath                  = "Return-Path"                   // [RFC5322]
	MailHeaderSender                      = "Sender"                        // [RFC5322][RFC6854]
	MailHeaderSubject                     = "Subject"                       // [RFC5322]
	MailHeaderTLSReportDomain             = "TLS-Report-Domain"             // [RFC8460]
	MailHeaderTLSReportSubmitter          = "TLS-Report-Submitter"          // [RFC8460]
	MailHeaderTLSRequired                 = "TLS-Required"                  // [RFC8689]
	MailHeaderTo                          = "To"                            // [RFC5322]
	MailHeaderVBRInfo                     = "VBR-Info"                      // [RFC5518]
)

// Status: (Empty).
const (
	MailHeaderAcceptLanguage                  = "Accept-Language"                    //	[RFC4021]
	MailHeaderAlternateRecipient              = "Alternate-Recipient"                //	[RFC4021]
	MailHeaderApparentlyTo                    = "Apparently-To"                      //	[RFC2076]
	MailHeaderAuthor                          = "Author"                             //	[RFC9057]
	MailHeaderAutoforwarded                   = "Autoforwarded"                      //	[RFC4021]
	MailHeaderAutosubmitted                   = "Autosubmitted"                      //	[RFC4021]
	MailHeaderCFBLAddress                     = "CFBL-Address"                       //	[RFC-benecke-cfbl-address-header-13]
	MailHeaderCFBLFeedbackID                  = "CFBL-Feedback-ID"                   //	[RFC-benecke-cfbl-address-header-13]
	MailHeaderContentIdentifier               = "Content-Identifier"                 //	[RFC4021]
	MailHeaderContentReturn                   = "Content-Return"                     //	[RFC4021]
	MailHeaderConversion                      = "Conversion"                         //	[RFC4021]
	MailHeaderConversionWithLoss              = "Conversion-With-Loss"               //	[RFC4021]
	MailHeaderDeferredDelivery                = "Deferred-Delivery"                  //	[RFC4021]
	MailHeaderDeliveredTo                     = "Delivered-To"                       //	[RFC9228]
	MailHeaderDeliveryDate                    = "Delivery-Date"                      //	[RFC4021]
	MailHeaderDiscardedX400IPMSExtensions     = "Discarded-X400-IPMS-Extensions"     //	[RFC4021]
	MailHeaderDiscardedX400MTSExtensions      = "Discarded-X400-MTS-Extensions"      //	[RFC4021]
	MailHeaderDiscloseRecipients              = "Disclose-Recipients"                //	[RFC4021]
	MailHeaderDispositionNotificationOptions  = "Disposition-Notification-Options"   //	[RFC4021]
	MailHeaderDispositionNotificationTo       = "Disposition-Notification-To"        //	[RFC4021]
	MailHeaderDLExpansionHistory              = "DL-Expansion-History"               //	[RFC4021]
	MailHeaderEDIINTFeatures                  = "EDIINT-Features"                    //	[RFC6017]
	MailHeaderEesstVersion                    = "Eesst-Version"                      //	[RFC7681]
	MailHeaderEncoding                        = "Encoding"                           //	[RFC4021]
	MailHeaderEncrypted                       = "Encrypted"                          //	[RFC4021]
	MailHeaderErrorsTo                        = "Errors-To"                          //	[RFC2076]
	MailHeaderExpires                         = "Expires"                            //	[RFC4021]
	MailHeaderExpiryDate                      = "Expiry-Date"                        //	[RFC4021]
	MailHeaderFormSub                         = "Form-Sub"                           //	[draft-levine-mailbomb-header]
	MailHeaderGenerateDeliveryReport          = "Generate-Delivery-Report"           //	[RFC4021]
	MailHeaderImportance                      = "Importance"                         //	[RFC4021]
	MailHeaderIncompleteCopy                  = "Incomplete-Copy"                    //	[RFC4021]
	MailHeaderJabberID                        = "Jabber-ID"                          //	[RFC7259]
	MailHeaderLanguage                        = "Language"                           //	[RFC4021]
	MailHeaderLatestDeliveryTime              = "Latest-Delivery-Time"               //	[RFC4021]
	MailHeaderListArchive                     = "List-Archive"                       //	[RFC4021]
	MailHeaderListHelp                        = "List-Help"                          //	[RFC4021]
	MailHeaderListID                          = "List-ID"                            //	[RFC4021]
	MailHeaderListOwner                       = "List-Owner"                         //	[RFC4021]
	MailHeaderListPost                        = "List-Post"                          //	[RFC4021]
	MailHeaderListSubscribe                   = "List-Subscribe"                     //	[RFC4021]
	MailHeaderListUnsubscribe                 = "List-Unsubscribe"                   //	[RFC4021]
	MailHeaderMessageContext                  = "Message-Context"                    //	[RFC4021]
	MailHeaderMessageType                     = "Message-Type"                       //	[RFC4021]
	MailHeaderMMHSAcp127MessageIdentifier     = "MMHS-Acp127-Message-Identifier"     //	[RFC6477][ACP123
	MailHeaderMMHSAuthorizingUsers            = "MMHS-Authorizing-Users"             //	[RFC7912]
	MailHeaderMMHSCodressMessageIndicator     = "MMHS-Codress-Message-Indicator"     //	[RFC6477][ACP123
	MailHeaderMMHSCopyPrecedence              = "MMHS-Copy-Precedence"               //	[RFC6477][ACP123
	MailHeaderMMHSExemptedAddress             = "MMHS-Exempted-Address"              //	[RFC6477][ACP123
	MailHeaderMMHSExtendedAuthorisationInfo   = "MMHS-Extended-Authorisation-Info"   //	[RFC6477][ACP123
	MailHeaderMMHSHandlingInstructions        = "MMHS-Handling-Instructions"         //	[RFC6477][ACP123
	MailHeaderMMHSMessageInstructions         = "MMHS-Message-Instructions"          //	[RFC6477][ACP123
	MailHeaderMMHSMessageType                 = "MMHS-Message-Type"                  //	[RFC6477][ACP123
	MailHeaderMMHSOriginatorPLAD              = "MMHS-Originator-PLAD"               //	[RFC6477][ACP123
	MailHeaderMMHSOriginatorReference         = "MMHS-Originator-Reference"          //	[RFC6477][ACP123
	MailHeaderMMHSOtherRecipientsIndicatorCC  = "MMHS-Other-Recipients-Indicator-CC" //	[RFC6477][ACP123
	MailHeaderMMHSOtherRecipientsIndicatorTo  = "MMHS-Other-Recipients-Indicator-To" //	[RFC6477][ACP123
	MailHeaderMMHSPrimaryPrecedence           = "MMHS-Primary-Precedence"            //	[RFC6477][ACP123
	MailHeaderMMHSSubjectIndicatorCodes       = "MMHS-Subject-Indicator-Codes"       //	[RFC6477][ACP123
	MailHeaderObsoletes                       = "Obsoletes"                          //	[RFC4021]
	MailHeaderOriginalEncodedInformationTypes = "Original-Encoded-Information-Types" //	[RFC4021]
	MailHeaderOriginalMessageID               = "Original-Message-ID"                //	[RFC4021]
	MailHeaderOriginatorReturnAddress         = "Originator-Return-Address"          //	[RFC4021]
	MailHeaderPICSLabel                       = "PICS-Label"                         //	[RFC4021]
	MailHeaderPreventNonDeliveryReport        = "Prevent-NonDelivery-Report"         //	[RFC4021]
	MailHeaderPriority                        = "Priority"                           //	[RFC4021]
	MailHeaderPrivicon                        = "Privicon"                           //	[draft-koenig-privicons]
	MailHeaderReplyBy                         = "Reply-By"                           //	[RFC4021]
	MailHeaderSensitivity                     = "Sensitivity"                        //	[RFC4021]
	MailHeaderSIOLabel                        = "SIO-Label"                          //	[RFC7444]
	MailHeaderSIOLabelHistory                 = "SIO-Label-History"                  //	[RFC7444]
	MailHeaderSolicitation                    = "Solicitation"                       //	[RFC3865]
	MailHeaderSupersedes                      = "Supersedes"                         //	[RFC4021]
	MailHeaderXMittente                       = "X-Mittente"                         //	[RFC6109]
	MailHeaderXRicevuta                       = "X-Ricevuta"                         //	[RFC6109]
	MailHeaderXRiferimentoMessageID           = "X-Riferimento-Message-ID"           //	[RFC6109]
	MailHeaderXTipoRicevuta                   = "X-TipoRicevuta"                     //	[RFC6109]
	MailHeaderXTrasporto                      = "X-Trasporto"                        //	[RFC6109]
	MailHeaderXVerificaSicurezza              = "X-VerificaSicurezza"                //	[RFC6109]
	MailHeaderX400ContentIdentifier           = "X400-Content-Identifier"            //	[RFC4021]
	MailHeaderX400ContentReturn               = "X400-Content-Return"                //	[RFC4021]
	MailHeaderX400ContentType                 = "X400-Content-Type"                  //	[RFC4021]
	MailHeaderX400MTSIdentifier               = "X400-MTS-Identifier"                //	[RFC4021]
	MailHeaderX400Originator                  = "X400-Originator"                    //	[RFC4021]
	MailHeaderX400Received                    = "X400-Received"                      //	[RFC4021]
	MailHeaderX400Recipients                  = "X400-Recipients"                    //	[RFC4021]
	MailHeaderX400Trace                       = "X400-Trace"                         //	[RFC4021]
)
