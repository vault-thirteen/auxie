package m

import "strings"

type MIME struct {
	Name      string
	Template  string
	Reference string
}

func (m *MIME) NameForFile() (out string) {
	out = strings.ReplaceAll(m.Name, "-", " ")
	out = strings.ReplaceAll(out, "_", " ")
	out = strings.ReplaceAll(out, "+", " ")
	out = strings.ReplaceAll(out, ".", " ")

	parts := strings.Split(out, " ")

	for i, _ := range parts {
		parts[i] = CapitalizeFirstLetter(parts[i])
	}

	out = strings.Join(parts, "")

	return out
}
