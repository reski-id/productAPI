package usecase

import (
	"errors"
	"productapi/domain"
)

type productImgUseCase struct {
	productImgData domain.ProductImgData
}

func New(model domain.ProductImgData) domain.ProductImgUseCase {
	return &productImgUseCase{
		productImgData: model,
	}
}

func (nu *productImgUseCase) AddProductImages(newData domain.ProductImages) (domain.ProductImages, error) {

	res := nu.productImgData.Insert(newData)

	if res.ID == 0 {
		return domain.ProductImages{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *productImgUseCase) UpProductImages(IDProductImages int, updateData domain.ProductImages) (domain.ProductImages, error) {
	if IDProductImages == -1 {
		return domain.ProductImages{}, errors.New("invalid data")
	}

	res := nu.productImgData.Update(IDProductImages, updateData)
	if res.ID == 0 {
		return domain.ProductImages{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *productImgUseCase) DelProductImages(IDProductImages int) (bool, error) {
	res := nu.productImgData.Delete(IDProductImages)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (nu *productImgUseCase) GetAllD() ([]domain.ProductImages, error) {
	res := nu.productImgData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (nu *productImgUseCase) GetSpecificProductImages(dataID int) ([]domain.ProductImages, error) {
	res := nu.productImgData.GetbyID(dataID)
	if dataID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}
