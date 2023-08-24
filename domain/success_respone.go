package domain

type SuccessResponse struct {
	Message string `json:"message"`
	Data    Movie  `json:"data"`
}
