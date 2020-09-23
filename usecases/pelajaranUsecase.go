package usecases

import (
	"github/coroo/enigma-school/models"
	"github/coroo/enigma-school/repositories"
)

type IPelajaranUseCase interface {
	GetPelajaranAllService() ([]*models.Pelajaran, error)
	GetPelajaranByIdService(id string) (*models.Pelajaran, error)
	DeletePelajaranByIdService(id string) (*models.Pelajaran, error)
	RegisterNewPelajaranService(pelajaran models.Pelajaran) (*models.Pelajaran, error)
	UpdateExistingPelajaranService(pelajaran models.Pelajaran) (*models.Pelajaran, error)
}

type PelajaranUseCase struct {
	repo repositories.IPelajaranRepository
}

func NewPelajaranUseCase(repo repositories.IPelajaranRepository) IPelajaranUseCase {
	return &PelajaranUseCase{repo}
}

func (p *PelajaranUseCase) RegisterNewPelajaranService(pelajaran models.Pelajaran) (*models.Pelajaran, error) {
	return p.repo.Insert(pelajaran)
}

func (p *PelajaranUseCase) UpdateExistingPelajaranService(pelajaran models.Pelajaran) (*models.Pelajaran, error) {
	return p.repo.Update(pelajaran)
}

func (p *PelajaranUseCase) GetPelajaranAllService() ([]*models.Pelajaran, error) {
	return p.repo.FindAll()
}

func (p *PelajaranUseCase) GetPelajaranByIdService(id string) (*models.Pelajaran, error) {
	return p.repo.FindOneById(id)
}

func (p *PelajaranUseCase) DeletePelajaranByIdService(id string) (*models.Pelajaran, error) {
	return p.repo.DeleteOneById(id)
}
