package rs

// Seek is a wrapper of the child's method.
func (rs *ReaderSeeker) Seek(offset int64, whence int) (int64, error) {
	return rs.seeker.Seek(offset, whence)
}
