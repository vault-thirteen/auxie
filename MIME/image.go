package mime

// Image.
const (
	TypeImageAces                      = "image/aces"                         // [SMPTE][Howard_Lukk]
	TypeImageApng                      = "image/apng"                         // [W3C][W3C_PNG_Working_Group]
	TypeImageAvci                      = "image/avci"                         // [ISO-IEC_JTC_1][David_Singer]
	TypeImageAvcs                      = "image/avcs"                         // [ISO-IEC_JTC_1][David_Singer]
	TypeImageAvif                      = "image/avif"                         // [Alliance_for_Open_Media][Cyril_Concolato]
	TypeImageBmp                       = "image/bmp"                          // [RFC7903]
	TypeImageCgm                       = "image/cgm"                          // [Alan_Francis]
	TypeImageDicomRle                  = "image/dicom-rle"                    // [DICOM_Standard_Committee][David_Clunie]
	TypeImageDpx                       = "image/dpx"                          // [SMPTE][SMPTE_Director_of_Standards_Development]
	TypeImageEmf                       = "image/emf"                          // [RFC7903]
	TypeImageExample                   = "image/example"                      // [RFC4735]
	TypeImageFits                      = "image/fits"                         // [RFC4047]
	TypeImageG3fax                     = "image/g3fax"                        // [RFC1494]
	TypeImageGif                       = "image/gif"                          // [RFC2045][RFC2046]
	TypeImageHeic                      = "image/heic"                         // [ISO-IEC_JTC_1][David_Singer]
	TypeImageHeicSequence              = "image/heic-sequence"                // [ISO-IEC_JTC_1][David_Singer]
	TypeImageHeif                      = "image/heif"                         // [ISO-IEC_JTC_1][David_Singer]
	TypeImageHeifSequence              = "image/heif-sequence"                // [ISO-IEC_JTC_1][David_Singer]
	TypeImageHej2k                     = "image/hej2k"                        // [ISO-IEC_JTC_1][ITU-T]
	TypeImageHsj2                      = "image/hsj2"                         // [ISO-IEC_JTC_1][ITU-T]
	TypeImageIef                       = "image/ief"                          // [RFC1314]
	TypeImageJ2c                       = "image/j2c"                          // [ISO-IEC_JTC_1_SC_29_WG_1][ISO-IEC_JTC_1][ITU-T]
	TypeImageJls                       = "image/jls"                          // [DICOM_Standard_Committee][David_Clunie]
	TypeImageJp2                       = "image/jp2"                          // [RFC3745]
	TypeImageJpeg                      = "image/jpeg"                         // [RFC2045][RFC2046]
	TypeImageJph                       = "image/jph"                          // [ISO-IEC_JTC_1][ITU-T]
	TypeImageJphc                      = "image/jphc"                         // [ISO-IEC_JTC_1][ITU-T]
	TypeImageJpm                       = "image/jpm"                          // [RFC3745]
	TypeImageJpx                       = "image/jpx"                          // [RFC3745][ISO-IEC_JTC_1_SC_29_WG_1]
	TypeImageJxl                       = "image/jxl"                          // [ISO-IEC_JTC_1_SC_29_WG_1][ISO-IEC_JTC_1]
	TypeImageJxr                       = "image/jxr"                          // [ISO-IEC_JTC_1][ITU-T]
	TypeImageJxrA                      = "image/jxrA"                         // [ISO-IEC_JTC_1][ITU-T]
	TypeImageJxrS                      = "image/jxrS"                         // [ISO-IEC_JTC_1][ITU-T]
	TypeImageJxs                       = "image/jxs"                          // [ISO-IEC_JTC_1]
	TypeImageJxsc                      = "image/jxsc"                         // [ISO-IEC_JTC_1]
	TypeImageJxsi                      = "image/jxsi"                         // [ISO-IEC_JTC_1]
	TypeImageJxss                      = "image/jxss"                         // [ISO-IEC_JTC_1]
	TypeImageKtx                       = "image/ktx"                          // [Khronos][Mark_Callow]
	TypeImageKtx2                      = "image/ktx2"                         // [Khronos][Mark_Callow]
	TypeImageNaplps                    = "image/naplps"                       // [Ilya_Ferber]
	TypeImagePng                       = "image/png"                          // [W3C][PNG_Working_Group]
	TypeImagePrsBtif                   = "image/prs.btif"                     // [Ben_Simon]
	TypeImagePrsPti                    = "image/prs.pti"                      // [Juern_Laun]
	TypeImagePwgRaster                 = "image/pwg-raster"                   // [Michael_Sweet]
	TypeImageSvgXml                    = "image/svg+xml"                      // [W3C][http://www.w3.org/TR/SVG/mimereg.html]
	TypeImageT38                       = "image/t38"                          // [RFC3362]
	TypeImageTiff                      = "image/tiff"                         // [RFC3302]
	TypeImageTiffFx                    = "image/tiff-fx"                      // [RFC3950]
	TypeImageVndAdobePhotoshop         = "image/vnd.adobe.photoshop"          // [Kim_Scarborough]
	TypeImageVndAirzipAcceleratorAzv   = "image/vnd.airzip.accelerator.azv"   // [Gary_Clueit]
	TypeImageVndCnsInf2                = "image/vnd.cns.inf2"                 // [Ann_McLaughlin]
	TypeImageVndDeceGraphic            = "image/vnd.dece.graphic"             // [Michael_A_Dolan]
	TypeImageVndDjvu                   = "image/vnd.djvu"                     // [Leon_Bottou]
	TypeImageVndDwg                    = "image/vnd.dwg"                      // [Jodi_Moline]
	TypeImageVndDxf                    = "image/vnd.dxf"                      // [Jodi_Moline]
	TypeImageVndDvbSubtitle            = "image/vnd.dvb.subtitle"             // [Peter_Siebert][Michael_Lagally]
	TypeImageVndFastbidsheet           = "image/vnd.fastbidsheet"             // [Scott_Becker]
	TypeImageVndFpx                    = "image/vnd.fpx"                      // [Marc_Douglas_Spencer]
	TypeImageVndFst                    = "image/vnd.fst"                      // [Arild_Fuldseth]
	TypeImageVndFujixeroxEdmicsMmr     = "image/vnd.fujixerox.edmics-mmr"     // [Masanori_Onda]
	TypeImageVndFujixeroxEdmicsRlc     = "image/vnd.fujixerox.edmics-rlc"     // [Masanori_Onda]
	TypeImageVndGlobalgraphicsPgb      = "image/vnd.globalgraphics.pgb"       // [Martin_Bailey]
	TypeImageVndMicrosoftIcon          = "image/vnd.microsoft.icon"           // [Simon_Butcher]
	TypeImageVndMix                    = "image/vnd.mix"                      // [Saveen_Reddy]
	TypeImageVndMsModi                 = "image/vnd.ms-modi"                  // [Gregory_Vaughan]
	TypeImageVndMozillaApng            = "image/vnd.mozilla.apng"             // [Stuart_Parmenter]
	TypeImageVndNetFpx                 = "image/vnd.net-fpx"                  // [Marc_Douglas_Spencer]
	TypeImageVndPcoB16                 = "image/vnd.pco.b16"                  // [PCO_AG][Jan_Zeman]
	TypeImageVndRadiance               = "image/vnd.radiance"                 // [Randolph_Fritz][Greg_Ward]
	TypeImageVndSealedPng              = "image/vnd.sealed.png"               // [David_Petersen]
	TypeImageVndSealedmediaSoftsealGif = "image/vnd.sealedmedia.softseal.gif" // [David_Petersen]
	TypeImageVndSealedmediaSoftsealJpg = "image/vnd.sealedmedia.softseal.jpg" // [David_Petersen]
	TypeImageVndSvf                    = "image/vnd.svf"                      // [Jodi_Moline]
	TypeImageVndTencentTap             = "image/vnd.tencent.tap"              // [Ni_Hui]
	TypeImageVndValveSourceTexture     = "image/vnd.valve.source.texture"     // [Henrik_Andersson]
	TypeImageVndWapWbmp                = "image/vnd.wap.wbmp"                 // [Peter_Stark]
	TypeImageVndXiff                   = "image/vnd.xiff"                     // [Steven_Martin]
	TypeImageVndZbrushPcx              = "image/vnd.zbrush.pcx"               // [Chris_Charabaruk]
	TypeImageWebp                      = "image/webp"                         // [RFC-zern-webp-15]
	TypeImageWmf                       = "image/wmf"                          // [RFC7903]
	TypeImageXEmf                      = "image/emf"                          // [RFC7903] DEPRECATED In Favor Of Image/Emf
	TypeImageXWmf                      = "image/wmf"                          // [RFC7903] DEPRECATED In Favor Of Image/Wmf
)

const (
	TypeImageAny = "image/*"
)
