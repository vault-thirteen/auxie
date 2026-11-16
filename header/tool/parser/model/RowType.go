package m

type RowType byte

const (
	RowType_Http               = RowType(1)
	RowType_PermanentMessage   = RowType(2)
	RowType_ProvisionalMessage = RowType(3)
)
