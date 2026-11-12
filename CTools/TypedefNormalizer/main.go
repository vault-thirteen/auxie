package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"time"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

const (
	Introduction = `Type C/C++ definitions, then press Ctrl+Z and ENTER.
Send capital letter Q or small letter q to quit.`
	GoTypeDefinition = "type %s = %s"
)

func main() {
	cio.PrintMessage(Introduction)
	var mustQuit = false
	var data []byte
	var err error
	var buf []byte
	for !mustQuit {
		time.Sleep(time.Second * 1)

		data, err = getInput()
		if err != nil {
			errors.LogError(err)
			continue
		}
		if data == nil {
			continue
		}
		if strings.ToLower(strings.TrimSpace(string(data))) == "q" {
			mustQuit = true
			continue
		}
		if len(data) == 0 {
			continue
		}

		buf, err = convertCTypedefsToGo(data)
		if err != nil {
			errors.LogError(err)
			continue
		}

		cio.PrintMessage(string(buf))
	}
}

// Example:
// typedef LPCH LPTCH, PTCH;
// typedef char TCHAR, *PTCHAR;

func convertCTypedefsToGo(in []byte) (out []byte, err error) {
	var lines []string
	lines, err = getTextLines(in)
	if err != nil {
		return nil, err
	}

	lines = removeEmptyLines(lines)
	lines = removeDoubleSpaces(lines)
	if !checkLineEnds(lines) {
		return nil, errors.ErrorFromText("no ';' at the end")
	}

	lines = cutLineEnds(lines, ";")
	lines, err = processLines(lines)
	if err != nil {
		return nil, err
	}

	return []byte(strings.Join(lines, "\r\n")), nil
}
func getTextLines(data []byte) (lines []string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	lines = make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	err = scanner.Err()
	if err != nil {
		return
	}

	return lines, nil
}
func removeEmptyLines(in []string) (out []string) {
	out = make([]string, 0)

	var tmp string
	for _, line := range in {
		tmp = strings.TrimSpace(line)
		if len(tmp) == 0 {
			continue
		}

		out = append(out, tmp)
	}

	return out
}
func removeDoubleSpaces(in []string) (out []string) {
	out = make([]string, 0)

	var tmp string
	for _, line := range in {
		tmp = strings.Join(strings.Fields(line), " ")
		if len(tmp) == 0 {
			continue
		}

		out = append(out, tmp)
	}

	return out
}
func checkLineEnds(lines []string) (ok bool) {
	for _, line := range lines {
		if !strings.HasSuffix(line, ";") {
			return false
		}
	}

	return true
}
func cutLineEnds(in []string, suffix string) (out []string) {
	out = make([]string, 0, len(in))

	var newLine string
	for _, line := range in {
		newLine = strings.TrimSuffix(line, suffix)
		out = append(out, newLine)
	}

	return out
}
func processLines(in []string) (out []string, err error) {
	out = make([]string, 0)

	var parts []string
	var oldType string
	var newTypes []string
	var goTypeDefLine string
	for _, line := range in {
		parts = strings.Split(line, " ")
		if len(parts) < 3 {
			return nil, errors.ErrorFromText("not enough parts")
		}
		if parts[0] != "typedef" {
			return nil, errors.ErrorFromText("no type definition")
		}
		oldType = parts[1]
		newTypes = cutLineEnds(parts[2:], ",")

		for _, newType := range newTypes {
			if strings.HasPrefix(newType, "*") {
				goTypeDefLine = cio.FormatMessage(GoTypeDefinition, strings.TrimPrefix(newType, "*"), "*"+oldType)
			} else {
				goTypeDefLine = cio.FormatMessage(GoTypeDefinition, newType, oldType)
			}

			out = append(out, goTypeDefLine)
		}
	}

	return out, nil
}
func getInput() (data []byte, err error) {
	return io.ReadAll(os.Stdin)
}
