package delivery

import "productapi/domain"

type DataResponse struct {
	ID         int `json:"id"  form:"id"`
	ProductID  int `json:"ProductID" form:"ProductID"`
	CategoryID int `json:"CategoryID" form:"CategoryID"`
}

func FromDomain(data domain.ProductCategory) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.ProductID = data.ProductID
	res.CategoryID = data.CategoryID
	return res
}
