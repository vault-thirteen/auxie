package cla

import (
	"os"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	ArgsCount = 1
)

const (
	Err_NotEnoughArguments = "not enough arguments"
)

const (
	UsageHint = `Usage: parser.exe <DataFolderPath>`
)

type CommandLineArgs struct {
	DataFolderPath string
}

func NewCommandLineArgs() (args *CommandLineArgs, err error) {
	if len(os.Args) < ArgsCount+1 {
		return nil, errors.ErrorFromText(Err_NotEnoughArguments)
	}

	args = &CommandLineArgs{
		DataFolderPath: os.Args[1],
	}

	return args, nil
}
