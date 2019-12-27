package testdata
	import(
		"database/sql"
		"context"
		"fmt"
		"strings"
				testdata "github.com/wwq1988/sqlutil/testdata/t"
			
	)
// User2Repo User2Repo
type User2Repo struct{
	db *sql.DB
}

// NewUser2Repo NewUser2Repo
func NewUser2Repo(db *sql.DB) *User2Repo {
	return &User2Repo{
		db: db,
	}
}


// FindByID FindByID
func (rp User2Repo) FindByID(ctx context.Context, id int64) ([]*testdata.User2, error) {
	rows, err := rp.db.QueryContext(ctx, "select id, username, password from user2 where id = ?",id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*testdata.User2
	for rows.Next() {
		result := &testdata.User2{}
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



// DeleteByID DeleteByID
func (rp User2Repo) DeleteByID(ctx context.Context, id int64) error {
	_, err := rp.db.ExecContext(ctx, "delete from user2 where id = ?",id)
	if err != nil {
		return err
	}
	return nil
}


// Create Create
func (rp User2Repo) Create(ctx context.Context,obj *testdata.User2) (int64, error) {
	result, err := rp.db.ExecContext(ctx, "insert into user2 (id, username, password) values(?, ?, ?)", obj.ID, obj.Username, obj.Password)
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
func (rp User2Repo) BatchCreate(ctx context.Context, objs []*testdata.User2) error {
	sqlBaseStr := "insert into user2 (id, username, password) values %s"
	sqlPlaceHolder := make([]string, 0, len(objs))
	sqlArgs := make([]interface{}, 0, len(objs)*3)
	for _, obj := range objs {
		sqlPlaceHolder = append(sqlPlaceHolder, "(?, ?, ?)")
		sqlArgs = append(sqlArgs, obj.ID, obj.Username, obj.Password)
	}
	sqlStr := fmt.Sprintf(sqlBaseStr, strings.Join(sqlPlaceHolder, ","))
	if _,err := rp.db.ExecContext(ctx, sqlStr, sqlArgs...); err != nil {
		return err
	}
	return nil
}
