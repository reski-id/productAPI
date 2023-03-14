package delivery

import "productapi/domain"

type DataResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	Enable bool   `json:"enable" form:"enable"`
}

func FromDomain(data domain.Category) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.Name = data.Name
	res.Enable = data.Enable
	return res
}
