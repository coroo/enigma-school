package models

import validation "github.com/go-ozzo/ozzo-validation"

type MahasiswaWithPelajaran struct {
	Id          string
	TahunAjaran string
	Mahasiswa
	Pelajaran
}

type Penempatan struct {
	Id            string
	Nim           string
	KodePelajaran string
	TahunAjaran   string
}

func (a Penempatan) Validate() error {
	return validation.ValidateStruct(&a,
		// Nim cannot be empty, and the length must between 3 and 10
		validation.Field(&a.TahunAjaran, validation.Required, validation.Length(3, 10)),
		// NamaPenempatan cannot be empty, and the length must between 5 and 50
		// validation.Field(&a.NamaPenempatan, validation.Required, validation.Length(5, 50)),
		// JurusanPenempatan cannot be empty
		// validation.Field(&a.JurusanPenempatan, validation.Required),
	)
}

// type MataKuliah struct {
// 	MataKuliahCount int64
// 	MataKuliahTotal int
// }
