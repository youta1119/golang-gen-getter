package generator

import (
	"bytes"
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/hori-ryota/go-strcase"
	"strings"
)

type GetterGenerator struct {
	PkgName  string
	TypeName string
	Fields   []*Field
}

type Field struct {
	Name     string
	TypeName string
}

func (g *GetterGenerator) Generate() ([]byte, error) {
	f := jen.NewFile(g.PkgName)
	receiverName := strings.ToLower(string(g.TypeName[0]))
	for _, field := range g.Fields {
		f.Func().
			Params(jen.Id(receiverName).Op("*").Id(g.TypeName)). // receiver
			Id(formatGetterName(field.Name)). // function name
			Params().
			Id(field.TypeName). // return type
			Block(jen.Return(jen.Id(receiverName).Dot(field.Name))) //function block
	}

	buf := &bytes.Buffer{}
	err := f.Render(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to generate code: %w", err)
	}
	return buf.Bytes(), nil
}

// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

func formatGetterName(name string) string {
	if u := strings.ToUpper(name); commonInitialisms[u] {
		return u
	}
	return strcase.ToUpperCamel(name)
}
