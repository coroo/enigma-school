package deliveries

import (
	"github/coroo/enigma-school/models"
	"github/coroo/enigma-school/usecases"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

type IMahasiswaDelivery interface {
	GetMahasiswaById(id string) (*models.Mahasiswa, error)
	DeleteMahasiswaById(id string) (*models.Mahasiswa, error)
	PrintOneMahasiswa(result *models.Mahasiswa)
	RegisterNewMahasiswa(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error)
	UpdateExistingMahasiswa(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error)
}

type MahasiswaDelivery struct {
	mahasiswaService usecases.IMahasiswaUseCase
}

func NewMahasiswaDelivery(service usecases.IMahasiswaUseCase) IMahasiswaDelivery {
	return &MahasiswaDelivery{service}
}

func (pd *MahasiswaDelivery) RegisterNewMahasiswa(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error) {
	var err error
	err = mahasiswaWithPelajaran.Mahasiswa.Validate()
	if err != nil {
		return nil, err
	}
	return pd.mahasiswaService.RegisterNewMahasiswaService(mahasiswaWithPelajaran)
}

func (pd *MahasiswaDelivery) UpdateExistingMahasiswa(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error) {
	var err error
	err = mahasiswaWithPelajaran.Mahasiswa.Validate()
	if err != nil {
		return nil, err
	}
	return pd.mahasiswaService.UpdateExistingMahasiswaService(mahasiswaWithPelajaran)
}

func (pd *MahasiswaDelivery) GetMahasiswaById(id string) (*models.Mahasiswa, error) {
	return pd.mahasiswaService.GetMahasiswaByIdService(id)
}

func (pd *MahasiswaDelivery) DeleteMahasiswaById(id string) (*models.Mahasiswa, error) {
	return pd.mahasiswaService.DeleteMahasiswaByIdService(id)
}

func (pd *MahasiswaDelivery) PrintOneMahasiswa(result *models.Mahasiswa) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "NIM", "Nama Mahasiswa", "Jurusan Mahasiswa"})
	t.AppendRow([]interface{}{"", result.Nim, result.NamaMahasiswa, result.JurusanMahasiswa})
	t.Render()
}
