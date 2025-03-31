package dtos

type UpdateMemeCoinRequest struct {
	Description string `json:"description"`
}

type CreateMemeCoinRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
