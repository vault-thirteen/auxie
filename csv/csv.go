package csv

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// Unfortunately standard CSV library in Go language is useless.
// It does not add quotes to strings when it writes them into a CSV file !

// Errors.
const (
	ErrFDataTypeIsUnsupported = "data type is unsupported: %s"
	ErrNoDataInRow            = "no data in row"
	ErrFRowSizeMismatch       = "row size mismatch: %v vs %v"
)

// Settings.
const (
	CommaStr   = `,`
	NewLineStr = "\r\n" // RFC 4180.
)

// Writer is a simple CSV writer.
type Writer struct {
	iow          io.Writer
	columnsCount int
	newLineBA    []byte
}

// NewWriter is a constructor of a CSV writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		iow:          w,
		columnsCount: -1,
		newLineBA:    []byte(NewLineStr),
	}
}

// WriteRow writes a row of cells.
func (w *Writer) WriteRow(row []any) (err error) {
	// 1. Check columns count.
	{
		// No data ?
		if len(row) == 0 {
			return errors.New(ErrNoDataInRow)
		}

		// Set the ruler when the first row arrives.
		if w.columnsCount == -1 {
			w.columnsCount = len(row)
		}

		// Short or long row ?
		if len(row) != w.columnsCount {
			return fmt.Errorf(ErrFRowSizeMismatch, w.columnsCount, len(row))
		}
	}

	var (
		buf     = make([]string, 0, len(row))
		cell    string
		intVar  int
		uintVar uint
		f32Var  float32
		f64Var  float64
		ba      []byte
	)

	// 2. Prepare text for writing.
	{
		for _, v := range row {
			switch v.(type) {
			case string:
				cell = v.(string)
				cell = addEdgeQuotes(escapeString(cell))

			case int:
				intVar = v.(int)
				cell = strconv.Itoa(intVar)

			case uint:
				uintVar = v.(uint)
				cell = strconv.FormatUint(uint64(uintVar), 10)

			case float32:
				f32Var = v.(float32)
				cell = strconv.FormatFloat(float64(f32Var), 'f', -1, 64)

			case float64:
				f64Var = v.(float64)
				cell = strconv.FormatFloat(float64(f64Var), 'f', -1, 64)

			case []byte:
				// Raw bytes are printed as a string in hexadecimal format.
				ba = v.([]byte)
				cell = addEdgeQuotes(strings.ToUpper(hex.EncodeToString(ba)))

			default:
				return fmt.Errorf(ErrFDataTypeIsUnsupported, reflect.TypeOf(v))
			}

			buf = append(buf, cell)
		}
	}

	// 3. Write the text.
	{
		iLast := len(row) - 1

		for i, s := range buf {
			if i != iLast {
				_, err = w.iow.Write([]byte(s + CommaStr))
			} else {
				_, err = w.iow.Write([]byte(s))
			}
			if err != nil {
				return err
			}
		}

		_, err = w.iow.Write(w.newLineBA)
		if err != nil {
			return err
		}
	}

	return nil
}

// escapeString screens double quote symbols in a string.
func escapeString(s string) string {
	return strings.ReplaceAll(s, `"`, `""`)
}

// addEdgeQuotes adds a double quote symbol at the string's edges.
func addEdgeQuotes(s string) string {
	return `"` + s + `"`
}
