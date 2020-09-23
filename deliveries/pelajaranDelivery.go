package deliveries

import (
	"fmt"
	"github/coroo/enigma-school/models"
	"github/coroo/enigma-school/usecases"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

type IPelajaranDelivery interface {
	GetPelajaranAll() ([]*models.Pelajaran, error)
	GetPelajaranById(id string) (*models.Pelajaran, error)
	DeletePelajaranById(id string) (*models.Pelajaran, error)
	PrintPelajaran(result []*models.Pelajaran)
	PrintOnePelajaran(result *models.Pelajaran)
	RegisterNewPelajaran(pelajaran models.Pelajaran) (*models.Pelajaran, error)
	UpdateExistingPelajaran(pelajaran models.Pelajaran) (*models.Pelajaran, error)
}

type PelajaranDelivery struct {
	pelajaranService usecases.IPelajaranUseCase
}

func NewPelajaranDelivery(service usecases.IPelajaranUseCase) IPelajaranDelivery {
	return &PelajaranDelivery{service}
}

func (pd *PelajaranDelivery) RegisterNewPelajaran(pelajaran models.Pelajaran) (*models.Pelajaran, error) {
	var err error
	err = pelajaran.Validate()
	if err != nil {
		return nil, err
	}
	return pd.pelajaranService.RegisterNewPelajaranService(pelajaran)
}

func (pd *PelajaranDelivery) UpdateExistingPelajaran(pelajaran models.Pelajaran) (*models.Pelajaran, error) {
	var err error
	err = pelajaran.Validate()
	if err != nil {
		return nil, err
	}
	return pd.pelajaranService.UpdateExistingPelajaranService(pelajaran)
}

func (pd *PelajaranDelivery) GetPelajaranAll() ([]*models.Pelajaran, error) {
	return pd.pelajaranService.GetPelajaranAllService()
}

func (pd *PelajaranDelivery) GetPelajaranById(id string) (*models.Pelajaran, error) {
	return pd.pelajaranService.GetPelajaranByIdService(id)
}

func (pd *PelajaranDelivery) DeletePelajaranById(id string) (*models.Pelajaran, error) {
	return pd.pelajaranService.DeletePelajaranByIdService(id)
}

func (pd *PelajaranDelivery) PrintOnePelajaran(result *models.Pelajaran) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Kode Pelajaran", "Nama Pelajaran", "Jumlah SKS"})
	t.AppendRow([]interface{}{"", result.KodePelajaran, result.NamaPelajaran, result.JumlahSks})
	t.Render()
}

func (pd *PelajaranDelivery) PrintPelajaran(result []*models.Pelajaran) {
	fmt.Println("\n ")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Kode Pelajaran", "Nama Pelajaran", "Jumlah SKS"})

	for a, p := range result {
		t.AppendRow([]interface{}{a + 1, p.KodePelajaran, p.NamaPelajaran, p.JumlahSks})
	}
	t.Render()
}
