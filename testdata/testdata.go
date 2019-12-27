package testdata

import _ "database/sql" //a

// User table:user a
type User struct {
	ID       int64 `sqlutil:"id,by"`
	Username int64 `sqlutil:"username"`
	Password int64 `sqlutil:"password"`
}

// User3 table:user3 a
type User3 struct {
	ID       int64 `sqlutil:"id,by"`
	Username int64 `sqlutil:"username"`
	Password int64 `sqlutil:"password"`
}

// A 2
type A int
