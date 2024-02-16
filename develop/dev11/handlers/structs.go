package handlers

type SuccessResponse struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
