package e

type CLSCTX uint32

const (
	CLSCTX_INPROC_SERVER                             CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER                            CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER                              CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16                           CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER                             CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16                          CLSCTX = 0x20
	CLSCTX_RESERVED1                                 CLSCTX = 0x40
	CLSCTX_RESERVED2                                 CLSCTX = 0x80
	CLSCTX_RESERVED3                                 CLSCTX = 0x100
	CLSCTX_RESERVED4                                 CLSCTX = 0x200
	CLSCTX_NO_CODE_DOWNLOAD                          CLSCTX = 0x400
	CLSCTX_RESERVED5                                 CLSCTX = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL                         CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD                      CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG                            CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA                               CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA                                CLSCTX = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT                      CLSCTX = 0x20000
	CLSCTX_ACTIVATE_X86_SERVER                       CLSCTX = 0x40000
	CLSCTX_ACTIVATE_32_BIT_SERVER                    CLSCTX = CLSCTX_ACTIVATE_X86_SERVER
	CLSCTX_ACTIVATE_64_BIT_SERVER                    CLSCTX = 0x80000
	CLSCTX_ENABLE_CLOAKING                           CLSCTX = 0x100000
	CLSCTX_APPCONTAINER                              CLSCTX = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU                        CLSCTX = 0x800000
	CLSCTX_RESERVED6                                 CLSCTX = 0x1000000
	CLSCTX_ACTIVATE_ARM32_SERVER                     CLSCTX = 0x2000000
	CLSCTX_ALLOW_LOWER_TRUST_REGISTRATION            CLSCTX = 0x4000000
	CLSCTX_SERVER_MUST_BE_EQUAL_OR_GREATER_PRIVILEGE CLSCTX = 0x8000000
	CLSCTX_DO_NOT_ELEVATE_SERVER                     CLSCTX = 0x10000000
	CLSCTX_PS_DLL                                    CLSCTX = 0x80000000
)
