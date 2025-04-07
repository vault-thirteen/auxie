package mime

// Message.
const (
	TypeMessageBhttp                         = "message/bhttp"                           // [RFC9292]
	TypeMessageCPIM                          = "message/CPIM"                            // [RFC3862]
	TypeMessageDeliveryStatus                = "message/delivery-status"                 // [RFC1894]
	TypeMessageDispositionNotification       = "message/disposition-notification"        // [RFC8098]
	TypeMessageExample                       = "message/example"                         // [RFC4735]
	TypeMessageExternalBody                  = "message/external-body"                   // [RFC2045][RFC2046]
	TypeMessageFeedbackReport                = "message/feedback-report"                 // [RFC5965]
	TypeMessageGlobal                        = "message/global"                          // [RFC6532]
	TypeMessageGlobalDeliveryStatus          = "message/global-delivery-status"          // [RFC6533]
	TypeMessageGlobalDispositionNotification = "message/global-disposition-notification" // [RFC6533]
	TypeMessageGlobalHeaders                 = "message/global-headers"                  // [RFC6533]
	TypeMessageHttp                          = "message/http"                            // [RFC9112]
	TypeMessageImdnXml                       = "message/imdn+xml"                        // [RFC5438]
	TypeMessageMls                           = "message/mls"                             // [RFC9420]
	TypeMessageNews                          = "message/news"                            // [RFC5537][Henry_Spencer]
	TypeMessageOhttpReq                      = "message/ohttp-req"                       // [RFC9458]
	TypeMessageOhttpRes                      = "message/ohttp-res"                       // [RFC9458]
	TypeMessagePartial                       = "message/partial"                         // [RFC2045][RFC2046]
	TypeMessageRfc822                        = "message/rfc822"                          // [RFC2045][RFC2046]
	TypeMessageSHttp                         = "message/s-http"                          // [RFC2660][Status change of HTTP experiments to Historic]
	TypeMessageSip                           = "message/sip"                             // [RFC3261]
	TypeMessageSipfrag                       = "message/sipfrag"                         // [RFC3420]
	TypeMessageTrackingStatus                = "message/tracking-status"                 // [RFC3886]
	TypeMessageVndSiSimp                     = "message/vnd.si.simp"                     // [Nicholas_Parks_Young]
	TypeMessageVndWfaWsc                     = "message/vnd.wfa.wsc"                     // [Mick_Conley]
)

const (
	TypeMessageAny = "message/*"
)
