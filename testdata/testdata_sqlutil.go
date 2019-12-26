package testdata

import (
	"context"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (rp UserRepo) FindById(ctx context.Context, id int64) ([]*User, error) {
	rows, err := rp.db.QueryContext(ctx, "select id, username, password from user where id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*User
	for rows.Next() {
		result := &User{}
		if err := rows.Scan(&result.ID, &result.Username, &result.Password); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
