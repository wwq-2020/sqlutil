package main

const tplStr = `
type {{.Name}}Repo struct{
	db *sql.DB
}

func New{{.Name}}Repo(db *sql.DB) *{{.Name}}Repo {
	return &{{.Name}}Repo{
		db: db,
	}
}

{{range $idx,$each := .Bys}}
func (rp {{$.Name}}Repo) FindBy{{$each.Name|title}}(ctx context.Context, {{$each.Name}} {{$each.Type}}) ([]*{{$.Name}}, error) {
	rows, err := rp.db.QueryContext(ctx, "select {{$.Column}} from {{$.TableName}} where {{$each.Name}} = ?",{{$each.Name}})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*{{$.Name}}
	for rows.Next() {
		result := &{{$.Name}}{}
		if err := rows.Scan({{$.Scan|raw}}); err !=nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
{{end}}


`
