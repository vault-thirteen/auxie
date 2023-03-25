package reader

import (
	"fmt"
	"io"

	"github.com/vault-thirteen/auxie/BOM"
)

const (
	ErrEncodingIsNotCertain = "encoding is not certain: %v"
)

// Reader is a reader capable of guessing the encoding of a stream.
// The reader is single-threaded, i.e. it does not support concurrency.
type Reader struct {
	// Internal reader.
	r io.Reader

	// shouldSkipBOM flag changes behaviour of the reader. When it is set to
	// true, reader does not include the BOM into the response. When set to
	// false, the reader includes the BOM into the response.
	//
	// Separate "blessings" come to the "inventors" of BOMs for UTF-16 (LE) and
	// UTF-32 (LE). The former has BOM consisting of {FF FE} and the latter has
	// BOM equal to {FF FE 00 00}. All this means that a {FF FE 00 00} sequence
	// of bytes can have two meanings: it is either a BOM of UTF-32 (LE)
	// encoding or it is a BOM of UTF-16 (LE) encoding with two NUL bytes
	// following the BOM. When there is more than one possible encoding,
	// we can not guarantee correct detection of encoding and thus can not skip
	// the BOM without damaging the data stream. When we see such a "mystery"
	// in the data stream, we return an error.
	shouldSkipBOM bool

	// A dynamic array for storing first bytes taken from the data stream to
	// detect its encoding.
	firstBytes []byte

	// List of possible encodings used in the data stream. Some encodings have
	// similar BOMs, that is why it is possible to detect more than a single
	// encoding.
	encodings []bom.Encoding
}

func NewReader(r io.Reader, shouldSkipBOM bool) (bomReader *Reader, err error) {
	bomReader = &Reader{
		r:             r,
		shouldSkipBOM: shouldSkipBOM,
		firstBytes:    nil,
		encodings:     make([]bom.Encoding, 0),
	}

	err = bomReader.detectEncoding()
	if err != nil {
		return nil, err
	}

	if bomReader.shouldSkipBOM {
		if len(bomReader.encodings) > 1 {
			return nil, fmt.Errorf(ErrEncodingIsNotCertain, bomReader.encodings)
		}
		if len(bomReader.encodings) == 1 {
			b := bom.BOMs()[bomReader.encodings[0]] // BOM.
			bomReader.firstBytes = bomReader.firstBytes[len(b):]
		}
	}

	return bomReader, nil
}

// detectEncoding tries to detect an encoding of the data stream.
// Bytes read by the detector are stored in a buffer for future usage.
func (br *Reader) detectEncoding() (err error) {
	br.encodings, br.firstBytes, err = bom.SearchForBOM(br.r)
	if err != nil {
		return err
	}

	return nil
}

// GetEncodings returns the list of encodings detected in the data stream.
func (br *Reader) GetEncodings() (encodings []bom.Encoding) {
	return br.encodings
}

// Read reads data from the stream.
// It is a standard method for the 'io.Reader' interface.
func (br *Reader) Read(dst []byte) (n int, err error) {
	if len(br.firstBytes) == 0 {
		return br.r.Read(dst)
	}

	if len(dst) <= len(br.firstBytes) {
		copy(dst, br.firstBytes)
		br.firstBytes = br.firstBytes[len(dst):]
		return len(dst), nil
	}

	// dst is longer than br.firstBytes.
	// Concatenate two pieces.
	copy(dst, br.firstBytes)
	n, err = br.r.Read(dst[len(br.firstBytes):])
	fbLen := len(br.firstBytes) // Store the length before trimming.
	br.firstBytes = br.firstBytes[len(br.firstBytes):]
	if err != nil {
		return fbLen + n, err
	}
	return fbLen + n, nil
}
