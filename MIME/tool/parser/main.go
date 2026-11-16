package main

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	cla "github.com/vault-thirteen/auxie/MIME/tool/parser/CommandLineArgs"
	m "github.com/vault-thirteen/auxie/MIME/tool/parser/model"
	"github.com/vault-thirteen/auxie/cio"
	ae "github.com/vault-thirteen/auxie/errors"
)

const (
	Type_Application = "application"
	Type_Audio       = "audio"
	Type_Font        = "font"
	Type_Haptics     = "haptics"
	Type_Image       = "image"
	Type_Message     = "message"
	Type_Model       = "model"
	Type_Multipart   = "multipart"
	Type_Text        = "text"
	Type_Video       = "video"

	CR               = "\r"
	LF               = "\n"
	CRLF             = CR + LF
	OutputLineFormat = `%s = "%s" // %s`
	FilePermissions  = 0777
)

const (
	Err_ColumnNamesMismatch = "column names do not match"
)

var (
	columns = []string{"name", "template", "reference"}

	supportedTypes = []string{
		Type_Application,
		Type_Audio,
		Type_Font,
		Type_Haptics,
		Type_Image,
		Type_Message,
		Type_Model,
		Type_Multipart,
		Type_Text,
		Type_Video,
	}
)

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	args, err := cla.NewCommandLineArgs()
	if err != nil {
		ae.LogError(err)
		cio.PrintMessage(cla.UsageHint)
		os.Exit(1)
		return
	}

	err = work(args.DataFolderPath)
	mustBeNoError(err)
}
func work(dataFolderPath string) (err error) {
	var buf []byte
	var srcFilePath string
	var dstFilePath string
	var types []m.MIME
	for _, t := range supportedTypes {
		srcFilePath = filepath.Join(dataFolderPath, t+".csv")
		dstFilePath = filepath.Join(dataFolderPath, t+".go")

		types, err = collectFileTypes(srcFilePath, t)
		if err != nil {
			return err
		}

		buf, err = composeFile(types, t)
		if err != nil {
			return err
		}

		err = os.WriteFile(dstFilePath, buf, FilePermissions)
		if err != nil {
			return err
		}
	}

	return nil
}
func collectFileTypes(filePath string, mimeType string) (types []m.MIME, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		derr := file.Close()
		err = ae.Combine(err, derr)
	}()

	reader := csv.NewReader(file)

	// Read the first (header) row and check it.
	var row []string
	row, err = reader.Read()
	if err != nil {
		return nil, err
	}
	if !compareRowValues(row, columns) {
		return nil, ae.ErrorFromText(Err_ColumnNamesMismatch)
	}

	// Read data lines.
	types = make([]m.MIME, 0)
	var t m.MIME
	for {
		row, err = reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		t, err = collectRowData(row)
		if err != nil {
			return nil, err
		}

		types = append(types, t)
	}

	return types, nil
}
func collectRowData(row m.Row) (t m.MIME, err error) {
	t = m.MIME{
		Name:      row[0],
		Template:  row[1],
		Reference: row[2],
	}

	return t, nil
}
func compareRowValues(r1, r2 m.Row) (ok bool) {
	if len(r1) != len(r2) {
		return false
	}

	var v1, v2 string
	for i := 0; i < len(r1); i++ {
		v1 = strings.TrimSpace(strings.ToLower(r1[i]))
		v2 = strings.TrimSpace(strings.ToLower(r2[i]))
		if v1 != v2 {
			return false
		}
	}

	return true
}
func composeFile(types []m.MIME, t string) (data []byte, err error) {
	var sb = new(strings.Builder)
	var typeName = m.CapitalizeFirstLetter(t)
	SB_AddLine(sb, `package mime`)
	SB_AddLine(sb, "")

	var line string
	line = `// ` + typeName + `.`
	SB_AddLine(sb, line)
	SB_AddLine(sb, `const (`)

	var varName string
	for _, typ := range types {
		varName = "Type" + typeName + typ.NameForFile()
		line = cio.FormatMessage(OutputLineFormat, varName, typ.Template, typ.Reference)
		SB_AddLine(sb, line)
	}
	SB_AddLine(sb, `)`)
	SB_AddLine(sb, "")

	line = `const (` + CRLF +
		`Type` + typeName + `Any = "` + t + `/*"` + CRLF +
		`)`
	SB_AddLine(sb, line)

	return []byte(sb.String()), nil
}
func SB_AddLine(sb *strings.Builder, text string) {
	sb.WriteString(text + CRLF)
}
