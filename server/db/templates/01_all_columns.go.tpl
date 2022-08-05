{{$alias := .Aliases.Table .Table.Name}}

var {{$alias.UpSingular}}AllColumns = []string{
	{{range $column := .Table.Columns -}}
	{{- $colAlias := $alias.Column $column.Name -}}
	{{$alias.UpSingular}}Columns.{{$colAlias}},
	{{end -}}
}
