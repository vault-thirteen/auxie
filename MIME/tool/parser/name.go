package main

import "strings"

var redundantSymbolsInName = []string{
	`-`,
	`+`,
	`.`,
	`_`,
}

func cleanName(in string) (out string) {
	out = in

	for _, rs := range redundantSymbolsInName {
		out = strings.ReplaceAll(out, rs, ``)
	}

	return out
}

// "Regular Expression" for searching normal symbols:
// [a-zA-Z0-9]+
