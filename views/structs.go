package structs

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequest struct {
	name string `json:"name"`
	cpf string `json:"cpf"`
	balance float64  `json:"balance"`
	datetime string `json:"datetime"`
}