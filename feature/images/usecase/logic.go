package usecase

import (
	"errors"
	"productapi/domain"
)

type imagesUseCase struct {
	imagesData domain.ImagesData
}

func New(model domain.ImagesData) domain.ImagesUseCase {
	return &imagesUseCase{
		imagesData: model,
	}
}

func (nu *imagesUseCase) AddImages(IDUser int, newImages domain.Images) (domain.Images, error) {
	if IDUser == -1 {
		return domain.Images{}, errors.New("invalid data")
	}

	res := nu.imagesData.Insert(newImages)

	if res.ID == 0 {
		return domain.Images{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *imagesUseCase) UpImages(IDImages int, updateData domain.Images) (domain.Images, error) {
	if IDImages == -1 {
		return domain.Images{}, errors.New("invalid data")
	}

	res := nu.imagesData.Update(IDImages, updateData)
	if res.ID == 0 {
		return domain.Images{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *imagesUseCase) DelImages(IDImages int) (bool, error) {
	res := nu.imagesData.Delete(IDImages)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (nu *imagesUseCase) GetSpecificImages(prodID int) ([]domain.Images, error) {
	res := nu.imagesData.GetImagesID(prodID)
	if prodID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}

func (nu *imagesUseCase) GetAllP() ([]domain.Images, error) {
	res := nu.imagesData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}
