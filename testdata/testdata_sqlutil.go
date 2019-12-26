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

type User3Repo struct {
	db *sql.DB
}

func NewUser3Repo(db *sql.DB) *User3Repo {
	return &User3Repo{
		db: db,
	}
}

func (rp User3Repo) FindById(ctx context.Context, id int64) ([]*User3, error) {
	rows, err := rp.db.QueryContext(ctx, "select id, username, password from user3 where id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*User3
	for rows.Next() {
		result := &User3{}
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
