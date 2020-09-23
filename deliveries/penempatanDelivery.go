package deliveries

import (
	"fmt"
	"github/coroo/enigma-school/models"
	"github/coroo/enigma-school/usecases"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

type IPenempatanDelivery interface {
	GetPenempatanByMahasiswaNim(id string) ([]*models.MahasiswaWithPelajaran, error)
	DeletePenempatanById(id string) (*models.Penempatan, error)
	PrintOnePenempatan(result *models.Penempatan)
	PrintPenempatan(result []*models.MahasiswaWithPelajaran)
	RegisterNewPenempatan(penempatan models.Penempatan) (*models.Penempatan, error)
}

type PenempatanDelivery struct {
	mahasiswaService usecases.IPenempatanUseCase
}

func NewPenempatanDelivery(service usecases.IPenempatanUseCase) IPenempatanDelivery {
	return &PenempatanDelivery{service}
}

func (pd *PenempatanDelivery) RegisterNewPenempatan(penempatan models.Penempatan) (*models.Penempatan, error) {
	var err error
	err = penempatan.Validate()
	if err != nil {
		return nil, err
	}
	return pd.mahasiswaService.RegisterNewPenempatanService(penempatan)
}

func (pd *PenempatanDelivery) GetPenempatanByMahasiswaNim(id string) ([]*models.MahasiswaWithPelajaran, error) {
	return pd.mahasiswaService.GetPenempatanByMahasiswaNimService(id)
}

func (pd *PenempatanDelivery) DeletePenempatanById(id string) (*models.Penempatan, error) {
	return pd.mahasiswaService.DeletePenempatanByIdService(id)
}

func (pd *PenempatanDelivery) PrintOnePenempatan(result *models.Penempatan) {
	fmt.Printf("%v\n", result)
}

func (pd *PenempatanDelivery) PrintPenempatan(result []*models.MahasiswaWithPelajaran) {
	fmt.Println("\n ")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Kode Pelajaran", "Nama Pelajaran", "Jumlah SKS", "Tahun Ajaran"})

	var totalSks int
	for a, p := range result {
		if a == 0 {
			fmt.Println("Nama Mahasiswa: ", p.Mahasiswa.NamaMahasiswa)
			fmt.Println("Jurusan Mahasiswa: ", p.Mahasiswa.JurusanMahasiswa)
		}

		t.AppendRow([]interface{}{a + 1, p.Pelajaran.KodePelajaran, p.Pelajaran.NamaPelajaran, p.Pelajaran.JumlahSks, p.TahunAjaran})
		totalSks += p.Pelajaran.JumlahSks
	}
	t.AppendFooter(table.Row{"", "Total SKS Diambil", "", totalSks, ""})
	t.Render()
}
