package m

import "strings"

type Header struct {
	FieldName string
	Template  string
	Protocol  string
	Status    string
	Reference string
	Comment   string
}

func (h *Header) CommentForFile() (out string) {
	out = h.Reference

	if len(h.Comment) > 0 {
		out += " " + h.Comment
	}

	return out
}

func (h *Header) NameForFile() (out string) {
	out = strings.ReplaceAll(h.FieldName, "-", "")
	out = strings.ReplaceAll(out, "_", "")
	return out
}
