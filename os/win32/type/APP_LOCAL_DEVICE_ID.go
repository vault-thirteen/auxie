package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

const APP_LOCAL_DEVICE_ID_SIZE = 32

type APP_LOCAL_DEVICE_ID struct {
	Value [APP_LOCAL_DEVICE_ID_SIZE]b.BYTE
}
