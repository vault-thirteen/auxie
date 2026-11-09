package htmldom

// HTML Tags.
// Source: https://html.spec.whatwg.org/multipage/indices.html#elements-3
const (
	TagA               = "a"
	TagAbbr            = "abbr"
	TagAddress         = "address"
	TagArea            = "area"
	TagArticle         = "article"
	TagAside           = "aside"
	TagAudio           = "audio"
	TagB               = "b"
	TagBase            = "base"
	TagBdi             = "bdi"
	TagBdo             = "bdo"
	TagBlockquote      = "blockquote"
	TagBody            = "body"
	TagBr              = "br"
	TagButton          = "button"
	TagCanvas          = "canvas"
	TagCaption         = "caption"
	TagCite            = "cite"
	TagCode            = "code"
	TagCol             = "col"
	TagColgroup        = "colgroup"
	TagData            = "data"
	TagDatalist        = "datalist"
	TagDd              = "dd"
	TagDel             = "del"
	TagDetails         = "details"
	TagDfn             = "dfn"
	TagDialog          = "dialog"
	TagDiv             = "div"
	TagDl              = "dl"
	TagDt              = "dt"
	TagEm              = "em"
	TagEmbed           = "embed"
	TagFieldset        = "fieldset"
	TagFigcaption      = "figcaption"
	TagFigure          = "figure"
	TagFooter          = "footer"
	TagForm            = "form"
	TagH1              = "h1"
	TagH2              = "h2"
	TagH3              = "h3"
	TagH4              = "h4"
	TagH5              = "h5"
	TagH6              = "h6"
	TagHead            = "head"
	TagHeader          = "header"
	TagHgroup          = "hgroup"
	TagHr              = "hr"
	TagHtml            = "html"
	TagI               = "i"
	TagIframe          = "iframe"
	TagImg             = "img"
	TagInput           = "input"
	TagIns             = "ins"
	TagKbd             = "kbd"
	TagLabel           = "label"
	TagLegend          = "legend"
	TagLi              = "li"
	TagLink            = "link"
	TagMain            = "main"
	TagMap             = "map"
	TagMark            = "mark"
	TagMath            = "math"
	TagMenu            = "menu"
	TagMeta            = "meta"
	TagMeter           = "meter"
	TagNav             = "nav"
	TagNoscript        = "noscript"
	TagObject          = "object"
	TagOl              = "ol"
	TagOptgroup        = "optgroup"
	TagOption          = "option"
	TagOutput          = "output"
	TagP               = "p"
	TagPicture         = "picture"
	TagPre             = "pre"
	TagProgress        = "progress"
	TagQ               = "q"
	TagRp              = "rp"
	TagRt              = "rt"
	TagRuby            = "ruby"
	TagS               = "s"
	TagSamp            = "samp"
	TagScript          = "script"
	TagSearch          = "search"
	TagSection         = "section"
	TagSelect          = "select"
	TagSelectedcontent = "selectedcontent"
	TagSlot            = "slot"
	TagSmall           = "small"
	TagSource          = "source"
	TagSpan            = "span"
	TagStrong          = "strong"
	TagStyle           = "style"
	TagSub             = "sub"
	TagSummary         = "summary"
	TagSup             = "sup"
	TagSvg             = "svg"
	TagTable           = "table"
	TagTbody           = "tbody"
	TagTd              = "td"
	TagTemplate        = "template"
	TagTextarea        = "textarea"
	TagTfoot           = "tfoot"
	TagTh              = "th"
	TagThead           = "thead"
	TagTime            = "time"
	TagTitle           = "title"
	TagTr              = "tr"
	TagTrack           = "track"
	TagU               = "u"
	TagUl              = "ul"
	TagVar             = "var"
	TagVideo           = "video"
	TagWbr             = "wbr"
)

// Non-standard HTML Tags with Limited Availability.
const (
	TagGeolocation = "geolocation" // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/geolocation
)

// Old & Deprecated HTML Tags.
const (
	TagAcronym     = "acronym"     // Not supported in HTML5.
	TagApplet      = "applet"      // Not supported in HTML5.
	TagBasefont    = "basefont"    // Not supported in HTML5.
	TagBig         = "big"         // Not supported in HTML5.
	TagCenter      = "center"      // Not supported in HTML5.
	TagDir         = "dir"         // Not supported in HTML5.
	TagFencedframe = "fencedframe" // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/fencedframe
	TagFont        = "font"        // Not supported in HTML5.
	TagFrame       = "frame"       // Not supported in HTML5.
	TagFrameset    = "frameset"    // Not supported in HTML5.
	TagMarquee     = "marquee"     // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/marquee
	TagNobr        = "nobr"        // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/nobr
	TagNoembed     = "noembed"     // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/noembed
	TagNoframes    = "noframes"    // Not supported in HTML5.
	TagParam       = "param"       // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/param
	TagPlaintext   = "plaintext"   // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/plaintext
	TagRb          = "rb"          // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/rb
	TagRtc         = "rtc"         // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/rtc
	TagStrike      = "strike"      // Not supported in HTML5.
	TagTt          = "tt"          // Not supported in HTML5.
	TagXmp         = "xmp"         // https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/xmp
)

// Source:
// https://www.w3schools.com/TAGS/default.asp
