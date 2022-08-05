{{$alias := .Aliases.Table .Table.Name}}
{{- $length := len .Table.Columns -}}
{{ if gt $length 1 }}
func BulkInsert{{$alias.UpPlural}}(ctx context.Context, exec boil.ContextExecutor, rows []*{{$alias.UpSingular}}) error {
    if len(rows) == 0 {
        return nil
    }
    
    var (
        noPKCols = {{$alias.DownSingular}}AllColumns[1:]
        noPKColsNum = len(noPKCols)
    )
    valueStrings := make([]string, len(rows))
    valueArgs := make([]interface{}, 0, noPKColsNum*len(rows))
    for i, row := range rows {
        valueStrings[i] = "("+strmangle.Placeholders(false, noPKColsNum, 1, 1)+")"
        {{range $idx, $column := .Table.Columns -}}
            {{- $colAlias := $alias.Column $column.Name -}}
            {{- if ne $idx 0 -}}
        valueArgs = append(valueArgs, row.{{$colAlias}})
            {{end -}}
        {{end -}}
    }

	stmt := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s",
		TableNames.{{titleCase $.Table.Name}},
		strings.Join(noPKCols, ","),
		strings.Join(valueStrings, ","),
	)

	_, err := exec.ExecContext(ctx, stmt, valueArgs...)
	return errors.WithStack(err)
}
{{ end }}
