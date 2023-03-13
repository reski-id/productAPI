package usecase

import (
	"errors"
	"productapi/domain"
)

type productUseCase struct {
	productData domain.ProductData
}

func New(model domain.ProductData) domain.ProductUseCase {
	return &productUseCase{
		productData: model,
	}
}

func (nu *productUseCase) AddProduct(IDUser int, newProduct domain.Product) (domain.Product, error) {
	if IDUser == -1 {
		return domain.Product{}, errors.New("invalid data")
	}

	res := nu.productData.Insert(newProduct)

	if res.ID == 0 {
		return domain.Product{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *productUseCase) UpProduct(IDProduct int, updateData domain.Product) (domain.Product, error) {
	if IDProduct == -1 {
		return domain.Product{}, errors.New("invalid data")
	}

	res := nu.productData.Update(IDProduct, updateData)
	if res.ID == 0 {
		return domain.Product{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *productUseCase) DelProduct(IDProduct int) (bool, error) {
	res := nu.productData.Delete(IDProduct)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (nu *productUseCase) GetSpecificProduct(prodID int) ([]domain.Product, error) {
	res := nu.productData.GetProductID(prodID)
	if prodID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}

func (nu *productUseCase) GetAllP() ([]domain.Product, error) {
	res := nu.productData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}
