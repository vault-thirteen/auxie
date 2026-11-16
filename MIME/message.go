package mime

// Message.
const (
	TypeMessageBhttp                         = "message/bhttp"                           // [RFC 9292]
	TypeMessageCPIM                          = "message/CPIM"                            // [RFC 3862]
	TypeMessageDeliveryStatus                = "message/delivery-status"                 // [RFC 1894]
	TypeMessageDispositionNotification       = "message/disposition-notification"        // [RFC 8098]
	TypeMessageExample                       = "message/example"                         // [RFC 4735]
	TypeMessageExternalBody                  = "message/external-body"                   // [RFC 2045]; [RFC 2046]
	TypeMessageFeedbackReport                = "message/feedback-report"                 // [RFC 5965]
	TypeMessageGlobal                        = "message/global"                          // [RFC 6532]
	TypeMessageGlobalDeliveryStatus          = "message/global-delivery-status"          // [RFC 6533]
	TypeMessageGlobalDispositionNotification = "message/global-disposition-notification" // [RFC 6533]
	TypeMessageGlobalHeaders                 = "message/global-headers"                  // [RFC 6533]
	TypeMessageHttp                          = "message/http"                            // [RFC 9112]
	TypeMessageImdnXml                       = "message/imdn+xml"                        // [RFC 5438]
	TypeMessageMls                           = "message/mls"                             // [RFC 9420]
	TypeMessageNews                          = "message/news"                            // [RFC 5537]; [Henry_Spencer]
	TypeMessageOhttpChunkedReq               = "message/ohttp-chunked-req"               // [RFC-ietf-ohai-chunked-ohttp-08]
	TypeMessageOhttpChunkedRes               = "message/ohttp-chunked-res"               // [RFC-ietf-ohai-chunked-ohttp-08]
	TypeMessageOhttpReq                      = "message/ohttp-req"                       // [RFC 9458]
	TypeMessageOhttpRes                      = "message/ohttp-res"                       // [RFC 9458]
	TypeMessagePartial                       = "message/partial"                         // [RFC 2045]; [RFC 2046]
	TypeMessageRfc822                        = "message/rfc822"                          // [RFC 2045]; [RFC 2046]
	TypeMessageSHttp                         = "message/s-http"                          // [RFC 2660]; [Status change of HTTP experiments to Historic]
	TypeMessageSip                           = "message/sip"                             // [RFC 3261]
	TypeMessageSipfrag                       = "message/sipfrag"                         // [RFC 3420]
	TypeMessageTrackingStatus                = "message/tracking-status"                 // [RFC 3886]
	TypeMessageVndSiSimp                     = "message/vnd.si.simp"                     // [Nicholas_Parks_Young]
	TypeMessageVndWfaWsc                     = "message/vnd.wfa.wsc"                     // [Mick_Conley]
)

const (
	TypeMessageAny = "message/*"
)
