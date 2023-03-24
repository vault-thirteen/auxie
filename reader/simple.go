package reader

import "io"

func ReadByte(r io.Reader) (b byte, err error) {
	buf := make([]byte, 1)
	_, err = r.Read(buf)
	if err != nil {
		return b, err
	}

	return buf[0], nil
}
