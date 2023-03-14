package delivery

import "productapi/domain"

type DataResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	File   string `json:"file" form:"file"`
	Enable bool   `json:"enable" form:"enable"`
}

func FromDomain(data domain.Images) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.Name = data.Name
	res.File = data.File
	res.Enable = data.Enable
	return res
}
