package number

import (
	"reflect"
	"strconv"
	"unsafe"
)

func ParseUint(s string) (u uint, err error) {
	var tmp uint64
	tmp, err = strconv.ParseUint(s, 10, int(8*unsafe.Sizeof(reflect.Uint)))
	if err != nil {
		return 0, err
	}

	return uint(tmp), nil
}

func ParseUint64(s string) (u uint64, err error) {
	return strconv.ParseUint(s, 10, 64)
}

func ParseUint32(s string) (u uint32, err error) {
	var tmp uint64
	tmp, err = strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(tmp), nil
}

func ParseUint16(s string) (u uint16, err error) {
	var tmp uint64
	tmp, err = strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}

	return uint16(tmp), nil
}

func ParseUint8(s string) (u uint8, err error) {
	var tmp uint64
	tmp, err = strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(tmp), nil
}

func ParseInt(s string) (i int, err error) {
	var tmp int64
	tmp, err = strconv.ParseInt(s, 10, int(8*unsafe.Sizeof(reflect.Int)))
	if err != nil {
		return 0, err
	}

	return int(tmp), nil
}

func ParseInt64(s string) (i int64, err error) {
	return strconv.ParseInt(s, 10, 64)
}

func ParseInt32(s string) (i int32, err error) {
	var tmp int64
	tmp, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(tmp), nil
}

func ParseInt16(s string) (i int16, err error) {
	var tmp int64
	tmp, err = strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, err
	}

	return int16(tmp), nil
}

func ParseInt8(s string) (i int8, err error) {
	var tmp int64
	tmp, err = strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, err
	}

	return int8(tmp), nil
}
