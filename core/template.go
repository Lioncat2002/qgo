package core

const CREATE_TABLE = `
CREATE TABLE {{.Name}} (
	{{- range $index, $field := .Fields }}
	{{ $field.Name }} {{ $field.FieldType }}{{ if ne (add $index 1) (len $.Fields) }},{{ end }}
	{{- end }}
);
`