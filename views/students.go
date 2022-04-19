package views

type GetStudentResponse struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type GetAllStudentsResponse struct {
	Students []GetStudentResponse `json:"students"`
}
