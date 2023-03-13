package delivery

import "productapi/domain"

type DataResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Enable      bool   `json:"enable" form:"enable"`
}

func FromDomain(data domain.Product) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.Name = data.Name
	res.Description = data.Description
	res.Enable = data.Enable
	return res
}
