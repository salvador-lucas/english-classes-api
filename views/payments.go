package views

type AddPaymentRequest struct {
	Amount   float32 `json:"amount"`
	StudenID uint    `json:"studentId"`
	Period   string  `json:"period"`
}

type AddPaymentResponse struct {
	Amount   float32 `json:"amount"`
	StudenID uint    `json:"studentId"`
	Period   string  `json:"period"`
}
