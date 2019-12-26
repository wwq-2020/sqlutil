package testdata
			import(
				"database/sql"
				"context"
			)
type User2Repo struct{
	db *sql.DB
}

func NewUser2Repo(db *sql.DB) *User2Repo {
	return &User2Repo{
		db: db,
	}
}


func (rp User2Repo) FindById(ctx context.Context, id int64) ([]*User2, error) {
	rows, err := rp.db.QueryContext(ctx, "select id, username, password from user2 where id = ?",id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*User2
	for rows.Next() {
		result := &User2{}
		if err := rows.Scan(&result.ID, &result.Username, &result.Password); err !=nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}



