package cio

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/sys/windows"
	"golang.org/x/term"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_StdInputIsNotTerminal = "standard input is not a terminal console"
)

func GetInputKey() (key Key, err error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return key, errors.ErrorFromText(Err_StdInputIsNotTerminal)
	}

	var termStateOld *term.State
	termStateOld, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return key, err
	}
	defer func() {
		derr := term.Restore(int(os.Stdin.Fd()), termStateOld)
		if derr != nil {
			err = errors.Combine(err, derr)
		}
	}()

	seq := make([]byte, 5)

	var nBytesRead int
	nBytesRead, err = os.Stdin.Read(seq)
	if err != nil {
		return key, err
	}

	seq = seq[:nBytesRead]

	return KeyFromTerminalByteSequence(seq)
}

func GetInputString(trimSpace bool) (s string, err error) {
	reader := bufio.NewReader(os.Stdin)

	s, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if trimSpace {
		return strings.TrimSpace(s), nil
	}

	return s, nil
}

func GetInputPassword() (s string, err error) {
	var buf []byte
	buf, err = term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func GetInputInt() (i int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := strings.TrimSpace(scanner.Text())

	i, err = strconv.Atoi(inputString)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func GetInputIntRetry() (i int) {
	var err error
	i, err = GetInputInt()
	for err != nil {
		i, err = GetInputInt()
	}

	return i
}

func FormatMessage(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

func LogMessage(args ...any) {
	log.Println(args...)
}

func Print(args ...any) {
	fmt.Print(args...)
}

func PrintMessage(args ...any) {
	fmt.Println(args...)
}

func EnableAnsiSymbolsInTerminal() (err error) {
	if runtime.GOOS == "windows" {
		stdout := windows.Handle(os.Stdout.Fd())
		var mode uint32

		err = windows.GetConsoleMode(stdout, &mode)
		if err != nil {
			return err
		}

		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

		err = windows.SetConsoleMode(stdout, mode)
		if err != nil {
			return err
		}
	}

	return nil
}
