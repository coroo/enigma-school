package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github/coroo/enigma-school/models"

	guuid "github.com/google/uuid"
)

type IPenempatanRepository interface {
	Insert(penempatan models.Penempatan) (*models.Penempatan, error)
	FindAllByNim(nim string) ([]*models.MahasiswaWithPelajaran, error)
	DeleteOneById(nim string) (*models.Penempatan, error)
}

var (
	penempatanQueries = map[string]string{
		"deleteOneById":          "DELETE FROM penempatan WHERE id=?",
		"penempatanFindAllByNim": "SELECT A.id,A.nim,A.kode_pelajaran,A.tahun_ajaran,B.nama_pelajaran,B.jumlah_sks,C.nama_mahasiswa,C.jurusan_mahasiswa FROM penempatan A LEFT JOIN m_pelajaran B ON A.kode_pelajaran=B.kode_pelajaran LEFT JOIN mahasiswa C ON A.nim=C.nim where A.nim=?",
		"insertPenempatan":       "INSERT into penempatan(id,nim,kode_pelajaran,tahun_ajaran) values(?,?,?,?)",
	}
)

type PenempatanRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewPenempatanRepository(db *sql.DB) IPenempatanRepository {
	ps := make(map[string]*sql.Stmt, len(penempatanQueries))
	for n, v := range penempatanQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &PenempatanRepository{
		db, ps,
	}
}

func (r *PenempatanRepository) Insert(penempatan models.Penempatan) (*models.Penempatan, error) {
	penempatanId := guuid.New().String()
	res, err := r.ps["insertPenempatan"].Exec(penempatanId, penempatan.Nim, penempatan.KodePelajaran, penempatan.TahunAjaran)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &penempatan, nil
}

func (r *PenempatanRepository) FindAllByNim(id string) ([]*models.MahasiswaWithPelajaran, error) {
	rows, err := r.ps["penempatanFindAllByNim"].Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.MahasiswaWithPelajaran, 0)
	for rows.Next() {
		res := new(models.MahasiswaWithPelajaran)
		err = rows.Scan(&res.Id, &res.Nim, &res.KodePelajaran, &res.TahunAjaran, &res.Pelajaran.NamaPelajaran, &res.Pelajaran.JumlahSks, &res.Mahasiswa.NamaMahasiswa, &res.Mahasiswa.JurusanMahasiswa)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}

func (r *PenempatanRepository) DeleteOneById(id string) (*models.Penempatan, error) {
	row := r.ps["deleteOneById"].QueryRow(id)
	res := new(models.Penempatan)
	err := row.Scan(&res.Id, &res.Nim, &res.KodePelajaran, &res.TahunAjaran)
	if err != nil {
		return res, err
	}
	return res, nil
}

// func (r *PenempatanRepository) FindDailyReportOder() (*models.PenempatanSummary, error) {
// 	currentTime := time.Now()
// 	row := r.ps["penempatanReportAllByToday"].QueryRow(currentTime.Format("2006-01-02"))
// 	res := new(models.PenempatanSummary)
// 	err := row.Scan(&res.PenempatanCount, &res.PenempatanTotal)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

// func (r *PenempatanRepository) FindMonthlyReportOder() (*models.PenempatanSummary, error) {
// 	currentTime := time.Now()
// 	row := r.ps["penempatanReportAllByMonthYear"].QueryRow(currentTime.Month(), currentTime.Year())
// 	res := new(models.PenempatanSummary)
// 	err := row.Scan(&res.PenempatanCount, &res.PenempatanTotal)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

// func (r *PenempatanRepository) FindReportPenempatan() (*models.PenempatanSummary, error) {
// 	row := r.ps["penempatanReportAllByMonthYear"].QueryRow()
// 	res := new(models.PenempatanSummary)
// 	err := row.Scan(&res.PenempatanCount, &res.PenempatanTotal)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }
