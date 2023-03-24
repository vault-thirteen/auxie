package bom

import "sort"

type encodingSorter []Encoding

func (bs encodingSorter) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs encodingSorter) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs encodingSorter) Len() int {
	return len(bs)
}

func sortEncodings(e []Encoding) {
	sort.Sort(encodingSorter(e))
}
