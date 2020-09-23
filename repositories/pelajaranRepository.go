package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github/coroo/enigma-school/models"
	"strings"
)

type IPelajaranRepository interface {
	Insert(pelajaran models.Pelajaran) (*models.Pelajaran, error)
	Update(pelajaran models.Pelajaran) (*models.Pelajaran, error)
	FindOneById(kodePelajaran string) (*models.Pelajaran, error)
	FindAll() ([]*models.Pelajaran, error)
	DeleteOneById(kodePelajaran string) (*models.Pelajaran, error)
}

var (
	pelajaranQueries = map[string]string{
		"deleteOneById":        "DELETE FROM m_pelajaran WHERE kode_pelajaran=?",
		"pelajaranFindAll":     "select kode_pelajaran,nama_pelajaran,jumlah_sks from m_pelajaran",
		"pelajaranFindOneById": "select kode_pelajaran,nama_pelajaran,jumlah_sks from m_pelajaran where kode_pelajaran=?",
		"insertPelajaran":      "insert into m_pelajaran(kode_pelajaran,nama_pelajaran,jumlah_sks) values(?,?,?)",
		"updatePelajaran":      "UPDATE m_pelajaran SET nama_pelajaran = ?,jumlah_sks = ? WHERE kode_pelajaran = ?",
	}
)

type PelajaranRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewPelajaranRepository(db *sql.DB) IPelajaranRepository {
	ps := make(map[string]*sql.Stmt, len(pelajaranQueries))
	for n, v := range pelajaranQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &PelajaranRepository{
		db, ps,
	}
}

func (r *PelajaranRepository) Insert(pelajaran models.Pelajaran) (*models.Pelajaran, error) {
	res, err := r.ps["insertPelajaran"].Exec(strings.ToUpper(pelajaran.KodePelajaran), pelajaran.NamaPelajaran, pelajaran.JumlahSks)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &pelajaran, nil
}

func (r *PelajaranRepository) Update(pelajaran models.Pelajaran) (*models.Pelajaran, error) {
	res, err := r.ps["updatePelajaran"].Exec(pelajaran.NamaPelajaran, pelajaran.JumlahSks, strings.ToUpper(pelajaran.KodePelajaran))
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Update failed", err))
	}
	return &pelajaran, nil
}

func (r *PelajaranRepository) FindOneById(kodePelajaran string) (*models.Pelajaran, error) {
	row := r.ps["pelajaranFindOneById"].QueryRow(strings.ToUpper(kodePelajaran))
	res := new(models.Pelajaran)
	err := row.Scan(&res.KodePelajaran, &res.NamaPelajaran, &res.JumlahSks)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *PelajaranRepository) FindAll() ([]*models.Pelajaran, error) {
	rows, err := r.ps["pelajaranFindAll"].Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Pelajaran, 0)
	for rows.Next() {
		res := new(models.Pelajaran)
		err := rows.Scan(&res.KodePelajaran, &res.NamaPelajaran, &res.JumlahSks)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}

func (r *PelajaranRepository) DeleteOneById(kodePelajaran string) (*models.Pelajaran, error) {
	row := r.ps["deleteOneById"].QueryRow(strings.ToUpper(kodePelajaran))
	res := new(models.Pelajaran)
	err := row.Scan(&res.KodePelajaran, &res.NamaPelajaran, &res.JumlahSks)
	if err != nil {
		return res, err
	}
	return res, nil
}
