package mime

// Multipart.
const (
	TypeMultipartAlternative    = "multipart/alternative"       // [RFC 2046]; [RFC 2045]
	TypeMultipartAppledouble    = "multipart/appledouble"       // [Patrik_Faltstrom]
	TypeMultipartByteranges     = "multipart/byteranges"        // [RFC 9110]
	TypeMultipartDigest         = "multipart/digest"            // [RFC 2046]; [RFC 2045]
	TypeMultipartEncrypted      = "multipart/encrypted"         // [RFC 1847]
	TypeMultipartExample        = "multipart/example"           // [RFC 4735]
	TypeMultipartFormData       = "multipart/form-data"         // [RFC 7578]
	TypeMultipartHeaderSet      = "multipart/header-set"        // [Dave_Crocker]
	TypeMultipartMixed          = "multipart/mixed"             // [RFC 2046]; [RFC 2045]
	TypeMultipartMultilingual   = "multipart/multilingual"      // [RFC 8255]
	TypeMultipartParallel       = "multipart/parallel"          // [RFC 2046]; [RFC 2045]
	TypeMultipartRelated        = "multipart/related"           // [RFC 2387]
	TypeMultipartReport         = "multipart/report"            // [RFC 6522]
	TypeMultipartSigned         = "multipart/signed"            // [RFC 1847]
	TypeMultipartVndBintMedPlus = "multipart/vnd.bint.med-plus" // [Heinz-Peter_Schütz]
	TypeMultipartVoiceMessage   = "multipart/voice-message"     // [RFC 3801]
	TypeMultipartXMixedReplace  = "multipart/x-mixed-replace"   // [W3C]; [Robin_Berjon]
)

const (
	TypeMultipartAny = "multipart/*"
)
