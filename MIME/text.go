package mime

// Text.
const (
	TypeText1dInterleavedParityfec        = "text/1d-interleaved-parityfec"         // [RFC6015]
	TypeTextCacheManifest                 = "text/cache-manifest"                   // [W3C][Robin_Berjon]
	TypeTextCalendar                      = "text/calendar"                         // [RFC5545]
	TypeTextCql                           = "text/cql"                              // [HL7][Bryn_Rhodes]
	TypeTextCqlExpression                 = "text/cql-expression"                   // [HL7][Bryn_Rhodes]
	TypeTextCqlIdentifier                 = "text/cql-identifier"                   // [HL7][Bryn_Rhodes]
	TypeTextCss                           = "text/css"                              // [RFC2318]
	TypeTextCsv                           = "text/csv"                              // [RFC4180][RFC7111]
	TypeTextCsvSchema                     = "text/csv-schema"                       // [National_Archives_UK][David_Underdown]
	TypeTextDirectory                     = "text/directory"                        // [RFC2425][RFC6350] N.B.: DEPRECATED by RFC6350.
	TypeTextDns                           = "text/dns"                              // [RFC4027]
	TypeTextEcmascript                    = "text/ecmascript"                       // [RFC9239] N.B.: OBSOLETED in favour of 'text/javascript'.
	TypeTextEncaprtp                      = "text/encaprtp"                         // [RFC6849]
	TypeTextEnriched                      = "text/enriched"                         // [RFC1896]
	TypeTextExample                       = "text/example"                          // [RFC4735]
	TypeTextFhirpath                      = "text/fhirpath"                         // [HL7][Bryn_Rhodes]
	TypeTextFlexfec                       = "text/flexfec"                          // [RFC8627]
	TypeTextFwdred                        = "text/fwdred"                           // [RFC6354]
	TypeTextGff3                          = "text/gff3"                             // [Sequence_Ontology]
	TypeTextGrammarRefList                = "text/grammar-ref-list"                 // [RFC6787]
	TypeTextHl7v2                         = "text/hl7v2"                            // [HL7][Marc_Duteau]
	TypeTextHtml                          = "text/html"                             // [W3C][Robin_Berjon]
	TypeTextJavascript                    = "text/javascript"                       // [RFC9239]
	TypeTextJcrCnd                        = "text/jcr-cnd"                          // [Peeter_Piegaze]
	TypeTextMarkdown                      = "text/markdown"                         // [RFC7763]
	TypeTextMizar                         = "text/mizar"                            // [Jesse_Alama]
	TypeTextN3                            = "text/n3"                               // [W3C][Eric_Prudhommeaux]
	TypeTextParameters                    = "text/parameters"                       // [RFC7826]
	TypeTextParityfec                     = "text/parityfec"                        // [RFC3009]
	TypeTextPlain                         = "text/plain"                            // [RFC2046][RFC3676][RFC5147]
	TypeTextProvenanceNotation            = "text/provenance-notation"              // [W3C][Ivan_Herman]
	TypeTextPrsFallensteinRst             = "text/prs.fallenstein.rst"              // [Benja_Fallenstein]
	TypeTextPrsLinesTag                   = "text/prs.lines.tag"                    // [John_Lines]
	TypeTextPrsPropLogic                  = "text/prs.prop.logic"                   // [Hans-Dieter_A._Hiep]
	TypeTextPrsTexi                       = "text/prs.texi"                         // [Matin_Bavardi]
	TypeTextRaptorfec                     = "text/raptorfec"                        // [RFC6682]
	TypeTextRED                           = "text/RED"                              // [RFC4102]
	TypeTextRfc822Headers                 = "text/rfc822-headers"                   // [RFC6522]
	TypeTextRichtext                      = "text/richtext"                         // [RFC2045][RFC2046]
	TypeTextRtf                           = "text/rtf"                              // [Paul_Lindner]
	TypeTextRtpEncAescm128                = "text/rtp-enc-aescm128"                 // [_3GPP]
	TypeTextRtploopback                   = "text/rtploopback"                      // [RFC6849]
	TypeTextRtx                           = "text/rtx"                              // [RFC4588]
	TypeTextSGML                          = "text/SGML"                             // [RFC1874]
	TypeTextShaclc                        = "text/shaclc"                           // [W3C_SHACL_Community_Group][Vladimir_Alexiev]
	TypeTextShex                          = "text/shex"                             // [W3C][Eric_Prudhommeaux]
	TypeTextSpdx                          = "text/spdx"                             // [Linux_Foundation][Rose_Judge]
	TypeTextStrings                       = "text/strings"                          // [IEEE-ISTO-PWG-PPP]
	TypeTextT140                          = "text/t140"                             // [RFC4103]
	TypeTextTabSeparatedValues            = "text/tab-separated-values"             // [Paul_Lindner]
	TypeTextTroff                         = "text/troff"                            // [RFC4263]
	TypeTextTurtle                        = "text/turtle"                           // [W3C][Eric_Prudhommeaux]
	TypeTextUlpfec                        = "text/ulpfec"                           // [RFC5109]
	TypeTextUriList                       = "text/uri-list"                         // [RFC2483]
	TypeTextVcard                         = "text/vcard"                            // [RFC6350]
	TypeTextVndA                          = "text/vnd.a"                            // [Regis_Dehoux]
	TypeTextVndAbc                        = "text/vnd.abc"                          // [Steve_Allen]
	TypeTextVndAsciiArt                   = "text/vnd.ascii-art"                    // [Kim_Scarborough]
	TypeTextVndCurl                       = "text/vnd.curl"                         // [Robert_Byrnes]
	TypeTextVndDebianCopyright            = "text/vnd.debian.copyright"             // [Charles_Plessy]
	TypeTextVndDMClientScript             = "text/vnd.DMClientScript"               // [Dan_Bradley]
	TypeTextVndDvbSubtitle                = "text/vnd.dvb.subtitle"                 // [Peter_Siebert][Michael_Lagally]
	TypeTextVndEsmertecThemeDescriptor    = "text/vnd.esmertec.theme-descriptor"    // [Stefan_Eilemann]
	TypeTextVndExchangeable               = "text/vnd.exchangeable"                 // [Martin_Cizek]
	TypeTextVndFamilysearchGedcom         = "text/vnd.familysearch.gedcom"          // [Gordon_Clarke]
	TypeTextVndFiclabFlt                  = "text/vnd.ficlab.flt"                   // [Steve_Gilberd]
	TypeTextVndFly                        = "text/vnd.fly"                          // [John-Mark_Gurney]
	TypeTextVndFmiFlexstor                = "text/vnd.fmi.flexstor"                 // [Kari_E._Hurtta]
	TypeTextVndGml                        = "text/vnd.gml"                          // [Mi_Tar]
	TypeTextVndGraphviz                   = "text/vnd.graphviz"                     // [John_Ellson]
	TypeTextVndHans                       = "text/vnd.hans"                         // [Hill_Hanxv]
	TypeTextVndHgl                        = "text/vnd.hgl"                          // [Heungsub_Lee]
	TypeTextVndIn3d3dml                   = "text/vnd.in3d.3dml"                    // [Michael_Powers]
	TypeTextVndIn3dSpot                   = "text/vnd.in3d.spot"                    // [Michael_Powers]
	TypeTextVndIPTCNewsML                 = "text/vnd.IPTC.NewsML"                  // [IPTC]
	TypeTextVndIPTCNITF                   = "text/vnd.IPTC.NITF"                    // [IPTC]
	TypeTextVndLatexZ                     = "text/vnd.latex-z"                      // [Mikusiak_Lubos]
	TypeTextVndMotorolaReflex             = "text/vnd.motorola.reflex"              // [Mark_Patton]
	TypeTextVndMsMediapackage             = "text/vnd.ms-mediapackage"              // [Jan_Nelson]
	TypeTextVndNet2phoneCommcenterCommand = "text/vnd.net2phone.commcenter.command" // [Feiyu_Xie]
	TypeTextVndRadisysMsmlBasicLayout     = "text/vnd.radisys.msml-basic-layout"    // [RFC5707]
	TypeTextVndSenxWarpscript             = "text/vnd.senx.warpscript"              // [Pierre_Papin]
	TypeTextVndSiUricatalogue             = "text/vnd.si.uricatalogue"              // [Nicholas_Parks_Young] N.B.: OBSOLETED by request.
	TypeTextVndSunJ2meAppDescriptor       = "text/vnd.sun.j2me.app-descriptor"      // [Gary_Adams]
	TypeTextVndSosi                       = "text/vnd.sosi"                         // [Petter_Reinholdtsen]
	TypeTextVndTrolltechLinguist          = "text/vnd.trolltech.linguist"           // [David_Lee_Lambert]
	TypeTextVndWapSi                      = "text/vnd.wap.si"                       // [WAP-Forum]
	TypeTextVndWapSl                      = "text/vnd.wap.sl"                       // [WAP-Forum]
	TypeTextVndWapWml                     = "text/vnd.wap.wml"                      // [Peter_Stark]
	TypeTextVndWapWmlscript               = "text/vnd.wap.wmlscript"                // [Peter_Stark]
	TypeTextVtt                           = "text/vtt"                              // [W3C][Silvia_Pfeiffer]
	TypeTextWgsl                          = "text/wgsl"                             // [W3C][David_Neto]
	TypeTextXml                           = "text/xml"                              // [RFC7303]
	TypeTextXmlExternalParsedEntity       = "text/xml-external-parsed-entity"       // [RFC7303]
)

const (
	TypeTextAny = "text/*"
)