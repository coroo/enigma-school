package usecases

import (
	"github/coroo/enigma-school/models"
	"github/coroo/enigma-school/repositories"
)

type IMahasiswaUseCase interface {
	GetMahasiswaByIdService(id string) (*models.Mahasiswa, error)
	DeleteMahasiswaByIdService(id string) (*models.Mahasiswa, error)
	RegisterNewMahasiswaService(MahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error)
	UpdateExistingMahasiswaService(MahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error)
}

type MahasiswaUseCase struct {
	repo repositories.IMahasiswaRepository
}

func NewMahasiswaUseCase(repo repositories.IMahasiswaRepository) IMahasiswaUseCase {
	return &MahasiswaUseCase{repo}
}

func (p *MahasiswaUseCase) RegisterNewMahasiswaService(MahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error) {
	return p.repo.Insert(MahasiswaWithPelajaran)
}

func (p *MahasiswaUseCase) UpdateExistingMahasiswaService(MahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error) {
	return p.repo.Update(MahasiswaWithPelajaran)
}

func (p *MahasiswaUseCase) GetMahasiswaByIdService(id string) (*models.Mahasiswa, error) {
	return p.repo.FindOneById(id)
}

func (p *MahasiswaUseCase) DeleteMahasiswaByIdService(id string) (*models.Mahasiswa, error) {
	return p.repo.DeleteOneById(id)
}
