package bom

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/vault-thirteen/auxie/reader"
)

const (
	ErrUnknownEncoding = "unknown encoding: %v"
	ErrBOMIsNotFound   = "byte order mark is not found"
	ErrDuplicateProbe  = "duplicate probe for encoding %v"
)

var (
	bomUTF8       = []byte{0xEF, 0xBB, 0xBF}
	bomUTF16BE    = []byte{0xFE, 0xFF}
	bomUTF16LE    = []byte{0xFF, 0xFE}
	bomUTF32BE    = []byte{0x00, 0x00, 0xFE, 0xFF}
	bomUTF32LE    = []byte{0xFF, 0xFE, 0x00, 0x00}
	bomUTF7       = []byte{0x2B, 0x2F, 0x76}       // +/v
	bomUTF1       = []byte{0xF7, 0x64, 0x4C}       //?dL
	bomUTF_EBCDIC = []byte{0xDD, 0x73, 0x66, 0x73} // ?sfs
	bomSCSU       = []byte{0x0E, 0xFE, 0xFF}
	bomBOCU1      = []byte{0xFB, 0xEE, 0x28} // ??(
	bomGB18030    = []byte{0x84, 0x31, 0x95, 0x33}
)

func BOMUTF8() []byte       { return bomUTF8 }
func BOMUTF16BE() []byte    { return bomUTF16BE }
func BOMUTF16LE() []byte    { return bomUTF16LE }
func BOMUTF32BE() []byte    { return bomUTF32BE }
func BOMUTF32LE() []byte    { return bomUTF32LE }
func BOMUTF7() []byte       { return bomUTF7 }
func BOMUTF1() []byte       { return bomUTF1 }
func BOMUTF_EBCDIC() []byte { return bomUTF_EBCDIC }
func BOMSCSU() []byte       { return bomSCSU }
func BOMBOCU1() []byte      { return bomBOCU1 }
func BOMGB18030() []byte    { return bomGB18030 }

var boms = map[Encoding][]byte{
	EncodingUTF8:       bomUTF8,
	EncodingUTF16BE:    bomUTF16BE,
	EncodingUTF16LE:    bomUTF16LE,
	EncodingUTF32BE:    bomUTF32BE,
	EncodingUTF32LE:    bomUTF32LE,
	EncodingUTF7:       bomUTF7,
	EncodingUTF1:       bomUTF1,
	EncodingUTF_EBCDIC: bomUTF_EBCDIC,
	EncodingSCSU:       bomSCSU,
	EncodingBOCU1:      bomBOCU1,
	EncodingGB18030:    bomGB18030,
}

func BOMs() map[Encoding][]byte {
	return boms
}

// ReadBOMOfEncoding tries to read the BOM of a specified encoding.
// The prefix which was read from the stream is always returned.
func ReadBOMOfEncoding(r io.Reader, enc Encoding) (prefix []byte, err error) {
	bom, ok := boms[enc]
	if !ok {
		return nil, fmt.Errorf(ErrUnknownEncoding, enc)
	}

	prefix = make([]byte, len(bom))
	var n int
	n, err = r.Read(prefix)
	prefix = prefix[:n]
	if err != nil {
		return prefix, err
	}

	if bytes.Equal(prefix, bom) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// SkipBOM tries to skip a BOM prefix of the specified encoding in the stream.
// It reads the BOM, shifting the reader's "cursor".
func SkipBOM(r io.Reader, enc Encoding) (err error) {
	_, err = ReadBOMOfEncoding(r, enc)
	if err != nil {
		return err
	}

	return nil
}

// SearchForBOM searches the stream for BOM.
// n = number of bytes read from the stream.
func SearchForBOM(r io.Reader) (encodings []Encoding, n int, err error) {
	encodings = make([]Encoding, 0)

	// Prepare the lists of encodings.
	accurateProbes := make(map[Encoding]*Probe) // Results.
	encodingsToProbe := make(map[Encoding]bool) // Task-list.
	for _, enc := range possibleEncodings {
		encodingsToProbe[enc] = true
	}

	// Increase the accumulator size by one byte per loop and get accurate
	// probes. Save accurate probes. Stop when all the probes are accurate.
	mbs := getMaximumBOMSize(possibleEncodings)
	var acc = make([]byte, 0, mbs)
	var report *Report
	var b byte
	probeSize := 1
	for ; probeSize <= mbs; probeSize++ {
		// Get a byte.
		b, err = reader.ReadByte(r)
		if err != nil {
			return encodings, probeSize - 1, err
		}
		acc = append(acc, b)

		// Get a report and extract accurate probes.
		report, err = GetEncodingsReport(acc, encodingsToProbe)
		if err != nil {
			return encodings, probeSize, err
		}
		tmp := report.GetAccurateProbes()
		for _, p := range tmp {
			_, isDuplicate := accurateProbes[p.Encoding]
			if isDuplicate {
				return encodings, probeSize, fmt.Errorf(ErrDuplicateProbe, p.Encoding)
			}
			accurateProbes[p.Encoding] = p
			delete(encodingsToProbe, p.Encoding)
		}
		if report.IsAccurate() && len(encodingsToProbe) == 0 {
			break
		}
	}

	// Process the list of accurate probes.
	for _, ap := range accurateProbes {
		if ap.Probability.IsYes() {
			encodings = append(encodings, ap.Encoding)
		}
	}

	return encodings, probeSize, nil
}

// getMaximumBOMSize returns the maximum BOM size of all possible encodings.
func getMaximumBOMSize(encodings []Encoding) (mbs int) {
	mbs = 0

	for _, pe := range encodings {
		if len(boms[pe]) > mbs {
			mbs = len(boms[pe])
		}
	}

	return mbs
}
