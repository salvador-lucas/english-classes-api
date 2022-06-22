package views

type GetStudentResponse struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type GetAllStudentsResponse struct {
	Students []GetStudentResponse `json:"students"`
}

type AddStudentRequest struct {
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Cellphone uint   `json:"cellphone"`
}

type AddStudentResponse struct {
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Cellphone uint   `json:"cellphone"`
}

//payments
type AddPaymentRequest struct {
	Amount float32 `json:"amount"`
	Period string  `json:"period"`
}

type AddPaymentResponse struct {
	Amount float32 `json:"amount"`
	Period string  `json:"period"`
}

type GetStudentPaymentsResponse struct {
	Payments []AddPaymentResponse `json:"payments"`
}
