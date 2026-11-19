package t

import b "github.com/vault-thirteen/auxie/os/win32/base"

type SECURITY_ATTRIBUTES struct {
	NLength              b.DWORD
	LpSecurityDescriptor b.LPVOID
	BInheritHandle       b.BOOL
}

type PSECURITY_ATTRIBUTES = *SECURITY_ATTRIBUTES
type LPSECURITY_ATTRIBUTES = *SECURITY_ATTRIBUTES
