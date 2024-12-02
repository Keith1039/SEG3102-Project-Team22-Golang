package structs

type Student struct {
	StudentNumber  int    `db:"student_number"`
	TeamID         int    `db:"team_id"`
	StudentProgram string `db:"student_program"`
	CourseSection  string `db:"course_section"`
}
