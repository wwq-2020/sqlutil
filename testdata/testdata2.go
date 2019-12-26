package testdata

// User2 table:user2 a
type User2 struct {
	ID       int64 `sqlutil:"id,by"`
	Username int64 `sqlutil:"username"`
	Password int64 `sqlutil:"password"`
}
