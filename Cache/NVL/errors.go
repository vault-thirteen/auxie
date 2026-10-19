package nvl

const (
	Err_BottomRecordDoesNotExist = "bottom record does not exist"
	Err_UidIsEmpty               = "UID is empty" //TODO
	Err_TtlIsZero                = "zero TTL will totally disable the cache"

	Errf_RecordIsNotFound = `record is not found, uid=%v`
	Errf_RecordIsOutdated = `record is outdated, uid=%v`
)
