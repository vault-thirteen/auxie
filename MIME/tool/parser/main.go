package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	e "github.com/vault-thirteen/auxie/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const OsArgsCount = 2

const (
	OsExitCode_BadArgsCount = 1
)

const (
	ErrTooManyMatches = "too many matches"
)

const (
	NewLine          = "\r\n"
	OutputFileExt    = "go"
	RecordLinePrefix = "\t"
)

// First argument is path to a folder with source files.
func main() {
	if len(os.Args) < 1+OsArgsCount {
		fmt.Println("Usage:  [tool.exe] <Input_Folder> <Output_Folder>")
		os.Exit(OsExitCode_BadArgsCount)
	}

	pathToSrcFolder := os.Args[1]
	outputFolder := os.Args[2]
	fmt.Println("Input folder: ", pathToSrcFolder)
	fmt.Println("Output folder: ", outputFolder)

	filesToProcess := []string{
		`application.csv`,
		`audio.csv`,
		`font.csv`,
		`haptics.csv`,
		`image.csv`,
		`message.csv`,
		`model.csv`,
		`multipart.csv`,
		`text.csv`,
		`video.csv`,
	}

	var inputPath string
	var res *Result
	var err error
	for _, file := range filesToProcess {
		inputPath = filepath.Join(pathToSrcFolder, file)
		res, err = processFile(inputPath, outputFolder)
		mustBeNoError(err)

		fmt.Println(fmt.Sprintf("%d records processed.", res.RowCount))
	}
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func processFile(inputFilePath string, outputFolder string) (res *Result, err error) {
	fileName := filepath.Base(inputFilePath)
	fileNameWOE := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	caser := cases.Title(language.English, cases.NoLower)
	category := caser.String(fileNameWOE)

	fmt.Print(category + ": ")

	var fIn *os.File
	fIn, err = os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		derr := fIn.Close()
		if derr != nil {
			err = e.Combine(err, derr)
		}
	}()

	var outputFilePath = filepath.Join(outputFolder, fileNameWOE+"."+OutputFileExt)

	var fOut *os.File
	fOut, err = os.OpenFile(outputFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	defer func() {
		derr := fOut.Close()
		if derr != nil {
			err = e.Combine(err, derr)
		}
	}()

	rdr := csv.NewReader(fIn)
	var rows [][]string
	rows, err = rdr.ReadAll()
	if err != nil {
		return nil, err
	}

	res = &Result{
		RawData:  rows,
		RowCount: len(rows),
		Records:  make([]Record, 0, len(rows)),
	}

	var re *regexp.Regexp
	re, err = regexp.Compile(`\(([^)]+)\)`)
	if err != nil {
		return nil, err
	}

	// Write header.
	{
		_, err = fOut.Write([]byte(`package mime` + NewLine))
		if err != nil {
			return nil, err
		}

		_, err = fOut.Write([]byte(NewLine))
		if err != nil {
			return nil, err
		}

		_, err = fOut.Write([]byte(`// ` + category + "." + NewLine))
		if err != nil {
			return nil, err
		}

		_, err = fOut.Write([]byte(`const (` + NewLine))
		if err != nil {
			return nil, err
		}
	}

	var rec Record
	for i, row := range rows {
		if i == 0 {
			continue
		}

		rec = Record{
			TypeValue: row[1],
			Reference: row[2],
		}

		// Unfortunately, caser does not recognize dot symbol as separator.
		// Also, unfortunately, people working in IANA organisation put notes
		// for the entry into a field where notes should never be, i.e. into
		// the name field. This is crazy.
		{
			x := strings.ReplaceAll(row[0], `.`, `-`)
			rec.TypeName = cleanName(`Type` + category + caser.String(x))

			matches := re.FindAllString(rec.TypeName, -1)
			if len(matches) > 0 {
				if len(matches) > 1 {
					return nil, errors.New(ErrTooManyMatches)
				}

				rec.Comment = matches[0]
				rec.TypeName = strings.TrimSpace(strings.ReplaceAll(rec.TypeName, rec.Comment, ``))
			}
		}

		res.Records = append(res.Records, rec)

		// Write line.
		if len(rec.Comment) == 0 {
			_, err = fOut.Write([]byte(RecordLinePrefix + rec.TypeName + ` = "` + rec.TypeValue + `"` + ` // ` + rec.Reference + NewLine))
		} else {
			_, err = fOut.Write([]byte(RecordLinePrefix + rec.TypeName + ` = "` + rec.TypeValue + `"` + ` // ` + rec.Reference + " " + rec.Comment + NewLine))
		}

		if err != nil {
			return nil, err
		}
	}

	// Write footer.
	{
		_, err = fOut.Write([]byte(`)` + NewLine))
		if err != nil {
			return nil, err
		}

		_, err = fOut.Write([]byte(NewLine))
		if err != nil {
			return nil, err
		}

		_, err = fOut.Write([]byte(`const (` + NewLine))
		if err != nil {
			return nil, err
		}

		x := `"` + strings.ToLower(category) + `/*"`
		_, err = fOut.Write([]byte(RecordLinePrefix + `Type` + category + "Any = " + x + NewLine))
		if err != nil {
			return nil, err
		}

		_, err = fOut.Write([]byte(`)` + NewLine))
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
