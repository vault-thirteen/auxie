package main

import (
	"cmp"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	"github.com/vault-thirteen/auxie/cio"
	ae "github.com/vault-thirteen/auxie/errors"
	cla "github.com/vault-thirteen/auxie/header/tool/parser/CommandLineArgs"
	m "github.com/vault-thirteen/auxie/header/tool/parser/model"
)

const (
	Folder_Http    = "HTTP"
	Folder_Message = "Message"

	CsvFile_HttpHeaders                = "field-names.csv"
	CsvFile_MessageHeaders_Permanent   = "perm-headers.csv"
	CsvFile_MessageHeaders_Provisional = "prov-headers.csv"

	Protocol_Http    = "http"
	Protocol_Mail    = "mail"
	Protocol_Mime    = "mime"
	Protocol_NetNews = "netnews"

	Status_None          = ""
	Status_Deprecated    = "deprecated"
	Status_Experimental  = "experimental"
	Status_Informational = "informational"
	Status_Obsoleted     = "obsoleted"
	Status_Permanent     = "permanent"
	Status_Provisional   = "provisional"
	Status_Reserved      = "reserved"
	Status_Standard      = "standard"

	CR                     = "\r"
	LF                     = "\n"
	CRLF                   = CR + LF
	HeaderOutputLineFormat = `%s = "%s" // %s`
	FilePermissions        = 0777
)

const (
	Err_ColumnNamesMismatch = "column names do not match"
	Err_RowTypeIsUnknown    = "unknown row type"

	Errf_ProtocolIsNotFound  = "protocol is not found: %s"
	Errf_StatusIsUnsupported = "status is not supported: %s"
)

var (
	columns_http                = []string{"field name", "status", "structured type", "reference", "comments"}
	columns_message_permanent   = []string{"header field name", "template", "protocol", "status", "trace", "reference"}
	columns_message_provisional = []string{"header field name", "template", "protocol", "status", "trace", "reference"}

	supportedStatuses = []string{
		Status_None,
		Status_Deprecated,
		Status_Experimental,
		Status_Informational,
		Status_Obsoleted,
		Status_Permanent,
		Status_Provisional,
		Status_Reserved,
		Status_Standard,
	}

	supportedProtocols = []string{
		Protocol_Http,
		Protocol_Mail,
		Protocol_Mime,
		Protocol_NetNews,
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
	var allHeaders []m.Header
	allHeaders, err = collectHeaderData(dataFolderPath)
	if err != nil {
		return err
	}

	var pHeadersByStatus map[string][]m.Header
	var buf []byte
	var dstFilePath string
	for _, proto := range supportedProtocols {
		pHeadersByStatus, err = getProtocolHeadersSorted(allHeaders, proto)
		if err != nil {
			return err
		}

		buf, err = composeFile(pHeadersByStatus, proto)
		if err != nil {
			return err
		}

		dstFilePath = filepath.Join(dataFolderPath, proto+".go")
		err = os.WriteFile(dstFilePath, buf, FilePermissions)
		if err != nil {
			return err
		}
	}

	return nil
}
func collectHeaderData(dataFolderPath string) (allHeaders []m.Header, err error) {
	allHeaders = make([]m.Header, 0)

	var headers []m.Header
	headers, err = collectHeadersFromFile(filepath.Join(dataFolderPath, Folder_Http, CsvFile_HttpHeaders), m.RowType_Http)
	if err != nil {
		return nil, err
	}
	allHeaders = append(allHeaders, headers...)

	headers, err = collectHeadersFromFile(filepath.Join(dataFolderPath, Folder_Message, CsvFile_MessageHeaders_Permanent), m.RowType_PermanentMessage)
	if err != nil {
		return nil, err
	}
	allHeaders = append(allHeaders, headers...)

	headers, err = collectHeadersFromFile(filepath.Join(dataFolderPath, Folder_Message, CsvFile_MessageHeaders_Provisional), m.RowType_ProvisionalMessage)
	if err != nil {
		return nil, err
	}
	allHeaders = append(allHeaders, headers...)

	return allHeaders, nil
}
func collectHeadersFromFile(filePath string, rowType m.RowType) (headers []m.Header, err error) {
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

	var ok bool
	ok, err = compareHeaderRow(row, rowType)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ae.ErrorFromText(Err_ColumnNamesMismatch)
	}

	// Read data lines.
	headers = make([]m.Header, 0)
	var header m.Header
	for {
		row, err = reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		header, err = collectRowData(row, rowType)
		if err != nil {
			return nil, err
		}

		headers = append(headers, header)
	}

	return headers, nil
}
func collectRowData(row m.Row, rowType m.RowType) (h m.Header, err error) {
	switch rowType {
	case m.RowType_Http:
		h = m.Header{
			FieldName: row[0],
			Protocol:  Protocol_Http,
			Status:    strings.ToLower(row[1]),
			Reference: row[3],
			Comment:   row[4],
		}

	case m.RowType_PermanentMessage:
		h = m.Header{
			FieldName: row[0],
			Protocol:  row[2],
			Status:    strings.ToLower(row[3]),
			Reference: row[5],
		}

	case m.RowType_ProvisionalMessage:
		h = m.Header{
			FieldName: row[0],
			Protocol:  row[2],
			Status:    strings.ToLower(row[3]),
			Reference: row[5],
		}

	default:
		return h, ae.ErrorFromText(Err_RowTypeIsUnknown)
	}

	return h, nil
}
func compareHeaderRow(row m.Row, rowType m.RowType) (ok bool, err error) {
	switch rowType {
	case m.RowType_Http:
		return compareRowValues(row, columns_http), nil

	case m.RowType_PermanentMessage:
		return compareRowValues(row, columns_message_permanent), nil

	case m.RowType_ProvisionalMessage:
		return compareRowValues(row, columns_message_provisional), nil

	default:
		return false, ae.ErrorFromText(Err_RowTypeIsUnknown)
	}
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
func getHeadersByProtocol(allHeaders []m.Header, protocol string) (out []m.Header) {
	out = make([]m.Header, 0)
	protocolLC := strings.ToLower(protocol)

	for _, header := range allHeaders {
		if strings.ToLower(header.Protocol) == protocolLC {
			out = append(out, header)
		}
	}

	return out
}
func getHeadersByStatus(allHeaders []m.Header, status string) (out []m.Header) {
	out = make([]m.Header, 0)
	statusLC := strings.ToLower(status)

	for _, header := range allHeaders {
		if strings.ToLower(header.Status) == statusLC {
			out = append(out, header)
		}
	}

	return out
}
func getProtocolHeadersSorted(allHeaders []m.Header, protocol string) (pHeadersByStatus map[string][]m.Header, err error) {
	var pHeaders []m.Header
	pHeaders = getHeadersByProtocol(allHeaders, protocol)
	if len(pHeaders) == 0 {
		return nil, ae.ErrorFromFormattedText(Errf_ProtocolIsNotFound, protocol)
	}

	err = searchForUnsupportedStatuses(pHeaders)
	if err != nil {
		return nil, err
	}

	pHeadersByStatus = make(map[string][]m.Header)
	var tmp []m.Header
	for _, ss := range supportedStatuses {
		tmp = getHeadersByStatus(pHeaders, ss)
		slices.SortFunc(tmp, func(a, b m.Header) int { return cmp.Compare(a.FieldName, b.FieldName) })
		pHeadersByStatus[ss] = tmp
	}

	return pHeadersByStatus, nil
}
func searchForUnsupportedStatuses(headers []m.Header) (err error) {
	for _, header := range headers {
		if !slices.Contains(supportedStatuses, header.Status) {
			return ae.ErrorFromFormattedText(Errf_StatusIsUnsupported, header.Status)
		}
	}

	return nil
}
func composeFile(pHeadersByStatus map[string][]m.Header, protocol string) (data []byte, err error) {
	var sb = new(strings.Builder)
	var protocolName = capitalizeFirstLetter(protocol)
	var line string
	line = `// ` + protocolName + ` protocol header field names.`
	SB_AddLine(sb, line)
	SB_AddLine(sb, "")

	var varName, statusName string
	for _, ss := range supportedStatuses {
		statusName = capitalizeFirstLetter(ss)
		if len(statusName) == 0 {
			statusName = `(Empty)`
		}

		SB_AddLine(sb, cio.FormatMessage(`// Status: %s.`, statusName))
		SB_AddLine(sb, `const (`)

		for _, h := range pHeadersByStatus[ss] {
			varName = protocolName + "Header" + h.NameForFile()
			line = cio.FormatMessage(HeaderOutputLineFormat, varName, h.FieldName, h.CommentForFile())
			SB_AddLine(sb, line)
		}

		SB_AddLine(sb, `)`)
		SB_AddLine(sb, "")
	}

	return []byte(sb.String()), nil
}
func capitalizeFirstLetter(in string) (out string) {
	if len(in) == 0 {
		return ""
	}

	runes := []rune(in)
	if len(runes) == 0 {
		return ""
	}

	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
func SB_AddLine(sb *strings.Builder, text string) {
	sb.WriteString(text + CRLF)
}
