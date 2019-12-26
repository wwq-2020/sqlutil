package main

const tplStr = `
// {{.Name}}Repo {{.Name}}Repo
type {{.Name}}Repo struct{
	db *sql.DB
}

// New{{.Name}}Repo New{{.Name}}Repo
func New{{.Name}}Repo(db *sql.DB) *{{.Name}}Repo {
	return &{{.Name}}Repo{
		db: db,
	}
}

{{range $idx,$each := .Bys}}
// FindBy{{$each.Name|title}} FindBy{{$each.Name|title}}
func (rp {{$.Name}}Repo) FindBy{{$each.Name|title}}(ctx context.Context, {{$each.Name}} {{$each.Type}}) ([]*{{$.Pkg}}{{$.Name}}, error) {
	rows, err := rp.db.QueryContext(ctx, "select {{$.Column}} from {{$.TableName}} where {{$each.Name}} = ?",{{$each.Name}})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*{{$.Pkg}}{{$.Name}}
	for rows.Next() {
		result := &{{$.Pkg}}{{$.Name}}{}
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

{{range $idx,$each := .Bys}}
// DeleteBy{{$each.Name|title}} DeleteBy{{$each.Name|title}}
func (rp {{$.Name}}Repo) DeleteBy{{$each.Name|title}}(ctx context.Context, {{$each.Name}} {{$each.Type}}) error {
	_, err := rp.db.ExecContext(ctx, "delete from {{$.TableName}} where {{$each.Name}} = ?",{{$each.Name}})
	if err != nil {
		return err
	}
	return nil
}
{{end}}

// Create Create
func (rp {{.Name}}Repo) Create(ctx context.Context,obj *{{.Pkg}}{{.Name}}) (int64, error) {
	result, err := rp.db.ExecContext(ctx, "insert into {{.TableName}} ({{.Column}}) values({{.PlaceHolder}})", {{.Value}})
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

// BatchCreate BatchCreate
func (rp {{.Name}}Repo) BatchCreate(ctx context.Context, objs []*{{.Pkg}}{{.Name}}) error {
	sqlBaseStr := "insert into {{.TableName}} ({{.Column}}) values %s"
	sqlPlaceHolder := make([]string, 0, len(objs))
	sqlArgs := make([]interface{}, 0, len(objs)*{{.ColumnCount}})
	for _, obj := range objs {
		sqlPlaceHolder = append(sqlPlaceHolder, "({{.PlaceHolder}})")
		sqlArgs = append(sqlArgs, {{.Value}})
	}
	sqlStr := fmt.Sprintf(sqlBaseStr, strings.Join(sqlPlaceHolder, ","))
	if _,err := rp.db.ExecContext(ctx, sqlStr, sqlArgs...); err != nil {
		return err
	}
	return nil
}
`
