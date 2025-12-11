package main

import "strings"

type CustomStringsBuilder struct {
	strings.Builder
}

func (sb *CustomStringsBuilder) WriteStrings(s ...string) {
	for _, el := range s {
		sb.WriteString(el)
	}
}
