package mime

// Multipart.
const (
	TypeMultipartAlternative    = "multipart/alternative"       // [RFC2046][RFC2045]
	TypeMultipartAppledouble    = "multipart/appledouble"       // [Patrik_Faltstrom]
	TypeMultipartByteranges     = "multipart/byteranges"        // [RFC9110]
	TypeMultipartDigest         = "multipart/digest"            // [RFC2046][RFC2045]
	TypeMultipartEncrypted      = "multipart/encrypted"         // [RFC1847]
	TypeMultipartExample        = "multipart/example"           // [RFC4735]
	TypeMultipartFormData       = "multipart/form-data"         // [RFC7578]
	TypeMultipartHeaderSet      = "multipart/header-set"        // [Dave_Crocker]
	TypeMultipartMixed          = "multipart/mixed"             // [RFC2046][RFC2045]
	TypeMultipartMultilingual   = "multipart/multilingual"      // [RFC8255]
	TypeMultipartParallel       = "multipart/parallel"          // [RFC2046][RFC2045]
	TypeMultipartRelated        = "multipart/related"           // [RFC2387]
	TypeMultipartReport         = "multipart/report"            // [RFC6522]
	TypeMultipartSigned         = "multipart/signed"            // [RFC1847]
	TypeMultipartVndBintMedPlus = "multipart/vnd.bint.med-plus" // [Heinz-Peter_Schütz]
	TypeMultipartVoiceMessage   = "multipart/voice-message"     // [RFC3801]
	TypeMultipartXMixedReplace  = "multipart/x-mixed-replace"   // [W3C][Robin_Berjon]
)

const (
	TypeMultipartAny = "multipart/*"
)
