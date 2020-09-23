package usecases

import (
	"github/coroo/enigma-school/models"
	"github/coroo/enigma-school/repositories"
)

type IPenempatanUseCase interface {
	GetPenempatanByMahasiswaNimService(id string) ([]*models.MahasiswaWithPelajaran, error)
	DeletePenempatanByIdService(id string) (*models.Penempatan, error)
	RegisterNewPenempatanService(penempatan models.Penempatan) (*models.Penempatan, error)
}

type PenempatanUseCase struct {
	repo repositories.IPenempatanRepository
}

func NewPenempatanUseCase(repo repositories.IPenempatanRepository) IPenempatanUseCase {
	return &PenempatanUseCase{repo}
}

func (p *PenempatanUseCase) RegisterNewPenempatanService(penempatan models.Penempatan) (*models.Penempatan, error) {
	return p.repo.Insert(penempatan)
}

func (p *PenempatanUseCase) GetPenempatanByMahasiswaNimService(id string) ([]*models.MahasiswaWithPelajaran, error) {
	return p.repo.FindAllByNim(id)
}

func (p *PenempatanUseCase) DeletePenempatanByIdService(id string) (*models.Penempatan, error) {
	return p.repo.DeleteOneById(id)
}
