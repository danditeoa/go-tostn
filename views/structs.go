package structs

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type Account struct {
	Id        int
	Name      string
	Cpf       string
	Balance   float64
	CreatedAt string
}

type Transfers struct {
	Id                   int `json:"id"`
	AccountOriginId      int
	AccountDestinationId int
	Amount               float64
	CreatedAt            string
}