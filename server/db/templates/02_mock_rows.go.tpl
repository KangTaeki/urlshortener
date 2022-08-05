{{$alias := .Aliases.Table .Table.Name}}

func Mock{{$alias.UpSingular}}Rows(m sqlmock.Sqlmock, data ...*{{$alias.UpSingular}}) *sqlmock.Rows {
    rows := m.NewRows({{$alias.UpSingular}}AllColumns)
    for _, d := range data {
        rows = rows.AddRow(
            {{range $column := .Table.Columns -}}
        	{{- $colAlias := $alias.Column $column.Name -}}
        	d.{{$colAlias}},
        	{{end -}}
        )
    }
    return rows
}
