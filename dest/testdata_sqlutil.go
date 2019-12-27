package testdata
	import(
		"database/sql"
		"context"
		"fmt"
		"strings"
				testdata "github.com/wwq1988/sqlutil/testdata"
			
	)
// UserRepo UserRepo
type UserRepo struct{
	db *sql.DB
}

// NewUserRepo NewUserRepo
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}


// FindByID FindByID
func (rp UserRepo) FindByID(ctx context.Context, id int64) ([]*testdata.User, error) {
	rows, err := rp.db.QueryContext(ctx, "select id, username, password from user where id = ?",id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*testdata.User
	for rows.Next() {
		result := &testdata.User{}
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
func (rp UserRepo) DeleteByID(ctx context.Context, id int64) error {
	_, err := rp.db.ExecContext(ctx, "delete from user where id = ?",id)
	if err != nil {
		return err
	}
	return nil
}


// Create Create
func (rp UserRepo) Create(ctx context.Context,obj *testdata.User) (int64, error) {
	result, err := rp.db.ExecContext(ctx, "insert into user (id, username, password) values(?, ?, ?)", obj.ID, obj.Username, obj.Password)
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
func (rp UserRepo) BatchCreate(ctx context.Context, objs []*testdata.User) error {
	sqlBaseStr := "insert into user (id, username, password) values %s"
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

// User3Repo User3Repo
type User3Repo struct{
	db *sql.DB
}

// NewUser3Repo NewUser3Repo
func NewUser3Repo(db *sql.DB) *User3Repo {
	return &User3Repo{
		db: db,
	}
}


// FindByID FindByID
func (rp User3Repo) FindByID(ctx context.Context, id int64) ([]*testdata.User3, error) {
	rows, err := rp.db.QueryContext(ctx, "select id, username, password from user3 where id = ?",id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*testdata.User3
	for rows.Next() {
		result := &testdata.User3{}
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
func (rp User3Repo) DeleteByID(ctx context.Context, id int64) error {
	_, err := rp.db.ExecContext(ctx, "delete from user3 where id = ?",id)
	if err != nil {
		return err
	}
	return nil
}


// Create Create
func (rp User3Repo) Create(ctx context.Context,obj *testdata.User3) (int64, error) {
	result, err := rp.db.ExecContext(ctx, "insert into user3 (id, username, password) values(?, ?, ?)", obj.ID, obj.Username, obj.Password)
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
func (rp User3Repo) BatchCreate(ctx context.Context, objs []*testdata.User3) error {
	sqlBaseStr := "insert into user3 (id, username, password) values %s"
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
