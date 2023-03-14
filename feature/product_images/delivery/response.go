package delivery

import "productapi/domain"

type DataResponse struct {
	ID        int `json:"id"  form:"id"`
	ProductID int `json:"ProductID" form:"ProductID"`
	ImagesID  int `json:"ImagesID" form:"ImagesID"`
}

func FromDomain(data domain.ProductImages) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.ProductID = data.ProductID
	res.ImagesID = data.ImagesID
	return res
}
