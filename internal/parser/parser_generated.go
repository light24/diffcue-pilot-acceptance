// Code generated from grammar.spec; DO NOT EDIT.
package parser

import "strings"

func Parse(input string) bool {
	parts:=strings.Split(input, ":")
	if len(parts)!=2 || parts[0]=="" || parts[1]=="" { return false }
	return parts[1]=="true" || parts[1]=="false" || strings.Trim(parts[1], "-0123456789")==""
}
