package ssc

import "sync"

type KeyType interface {
	string | int | uint
}

type ValueType interface {
	any
}

// SSC is a Simple Stupid Cache.
type SSC[KT KeyType, VT ValueType] struct {
	values  map[KT]VT
	maxSize int
	lock    sync.RWMutex
}

func NewSSC[KT KeyType, VT ValueType](maxSize int) (ssc *SSC[KT, VT]) {
	if maxSize <= 0 {
		return nil
	}

	return &SSC[KT, VT]{
		values:  map[KT]VT{},
		lock:    sync.RWMutex{},
		maxSize: maxSize,
	}
}

// GetSize returns the cache size.
func (ssc *SSC[KT, VT]) GetSize() (cacheSize int) {
	ssc.lock.RLock()
	defer ssc.lock.RUnlock()

	return ssc.getSize()
}

func (ssc *SSC[KT, VT]) getSize() (cacheSize int) {
	return len(ssc.values)
}

// GetValue gets the cache record.
func (ssc *SSC[KT, VT]) GetValue(key KT) (value VT, recordExists bool) {
	ssc.lock.RLock()
	defer ssc.lock.RUnlock()

	return ssc.getValue(key)
}

func (ssc *SSC[KT, VT]) getValue(key KT) (value VT, recordExists bool) {
	value, recordExists = ssc.values[key]
	return value, recordExists
}

// SetValue sets the cache record value.
func (ssc *SSC[KT, VT]) SetValue(key KT, value VT) (recordAlreadyExists bool) {
	ssc.lock.Lock()
	defer ssc.lock.Unlock()

	return ssc.setValue(key, value)
}

func (ssc *SSC[KT, VT]) setValue(key KT, value VT) (recordAlreadyExists bool) {
	_, recordAlreadyExists = ssc.values[key]

	if recordAlreadyExists {
		ssc.changeValue(key, value)
		return true
	} else {
		ssc.addValue(key, value)
		return false
	}
}

func (ssc *SSC[KT, VT]) changeValue(existingKey KT, newValue VT) {
	ssc.values[existingKey] = newValue
}

func (ssc *SSC[KT, VT]) addValue(newKey KT, newValue VT) {
	if ssc.isFull() {
		ssc.removeRandomRecord()
	}

	ssc.values[newKey] = newValue
}

func (ssc *SSC[KT, VT]) isFull() bool {
	return ssc.getSize() == ssc.maxSize
}

func (ssc *SSC[KT, VT]) removeRandomRecord() {
	l := ssc.getSize()
	if l == 0 {
		return
	}

	iX := (l - 1) / 2
	i := 0
	for k, _ := range ssc.values {
		if i == iX {
			delete(ssc.values, k)
			return
		}

		i++
	}
}

// Reset clears the cache.
func (ssc *SSC[KT, VT]) Reset() {
	ssc.lock.Lock()
	defer ssc.lock.Unlock()

	ssc.reset()
}

func (ssc *SSC[KT, VT]) reset() {
	ssc.values = map[KT]VT{}
}

// Delete deletes a record from cache.
func (ssc *SSC[KT, VT]) Delete(key KT) {
	ssc.lock.Lock()
	defer ssc.lock.Unlock()

	ssc.delete(key)
}

func (ssc *SSC[KT, VT]) delete(key KT) {
	delete(ssc.values, key)
}
