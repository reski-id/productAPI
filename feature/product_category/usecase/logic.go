package usecase

import (
	"errors"
	"productapi/domain"
)

type productCatUseCase struct {
	productCatData domain.ProductCatData
}

func New(model domain.ProductCatData) domain.ProductCatUseCase {
	return &productCatUseCase{
		productCatData: model,
	}
}

func (nu *productCatUseCase) AddProductCategory(newData domain.ProductCategory) (domain.ProductCategory, error) {

	res := nu.productCatData.Insert(newData)

	if res.ID == 0 {
		return domain.ProductCategory{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *productCatUseCase) UpProductCategory(IDProductCategory int, updateData domain.ProductCategory) (domain.ProductCategory, error) {
	if IDProductCategory == -1 {
		return domain.ProductCategory{}, errors.New("invalid data")
	}

	res := nu.productCatData.Update(IDProductCategory, updateData)
	if res.ID == 0 {
		return domain.ProductCategory{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *productCatUseCase) DelProductCategory(IDProductCategory int) (bool, error) {
	res := nu.productCatData.Delete(IDProductCategory)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (nu *productCatUseCase) GetAllD() ([]domain.ProductCategory, error) {
	res := nu.productCatData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (nu *productCatUseCase) GetSpecificProductCategory(dataID int) ([]domain.ProductCategory, error) {
	res := nu.productCatData.GetbyID(dataID)
	if dataID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}
