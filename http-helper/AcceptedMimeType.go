package httphelper

import (
	"fmt"
	"strings"

	mime "github.com/vault-thirteen/MIME"
	"github.com/vault-thirteen/auxie/number"
)

type AcceptedMimeType struct {
	MimeType string
	Weight   float32
}

func ParseRecord(rec string) (amt *AcceptedMimeType, err error) {
	parts := strings.Split(strings.TrimSpace(rec), ";")

	if len(parts) == 0 {
		return nil, fmt.Errorf(ErrSyntaxErrorInRecord, strings.TrimSpace(rec))
	}

	if len(parts[0]) == 0 {
		return nil, fmt.Errorf(ErrSyntaxErrorInRecord, strings.TrimSpace(rec))
	}

	err = checkMimeType(parts[0])
	if err != nil {
		return nil, err
	}

	amt = &AcceptedMimeType{
		MimeType: parts[0],
	}

	switch len(parts) {
	case 1:
		// Example: `text/html`.
		amt.Weight = 1.0

		return amt, nil

	case 2:
		// Example: `application/xml;q=0.9`.
		amt.Weight, err = ParseRecordWeight(parts[1])
		if err != nil {
			return nil, err
		}

		return amt, nil

	default:
		return nil, fmt.Errorf(ErrSyntaxErrorInRecord, strings.TrimSpace(rec))
	}
}

func ParseRecordWeight(rwt string) (weight float32, err error) {
	// Example: `q=0.9`.
	parts := strings.Split(strings.TrimSpace(rwt), "=")

	if (len(parts) != 2) || (parts[0] != `q`) {
		return 0, fmt.Errorf(ErrSyntaxErrorInWeight, rwt)
	}

	return number.ParseFloat32(parts[1])
}

func checkMimeType(mt string) (err error) {
	_, _, err = mime.GetMimeTypeParts(mt)
	if err != nil {
		return err
	}

	return nil
}
