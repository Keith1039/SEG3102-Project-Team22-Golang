package structs

type User struct {
	UserID    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	StudentID int    `db:"student_id"`
	Email     string `db:"email"`
	Role      string `db:"role"`
}
