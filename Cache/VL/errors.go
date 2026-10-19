package vl

const (
	Err_BottomRecordDoesNotExist = "bottom record does not exist"
	Err_UidIsEmpty               = "UID is empty" //TODO
	Err_DataIsEmpty              = "data is empty"
	Err_RecordIsTooBig           = "record is too big"
	Err_TtlIsZero                = "zero TTL will totally disable the cache"

	Errf_RecordIsNotFound = `record is not found, uid=%v`
	Errf_RecordIsOutdated = `record is outdated, uid=%v`
)
