package structs

type UserAuth struct {
	UserID   int    `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
