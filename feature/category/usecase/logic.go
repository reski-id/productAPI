package usecase

import (
	"errors"
	"productapi/domain"
)

type categoryUseCase struct {
	categoryData domain.CategoryData
}

func New(model domain.CategoryData) domain.CategoryUseCase {
	return &categoryUseCase{
		categoryData: model,
	}
}

func (nu *categoryUseCase) AddCategory(IDUser int, newCategory domain.Category) (domain.Category, error) {
	if IDUser == -1 {
		return domain.Category{}, errors.New("invalid data")
	}

	res := nu.categoryData.Insert(newCategory)

	if res.ID == 0 {
		return domain.Category{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *categoryUseCase) UpCategory(IDCategory int, updateData domain.Category) (domain.Category, error) {
	if IDCategory == -1 {
		return domain.Category{}, errors.New("invalid data")
	}

	res := nu.categoryData.Update(IDCategory, updateData)
	if res.ID == 0 {
		return domain.Category{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *categoryUseCase) DelCategory(IDCategory int) (bool, error) {
	res := nu.categoryData.Delete(IDCategory)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (nu *categoryUseCase) GetSpecificCategory(catID int) ([]domain.Category, error) {
	res := nu.categoryData.GetCategoryID(catID)
	if catID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}

func (nu *categoryUseCase) GetAllP() ([]domain.Category, error) {
	res := nu.categoryData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}
