package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github/coroo/enigma-school/models"
)

type IMahasiswaRepository interface {
	Insert(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error)
	Update(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error)
	FindOneById(nim string) (*models.Mahasiswa, error)
	DeleteOneById(nim string) (*models.Mahasiswa, error)

	// FindDailyReportOder() (*models.MahasiswaSummary, error)
	// FindMonthlyReportOder() (*models.MahasiswaSummary, error)
	// FindReportMahasiswa() (*models.MahasiswaSummary, error)
}

var (
	mahasiswaQueries = map[string]string{
		"deleteOneById":        "DELETE FROM mahasiswa WHERE nim=?",
		"mahasiswaFindOneById": "SELECT nim,nama_mahasiswa,jurusan_mahasiswa FROM mahasiswa where nim=?",
		"insertMahasiswa":      "INSERT into mahasiswa(nim,nama_mahasiswa,jurusan_mahasiswa) values(?,?,?)",
		"updateMahasiswa":      "UPDATE mahasiswa SET nama_mahasiswa = ?,jurusan_mahasiswa = ? WHERE nim = ?",
		// "mahasiswaReportAll":            "select sum(mahasiswa_quantity),sum(total_price) from mahasiswa",
		// "mahasiswaReportAllByMonthYear": "select sum(mahasiswa_quantity),sum(total_price) from mahasiswa WHERE month(created_at)=? AND year(created_at)=?",
		// "mahasiswaReportAllByToday":     "select sum(mahasiswa_quantity),sum(total_price) from mahasiswa WHERE created_at>=?",
	}
)

type MahasiswaRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewMahasiswaRepository(db *sql.DB) IMahasiswaRepository {
	ps := make(map[string]*sql.Stmt, len(mahasiswaQueries))
	for n, v := range mahasiswaQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &MahasiswaRepository{
		db, ps,
	}
}

func (r *MahasiswaRepository) Insert(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error) {
	// productRow := r.db.QueryRow("select count(nim), sum(product_price) from products where nama_mahasiswa=?", mahasiswaWithPelajaran.Mahasiswa.PelajaranCode)
	// checkPelajaran := new(models.TotalPelajaran)
	// err := productRow.Scan(&checkPelajaran.Count, &checkPelajaran.PelajaranSks)
	// if err != nil {
	// 	return nil, err
	// }
	// if checkPelajaran.Count < 1 {
	// 	fmt.Println("\nErr: Pelajaran code not found")
	// 	return nil, nil
	// }

	// mahasiswaId := guunim.New().String()
	// calculateTotalPrice := checkPelajaran.PelajaranSks
	res, err := r.ps["insertMahasiswa"].Exec(mahasiswaWithPelajaran.Mahasiswa.Nim, mahasiswaWithPelajaran.Mahasiswa.NamaMahasiswa, mahasiswaWithPelajaran.JurusanMahasiswa)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &mahasiswaWithPelajaran, nil
}

func (r *MahasiswaRepository) Update(mahasiswaWithPelajaran models.MahasiswaWithPelajaran) (*models.MahasiswaWithPelajaran, error) {
	res, err := r.ps["updateMahasiswa"].Exec(mahasiswaWithPelajaran.Mahasiswa.NamaMahasiswa, mahasiswaWithPelajaran.JurusanMahasiswa, mahasiswaWithPelajaran.Mahasiswa.Nim)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Update failed", err))
	}
	return &mahasiswaWithPelajaran, nil
}

func (r *MahasiswaRepository) FindOneById(nim string) (*models.Mahasiswa, error) {
	row := r.ps["mahasiswaFindOneById"].QueryRow(nim)
	res := new(models.Mahasiswa)
	err := row.Scan(&res.Nim, &res.NamaMahasiswa, &res.JurusanMahasiswa)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *MahasiswaRepository) DeleteOneById(nim string) (*models.Mahasiswa, error) {
	row := r.ps["deleteOneById"].QueryRow(nim)
	res := new(models.Mahasiswa)
	err := row.Scan(&res.Nim, &res.NamaMahasiswa, &res.JurusanMahasiswa)
	if err != nil {
		return res, err
	}
	return res, nil
}

// func (r *MahasiswaRepository) FindDailyReportOder() (*models.MahasiswaSummary, error) {
// 	currentTime := time.Now()
// 	row := r.ps["mahasiswaReportAllByToday"].QueryRow(currentTime.Format("2006-01-02"))
// 	res := new(models.MahasiswaSummary)
// 	err := row.Scan(&res.MahasiswaCount, &res.MahasiswaTotal)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

// func (r *MahasiswaRepository) FindMonthlyReportOder() (*models.MahasiswaSummary, error) {
// 	currentTime := time.Now()
// 	row := r.ps["mahasiswaReportAllByMonthYear"].QueryRow(currentTime.Month(), currentTime.Year())
// 	res := new(models.MahasiswaSummary)
// 	err := row.Scan(&res.MahasiswaCount, &res.MahasiswaTotal)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

// func (r *MahasiswaRepository) FindReportMahasiswa() (*models.MahasiswaSummary, error) {
// 	row := r.ps["mahasiswaReportAllByMonthYear"].QueryRow()
// 	res := new(models.MahasiswaSummary)
// 	err := row.Scan(&res.MahasiswaCount, &res.MahasiswaTotal)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }
