package testdata

// User table:user a
type User struct {
	ID       int64 `sqlutil:"id,by"`
	Username int64 `sqlutil:"username"`
	Password int64 `sqlutil:"password"`
}
