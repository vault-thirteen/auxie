package dll

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

func (dc *DllController) EnableConsoleColours() (err error) {
	var hConsole b.HANDLE
	hConsole, err = dc.GetStdHandle(b.STD_OUTPUT_HANDLE)
	if err != nil {
		return err
	}

	var dwMode b.DWORD
	_, err = dc.GetConsoleMode(hConsole, &dwMode)
	if err != nil {
		return err
	}

	dwMode |= b.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	_, err = dc.SetConsoleMode(hConsole, dwMode)
	if err != nil {
		return err
	}

	return nil
}
